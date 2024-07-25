package geometric

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

var chainId = big.NewInt(31337)

type testHarness struct {
	ethBackend *fakeEthBackend
	txmgr      *GeometricTxManager
}

func newTestHarness(t *testing.T) *testHarness {
	logger := testutils.NewTestLogger()
	ethBackend := NewFakeEthBackend()

	ecdsaSk, ecdsaAddr, err := testutils.NewEcdsaSkAndAddress()
	require.NoError(t, err)

	signerFn, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaSk}, chainId)
	require.NoError(t, err)

	skWallet, err := wallet.NewPrivateKeyWallet(ethBackend, signerFn, ecdsaAddr, logger)
	require.NoError(t, err)

	txmgr := NewGeometricTxnManager(ethBackend, skWallet, logger, NewNoopMetrics(), GeometricTxnManagerParams{})

	return &testHarness{
		ethBackend: ethBackend,
		txmgr:      txmgr,
	}
}

func TestGeometricTxManager(t *testing.T) {
	t.Run("Send", func(t *testing.T) {
		h := newTestHarness(t)

		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainId,
			Nonce:     0,
			GasTipCap: big.NewInt(1),
			GasFeeCap: big.NewInt(1_000_000_000),
			Gas:       21000,
			To:        testutils.ZeroAddress(),
			Value:     big.NewInt(1),
		})
		ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5)
		defer cancel()
		_, err := h.txmgr.Send(ctxWithTimeout, tx)
		require.NoError(t, err)

	})
}

type fakeEthBackend struct {
	blockNumber uint64
	minedTxs    map[common.Hash]*types.Receipt
}

func NewFakeEthBackend() *fakeEthBackend {
	return &fakeEthBackend{
		blockNumber: 0,
		minedTxs:    make(map[common.Hash]*types.Receipt),
	}
}

func (s *fakeEthBackend) BlockNumber(context.Context) (uint64, error) {
	return s.blockNumber, nil
}

func (s *fakeEthBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (s *fakeEthBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, nil

}
func (s *fakeEthBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{BaseFee: big.NewInt(0)}, nil
}
func (s *fakeEthBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	s.minedTxs[tx.Hash()] = &types.Receipt{
		BlockNumber: big.NewInt(int64(s.blockNumber)),
	}
	s.blockNumber++
	return nil
}
func (s *fakeEthBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return s.minedTxs[txHash], nil
}
