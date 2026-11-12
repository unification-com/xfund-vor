package service

import (
	"context"

	"oracle/chaincall"
	"oracle/config"
	"oracle/store"
)

// Service is the oracle's work on one chain.
type Service struct {
	ctx                  context.Context
	Chain                config.ChainConfig
	Store                *store.Store
	VORCoordinatorCaller *chaincall.VORCoordinatorCaller
}

func NewService(ctx context.Context, store *store.Store, chain config.ChainConfig) (*Service, error) {
	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(chain, []byte(store.Keystorage.GetSelectedPrivateKey()))
	if err != nil {
		return nil, err
	}
	return &Service{ctx: ctx, Chain: chain, Store: store, VORCoordinatorCaller: VORCoordinatorCaller}, nil
}
