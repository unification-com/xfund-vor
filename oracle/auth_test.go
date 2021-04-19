package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/config"
	"oracle/service"
	store2 "oracle/store"
	"oracle/store/keystorage"
	"os"
	"path/filepath"
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

func InitKeystore() (err error) {
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	err = Keystore.CheckToken("dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	err = Keystore.SelectPrivateKey(Config.Keystorage.Account)
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
	err = InitKeystore()
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
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "test_data", "generic_test_config.json")
	err := Init(configPath)
	if err != nil {
		t.Error(err)
	}
}
