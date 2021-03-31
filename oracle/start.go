package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sevlyar/go-daemon"
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
	"time"
)

func start() (err error) {
	var ctx = context.Background()
	var fee int64
	var paysGas bool

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
		FirstRun(keystore)
	}
	err = store.RandomnessRequest.Migrate()
	if err != nil {
		return err
	}
	err = auth(keystore)
	if err != nil {
		return err
	}

	daemonContext := &daemon.Context{
		PidFileName: "oracled.pid",
		PidFilePerm: 0644,
		LogFileName: config.Conf.LogFile,
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	daemon, err := daemonContext.Reborn()
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "main",
			"action":   "start oracle daemon",
			"result":   err,
		}).Error()
		return
	}
	if daemon != nil {
		return
	}
	defer daemonContext.Release()

	log.WithFields(logrus.Fields{
		"package":  "main",
		"function": "main",
		"action":   "start oracle daemon",
		"result":   "daemon started",
	}).Info()

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
		tx, err := oracleService.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(fee), paysGas)
		if tx != nil || err == nil {
			keystore.SetRegistered(keystore.KeyStore.PrivateKey)
		}
	}

	oracleController, err := controller.NewOracle(ctx, log, oracleService)
	oracleListener, err := chainlisten.NewVORCoordinatorListener(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost, oracleService, ctx)
	go oracleListener.StartPoll()

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == keystore.KeyStore.Token, nil
	}))
	e.POST("/withdraw", oracleController.Withdraw)
	e.POST("/register", oracleController.Register)
	e.POST("/changefee", oracleController.ChangeFee)
	e.POST("/withdraw", oracleController.Withdraw)
	e.POST("/stop", func(c echo.Context) error {
		if stop1 {
			log.WithFields(logrus.Fields{
				"package":  "main",
				"function": "start",
				"action":   "stop service",
				"result":   fmt.Sprintf("stopped %d %s", PID, time.Now().String()),
			}).Warning()
		}
		stop1 = true
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "stop service",
			"result":   fmt.Sprintf("stopped %d %s", PID, time.Now().String()),
		}).Warning()
		go func() {
			listener1.Close()
			if file1 != nil {
				file1.Close()
			}

			exit1 <- 1
		}()
		return err
	})
	e.GET("/about", oracleController.About)
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Conf.Serve.Host, config.Conf.Serve.Port)))

	return err
}
