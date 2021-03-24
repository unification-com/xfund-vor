package chaincall_test

import (
	"math/big"
	"oracle/chaincall"
	"oracle/config"
	"oracle/store/keystorage"
	"os"
	"testing"
)

var VORD20 *chaincall.VORD20Caller

func InitVORD20(configAddress string) (err error) {
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
	VORD20, err = chaincall.NewVORD20Caller(VORD20CallerTestValues())
	return err
}

func VORD20CallerTestValues() (string, string, *big.Int, []byte) {
	return Config.VORD20ContractAddress, Config.EthHTTPHost, big.NewInt(Config.NetworkID), []byte(Keystore.GetFirst().CipherPrivate)
}

func TestVORD20Caller_RollDice(t *testing.T) {
	err := InitVORD20(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20.RollDice(2)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}
