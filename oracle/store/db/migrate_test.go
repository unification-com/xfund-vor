package db_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"oracle/models/database"
	"oracle/store/db"
)

// The v0.1 schema, as v0.1's models declared it: no chain ids, request ids unique across the table.
type RandomnessRequest struct {
	gorm.Model
	KeyHash            string
	Sender             string `gorm:"index"`
	RequestId          string `gorm:"uniqueIndex"`
	RequestBlockNumber uint64 `gorm:"index"`
	RequestTxHash      string `gorm:"index"`
	Fee                uint64
	Status             int `gorm:"index"`
}

type FailedFulfilment struct {
	gorm.Model
	RequestId  string `gorm:"index"`
	TxHash     string `gorm:"index"`
	FailReason string
}

type BlocksStored struct {
	gorm.Model
	BlockHash   string `gorm:"index"`
	BlockNumber uint64 `gorm:"index"`
	TxHash      string `gorm:"index"`
}

func (BlocksStored) TableName() string { return "blocks_stored" }

// databases are the databases the migration is tested on: SQLite, and PostgreSQL when VOR_TEST_POSTGRES_DSN
// names an empty one.
func databases(t *testing.T) map[string]func(t *testing.T) *gorm.DB {
	open := func(t *testing.T, dialector gorm.Dialector) *gorm.DB {
		g, err := gorm.Open(dialector, &gorm.Config{Logger: logger.Discard})
		require.NoError(t, err)
		return g
	}
	dbs := map[string]func(t *testing.T) *gorm.DB{
		"sqlite": func(t *testing.T) *gorm.DB { return open(t, sqlite.Open(filepath.Join(t.TempDir(), "oracle.db"))) },
	}
	if dsn := os.Getenv("VOR_TEST_POSTGRES_DSN"); dsn != "" {
		dbs["postgres"] = func(t *testing.T) *gorm.DB {
			g := open(t, postgres.Open(dsn))
			t.Cleanup(func() {
				_ = g.Migrator().DropTable(&database.RandomnessRequest{}, &database.FailedFulfilment{}, &database.BlocksStored{}, &database.ChainCursor{})
			})
			return g
		}
	}
	return dbs
}

// v01Database is a database a v0.1 oracle used: its schema and some records.
func v01Database(t *testing.T, g *gorm.DB) {
	require.NoError(t, g.AutoMigrate(&RandomnessRequest{}, &FailedFulfilment{}, &BlocksStored{}))
	require.NoError(t, g.Create(&RandomnessRequest{RequestId: "r1", Sender: "0xa", Fee: 25, Status: 3, RequestBlockNumber: 100}).Error)
	require.NoError(t, g.Create(&RandomnessRequest{RequestId: "r2", Sender: "0xb", Fee: 25, Status: 1, RequestBlockNumber: 200}).Error)
	deleted := &RandomnessRequest{RequestId: "r3", Status: 3}
	require.NoError(t, g.Create(deleted).Error)
	require.NoError(t, g.Delete(deleted).Error) // soft-deleted rows are assigned too
	require.NoError(t, g.Create(&FailedFulfilment{RequestId: "r2", TxHash: "0xf", FailReason: "reverted"}).Error)
	require.NoError(t, g.Create(&BlocksStored{BlockHash: "0xbh", BlockNumber: 150, TxHash: "0xs"}).Error)
}

func TestMigrateV01Database(t *testing.T) {
	for name, newDB := range databases(t) {
		t.Run(name, func(t *testing.T) {
			g := newDB(t)
			v01Database(t, g)
			store := &db.DB{DB: g}

			require.NoError(t, store.Migrate([]int64{137}))

			// every record is assigned to the one chain, soft-deleted ones included
			for table, want := range map[interface{}]int64{&database.RandomnessRequest{}: 3, &database.FailedFulfilment{}: 1, &database.BlocksStored{}: 1} {
				var n int64
				require.NoError(t, g.Unscoped().Model(table).Where("chain_id = ?", 137).Count(&n).Error)
				assert.Equal(t, want, n, "%T", table)
			}

			migrator := g.Migrator()
			assert.False(t, migrator.HasIndex(&database.RandomnessRequest{}, "idx_randomness_requests_request_id"), "v0.1's table-wide unique index is gone")
			assert.True(t, migrator.HasIndex(&database.RandomnessRequest{}, "idx_randomness_requests_chain_request"))

			polygon, ethereum := store.ForChain(137), store.ForChain(1)
			found, err := polygon.FindByRequestId("r1")
			require.NoError(t, err)
			assert.Equal(t, uint64(25), found.Fee)
			_, err = ethereum.FindByRequestId("r1")
			assert.True(t, db.IsNotFound(err), "another chain's request is invisible")

			// request ids are unique per chain, not across chains
			require.NoError(t, ethereum.InsertNewRequest("kh", "0xc", "r1", database.REQUEST_STATUS_INITIALISED, "0xtx", 0, 0, 25))
			assert.Error(t, polygon.InsertNewRequest("kh", "0xc", "r1", database.REQUEST_STATUS_INITIALISED, "0xtx", 0, 0, 25))
			jobs, err := ethereum.GetJobs()
			require.NoError(t, err)
			require.Len(t, jobs, 1)
			assert.Equal(t, int64(1), jobs[0].ChainId)
			require.NoError(t, ethereum.UpdateRequestStatus("r1", database.REQUEST_STATUS_TX_FAILED, "nonce too low"))
			found, err = polygon.FindByRequestId("r1")
			require.NoError(t, err)
			assert.Equal(t, 3, found.Status, "updating one chain's request leaves the other's alone")

			// each chain has its own cursor
			next, err := polygon.NextBlock()
			require.NoError(t, err)
			assert.Zero(t, next)
			require.NoError(t, polygon.SetNextBlock(500))
			require.NoError(t, polygon.SetNextBlock(600))
			require.NoError(t, ethereum.SetNextBlock(9000))
			next, err = polygon.NextBlock()
			require.NoError(t, err)
			assert.Equal(t, uint64(600), next)

			// a second start changes nothing
			require.NoError(t, store.Migrate([]int64{137, 1}))
			var n int64
			require.NoError(t, g.Unscoped().Model(&database.RandomnessRequest{}).Count(&n).Error)
			assert.Equal(t, int64(4), n)
		})
	}
}

// Records from before chain ids can't be assigned when the oracle serves several chains.
func TestMigrateRefusesToGuessTheChain(t *testing.T) {
	for name, newDB := range databases(t) {
		t.Run(name, func(t *testing.T) {
			g := newDB(t)
			v01Database(t, g)
			err := (&db.DB{DB: g}).Migrate([]int64{137, 1})
			require.Error(t, err)
			assert.Contains(t, err.Error(), "predate chain ids")
		})
	}
}
