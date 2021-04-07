package store

import (
	"context"
	"oracle/store/db"
)

type Store struct {
	Db                *db.DB
	RandomnessRequest IRandomnessRequestStore
	Keystorage        IKeystorageStore
}

func NewStore(ctx context.Context, keystorage IKeystorageStore) (*Store, error) {
	var dbConn, err = db.NewDb()
	if err != nil {
		return nil, err
	}
	var store Store
	if dbConn != nil {
		store.Db = dbConn
		store.RandomnessRequest = db.NewRandomnessRequestStore(dbConn)
		store.Keystorage = keystorage
	}
	return &store, err
}
