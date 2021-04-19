package chaincall_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"oracle/chaincall"
	"os"
	"path/filepath"
	"testing"
)

var VORRandomnessRequestMockCaller *chaincall.VORRandomnessRequestMockCaller

func InitVORRandomnessRequestMockCaller() (err error) {
	VORRandomnessRequestMockCaller, err = chaincall.NewVORRandomnessRequestMockCaller(VORRandomnessRequestMockCallerTestValues())
	return err
}

func VORRandomnessRequestMockCallerTestValues() (string, string, *big.Int, []byte) {
	return Config.MockContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetFirst().Private)
}

func TestVORRandomnessRequestMockCaller_RandomnessRequest(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "request_mock_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
	if err != nil {
		t.Error(err)
	}

	err = InitVORRandomnessRequestMockCaller()
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
