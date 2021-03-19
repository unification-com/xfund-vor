package service

import (
	"context"
	"oracle/chaincall"
)

type Service struct {
	Oracle *Oracle
}

func NewService(ctx context.Context, caller *chaincall.VORCoordinatorCaller) *Service {
	return &Service{
		Oracle: NewOracle(ctx, caller),
	}
}
