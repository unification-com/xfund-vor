package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"oracle/config"
	controller "oracle/controller/api"
	"oracle/controller/chainlisten"
	"oracle/service"
	store2 "oracle/store"
	"oracle/store/keystorage"
)

// chainRetryDelay is how long a chain whose worker couldn't start waits before trying again.
const chainRetryDelay = time.Minute

// supervisor runs a worker per chain: its own service (so its own sender and nonce) and listener. A chain
// that can't start (its RPC is down, say) is retried without holding up the others.
type supervisor struct {
	ctx      context.Context
	store    *store2.Store
	keystore *keystorage.Keystorage
	fee      int64 // the fee entered on the first run: registers the proving key where a chain sets none

	mu       sync.Mutex
	services map[int64]*service.Service // running chains, by network id
}

func newSupervisor(ctx context.Context, store *store2.Store, keystore *keystorage.Keystorage, fee int64) *supervisor {
	return &supervisor{ctx: ctx, store: store, keystore: keystore, fee: fee, services: map[int64]*service.Service{}}
}

// run starts a worker for every chain.
func (s *supervisor) run(chains []config.ChainConfig) {
	for _, chain := range chains {
		go s.runChain(chain)
	}
}

// runChain starts the chain's worker, retrying until it starts or the oracle stops, then polls until the
// oracle stops.
func (s *supervisor) runChain(chain config.ChainConfig) {
	entry := log.WithFields(logrus.Fields{"package": "main", "function": "runChain", "chain": chain.Name})
	for {
		svc, listener, err := s.startChain(chain)
		if err == nil {
			svc.Scanner = listener
			s.mu.Lock()
			s.services[chain.NetworkID] = svc
			s.mu.Unlock()
			entry.Info("chain running")
			_ = listener.StartPoll()
			return
		}
		entry.WithField("retry_in", chainRetryDelay.String()).Error(err.Error())
		select {
		case <-s.ctx.Done():
			return
		case <-time.After(chainRetryDelay):
		}
	}
}

func (s *supervisor) startChain(chain config.ChainConfig) (*service.Service, *chainlisten.VORCoordinatorListener, error) {
	svc, err := service.NewService(s.ctx, s.store, chain)
	if err != nil {
		return nil, nil, err
	}
	fee := chain.Fee
	if fee == 0 {
		fee = s.fee
	}
	if err = ensureRegistered(svc.VORCoordinatorCaller, s.keystore, fee); err != nil {
		return nil, nil, err
	}
	listener, err := chainlisten.NewVORCoordinatorListener(svc, log, s.ctx)
	if err != nil {
		return nil, nil, err
	}
	return svc, listener, nil
}

// service finds a running chain by name or network id ("" when the oracle serves one chain), for the API.
func (s *supervisor) service(selector string) (*service.Service, error) {
	chain, err := config.Conf.ResolveChain(selector)
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if svc := s.services[chain.NetworkID]; svc != nil {
		return svc, nil
	}
	return nil, fmt.Errorf("%s: %w", chain.Name, controller.ErrChainNotRunning)
}
