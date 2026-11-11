package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"math/big"
	"net/http"
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

	err = store.Db.Migrate()

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

	oracleService, err := service.NewService(ctx, store)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "init service",
		}).Error(err.Error())
		return err
	}
	if !keystore.IsRegistered() {
		tx, err := oracleService.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(fee))
		if tx != nil || err == nil {
			keystore.SetRegistered()
		}
	}

	oracleController, err := controller.NewOracle(ctx, log, oracleService)
	oracleListener, err = chainlisten.NewVORCoordinatorListener(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost, oracleService, log, ctx)
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
