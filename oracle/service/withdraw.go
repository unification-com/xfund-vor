package service

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (d *Oracle) Withdraw(address string, amount int64) (*types.Transaction, error) {
	return d.VORCoordinatorCaller.Withdraw(address, big.NewInt(amount))
}
