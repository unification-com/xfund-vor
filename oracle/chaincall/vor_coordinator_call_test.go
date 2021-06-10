package chaincall_test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/store/keystorage"
	"os"
	"path/filepath"
	"testing"
)

var VORCoordinator *chaincall.VORCoordinatorCaller
var VORD20Caller *chaincall.VORD20Caller
var MockERC20Caller *chaincall.MockERC20Caller
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func Init(configAddress string, pass string) (err error) {
	Config, err = config.NewConfig(configAddress)
	if err != nil {
		return err
	}
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	if err != nil {
		return err
	}
	Keystore.CheckToken(pass)
	VORCoordinator, err = chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	VORD20Caller, err = chaincall.NewVORD20Caller(VORD20CallerTestValues())
	MockERC20Caller, err = chaincall.NewMockERC20Caller(MockERC20CallerTestValues())
	return err
}

func VORCoordinatorCallerTestValues() (string, string, string, *big.Int, []byte) {
	return Config.VORCoordinatorContractAddress, Config.BlockHashStoreContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func MockERC20CallerTestValues() (string, string, *big.Int, []byte) {
	return Config.MockContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func VORD20CallerTestValues() (string, string, *big.Int, []byte) {
	return Config.ContractCallerAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func TestVORCoordinatorCaller_HashOfKey(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "generic_test_config.json")
	err := Init(configPath, "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	if err != nil {
		t.Error(err)
	}

	HashOfKey, err := VORCoordinator.HashOfKey()
	if err != nil {
		t.Error(err)
	}
	t.Log(common.BytesToHash(HashOfKey[:]))
	t.Log(HashOfKey[:])
	t.Log(hexutil.Encode(HashOfKey[:]))
}

func TestVORCoordinatorCaller_Withdraw(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "generic_test_config.json")
	err := Init(configPath, "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	if err != nil {
		t.Error(err)
	}

	TransactOut, err := VORCoordinator.Withdraw("0x04FBC34DCf60c88e701a8B3161154451e33Eef75", big.NewInt(1))
	assert.Equal(t, "VM Exception while processing transaction: revert can't withdraw more than balance", err.Error())
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_ChangeFee(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "generic_test_config.json")
	err := Init(configPath, "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	if err != nil {
		t.Error(err)
	}

	_, _ = VORCoordinator.RegisterProvingKey(big.NewInt(1))

	TransactOut, err := VORCoordinator.ChangeFee(big.NewInt(2))
	if err != nil {
		t.Error(err)
	}
	fmt.Print(TransactOut.To().Hex())

	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_RegisterProvingKey(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "new_proving_key_test_config.json")
	err := Init(configPath, "62jcip8sx41d5hgyz2sqkk452tnskukh")
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.RegisterProvingKey(big.NewInt(1))
	//debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
	transactJson, err := json.Marshal(TransactOut)
	fmt.Print(string(transactJson))
}

func TestVORCoordinatorCaller_RandomnessRequest(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "generic_test_config.json")
	err := Init(configPath, "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	if err != nil {
		t.Error(err)
	}
	keyHash, err := VORCoordinator.HashOfKey()
	fee, _ := VORCoordinator.QueryFees("")
	TransactOut, err := VORCoordinator.RandomnessRequest(keyHash, big.NewInt(10), fee)
	//debug.PrintStack()
	t.Log(TransactOut)
	assert.Equal(t, "VM Exception while processing transaction: revert request can only be made by a contract", err.Error())
	transactJson, err := json.Marshal(TransactOut)
	fmt.Println(string(transactJson))
}

func TestVORCoordinatorCaller_FulfillRandomnessRequest(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "generic_test_config.json")
	err := Init(configPath, "dwkxnzn3kl1dlndvtdtvqko9gpaay5vj")
	if err != nil {
		t.Error(err)
	}

	TransactOut, err := VORCoordinator.FulfillRandomnessRequest([]byte("hfdjkhgldfjk"))
	//debug.PrintStack()
	t.Log(TransactOut)
	assert.Equal(t, "VM Exception while processing transaction: revert wrong proof length", err.Error())
}
