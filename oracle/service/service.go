package service

import (
	"context"
	"fmt"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/store"
)

type Service struct {
	ctx                  context.Context
	Store                *store.Store
	VORCoordinatorCaller *chaincall.VORCoordinatorCaller
}

func NewService(ctx context.Context, store *store.Store) (*Service, error) {
	fmt.Print(store.Keystorage.GetSelectedPrivateKey())
	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(config.Conf.MockContractAddress, config.Conf.EthHTTPHost, big.NewInt(config.Conf.NetworkID), []byte(store.Keystorage.GetSelectedPrivateKey()))
	if err != nil {
		return nil, err
	}
	return &Service{ctx: ctx, Store: store, VORCoordinatorCaller: VORCoordinatorCaller}, err
}
