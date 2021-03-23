package chainlisten_test

import (
	"github.com/sirupsen/logrus"
	"oracle/config"
	"oracle/controller/chainlisten"
	"oracle/store/keystorage"
	"os"
	"testing"
)

var VORCoordinator *chainlisten.VORCoordinatorListener
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func Init(configAddres string) (err error) {
	Config, err = config.NewConfig(configAddres)
	if err != nil {
		return err
	}
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	if err != nil {
		return err
	}
	VORCoordinator, err = chainlisten.NewVORCoordinatorListener(VORCoordinatorCallerTestValues())
	return err
}

func VORCoordinatorCallerTestValues() (string, string) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost
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
