package txmgr

import (
	"testing"
)

func TestGeometricTxManager(t *testing.T) {

}

type StubEthBackend struct {
	blockNumber uint64
}

func (s *StubEthBackend) BlockNumber() (uint64, error) {
	return s.blockNumber, nil
}
