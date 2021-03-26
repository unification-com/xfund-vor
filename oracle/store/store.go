package store

import (
	"context"
	"oracle/store/postgresql"
	"oracle/store/sqlite"
	"time"
)

const KeepAlivePeriod = 3

type Store struct {
	Pg     *postgresql.DB
	SQLite *sqlite.DB
}

func NewStore(ctx context.Context) (*Store, error) {
	var sqliteConn, err = sqlite.NewDB()
	if err != nil {
		return nil, err
	}
	var store Store
	if sqliteConn != nil {
		store.SQLite = sqliteConn
	}
	return &store, err
}

func (d Store) KeepAlivePostgres() {
	var err error
	for {
		time.Sleep(time.Second * KeepAlivePeriod)
		var lostConnection = false
		if d.Pg == nil {
			lostConnection = true
		} else if err = d.Pg.Exec("SELECT 1").Error; err != nil {
			lostConnection = true
		}
		if !lostConnection {
			continue
		}
		d.Pg, err = postgresql.Dial()
		if err != nil {
			continue
		}
	}
}
