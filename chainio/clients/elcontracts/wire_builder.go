//go:build wireinject
// +build wireinject

package elcontracts

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/google/wire"
)

func NewReaderFromConfig(
	cfg Config,
	ethClient eth.Client,
	logger logging.Logger,
) (*ChainReader, error) {
	wire.Build(NewBindingsFromConfig, NewChainReaderFromBindings)
	return &ChainReader{}, nil
}

func NewWriterFromConfig(
	cfg Config,
	ethClient eth.Client,
	logger logging.Logger,
	eigenMetrics metrics.Metrics,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	wire.Build(
		NewBindingsFromConfig,
		NewChainReaderFromBindings,
		wire.Bind(new(Reader), new(*ChainReader)),
		NewChainWriterFromBindings,
	)
	return &ChainWriter{}, nil
}

type ClientsAndBindings struct {
	ChainReader      *ChainReader
	ChainWriter      *ChainWriter
	ContractBindings *ContractBindings
}

func buildClientsFromConfig(
	cfg Config,
	ethClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*ClientsAndBindings, error) {
	wire.Build(
		NewBindingsFromConfig,
		NewChainReaderFromBindings,
		wire.Bind(new(Reader), new(*ChainReader)),
		wire.Bind(new(metrics.Metrics), new(*metrics.EigenMetrics)),
		NewChainWriterFromBindings,
		wire.Struct(new(ClientsAndBindings), "*"),
	)
	return &ClientsAndBindings{}, nil
}

func BuildClientsFromConfig(
	cfg Config,
	ethClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*ChainReader, *ChainWriter, *ContractBindings, error) {
	clientsAndBindings, err := buildClientsFromConfig(cfg, ethClient, txMgr, logger, eigenMetrics)
	if err != nil {
		return nil, nil, nil, err
	}
	return clientsAndBindings.ChainReader, clientsAndBindings.ChainWriter, clientsAndBindings.ContractBindings, nil
}