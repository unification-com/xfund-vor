package store

type IKeystorageStore interface {
	ExistsByUsername(account string) bool
	GeneratePrivate(username string) (string, error)
	AddExisting(username string, privateKey string) (err error)
	SelectPrivateKey(account string) (err error)
	GetSelectedPrivateKey() string
	SetBlockNumber(blockNumber int64) (err error)
	GetBlockNumber() (blockNumber int64, err error)
}
