package keystorage_test

import (
	"github.com/sirupsen/logrus"
	"oracle/store/keystorage"
	"os"
	"testing"
)

var Log = logrus.New()

func TestNewKeyStorage(t *testing.T) {
	keystoragePath := os.Args[len(os.Args)-1]

	keystore, err := keystorage.NewKeyStorage(Log, keystoragePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(*keystore)
}
