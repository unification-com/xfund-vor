package service

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (d *Oracle) ChangeFee(amount int64) (*types.Transaction, error) {
	return d.VORCoordinatorCaller.ChangeFee(big.NewInt(amount))
}
