package chainlisten_test

import (
	"github.com/sirupsen/logrus"
	"oracle/chaincall"
	"oracle/config"
	"oracle/controller/chainlisten"
	"oracle/service"
	"oracle/store/keystorage"
)

var VORCoordinator *chainlisten.VORCoordinatorListener
var VORCoordinatorCaller *chaincall.VORCoordinatorCaller
var Service *service.Service
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func InitConfig(configAddres string) (err error) {
	Config, err = config.NewConfig(configAddres)
	return err
}

func InitKeystore(configAddres string) (err error) {
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	return err
}
