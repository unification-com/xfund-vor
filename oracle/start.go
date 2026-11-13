package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"oracle/chaincall"
	"oracle/config"
	controller "oracle/controller/api"
	store2 "oracle/store"
	"oracle/store/keystorage"
	"runtime/debug"
)

var e = echo.New()

func start() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	stopWorkers = cancel
	var fee int64
	var passphrase string

	keystore, err := keystorage.NewKeyStorage(log, config.Conf.Keystorage.File)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "open keystorage",
		}).Error(err.Error())
		return err
	}

	store, err := store2.NewStore(context.Background(), keystore)

	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "new store",
		}).Error(err.Error())
		return err
	}

	if !keystore.Exists() {
		fee, passphrase, err = FirstRun(keystore)
		if err != nil {
			log.WithFields(logrus.Fields{
				"package":  "main",
				"function": "start",
				"action":   "first run",
			}).Error(err.Error())
			return err
		}
	}

	chains := config.Conf.ChainList()
	chainIDs := make([]int64, len(chains))
	for i, chain := range chains {
		chainIDs[i] = chain.NetworkID
	}
	err = store.Migrate(chainIDs)

	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "migrate db",
		}).Error(err.Error())
		return err
	}

	if passphrase == "" && options.PasswordFile != "" {
		passphrase = getPasswordFromFileOrFlag(options.PasswordFile)
	}
	if err = unlock(keystore, config.Conf.Keystorage.Account, passphrase); err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "unlock keystore",
		}).Error(err.Error())
		return err
	}
	// scrypt needed ~256 MB to unlock the key: hand it back now, as the oracle shares its host
	debug.FreeOSMemory()
	if !keystore.HasAPIKey() {
		log.WithFields(logrus.Fields{"package": "main", "function": "start"}).
			Warn("no API key is set, so the HTTP API refuses every request: issue one with `oracle api-key`")
	}

	// v0.1 kept a single chain's scan position in the keystore
	if len(chains) == 1 {
		if err = seedCursor(store.ForChain(chains[0].NetworkID), keystore); err != nil {
			return err
		}
	}
	workers := newSupervisor(ctx, store, keystore, fee)
	workers.run(chains)
	oracleController, err := controller.NewOracle(ctx, log, workers.service)
	if err != nil {
		return err
	}

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return keystore.IsAPIKey(key), nil
	}))

	chain := oracleController.Chain
	e.POST("/withdraw", chain(oracleController.Withdraw))
	e.POST("/register", chain(oracleController.Register))
	e.POST("/changefee", chain(oracleController.ChangeFee))
	e.POST("/changegranularfee", chain(oracleController.ChangeGranularFee))
	e.POST("/stop", func(c echo.Context) error {
		// reply first: the graceful shutdown waits for this response to go out
		go func() { _ = Stop() }()
		return c.String(http.StatusOK, "stopping")
	})
	e.POST("/queryfees", chain(oracleController.QueryFees))
	e.GET("/about", chain(oracleController.About))
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive")
	})
	e.GET("/querywithdrawable", chain(oracleController.QueryWithdrawableTokens))
	e.GET("/requests", chain(oracleController.QueryRequests))
	e.GET("/analytics", chain(oracleController.Analytics))
	e.GET("/consumers", chain(oracleController.Consumers))
	e.GET("/tx", chain(oracleController.GetTxInfo))
	e.POST("/scanfrom", chain(oracleController.ScanFrom))

	// Stop shuts the server down, which ends Start with ErrServerClosed: a clean exit
	if err = e.Start(fmt.Sprintf("%s:%d", config.Conf.Serve.Host, config.Conf.Serve.Port)); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// ensureRegistered makes sure the coordinator holds the oracle's proving key: the chain, not the keystore's
// flag, is the authority. An unregistered key is registered at fee (the chain's `fee`, or the first run's),
// waiting for the transaction to succeed.
func ensureRegistered(caller *chaincall.VORCoordinatorCaller, keystore *keystorage.Keystorage, fee int64) error {
	provider, err := caller.ProvingKeyProvider()
	if err != nil {
		return fmt.Errorf("checking the proving key's registration: %w", err)
	}
	entry := log.WithFields(logrus.Fields{"package": "main", "function": "ensureRegistered"})
	if provider == (common.Address{}) {
		if fee <= 0 {
			// carry on: the API still works, and a restart loop would help nobody
			entry.Error("the proving key isn't registered with the coordinator, so the oracle can't fulfil requests: set the chain's fee to register it")
			return nil
		}
		if err = caller.RegisterProvingKeyAndWait(big.NewInt(fee)); err != nil {
			return err
		}
		entry.Info("registered the proving key")
	} else if provider != caller.OracleAddress() {
		entry.WithField("provider", provider.Hex()).Warn("the proving key is registered to another provider address, which its fees are paid to")
	}
	return keystore.SetRegistered()
}

// seedCursor carries a single-chain oracle's scan position from the keystore, where v0.1 kept it, into the
// chain's cursor in the database, once.
func seedCursor(chainStore *store2.Store, keystore *keystorage.Keystorage) error {
	next, err := chainStore.Db.NextBlock()
	if err != nil || next > 0 {
		return err
	}
	if block, err := keystore.GetBlockNumber(); err == nil && block > 0 {
		return chainStore.Db.SetNextBlock(uint64(block))
	}
	return nil
}
