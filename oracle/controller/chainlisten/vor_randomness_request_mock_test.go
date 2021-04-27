package chainlisten_test

import (
	"context"
	"oracle/controller/chainlisten"
	"oracle/service"
	"os"
	"path/filepath"
	"testing"
)

var VORRandomnessRequestMockListener *chainlisten.VORRandomnessRequestMockListener

func InitVORRandomnessRequestMockListener(configAddres string) (err error) {
	InitConfig(configAddres)
	InitKeystore("dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	InitStore()
	InitCaller()
	VORRandomnessRequestMockListener, err = chainlisten.NewVORRandomnessRequestMockListener(VORRandomnessRequestMockListenerCallerTestValues())
	return err
}

func VORRandomnessRequestMockListenerCallerTestValues() (string, string, *service.Service, context.Context) {
	return Config.MockContractAddress, Config.EthHTTPHost, Service, context.Background()
}

func TestVORRandomnessRequestMockListener_Request(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "..", "test_data", "generic_test_config.json")
	err := InitVORRandomnessRequestMockListener(configPath)
	if err != nil {
		if err.Error() != "record not found" {
			t.Error(err)
		}
	}
	err = VORRandomnessRequestMockListener.Request()

	if err != nil {
		t.Error(err)
	}
}
