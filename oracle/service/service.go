package service

import (
	"context"
	"errors"

	"oracle/chaincall"
	"oracle/config"
	"oracle/store"
)

// ErrInvalidScanTarget: a scan position that isn't a block, "head" or "head-N" on the chain.
var ErrInvalidScanTarget = errors.New("give a block number, head or head-N, no later than the block after the chain's head")

// Scanner is a running chain's event scan.
type Scanner interface {
	// ScanFrom moves where the scan reads next, returning the old and new positions.
	ScanFrom(target string) (from, to uint64, err error)
}

// Service is the oracle's work on one chain.
type Service struct {
	ctx                  context.Context
	Chain                config.ChainConfig
	Store                *store.Store
	VORCoordinatorCaller *chaincall.VORCoordinatorCaller
	Scanner              Scanner // set once the chain's worker runs
}

// NewService serves chain, with the store scoped to it.
func NewService(ctx context.Context, store *store.Store, chain config.ChainConfig) (*Service, error) {
	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(chain, []byte(store.Keystorage.GetSelectedPrivateKey()))
	if err != nil {
		return nil, err
	}
	return &Service{ctx: ctx, Chain: chain, Store: store.ForChain(chain.NetworkID), VORCoordinatorCaller: VORCoordinatorCaller}, nil
}
