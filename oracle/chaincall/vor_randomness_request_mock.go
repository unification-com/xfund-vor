package chaincall

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"oracle/contracts/vor_randomness_request_mock"
)

// VORRandomnessRequestMockCaller drives the mock request contract in the integration tests.
type VORRandomnessRequestMockCaller struct {
	*transactor
	instance *vor_randomness_request_mock.VorRandomnessRequestMock
}

func NewVORRandomnessRequestMockCaller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*VORRandomnessRequestMockCaller, error) {
	t, err := connect(ethHostAddress, oraclePrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	instance, err := vor_randomness_request_mock.NewVorRandomnessRequestMock(common.HexToAddress(contractStringAddress), t.client)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockCaller{transactor: t, instance: instance}, nil
}

func (d *VORRandomnessRequestMockCaller) RandomnessRequest(keyHash [32]byte, consumerSeed *big.Int, feePaid *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.instance.RandomnessRequest(opts, keyHash, consumerSeed, feePaid)
	})
}
