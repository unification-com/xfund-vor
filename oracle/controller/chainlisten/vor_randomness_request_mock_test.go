package chainlisten_test

import (
	"context"
	"oracle/controller/chainlisten"
	"oracle/service"
	"os"
	"runtime/debug"
	"testing"
)

var VORRandomnessRequestMockListener *chainlisten.VORRandomnessRequestMockListener

func InitVORRandomnessRequestMockListener(configAddres string) (err error) {
	InitConfig(configAddres)
	InitKeystore(configAddres)
	InitCaller(configAddres)
	VORRandomnessRequestMockListener, err = chainlisten.NewVORRandomnessRequestMockListener(VORRandomnessRequestMockListenerCallerTestValues())
	return err
}

func VORRandomnessRequestMockListenerCallerTestValues() (string, string, *service.Service, context.Context) {
	return Config.MockContractAddress, Config.EthHTTPHost, Service, context.Background()
}

func TestVORRandomnessRequestMockListener_Request(t *testing.T) {
	err := InitVORRandomnessRequestMockListener(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	err = VORRandomnessRequestMockListener.Request()
	debug.PrintStack()
	if err != nil {
		t.Error(err)
	}
}
