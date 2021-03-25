package chaincall_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"oracle/chaincall"
	"os"
	"testing"
)

var VORRandomnessRequestMockCaller *chaincall.VORRandomnessRequestMockCaller

func InitVORRandomnessRequestMockCaller(configAddress string) (err error) {
	VORRandomnessRequestMockCaller, err = chaincall.NewVORRandomnessRequestMockCaller(VORRandomnessRequestMockCallerTestValues())
	return err
}

func VORRandomnessRequestMockCallerTestValues() (string, string, *big.Int, []byte) {
	return Config.MockContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetFirst().CipherPrivate)
}

func TestVORRandomnessRequestMockCaller_RandomnessRequest(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	err = InitVORRandomnessRequestMockCaller(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	keyHash, err := VORCoordinator.HashOfKey()
	TransactOut, err := VORRandomnessRequestMockCaller.RandomnessRequest(keyHash, big.NewInt(10), big.NewInt(100))
	//debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
	transactJson, err := json.Marshal(TransactOut)
	fmt.Println(string(transactJson))
}

//    vor_randomness_request_mock_test.go:35: &{0xc00046f9e0 {13839281513874783560 212112001 0x1b1dc40} {<nil>} {<nil>} {<nil>}}
//    {"type":"0x0","nonce":"0x5","gasPrice":"0x4a817c800","gas":"0x186a0","value":"0x0","input":"0x264eb1cf77c7d6c3e6ea3b1fa790608a3babe330dbaef9766c5a5496ea9e0d3438f2d843000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000064","v":"0x2d46","r":"0xda4cd5a4e5683bf2b24535dde5db3181ef98e728b89b063f1d1d47341a3b143b","s":"0x3675da56e720adda0c65b35a22bb2f5dfd908ec5fea524ca0c398aa9ee2b4e9","to":"0xf196ea87fc3905d555c3056e91d574ad2f474d35","hash":"0x1057ad13d9e15fb9b2bfc396a6032269da84330bcd13fbc3689bc7973c712bce"}
//
//    vor_randomness_request_mock_test.go:35: &{0xc0000866c0 {13839282143702750224 240994101 0x18bdc40} {<nil>} {<nil>} {<nil>}}
//	  {"type":"0x0","nonce":"0x6","gasPrice":"0x4a817c800","gas":"0x186a0","value":"0x0","input":"0x264eb1cf77c7d6c3e6ea3b1fa790608a3babe330dbaef9766c5a5496ea9e0d3438f2d843000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000064","v":"0x2d46","r":"0x73f370f906cab0efc3c7bbd84a10d5aaaeb1510767200ef1f2438d951546dfe2","s":"0x1cba3e21e12b5f4c07508224faa2f052be4e1d393b39a895eb5b09c6ade61ee2","to":"0xf196ea87fc3905d555c3056e91d574ad2f474d35","hash":"0x59cdaccd5b27eb6271e1016058ae29ddec68e4d97126cefbb0dd0560353c8b54"}
