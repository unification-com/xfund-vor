package chaincall_test

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	if !Keystore.Exists() {
		err = Keystore.AddGenerated(Config.Keystorage.Account)
	}
	VORCoordinator, err = chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	return err
}

func VORCoordinatorCallerTestValues() (string, string, *big.Int, []byte) {
	return Config.VORCoordinatorContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetFirst().CipherPrivate)
}

func TestVORCoordinatorCaller_GetTotalGasDeposits(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	GasDeposits, err := VORCoordinator.GetTotalGasDeposits(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
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
	t.Log(string([]byte(string(HashOfKey[:]))))
}

func TestVORCoordinatorCaller_GetGasTopUpLimit(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	GasDeposits, err := VORCoordinator.GetGasTopUpLimit(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
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

	TransactOut, err := VORCoordinator.ChangeFee(big.NewInt(1))
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
	TransactOut, err := VORCoordinator.RegisterProvingKey(*big.NewInt(100000), false)
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

func TestVORCoordinatorCaller_SetProviderPaysGas(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	TransactOut, err := VORCoordinator.SetProviderPaysGas(false)
	//debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
	fmt.Print(TransactOut.To().Hex())
	transactJson, err := json.Marshal(TransactOut)
	fmt.Print(string(transactJson))
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
