package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/gasoracle"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

type GeometricTxManager struct {
	// FIXME: is this mutex still needed?
	mu sync.Mutex

	ethClient eth.Client
	gasOracle gasoracle.GasOracle
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
}

var defaultParams = GeometricTxnManagerParams{
	confirmationBlocks:           0,                    // tx mined is considered confirmed
	txnBroadcastTimeout:          2 * time.Minute,      // fireblocks has had issues so we give it a long time
	txnConfirmationTimeout:       5 * 12 * time.Second, // 5 blocks
	maxSendTransactionRetry:      3,                    // arbitrary
	gasPricePercentageMultiplier: big.NewInt(10),       // 10%
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
}

func NewGeometricTxnManager(
	ethClient eth.Client,
	gasOracle gasoracle.GasOracle,
	wallet wallet.Wallet,
	logger logging.Logger,
	metrics *Metrics,
	params GeometricTxnManagerParams,
) *GeometricTxManager {
	fillParamsWithDefaultValues(&params)
	return &GeometricTxManager{
		ethClient: ethClient,
		gasOracle: gasOracle,
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

// ProcessTransaction sends the transaction and runs the monitoring loop which will bump the gasPrice until the tx get included.
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
	for retryFromFailure < t.params.maxSendTransactionRetry {
		gasTipCap, gasFeeCap, err := t.gasOracle.GetLatestGasCaps(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get latest gas caps: %w", err)
		}

		from, err := t.wallet.SenderAddress(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get sender address: %w", err)
		}
		txn, err = t.gasOracle.UpdateGas(ctx, req.tx, req.tx.Value(), gasTipCap, gasFeeCap, from)
		if err != nil {
			return nil, fmt.Errorf("failed to update gas price: %w", err)
		}
		txID, err = t.wallet.SendTransaction(ctx, txn)
		// TODO: what is this...? does it come from the fireblocks wallet?
		var urlErr *url.Error
		didTimeout := false
		if errors.As(err, &urlErr) {
			didTimeout = urlErr.Timeout()
		}
		if didTimeout || errors.Is(err, context.DeadlineExceeded) {
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

			// TODO(samlaf): how to maintain these better? How do we know which errors to use and where they are returned from?
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
	prevGasTipCap := tx.GasTipCap()
	prevGasFeeCap := tx.GasFeeCap()
	// get the gas tip cap and gas fee cap based on current network condition
	currentGasTipCap, currentGasFeeCap, err := t.gasOracle.GetLatestGasCaps(ctx)
	if err != nil {
		return nil, err
	}
	increasedGasTipCap := increaseGasPrice(prevGasTipCap, t.params.gasPricePercentageMultiplier)
	increasedGasFeeCap := increaseGasPrice(prevGasFeeCap, t.params.gasPricePercentageMultiplier)
	// make sure increased gas prices are not lower than current gas prices
	var newGasTipCap, newGasFeeCap *big.Int
	if currentGasTipCap.Cmp(increasedGasTipCap) > 0 {
		newGasTipCap = currentGasTipCap
	} else {
		newGasTipCap = increasedGasTipCap
	}
	if currentGasFeeCap.Cmp(increasedGasFeeCap) > 0 {
		newGasFeeCap = currentGasFeeCap
	} else {
		newGasFeeCap = increasedGasFeeCap
	}

	t.logger.Info(
		"increasing gas price",
		"txHash",
		tx.Hash().Hex(),
		"nonce",
		tx.Nonce(),
		"prevGasTipCap",
		prevGasTipCap,
		"prevGasFeeCap",
		prevGasFeeCap,
		"newGasTipCap",
		newGasTipCap,
		"newGasFeeCap",
		newGasFeeCap,
	)
	from, err := t.wallet.SenderAddress(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sender address: %w", err)
	}
	return t.gasOracle.UpdateGas(ctx, tx, tx.Value(), newGasTipCap, newGasFeeCap, from)
}

// increaseGasPrice increases the gas price by specified percentage.
// i.e. gasPrice + ((gasPrice * gasPricePercentageMultiplier + 99) / 100)
func increaseGasPrice(gasPrice, gasPricePercentageMultiplier *big.Int) *big.Int {
	if gasPrice == nil {
		return nil
	}
	bump := new(big.Int).Mul(gasPrice, gasPricePercentageMultiplier)
	bump = roundUpDivideBig(bump, hundred)
	return new(big.Int).Add(gasPrice, bump)
}

func roundUpDivideBig(a, b *big.Int) *big.Int {
	if a == nil || b == nil || b.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	one := new(big.Int).SetUint64(1)
	num := new(big.Int).Sub(new(big.Int).Add(a, b), one) // a + b - 1
	res := new(big.Int).Div(num, b)                      // (a + b - 1) / b
	return res
}
