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
	"os"
)

var e = echo.New()

func start() (err error) {
	var ctx = context.Background()
	var fee int64

	if config.Conf.LogFile != "" {
		logFile, err := os.OpenFile(config.Conf.LogFile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.WithFields(logrus.Fields{
				"package":  "main",
				"function": "start",
				"action":   "open log file",
				"result":   "can't open log file",
			}).Warning()
		} else {
			log.SetOutput(logFile)
		}
	} else {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "open log file",
			"result":   "log file is not specified",
		}).Warning()
	}

	keystore, err := keystorage.NewKeyStorage(log, config.Conf.Keystorage.File)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "open keystorage",
			"result":   "can't read keystorage, creating a new one...",
		}).Warning()
	}

	store, err := store2.NewStore(context.Background(), keystore)
	if !keystore.Exists() {
		fee, err = FirstRun(keystore)
	}
	err = store.Db.Migrate()
	if err != nil {
		return err
	}
	if options.Password == "" || (keystore.CheckToken(options.Password) != nil) {
		err = auth(keystore)
		if err != nil {
			return err
		}
	}

	err = keystore.SelectPrivateKey(config.Conf.Keystorage.Account)
	if err != nil {
		return err
	}

	oracleService, err := service.NewService(ctx, store)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "init service",
			"result":   err,
		}).Error()
		return err
	}
	if !keystore.IsRegisteredByPrivate(keystore.KeyStore.PrivateKey) {
		tx, err := oracleService.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(fee))
		if tx != nil || err == nil {
			keystore.SetRegistered(keystore.KeyStore.PrivateKey)
		}
	}

	oracleController, err := controller.NewOracle(ctx, log, oracleService)
	oracleListener, err = chainlisten.NewVORCoordinatorListener(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost, oracleService, ctx)
	go oracleListener.StartPoll()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == keystore.KeyStore.Token, nil
	}))
	e.POST("/withdraw", oracleController.Withdraw)
	e.POST("/register", oracleController.Register)
	e.POST("/changefee", oracleController.ChangeFee)
	e.POST("/stop", func(c echo.Context) error {
		err = Stop()
		return err
	})
	e.GET("/about", oracleController.About)
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Conf.Serve.Host, config.Conf.Serve.Port)))

	return err
}
