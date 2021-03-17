package walletworker

import (
	crypt "crypto"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func GeneratePrivate() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	return privateKey, err
}

func StringToPrivate(bytePrivateKey []byte) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.ToECDSA(bytePrivateKey)
	return privateKey, err
}

func GeneratePublic(privateKey *ecdsa.PrivateKey) crypt.PublicKey {
	publicKey := privateKey.Public()
	return publicKey
}
