package chaincall

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"oracle/contracts/mock_erc20"
)

// MockERC20Caller drives the mock xFUND token in the integration tests.
type MockERC20Caller struct {
	*transactor
	instance *mock_erc20.MockErc20
}

func NewMockERC20Caller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*MockERC20Caller, error) {
	t, err := connect(testChain(ethHostAddress, chainID), oraclePrivateKey)
	if err != nil {
		return nil, err
	}
	instance, err := mock_erc20.NewMockErc20(common.HexToAddress(contractStringAddress), t.client)
	if err != nil {
		return nil, err
	}
	return &MockERC20Caller{transactor: t, instance: instance}, nil
}

func (d *MockERC20Caller) Transfer(recipientAddress string, amount *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.instance.Transfer(opts, common.HexToAddress(recipientAddress), amount)
	})
}
