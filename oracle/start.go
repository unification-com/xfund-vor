package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/big"
	"oracle/chaincall"
	"oracle/service"
	"oracle/store/keystorage"
	"os"
)

func start() error {
	var err error
	var ctx = context.Background()

	if configuration.LogFile != "" {
		logFile, err := os.OpenFile(configuration.LogFile, os.O_WRONLY|os.O_CREATE, 0755)
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

	keystore, err := keystorage.NewKeyStorage(log, configuration.Keystorage.File)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "open keystorage",
			"result":   "can't read keystorage, creating a new one...",
		}).Warning()
	}
	VORCoordinatorCaller, err := chaincall.NewVORCoordinatorCaller(configuration.VORCoordinatorContractAddress, configuration.EthHTTPHost, big.NewInt(configuration.NetworkID), []byte(keystore.GetFirst().CipherPrivate))
	if err != nil || VORCoordinatorCaller == nil {
		log.WithFields(logrus.Fields{
			"package":  "main",
			"function": "start",
			"action":   "connect to VORCoordinator",
			"result":   "can't connect to VORCoordinator",
		}).Error()
		return fmt.Errorf("can't connect to VORCoordinator")
	}
	service.NewService(ctx, VORCoordinatorCaller)
	return err
}
