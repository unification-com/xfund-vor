package test

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/controller/chainlisten"
	"oracle/service"
	store2 "oracle/store"
	"oracle/store/keystorage"
	"os"
	"runtime/debug"
	"testing"
)

var VORCoordinator *chainlisten.VORCoordinatorListener
var VORCoordinatorCaller *chaincall.VORCoordinatorCaller
var Service *service.Service
var Keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()
var Store *store2.Store

func InitConfig(configAddres string) (err error) {
	Config, err = config.NewConfig(configAddres)
	return err
}

func InitKeystore() (err error) {
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	Keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	return err
}

func InitStore() (err error) {
	Store, err = store2.NewStore(context.Background(), Keystore)
	err = Store.RandomnessRequest.Migrate()
	return
}

func Init(configAddres string) (err error) {
	err = InitConfig(configAddres)
	err = InitKeystore()
	err = InitStore()
	return
}

func TestStartEverything(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	config.Conf = Config

	err = Keystore.CheckToken("rod0gbc3mhyxdiah2vwialx1q3osk5cw")
	if err != nil {
		t.Error(err)
	}
	//err = Keystore.AddExisting("rootuser", "25075fd6bca594a1d8f6f2643879c779d61b34cafad0d242478d068fb29eac3f")
	//if err != nil {
	//	t.Error(err)
	//}
	//err = Keystore.AddExisting("oracle", "7f5a703bcfa0405275f097e9b7c9d26450680c2f5dd5ae4daa4102331c057f18")
	//if err != nil {
	//	t.Error(err)
	//}
	Keystore.SelectPrivateKey(Config.Keystorage.Account)

	debug.PrintStack()
	Keystore.SelectPrivateKey("oracle")
	oracleService, err := service.NewService(context.Background(), Store)
	if err != nil || oracleService == nil {
		t.Error(err)
	}

	oracleVORCoordinatorListener, err := chainlisten.NewVORCoordinatorListener(Config.VORCoordinatorContractAddress, Config.EthHTTPHost, oracleService, context.Background())
	if err != nil || oracleVORCoordinatorListener == nil {
		t.Error(err)
	}

	//	register proving key
	tx, err := oracleService.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(100), false)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set fee
	tx, err = oracleService.VORCoordinatorCaller.ChangeFee(big.NewInt(1000000000000000000))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	get keyhash
	oracleKeyHash, err := oracleService.VORCoordinatorCaller.HashOfKey()

	//	create requestRandomness
	rootVORD20Caller, err := chaincall.NewVORD20Caller(Config.ContractCallerAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername("rootuser").Private))
	if err != nil || rootVORD20Caller == nil {
		t.Error(err)
	}

	rootMockERC20, err := chaincall.NewMockERC20Caller(Config.MockContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetByUsername("rootuser").Private))
	if err != nil {
		t.Error(err)
	}

	//	send money to VORD20
	amount := big.NewInt(1000000000000000000)
	amount.Mul(amount, big.NewInt(10))
	tx, err = rootMockERC20.Transfer(Config.ContractCallerAddress, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	tx, err = rootVORD20Caller.RollDice(big.NewInt(3))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set VORD20 keyHash
	tx, err = rootVORD20Caller.SetKeyHash(oracleKeyHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set VORD20 fee
	tx, err = rootVORD20Caller.SetFee(big.NewInt(1000000000000000000))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	create requestRandomness for oracle
	tx, err = rootVORD20Caller.RollDice(big.NewInt(4))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	run check
	err = oracleVORCoordinatorListener.Request()
	if err != nil {
		t.Error(err)
	}

}
