package sqlite_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/config"
	"oracle/store"
	"oracle/store/keystorage"
	"oracle/store/sqlite"
	"runtime/debug"
	"testing"
)

var sqliteDB *sqlite.DB
var keystore *keystorage.Keystorage
var Config *config.Config
var Log = logrus.New()

func TestRandomnessRequestStore_Insert(t *testing.T) {
	var err error
	keystore, err = keystorage.NewKeyStorage(Log, "../../keystore.json")
	if err != nil || keystore == nil {
		t.Error(err)
	}
	thestore, err := store.NewStore(context.Background(), keystore)
	if err != nil || thestore == nil {
		t.Error(err)
	}
	err = thestore.RandomnessRequest.Migrate()
	if err != nil {
		t.Error(err)
	}
	err = thestore.RandomnessRequest.Insert("keyHashStore", "SeedHex", "SenderHex", "RequestIDHex", "BlockHashHex", 1, "TransactionID", "Status")
	debug.PrintStack()
	if err != nil {
		t.Error(err)
	}
}
