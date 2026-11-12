package main

import (
	"context"
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
	"oracle/controller/chainlisten"
	"oracle/service"
	store2 "oracle/store"
	"oracle/store/keystorage"
	"runtime/debug"
)

var e = echo.New()

func start() (err error) {
	var ctx = context.Background()
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

	// one chain per process until the supervisor runs a worker per chain
	if len(chains) != 1 {
		return fmt.Errorf("this version runs one chain per process, and the config lists %d", len(chains))
	}
	if err = seedCursor(store.ForChain(chains[0].NetworkID), keystore); err != nil {
		return err
	}
	oracleService, err := service.NewService(ctx, store, chains[0])
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "init service",
		}).Error(err.Error())
		return err
	}
	if err = ensureRegistered(oracleService.VORCoordinatorCaller, keystore, fee); err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "register proving key",
		}).Error(err.Error())
		return err
	}

	oracleController, err := controller.NewOracle(ctx, log, oracleService)
	if err != nil {
		return err
	}
	oracleListener, err = chainlisten.NewVORCoordinatorListener(oracleService, log, ctx)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "init listener",
		}).Error(err.Error())
		return err
	}
	go oracleListener.StartPoll()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return keystore.IsAPIKey(key), nil
	}))

	e.POST("/withdraw", oracleController.Withdraw)
	e.POST("/register", oracleController.Register)
	e.POST("/changefee", oracleController.ChangeFee)
	e.POST("/changegranularfee", oracleController.ChangeGranularFee)
	e.POST("/stop", func(c echo.Context) error {
		err = Stop()
		return err
	})
	e.POST("/queryfees", oracleController.QueryFees)
	e.GET("/about", oracleController.About)
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive")
	})
	e.GET("/querywithdrawable", oracleController.QueryWithdrawableTokens)
	e.GET("/requests", oracleController.QueryRequests)
	e.GET("/analytics", oracleController.Analytics)
	e.GET("/consumers", oracleController.Consumers)
	e.GET("/tx", oracleController.GetTxInfo)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Conf.Serve.Host, config.Conf.Serve.Port)))

	return err
}

// ensureRegistered makes sure the coordinator holds the oracle's proving key: the chain, not the keystore's
// flag, is the authority. An unregistered key is registered at fee, waiting for the transaction to succeed;
// fee is only known on the first run.
func ensureRegistered(caller *chaincall.VORCoordinatorCaller, keystore *keystorage.Keystorage, fee int64) error {
	provider, err := caller.ProvingKeyProvider()
	if err != nil {
		return fmt.Errorf("checking the proving key's registration: %w", err)
	}
	entry := log.WithFields(logrus.Fields{"package": "main", "function": "ensureRegistered"})
	if provider == (common.Address{}) {
		if fee <= 0 {
			// carry on: the API still works, and a restart loop would help nobody
			entry.Error("the proving key isn't registered with the coordinator, so the oracle can't fulfil requests")
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
