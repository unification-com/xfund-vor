package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/config"
	"oracle/service"
	store2 "oracle/store"
	"oracle/store/keystorage"
	"os"
	"testing"
)

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

func Init(configAddres string) (err error) {
	err = InitConfig(configAddres)
	if err != nil {
		return err
	}
	err = InitKeystore(configAddres)
	if err != nil {
		return err
	}
	err = InitStore()
	if err != nil {
		return err
	}
	return err
}

func TestMain_Auth(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	err = auth(Keystore)
}
