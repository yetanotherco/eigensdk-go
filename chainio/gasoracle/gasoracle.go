package gasoracle

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Params struct {
	FallbackGasTipCap          uint64
	GasMultiplierPercentage    uint64
	GasTipMultiplierPercentage uint64
}

var defaultParams = Params{
	FallbackGasTipCap:          uint64(5_000_000_000), // 5 gwei
	GasMultiplierPercentage:    uint64(120),           // add an extra 20% gas buffer to the gas limit
	GasTipMultiplierPercentage: uint64(125),           // add an extra 25% to the gas tip
}

type GasOracle struct {
	params Params
	client eth.Client
	logger logging.Logger
}

// params are optional gas parameters any of which will be filled with default values if not provided
func New(client eth.Client, logger logging.Logger, params Params) *GasOracle {
	if params.FallbackGasTipCap == 0 {
		params.FallbackGasTipCap = defaultParams.FallbackGasTipCap
	}
	if params.GasMultiplierPercentage == 0 {
		params.GasMultiplierPercentage = defaultParams.GasMultiplierPercentage
	}
	if params.GasTipMultiplierPercentage == 0 {
		params.GasTipMultiplierPercentage = defaultParams.GasTipMultiplierPercentage
	}
	return &GasOracle{
		params: params,
		client: client,
		logger: logger,
	}
}

// TODO: should this be part of the public API?
func (o *GasOracle) GetLatestGasCaps(ctx context.Context) (gasTipCap, gasFeeCap *big.Int, err error) {
	gasTipCap, err = o.client.SuggestGasTipCap(ctx)
	if err != nil {
		// If the transaction failed because the backend does not support
		// eth_maxPriorityFeePerGas, fallback to using the default constant.
		// Currently Alchemy is the only backend provider that exposes this
		// method, so in the event their API is unreachable we can fallback to a
		// degraded mode of operation. This also applies to our test
		// environments, as hardhat doesn't support the query either.
		o.logger.Info("eth_maxPriorityFeePerGas is unsupported by current backend, using fallback gasTipCap")
		gasTipCap = big.NewInt(0).SetUint64(o.params.FallbackGasTipCap)
	}

	gasTipCap.Mul(gasTipCap, big.NewInt(int64(o.params.GasTipMultiplierPercentage))).Div(gasTipCap, big.NewInt(100))

	header, err := o.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, utils.WrapError("failed to get latest header", err)
	}
	gasFeeCap = getGasFeeCap(gasTipCap, header.BaseFee)
	return
}

// UpdateGasParams updates the three gas related parameters of a transaction:
// gasTipCap, gasFeeCap, and gasLimit. It uses the provided gasTipCap and gasFeeCap
// which could either come from o.GetLatestGasCaps, or be manually bumped by the caller,
// and estimates the gas limit based on the transaction.
// TODO: should we add a public method to also bump the gasTipCap and gasFeeCap, instead of forcing client to do it
// themselves?
func (o *GasOracle) UpdateGasParams(
	ctx context.Context,
	tx *types.Transaction,
	gasTipCap, gasFeeCap *big.Int,
	from common.Address,
) (*types.Transaction, error) {

	// we reestimate the gas limit because the state of the chain may have changed,
	// which could cause the previous gas limit to be insufficient
	gasLimit, err := o.client.EstimateGas(ctx, ethereum.CallMsg{
		From:      from,
		To:        tx.To(),
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     tx.Value(),
		Data:      tx.Data(),
	})
	if err != nil {
		return nil, utils.WrapError("failed to estimate gas", err)
	}
	// we also add a buffer to the gas limit to account for potential changes in the state of the chain
	// between the time of estimation and the time the transaction is mined
	bufferedGasLimit := o.addGasBuffer(gasLimit)

	return types.NewTx(&types.DynamicFeeTx{
		ChainID:    tx.ChainId(),
		Nonce:      tx.Nonce(),
		GasTipCap:  gasTipCap,
		GasFeeCap:  gasFeeCap,
		Gas:        bufferedGasLimit,
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	}), nil
}

func (o *GasOracle) addGasBuffer(gasLimit uint64) uint64 {
	return o.params.GasMultiplierPercentage * gasLimit / 100
}

// getGasFeeCap returns the gas fee cap for a transaction, calculated as:
// gasFeeCap = 2 * baseFee + gasTipCap
// Rationale: https://www.blocknative.com/blog/eip-1559-fees
func getGasFeeCap(gasTipCap *big.Int, baseFee *big.Int) *big.Int {
	return new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), gasTipCap)
}
