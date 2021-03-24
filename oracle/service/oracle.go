package service

import (
	"context"
	"oracle/chaincall"
)

type Oracle struct {
	ctx                  context.Context
	VORCoordinatorCaller *chaincall.VORCoordinatorCaller
}

func NewOracle(ctx context.Context, caller *chaincall.VORCoordinatorCaller) *Oracle {
	return &Oracle{ctx: ctx, VORCoordinatorCaller: caller}
}
