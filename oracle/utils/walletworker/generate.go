package walletworker

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GeneratePrivate() (*ecdsa.PrivateKey, string, error) {
	privateKey, err := crypto.GenerateKey()
	return privateKey, hexutil.Encode(crypto.FromECDSA(privateKey)), err
}

func StringToPrivate(bytePrivateKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(bytePrivateKey)
	return privateKey, err
}

func GeneratePublic(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, string) {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	return publicKeyECDSA, hexutil.Encode(crypto.FromECDSAPub(publicKey.(*ecdsa.PublicKey)))
}

func GenerateAddress(publicKeyECDSA *ecdsa.PublicKey) (common.Address, string) {
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, address.Hex()
}
