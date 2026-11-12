package store

import (
	"context"

	"oracle/store/db"
)

// Store is the oracle's database and keystore. Db is one chain's view of the database: nil on the Store
// NewStore returns, set on the one ForChain returns.
type Store struct {
	root       *db.DB
	Db         *db.ChainDB
	Keystorage IKeystorageStore
}

func NewStore(ctx context.Context, keystorage IKeystorageStore) (*Store, error) {
	root, err := db.NewDb()
	if err != nil {
		return nil, err
	}
	return &Store{root: root, Keystorage: keystorage}, nil
}

// Migrate brings the database's schema up to date for an oracle serving the chains with network ids chainIDs.
func (s *Store) Migrate(chainIDs []int64) error {
	return s.root.Migrate(chainIDs)
}

// ForChain is the store as the chain with network id chainID sees it.
func (s *Store) ForChain(chainID int64) *Store {
	return &Store{root: s.root, Db: s.root.ForChain(chainID), Keystorage: s.Keystorage}
}
