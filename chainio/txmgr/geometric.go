package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	hundred             = big.NewInt(100)
	queryTickerDuration = 3 * time.Second
)

type transaction struct {
	*types.Transaction
	TxID        wallet.TxID
	requestedAt time.Time
}

type txnRequest struct {
	tx *types.Transaction

	requestedAt time.Time
	// txAttempts are the transactions that have been attempted to be mined for this request.
	// If a transaction hasn't been confirmed within the timeout and a replacement transaction is sent,
	// the original transaction hash will be kept in this slice
	txAttempts []*transaction
}

type EthBackend interface {
	BlockNumber(ctx context.Context) (uint64, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

type GeometricTxManager struct {
	// FIXME: is this mutex still needed?
	mu sync.Mutex

	ethClient EthBackend
	wallet    wallet.Wallet
	logger    logging.Logger
	metrics   *Metrics

	// consts
	params GeometricTxnManagerParams
}

var _ TxManager = (*GeometricTxManager)(nil)

type GeometricTxnManagerParams struct {
	// number of blocks to wait for a transaction to be confirmed
	confirmationBlocks int
	// time to wait for a transaction to be broadcasted to the network
	// could be direct via eth_sendRawTransaction or indirect via a wallet service such as fireblocks
	txnBroadcastTimeout time.Duration
	// time to wait for a transaction to be confirmed (mined + confirmationBlocks blocks)
	txnConfirmationTimeout time.Duration
	// max number of times to retry sending a transaction before failing
	// this applies to every transaction attempt when a nonce is bumped
	maxSendTransactionRetry int
	// percentage multiplier for gas price. It needs to be >= 10 to properly replace existing transaction
	// e.g. 10 means 10% increase
	gasPricePercentageMultiplier *big.Int
	// default gas tip cap to use when eth_maxPriorityFeePerGas is not available
	FallbackGasTipCap uint64
	// percentage multiplier for gas limit. Should be >= 100
	GasMultiplierPercentage uint64
	// percentage multiplier for gas tip. Should be >= 100
	GasTipMultiplierPercentage uint64
}

var defaultParams = GeometricTxnManagerParams{
	confirmationBlocks:           0,                     // tx mined is considered confirmed
	txnBroadcastTimeout:          2 * time.Minute,       // fireblocks has had issues so we give it a long time
	txnConfirmationTimeout:       5 * 12 * time.Second,  // 5 blocks
	maxSendTransactionRetry:      3,                     // arbitrary
	gasPricePercentageMultiplier: big.NewInt(10),        // 10%
	FallbackGasTipCap:            uint64(5_000_000_000), // 5 gwei
	GasMultiplierPercentage:      uint64(120),           // add an extra 20% gas buffer to the gas limit
	GasTipMultiplierPercentage:   uint64(125),           // add an extra 25% to the gas tip
}

func fillParamsWithDefaultValues(params *GeometricTxnManagerParams) {
	if params.confirmationBlocks == 0 {
		params.confirmationBlocks = defaultParams.confirmationBlocks
	}
	if params.txnBroadcastTimeout == 0 {
		params.txnBroadcastTimeout = defaultParams.txnBroadcastTimeout
	}
	if params.txnConfirmationTimeout == 0 {
		params.txnConfirmationTimeout = defaultParams.txnConfirmationTimeout
	}
	if params.maxSendTransactionRetry == 0 {
		params.maxSendTransactionRetry = defaultParams.maxSendTransactionRetry
	}
	if params.gasPricePercentageMultiplier == nil {
		params.gasPricePercentageMultiplier = defaultParams.gasPricePercentageMultiplier
	}
	if params.FallbackGasTipCap == 0 {
		params.FallbackGasTipCap = defaultParams.FallbackGasTipCap
	}
	if params.GasMultiplierPercentage == 0 {
		params.GasMultiplierPercentage = defaultParams.GasMultiplierPercentage
	}
	if params.GasTipMultiplierPercentage == 0 {
		params.GasTipMultiplierPercentage = defaultParams.GasTipMultiplierPercentage
	}
}

func NewGeometricTxnManager(
	ethClient EthBackend,
	wallet wallet.Wallet,
	logger logging.Logger,
	metrics *Metrics,
	params GeometricTxnManagerParams,
) *GeometricTxManager {
	fillParamsWithDefaultValues(&params)
	return &GeometricTxManager{
		ethClient: ethClient,
		wallet:    wallet,
		logger:    logger.With("component", "GeometricTxManager"),
		metrics:   metrics,
		params:    params,
	}
}

// GetNoSendTxOpts This generates a noSend TransactOpts so that we can use
// this to generate the transaction without actually sending it
func (m *GeometricTxManager) GetNoSendTxOpts() (*bind.TransactOpts, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	from, err := m.wallet.SenderAddress(ctxWithTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to get sender address: %w", err)
	}
	return &bind.TransactOpts{
		From:   from,
		NoSend: true,
		Signer: NoopSigner,
	}, nil
}

func newTxnRequest(tx *types.Transaction) *txnRequest {
	return &txnRequest{
		tx:          tx,
		requestedAt: time.Now(),
		txAttempts:  make([]*transaction, 0),
	}
}

func (t *GeometricTxManager) Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return t.processTransaction(ctx, newTxnRequest(tx))
}

// ProcessTransaction sends the transaction and runs the monitoring loop which will bump the gasPrice until the tx get
// included.
func (t *GeometricTxManager) processTransaction(ctx context.Context, req *txnRequest) (*types.Receipt, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.logger.Debug("new transaction",
		"nonce", req.tx.Nonce(), "gasFeeCap", req.tx.GasFeeCap(), "gasTipCap", req.tx.GasTipCap(),
	)
	t.metrics.IncrementProcessingTxCount()
	defer t.metrics.DecrementProcessingTxCount()

	var txn *types.Transaction
	var txID wallet.TxID
	var err error
	retryFromFailure := 0
	from, err := t.wallet.SenderAddress(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sender address: %w", err)
	}
	for retryFromFailure < t.params.maxSendTransactionRetry {
		gasTipCap, err := t.estimateGasTipCap(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas tip cap: %w", err)
		}
		txn, err = t.updateGasTipCap(ctx, req.tx, gasTipCap, from)
		if err != nil {
			return nil, fmt.Errorf("failed to update gas price: %w", err)
		}
		txID, err = t.wallet.SendTransaction(ctx, txn)
		// TODO: what is this...? does it come from the fireblocks wallet?
		var timeout interface{ Timeout() bool }
		if errors.As(err, &timeout); timeout != nil && timeout.Timeout() {
			t.logger.Warn(
				"failed to send txn due to timeout",
				"hash",
				txn.Hash().Hex(),
				"numRetries",
				retryFromFailure,
				"maxRetry",
				t.params.maxSendTransactionRetry,
				"err",
				err,
			)
			// TODO: why do we only retry on client or server timeouts? what about other errors?
			retryFromFailure++
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to send txn %s: %w", txn.Hash().Hex(), err)
		} else {
			t.logger.Debug("successfully sent txn", "txID", txID, "txHash", txn.Hash().Hex())
			break
		}
	}

	if txn == nil || txID == "" {
		return nil, fmt.Errorf("failed to send txn %s: %w", req.tx.Hash().Hex(), err)
	}

	req.tx = txn
	req.txAttempts = append(req.txAttempts, &transaction{
		TxID:        txID,
		Transaction: txn,
		requestedAt: time.Now(),
	})

	receipt, err := t.monitorTransaction(ctx, req)
	if err == nil {
		if receipt.GasUsed > 0 {
			t.metrics.ObserveGasUsedWei(receipt.GasUsed)
		}
		t.metrics.ObserveConfirmationLatencyMs(time.Since(req.requestedAt).Milliseconds())
	}
	return receipt, err
}

// ensureAnyTransactionBroadcasted waits until all given transactions are broadcasted to the network.
func (t *GeometricTxManager) ensureAnyTransactionBroadcasted(ctx context.Context, txs []*transaction) error {
	queryTicker := time.NewTicker(queryTickerDuration)
	defer queryTicker.Stop()

	for {
		for _, tx := range txs {
			_, err := t.wallet.GetTransactionReceipt(ctx, tx.TxID)
			if err == nil || errors.Is(err, wallet.ErrReceiptNotYetAvailable) {
				t.metrics.ObserveBroadcastLatencyMs(time.Since(tx.requestedAt).Milliseconds())
				return nil
			}
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (t *GeometricTxManager) ensureAnyTransactionConfirmed(
	ctx context.Context,
	txs []*transaction,
) (*types.Receipt, error) {
	queryTicker := time.NewTicker(queryTickerDuration)
	defer queryTicker.Stop()
	var receipt *types.Receipt
	var err error
	// transactions that need to be queried. Some transactions will be removed from this map depending on their status.
	txnsToQuery := make(map[wallet.TxID]*types.Transaction, len(txs))
	for _, tx := range txs {
		txnsToQuery[tx.TxID] = tx.Transaction
	}

	for {
		for txID, tx := range txnsToQuery {
			receipt, err = t.wallet.GetTransactionReceipt(ctx, txID)
			if err == nil {
				chainTip, err := t.ethClient.BlockNumber(ctx)
				if err == nil {
					if receipt.BlockNumber.Uint64()+uint64(t.params.confirmationBlocks) > chainTip {
						t.logger.Debug(
							"transaction has been mined but don't have enough confirmations at current chain tip",
							"txnBlockNumber",
							receipt.BlockNumber.Uint64(),
							"numConfirmations",
							t.params.confirmationBlocks,
							"chainTip",
							chainTip,
						)
						break
					} else {
						return receipt, nil
					}
				} else {
					t.logger.Debug("failed to get chain tip while waiting for transaction to mine", "err", err)
				}
			}

			// TODO(samlaf): how to maintain these better? How do we know which errors to use and where they are
			// returned from?
			if errors.Is(err, ethereum.NotFound) || errors.Is(err, wallet.ErrReceiptNotYetAvailable) {
				t.logger.Debug("Transaction not yet mined", "txID", txID, "txHash", tx.Hash().Hex(), "err", err)
			} else if errors.Is(err, wallet.ErrTransactionFailed) {
				t.logger.Debug("Transaction failed", "txID", txID, "txHash", tx.Hash().Hex(), "err", err)
				delete(txnsToQuery, txID)
			} else if errors.Is(err, wallet.ErrNotYetBroadcasted) {
				t.logger.Error("Transaction has not been broadcasted to network but attempted to retrieve receipt", "err", err)
			} else {
				t.logger.Debug("Transaction receipt retrieval failed", "err", err)
			}
		}

		if len(txnsToQuery) == 0 {
			return nil, fmt.Errorf("all transactions failed")
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return receipt, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

// monitorTransaction waits until the transaction is confirmed (or failed) and resends it with a higher gas price if it
// is not mined without a timeout.
// It returns the receipt once the transaction has been confirmed.
// It returns an error if the transaction fails to be sent for reasons other than timeouts.
func (t *GeometricTxManager) monitorTransaction(ctx context.Context, req *txnRequest) (*types.Receipt, error) {
	numSpeedUps := 0
	retryFromFailure := 0

	var receipt *types.Receipt
	var err error

	rpcCallAttempt := func() error {
		t.logger.Debug("monitoring transaction", "txHash", req.tx.Hash().Hex(), "nonce", req.tx.Nonce())

		ctxWithTimeout, cancelBroadcastTimeout := context.WithTimeout(ctx, t.params.txnBroadcastTimeout)
		defer cancelBroadcastTimeout()

		// Ensure transactions are broadcasted to the network before querying the receipt.
		// This is to avoid querying the receipt of a transaction that hasn't been broadcasted yet.
		// For example, when Fireblocks wallet is used, there may be delays in broadcasting the transaction due to
		// latency from cosigning and MPC operations.
		err = t.ensureAnyTransactionBroadcasted(ctxWithTimeout, req.txAttempts)
		if err != nil && errors.Is(err, context.DeadlineExceeded) {
			t.logger.Warn(
				"transaction not broadcasted within timeout",
				"txHash",
				req.tx.Hash().Hex(),
				"nonce",
				req.tx.Nonce(),
			)
			fireblocksWallet, ok := t.wallet.(interface {
				CancelTransactionBroadcast(ctx context.Context, txID wallet.TxID) (bool, error)
			})
			if ok {
				// Consider these transactions failed as they haven't been broadcasted within timeout.
				// Cancel these transactions to avoid blocking the next transactions.
				for _, tx := range req.txAttempts {
					cancelled, err := fireblocksWallet.CancelTransactionBroadcast(ctx, tx.TxID)
					if err != nil {
						t.logger.Warn("failed to cancel Fireblocks transaction broadcast", "txID", tx.TxID, "err", err)
					} else if cancelled {
						t.logger.Info("cancelled Fireblocks transaction broadcast because it didn't get broadcasted within timeout", "txID", tx.TxID, "timeout", t.params.txnBroadcastTimeout.String())
					}
				}
			}
			return errors.New("transaction not broadcasted")
		} else if err != nil {
			t.logger.Error("unexpected error while waiting for Fireblocks transaction to broadcast", "txHash", req.tx.Hash().Hex(), "err", err)
			return err
		}

		ctxWithTimeout, cancelEvaluationTimeout := context.WithTimeout(ctx, t.params.txnConfirmationTimeout)
		defer cancelEvaluationTimeout()
		receipt, err = t.ensureAnyTransactionConfirmed(
			ctxWithTimeout,
			req.txAttempts,
		)
		return err
	}

	for {
		err = rpcCallAttempt()
		if err == nil {
			t.metrics.ObserveSpeedups(numSpeedUps)
			t.metrics.IncrementProcessedTxsTotal("success")
			return receipt, nil
		}

		if errors.Is(err, context.DeadlineExceeded) {
			if receipt != nil {
				t.logger.Warn(
					"transaction has been mined, but hasn't accumulated the required number of confirmations",
					"txHash",
					req.tx.Hash().Hex(),
					"nonce",
					req.tx.Nonce(),
				)
				continue
			}
			t.logger.Warn(
				"transaction not mined within timeout, resending with higher gas price",
				"txHash",
				req.tx.Hash().Hex(),
				"nonce",
				req.tx.Nonce(),
			)
			newTx, err := t.speedUpTxn(ctx, req.tx)
			if err != nil {
				t.logger.Error("failed to speed up transaction", "err", err)
				t.metrics.IncrementProcessedTxsTotal("failure")
				return nil, err
			}
			txID, err := t.wallet.SendTransaction(ctx, newTx)
			if err != nil {
				if retryFromFailure >= t.params.maxSendTransactionRetry {
					t.logger.Warn(
						"failed to send txn - retries exhausted",
						"txn",
						req.tx.Hash().Hex(),
						"attempt",
						retryFromFailure,
						"maxRetry",
						t.params.maxSendTransactionRetry,
						"err",
						err,
					)
					t.metrics.IncrementProcessedTxsTotal("failure")
					return nil, err
				} else {
					t.logger.Warn("failed to send txn - retrying", "txn", req.tx.Hash().Hex(), "attempt", retryFromFailure, "maxRetry", t.params.maxSendTransactionRetry, "err", err)
				}
				retryFromFailure++
				continue
			}

			t.logger.Debug("successfully sent txn", "txID", txID, "txHash", newTx.Hash().Hex())
			req.tx = newTx
			req.txAttempts = append(req.txAttempts, &transaction{
				TxID:        txID,
				Transaction: newTx,
			})
			numSpeedUps++
		} else {
			t.logger.Error("transaction failed", "txHash", req.tx.Hash().Hex(), "err", err)
			t.metrics.IncrementProcessedTxsTotal("failure")
			return nil, err
		}
	}
}

// speedUpTxn increases the gas price of the existing transaction by specified percentage.
// It makes sure the new gas price is not lower than the current gas price.
func (t *GeometricTxManager) speedUpTxn(
	ctx context.Context,
	tx *types.Transaction,
) (*types.Transaction, error) {
	// bump the current gasTip, and also reestimate it from the node, and take the highest value
	var newGasTipCap *big.Int
	{
		estimatedGasTipCap, err := t.estimateGasTipCap(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate gas tip cap: %w", err)
		}
		bumpedGasTipCap := t.addGasTipCapBuffer(tx.GasTipCap())
		if estimatedGasTipCap.Cmp(bumpedGasTipCap) > 0 {
			newGasTipCap = estimatedGasTipCap
		} else {
			newGasTipCap = bumpedGasTipCap
		}
	}

	from, err := t.wallet.SenderAddress(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sender address: %w", err)
	}
	newTx, err := t.updateGasTipCap(ctx, tx, newGasTipCap, from)
	if err != nil {
		return nil, fmt.Errorf("failed to update gas price: %w", err)
	}
	t.logger.Info(
		"increasing gas price",
		"prevTxHash", tx.Hash().Hex(), "newTxHash", newTx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"prevGasTipCap", tx.GasTipCap(), "newGasTipCap", newGasTipCap,
		"prevGasFeeCap", tx.GasFeeCap(), "newGasFeeCap", newTx.GasFeeCap(),
	)
	return newTx, nil
}

// UpdateGasParams updates the three gas related parameters of a transaction:
// - gasTipCap: calls the json-rpc method eth_maxPriorityFeePerGas and
// adds a extra buffer based on o.params.GasTipMultiplierPercentage
// - gasFeeCap: calculates the gas fee cap as 2 * baseFee + gasTipCap
// - gasLimit: calls the json-rpc method eth_estimateGas and
// adds a extra buffer based on o.params.GasMultiplierPercentage
func (t *GeometricTxManager) updateGasTipCap(
	ctx context.Context,
	tx *types.Transaction,
	newGasTipCap *big.Int,
	from common.Address,
) (*types.Transaction, error) {
	gasFeeCap, err := t.estimateGasFeeCap(ctx, newGasTipCap)
	if err != nil {
		return nil, utils.WrapError("failed to estimate gas fee cap", err)
	}

	// we reestimate the gas limit because the state of the chain may have changed,
	// which could cause the previous gas limit to be insufficient
	gasLimit, err := t.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From:      from,
		To:        tx.To(),
		GasTipCap: newGasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     tx.Value(),
		Data:      tx.Data(),
	})
	if err != nil {
		return nil, utils.WrapError("failed to estimate gas", err)
	}
	// we also add a buffer to the gas limit to account for potential changes in the state of the chain
	// between the time of estimation and the time the transaction is mined
	bufferedGasLimit := t.addGasBuffer(gasLimit)

	return types.NewTx(&types.DynamicFeeTx{
		ChainID:    tx.ChainId(),
		Nonce:      tx.Nonce(),
		GasTipCap:  newGasTipCap,
		GasFeeCap:  gasFeeCap,
		Gas:        bufferedGasLimit,
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	}), nil
}

func (t *GeometricTxManager) estimateGasTipCap(ctx context.Context) (gasTipCap *big.Int, err error) {
	gasTipCap, err = t.ethClient.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		// TODO: error could actually come from node not being down or network being slow, etc.
		t.logger.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = big.NewInt(0).SetUint64(t.params.FallbackGasTipCap)
	}
	return t.addGasTipCapBuffer(gasTipCap), nil
}

// addGasTipCapBuffer adds a buffer to the gas tip cap to account for potential changes in the state of the chain
// The result is returned in a new big.Int to avoid modifying the input gasTipCap.
func (t *GeometricTxManager) addGasTipCapBuffer(gasTipCap *big.Int) *big.Int {
	bumpedGasTipCap := new(big.Int).Set(gasTipCap)
	return bumpedGasTipCap.Mul(bumpedGasTipCap, big.NewInt(int64(t.params.GasTipMultiplierPercentage))).
		Div(bumpedGasTipCap, big.NewInt(100))
}

// estimateGasFeeCap returns the gas fee cap for a transaction, calculated as:
// gasFeeCap = 2 * baseFee + gasTipCap
// Rationale: https://www.blocknative.com/blog/eip-1559-fees
// The result is returned in a new big.Int to avoid modifying gasTipCap.
func (t *GeometricTxManager) estimateGasFeeCap(ctx context.Context, gasTipCap *big.Int) (*big.Int, error) {
	header, err := t.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, utils.WrapError("failed to get latest header", err)
	}
	return new(big.Int).Add(new(big.Int).Mul(header.BaseFee, big.NewInt(2)), gasTipCap), nil
}

func (t *GeometricTxManager) addGasBuffer(gasLimit uint64) uint64 {
	return t.params.GasMultiplierPercentage * gasLimit / 100
}
