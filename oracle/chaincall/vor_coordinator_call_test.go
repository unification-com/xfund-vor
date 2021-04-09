package chaincall_test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/store/keystorage"
	"os"
	"runtime/debug"
	"testing"
)

var VORCoordinator *chaincall.VORCoordinatorCaller
var VORD20Caller *chaincall.VORD20Caller
var MockERC20Caller *chaincall.MockERC20Caller
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func Init(configAddress string) (err error) {
	Config, err = config.NewConfig(configAddress)
	if err != nil {
		return err
	}
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	if err != nil {
		return err
	}
	Keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	VORCoordinator, err = chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	VORD20Caller, err = chaincall.NewVORD20Caller(VORD20CallerTestValues())
	MockERC20Caller, err = chaincall.NewMockERC20Caller(MockERC20CallerTestValues())
	return err
}

func VORCoordinatorCallerTestValues() (string, string, *big.Int, []byte) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func MockERC20CallerTestValues() (string, string, *big.Int, []byte) {
	return Config.MockContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func VORD20CallerTestValues() (string, string, *big.Int, []byte) {
	return Config.ContractCallerAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername(Config.Keystorage.Account).Private)
}

func TestVORCoordinatorCaller_HashOfKey(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
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
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	TransactOut, err := VORCoordinator.Withdraw("0x04FBC34DCf60c88e701a8B3161154451e33Eef75", big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_ChangeFee(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	TransactOut, err := VORCoordinator.ChangeFee(big.NewInt(1000000000000000000))
	if err != nil {
		t.Error(err)
	}
	fmt.Print(TransactOut.To().Hex())

	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_RegisterProvingKey(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.RegisterProvingKey(big.NewInt(100))
	//debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
	transactJson, err := json.Marshal(TransactOut)
	fmt.Print(string(transactJson))
}

func TestVORCoordinatorCaller_RandomnessRequest(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	keyHash, err := VORCoordinator.HashOfKey()
	TransactOut, err := VORCoordinator.RandomnessRequest(keyHash, big.NewInt(10), big.NewInt(100))
	//debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
	transactJson, err := json.Marshal(TransactOut)
	fmt.Println(string(transactJson))
}

func TestVORCoordinatorCaller_FulfillRandomnessRequest(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	TransactOut, err := VORCoordinator.FulfillRandomnessRequest([]byte("hfdjkhgldfjk"))
	debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
}
