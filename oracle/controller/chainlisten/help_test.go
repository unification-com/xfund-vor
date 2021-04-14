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
var Store *store2.Store

func InitConfig(configAddres string) (err error) {
	Config, err = config.NewConfig(configAddres)
	return err
}

func InitKeystore(configAddres string) (err error) {
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	Keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	Keystore.SelectPrivateKey(Config.Keystorage.Account)
	return err
}

func InitStore() (err error) {
	Store, err := store2.NewStore(context.Background(), Keystore)
	err = Store.Db.Migrate()
	return
}
