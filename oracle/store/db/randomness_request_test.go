package db_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/config"
	"oracle/store"
	"oracle/store/keystorage"
	"path/filepath"
	"runtime/debug"
	"testing"
)

var keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func TestRandomnessRequestStore_Insert(t *testing.T) {
	var err error
	// a fresh database per run: the insert below uses a fixed request id, so a database left by an earlier
	// run fails it on the unique index
	config.Conf.Database.Dialect = "sqlite"
	config.Conf.Database.Storage = filepath.Join(t.TempDir(), "oracle.db")
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
