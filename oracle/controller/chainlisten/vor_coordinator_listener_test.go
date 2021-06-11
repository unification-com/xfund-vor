package chainlisten_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/big"
	"oracle/controller/chainlisten"
	"oracle/service"
	"os"
	"path/filepath"
	"testing"
)

func InitCaller() (err error) {
	ctx := context.Background()
	Service, err = service.NewServiceFromPassedConfig(ctx, TestStore, Config)
	return err
}

func VORCoordinatorCallerTestValues() (string, string, *big.Int, []byte) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func Init(configAddres string, pass string) (err error) {
	err = InitConfig(configAddres)
	if err != nil {
		return err
	}
	err = InitKeystore(pass)
	if err != nil {
		return err
	}
	err = InitStore()
	if err != nil {
		return err
	}
	err = InitCaller()
	if err != nil {
		return err
	}
	VORCoordinator, err = chainlisten.NewVORCoordinatorListener(VORCoordinatorListenerTestValues())
	return err
}

func VORCoordinatorListenerTestValues() (string, string, *service.Service, *logrus.Logger, context.Context) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost, Service, logrus.New(), context.Background()
}

func TestVORCoordinatorListener_Request(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "..", "test_data", "generic_test_config.json")
	err := Init(configPath, "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	if err != nil {
		t.Error(err)
	}
	err = VORCoordinator.ProcessIncommingEvents()
	if err != nil {
		t.Error(err)
	}
}
