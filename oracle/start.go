package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"oracle/chaincall"
	"oracle/config"
	controller "oracle/controller/api"
	"oracle/controller/chainlisten"
	"oracle/service"
	"oracle/store/keystorage"
	"os"
	"time"
)

func start() error {
	var err error
	var ctx = context.Background()

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

	var oraclePrivateKey string
	if !keystore.Exists() {
		err = keystore.AddGenerated(config.Conf.Keystorage.Account)
		if err != nil {
			return err
		}
		oraclePrivateKey = keystore.GetFirst().CipherPrivate
	} else {
		oraclePrivateKeyModel, err := keystore.GetByAccount(config.Conf.Keystorage.Account)
		if err != nil {
			return err
		}
		oraclePrivateKey = oraclePrivateKeyModel.CipherPrivate
	}

	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost, big.NewInt(config.Conf.NetworkID), []byte(oraclePrivateKey))
	if err != nil || VORCoordinatorCaller == nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "connect to VORCoordinator",
			"result":   "can't connect to VORCoordinator",
		}).Error()
		return fmt.Errorf("can't connect to VORCoordinator")
	}
	oracleService := service.NewService(ctx, VORCoordinatorCaller)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "init service",
			"result":   "can't create oracle service",
		}).Error()
		return fmt.Errorf("can't create oracle service")
	}

	oracleController, err := controller.NewOracle(ctx, log, oracleService)
	oracleListener, err := chainlisten.NewVORCoordinatorListener(config.Conf.VORCoordinatorContractAddress, config.Conf.EthHTTPHost)
	go oracleListener.StartPoll()

	e := echo.New()
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
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "alive")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Conf.Serve.Host, config.Conf.Serve.Port)))

	return err
}
