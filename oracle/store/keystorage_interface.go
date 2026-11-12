package store

// IKeystorageStore is the keystore as the service and listeners use it, once start-up has unlocked it.
type IKeystorageStore interface {
	ExistsByUsername(account string) bool
	AddExisting(username string, privateKey string) (err error)
	GetSelectedPrivateKey() string
}
