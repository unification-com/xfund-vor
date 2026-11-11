package chaincall

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"oracle/contracts/vord_20"
)

// VORD20Caller drives the example dice-roll consumer in the integration tests.
type VORD20Caller struct {
	*transactor
	instance *vord_20.Vord20
}

func NewVORD20Caller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*VORD20Caller, error) {
	t, err := connect(ethHostAddress, oraclePrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	instance, err := vord_20.NewVord20(common.HexToAddress(contractStringAddress), t.client)
	if err != nil {
		return nil, err
	}
	return &VORD20Caller{transactor: t, instance: instance}, nil
}

func (d *VORD20Caller) IncreaseVORCoordinatorAllowance(amount *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.instance.IncreaseVorAllowance(opts, amount)
	})
}

func (d *VORD20Caller) RollDice(seed *big.Int, roller string) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.instance.RollDice(opts, seed, common.HexToAddress(roller))
	})
}

func (d *VORD20Caller) SetFee(fee *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.instance.SetFee(opts, fee)
	})
}

func (d *VORD20Caller) SetKeyHash(keyHash [32]byte) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.instance.SetKeyHash(opts, keyHash)
	})
}

func (d *VORD20Caller) Owner() (common.Address, error) {
	return d.instance.Owner(d.callOpts())
}

func (d *VORD20Caller) KeyHash() ([32]byte, error) {
	return d.instance.KeyHash(d.callOpts())
}

func (d *VORD20Caller) Fee() (*big.Int, error) {
	return d.instance.Fee(d.callOpts())
}

func (d *VORD20Caller) House(roller string) (string, error) {
	return d.instance.House(d.callOpts(), common.HexToAddress(roller))
}
