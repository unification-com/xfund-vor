package chaincall_test

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/store/keystorage"
	"oracle/walletworker"
	"os"
	"runtime/debug"
	"testing"
)

const (
	VORCoordinatorAddress = "0x22F043993312CB050E7F7A5C1207f68a05D3ef66"
)

var oraclePrivateKey = []byte("0xf54ca099a480e75a417a676855aed602f559d27f6f461f3754667b0b8af11ba6")
var VORCoordinator *chaincall.VORCoordinatorCaller
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func Init(configAddres string) {
	Config, _ = config.NewConfig(configAddres)
	Keystore, _ = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	VORCoordinator, _ = chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
}

func VORCoordinatorCallerTestValues() (string, string, *big.Int, []byte) {
	return VORCoordinatorAddress, fmt.Sprintf("http://%s:%d", Config.Serve.Host, Config.Serve.Port), big.NewInt(Config.NetworkID), []byte(Keystore.GetFirst().CipherPrivate)
}

func TestVORCoordinatorCaller_GetTotalGasDeposits(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	GasDeposits, err := VORCoordinator.GetTotalGasDeposits(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
}

func TestVORCoordinatorCaller_GetGasTopUpLimit(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	GasDeposits, err := VORCoordinator.GetGasTopUpLimit(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
}

func TestVORCoordinatorCaller_Withdraw(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	TransactOut, err := VORCoordinator.Withdraw("0x04FBC34DCf60c88e701a8B3161154451e33Eef75", big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_ChangeFee(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	TransactOut, err := VORCoordinator.ChangeFee(big.NewInt(1))
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_RegisterProvingKey(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	oraclePrivateKeyECDSA, err := crypto.HexToECDSA(string(oraclePrivateKey[2:]))
	if err != nil {
		return
	}
	oraclePublic := oraclePrivateKeyECDSA.Public().(*ecdsa.PublicKey)
	oracleAddress := walletworker.GenerateAddress(oraclePublic)
	TransactOut, err := VORCoordinator.RegisterProvingKey(*big.NewInt(10000), string(oracleAddress), false)
	debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
}

func TestVORCoordinatorCaller_SetProviderPaysGas(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	TransactOut, err := VORCoordinator.SetProviderPaysGas(false)
	debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
}
func TestVORCoordinatorCaller_FulfillRandomnessRequest(t *testing.T) {
	Init(os.Args[len(os.Args)-1])

	TransactOut, err := VORCoordinator.FulfillRandomnessRequest([]byte("hfdjkhgldfjk"))
	debug.PrintStack()
	t.Log(TransactOut)
	if err != nil {
		t.Error(err)
	}
}
