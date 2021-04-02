package chainlisten_test

import (
	"context"
	"math/big"
	"oracle/controller/chainlisten"
	"oracle/service"
	"os"
	"testing"
)

func InitCaller(configAddress string) (err error) {
	Service, err = service.NewService(context.Background(), Store)
	return err
}

func VORCoordinatorCallerTestValues() (string, string, *big.Int, []byte) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
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
	err = InitCaller(configAddres)
	if err != nil {
		return err
	}
	VORCoordinator, err = chainlisten.NewVORCoordinatorListener(VORCoordinatorListenerTestValues())
	return err
}

func VORCoordinatorListenerTestValues() (string, string, *service.Service, context.Context) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost, Service, context.Background()
}

func TestVORCoordinatorListener_Request(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	err = VORCoordinator.Request()
	if err != nil {
		t.Error(err)
	}
}
