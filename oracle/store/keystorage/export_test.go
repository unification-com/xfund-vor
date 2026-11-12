package keystorage

import gethkeystore "github.com/ethereum/go-ethereum/accounts/keystore"

// Tests encrypt with go-ethereum's light scrypt parameters (4 MB, milliseconds) instead of the standard ones.
func init() {
	scryptN, scryptP = gethkeystore.LightScryptN, gethkeystore.LightScryptP
}
