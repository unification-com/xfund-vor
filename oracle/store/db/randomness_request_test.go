package db_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/config"
	"oracle/store"
	"oracle/store/keystorage"
	"runtime/debug"
	"testing"
)

var keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func TestRandomnessRequestStore_Insert(t *testing.T) {
	var err error
	keystore, err = keystorage.NewKeyStorage(Log, "../../test_data/generic_keystore.json")
	if err != nil || keystore == nil {
		t.Error(err)
	}
	thestore, err := store.NewStore(context.Background(), keystore)
	if err != nil || thestore == nil {
		t.Error(err)
	}
	err = thestore.Db.Migrate()
	if err != nil {
		t.Error(err)
	}
	err = thestore.Db.InsertNewRequest(
		"keyHashStore",
		"Sender",
		"RequestId",
		1,
		"txHash",
		1, 1, 1)
	debug.PrintStack()
	if err != nil {
		t.Error(err)
	}
}
