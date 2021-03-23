package keystorage_test

import (
	"github.com/sirupsen/logrus"
	"oracle/store/keystorage"
	"os"
	"testing"
)

var Log = logrus.New()

func TestKeystorage_NewKeyStorage(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(*(keystore.KeyStore.GetKey()))
}

func TestKeystorage_Exists(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(keystore.Exists())
}

func TestKeystorage_AddGenerated(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	err = keystore.AddGenerated("testaccount")
	if err != nil {
		t.Error(err)
	}
}
