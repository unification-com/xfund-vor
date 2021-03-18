package walletworker

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func GeneratePrivate() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.GenerateKey()
	return privateKey, err
}

func StringToPrivate(bytePrivateKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(bytePrivateKey)
	return privateKey, err
}

func GeneratePublic(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	return publicKeyECDSA
}

func GenerateAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}
