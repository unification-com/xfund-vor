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
	"path/filepath"
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

const (
	// setupToken unlocks test_data/setup_test_keystore.json
	setupToken = "fq516b1boc8vrm7nasnb8fy7u5rb6zhh"
	// ownerKey is ganache's deterministic account 0
	ownerKey = "4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
)

func InitKeystore() (err error) {
	Keystore, err = keystorage.NewKeyStorage(Log, Config.Keystorage.File)
	if err != nil {
		return err
	}
	return Keystore.Unlock(Config.Keystorage.Account, setupToken)
}

func InitStore() (err error) {
	Store, err = store2.NewStore(context.Background(), Keystore)
	err = Store.Migrate([]int64{Config.ChainList()[0].NetworkID})
	return
}

func Init(configAddres string) (err error) {
	err = InitConfig(configAddres)
	err = InitKeystore()
	err = InitStore()
	return
}

func TestStartEverything(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "setup_test_config.json")
	err := Init(configPath)
	config.Conf = Config

	err = Keystore.AddExisting("owner", ownerKey)
	if err != nil {
		t.Error(err)
	}
	err = Keystore.AddExisting("oracle", "ae9a2e131e9b359b198fa280de53ddbe2247730b881faae7af08e567e58915bd")
	if err != nil {
		t.Error(err)
	}
	err = Keystore.AddExisting("roller", "2e114163041d2fb8d45f9251db259a68ee6bdbfd6d10fe1ae87c5c4bcd6ba491")
	if err != nil {
		t.Error(err)
	}
	if err = Keystore.Unlock(Config.Keystorage.Account, setupToken); err != nil {
		t.Error(err)
	}

	oracleService, err := service.NewService(context.Background(), Store, Config.ChainList()[0])
	if err != nil || oracleService == nil {
		t.Error(err)
	}

	oracleVORCoordinatorListener, err := chainlisten.NewVORCoordinatorListener(oracleService, Log, context.Background())
	if err != nil || oracleVORCoordinatorListener == nil {
		t.Error(err)
	}

	//	register proving key
	tx, err := oracleService.VORCoordinatorCaller.RegisterProvingKey(big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set fee
	tx, err = oracleService.VORCoordinatorCaller.ChangeFee(big.NewInt(200))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	get keyhash
	oracleKeyHash, err := oracleService.VORCoordinatorCaller.HashOfKey()

	//	create requestRandomness
	rootVORD20Owner, err := chaincall.NewVORD20Caller(Config.ContractCallerAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(ownerKey))
	if err != nil || rootVORD20Owner == nil {
		t.Error(err)
	}

	rootMockERC20, err := chaincall.NewMockERC20Caller(Config.MockContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(ownerKey))
	if err != nil {
		t.Error(err)
	}

	//	send money to VORD20
	amount := big.NewInt(10000)
	tx, err = rootMockERC20.Transfer(Config.ContractCallerAddress, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set VORD20 keyHash
	tx, err = rootVORD20Owner.SetKeyHash(oracleKeyHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	set VORD20 fee
	tx, err = rootVORD20Owner.SetFee(big.NewInt(200))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	create requestRandomness for oracle
	tx, err = rootVORD20Owner.RollDice(big.NewInt(24), "0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe")
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)

	//	run check
	time.Sleep(time.Second * 1)
	err = oracleVORCoordinatorListener.ProcessIncommingEvents()
	if err != nil {
		t.Error(err)
	}

	//	run jobs
	time.Sleep(time.Second * 1)
	err = oracleVORCoordinatorListener.CheckJobs()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 1)
	house, err := rootVORD20Owner.House("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe")
	if err != nil {
		t.Error(err)
	}
	t.Log(house)
}
