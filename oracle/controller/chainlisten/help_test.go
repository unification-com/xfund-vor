package chainlisten_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/chaincall"
	"oracle/config"
	"oracle/controller/chainlisten"
	"oracle/service"
	store2 "oracle/store"
	"oracle/store/keystorage"
)

var VORCoordinator *chainlisten.VORCoordinatorListener
var VORCoordinatorCaller *chaincall.VORCoordinatorCaller
var Service *service.Service
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()
var TestStore *store2.Store

func InitConfig(configAddres string) (err error) {
	Config, err = config.NewConfig(configAddres)
	return err
}

func InitKeystore(pass string) (err error) {
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	if err != nil {
		return err
	}
	return Keystore.Unlock(Config.Keystorage.Account, pass)
}

func InitStore() (err error) {
	TestStore, err = store2.NewStore(context.Background(), Keystore)
	err = TestStore.Migrate([]int64{Config.ChainList()[0].NetworkID})
	return
}
