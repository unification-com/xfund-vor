package walletworker_test

import (
	"github.com/ethereum/go-ethereum/crypto"
	"oracle/walletworker"
	"testing"
)

const privateKey = ""

func TestGeneratePrivate(t *testing.T) {
	privateKey, err := walletworker.GeneratePrivate()
	t.Error(err)
	t.Log(privateKey)
	t.Log(crypto.FromECDSA(privateKey))
}

func TestGeneratePublic(t *testing.T) {
	privateKey, err := walletworker.GeneratePrivate()
	t.Error(err)
	t.Log(privateKey)
	publicKey := walletworker.GeneratePublic(privateKey)
	t.Log(publicKey)
}

func TestGeneratePublic2(t *testing.T) {
	privateKey, err := walletworker.StringToPrivate([]byte(privateKey))
	t.Error(err)
	t.Log(privateKey)
	publicKey := walletworker.GeneratePublic(privateKey)
	t.Log(publicKey)
}
