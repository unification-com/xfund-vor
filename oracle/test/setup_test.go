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
	"time"
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
	err = Keystore.AddExisting("rootuser", "9b2e913979a8e390c4ba2ee74f3db065ad2b9b21578646785471cb3de1819ce0")
	if err != nil {
		t.Error(err)
	}
	err = Keystore.AddExisting("oracle", "b1f07ac06b3e581bb499009b51850b04800720e7a302fe44b531a2eee11ae274")
	if err != nil {
		t.Error(err)
	}
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
	//tx, err := oracleService.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(100), false)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(tx)

	//	set fee
	tx, err := oracleService.VORCoordinatorCaller.ChangeFee(big.NewInt(1000000000000000000))
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

	//rootVORD20Caller.RenewTransactOpts()
	//tx, err = rootVORD20Caller.RollDice(big.NewInt(7))
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(tx)

	//	set VORD20 keyHash
	rootVORD20Caller.RenewTransactOpts()
	tx, err = rootVORD20Caller.SetKeyHash(oracleKeyHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set VORD20 fee
	rootVORD20Caller.RenewTransactOpts()
	tx, err = rootVORD20Caller.SetFee(big.NewInt(1000000000000000000))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	create requestRandomness for oracle
	err = rootVORD20Caller.RenewTransactOpts()
	if err != nil {
		t.Error(err)
	}

	tx, err = rootVORD20Caller.RollDice(big.NewInt(10))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	run check
	time.Sleep(time.Second * 5)
	err = oracleVORCoordinatorListener.Request()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 5)
	house, err := rootVORD20Caller.House()
	if err != nil {
		t.Error(err)
	}
	t.Log(house)
}
