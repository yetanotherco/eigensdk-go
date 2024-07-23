package geometric

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

var chainId = big.NewInt(31337)

func TestGeometricTxManager(t *testing.T) {
	logger := testutils.NewTestLogger()
	ethBackend := &StubEthBackend{}
	ecdsaSk, ecdsaAddr, error := testutils.NewEcdsaSkAndAddress()
	if error != nil {
		t.Fatal(error)
	}

	signerFn, error := signerv2.PrivateKeySignerFn(ecdsaSk, chainId)
	if error != nil {
		t.Fatal(error)
	}
	wallet.NewPrivateKeyWallet(ethBackend, signerFn, ecdsaAddr, logger)
	txmgr := NewGeometricTxnManager(ethBackend)
}

type StubEthBackend struct {
	blockNumber uint64
}

func (s *StubEthBackend) BlockNumber() (uint64, error) {
	return s.blockNumber, nil
}

func (s *StubEthBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (s *StubEthBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, nil

}
func (s *StubEthBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return &types.Header{BaseFee: big.NewInt(0)}, nil
}
