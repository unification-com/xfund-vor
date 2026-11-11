package keystorage

import (
	"crypto/ecdsa"
	"fmt"

	gethkeystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

// scrypt cost of every key this oracle encrypts: go-ethereum's standard (256 MB, about a second per key). Only
// this package's tests lower it.
var scryptN, scryptP = gethkeystore.StandardScryptN, gethkeystore.StandardScryptP

// encryptV3 encrypts privateKey as a go-ethereum v3 keystore and checks the result decrypts back to the same,
// working key before returning it: a key that can't be recovered must never be written.
func encryptV3(privateKey *ecdsa.PrivateKey, passphrase string) ([]byte, common.Address, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("generating the keystore id: %w", err)
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	blob, err := gethkeystore.EncryptKey(&gethkeystore.Key{Id: id, Address: address, PrivateKey: privateKey}, passphrase, scryptN, scryptP)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("encrypting the key: %w", err)
	}
	if err = verifyV3(blob, passphrase, address); err != nil {
		return nil, common.Address{}, fmt.Errorf("the newly encrypted key does not decrypt back to the same key: %w", err)
	}
	return blob, address, nil
}

// decryptV3 decrypts a v3 keystore and checks the key matches the address it is stored under, which catches a
// corrupted or tampered keystore.
func decryptV3(blob []byte, passphrase string, address common.Address) (*ecdsa.PrivateKey, error) {
	key, err := gethkeystore.DecryptKey(blob, passphrase)
	if err != nil {
		return nil, err
	}
	if derived := crypto.PubkeyToAddress(key.PrivateKey.PublicKey); derived != address || key.Address != address {
		return nil, fmt.Errorf("the key decrypts to %s, but is stored as %s", derived.Hex(), address.Hex())
	}
	return key.PrivateKey, nil
}

// verifyV3 proves a v3 keystore yields a key that signs as address: it decrypts it, signs a probe digest and
// recovers the signer.
func verifyV3(blob []byte, passphrase string, address common.Address) error {
	privateKey, err := decryptV3(blob, passphrase, address)
	if err != nil {
		return err
	}
	digest := crypto.Keccak256([]byte("xfund-vor keystore verification probe"))
	signature, err := crypto.Sign(digest, privateKey)
	if err != nil {
		return fmt.Errorf("signing the probe: %w", err)
	}
	signer, err := crypto.SigToPub(digest, signature)
	if err != nil {
		return fmt.Errorf("recovering the probe's signer: %w", err)
	}
	if crypto.PubkeyToAddress(*signer) != address {
		return fmt.Errorf("the probe recovers a signer other than %s", address.Hex())
	}
	return nil
}
