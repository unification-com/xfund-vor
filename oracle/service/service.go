package service

import (
	"context"
	"log"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/store"
)

type Service struct {
	ctx                  context.Context
	Store                *store.Store
	VORCoordinatorCaller *chaincall.VORCoordinatorCaller
	log                  *log.Logger
}

func NewService(ctx context.Context, store *store.Store) (*Service, error) {
	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(
		config.Conf.VORCoordinatorContractAddress,
		config.Conf.BlockHashStoreContractAddress,
		config.Conf.EthHTTPHost,
		big.NewInt(config.Conf.NetworkID),
		[]byte(store.Keystorage.GetSelectedPrivateKey()),
	)
	if err != nil {
		return nil, err
	}
	return &Service{ctx: ctx, Store: store, VORCoordinatorCaller: VORCoordinatorCaller}, err
}

func NewServiceFromPassedConfig(ctx context.Context, store *store.Store, conf *config.Config) (*Service, error) {
	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(
		conf.VORCoordinatorContractAddress,
		conf.BlockHashStoreContractAddress,
		conf.EthHTTPHost,
		big.NewInt(conf.NetworkID),
		[]byte(store.Keystorage.GetSelectedPrivateKey()),
	)
	if err != nil {
		return nil, err
	}
	return &Service{ctx: ctx, Store: store, VORCoordinatorCaller: VORCoordinatorCaller}, err
}
