package service

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"oracle/chaincall"
)

type Oracle struct {
	ctx                  context.Context
	VORCoordinatorCaller *chaincall.VORCoordinatorCaller
}

func NewOracle(ctx context.Context, caller *chaincall.VORCoordinatorCaller) *Oracle {
	return &Oracle{ctx: ctx, VORCoordinatorCaller: caller}
}

func (d *Oracle) Withdraw(address string, amount int64) (*types.Transaction, error) {
	return d.VORCoordinatorCaller.Withdraw(address, big.NewInt(amount))
}
