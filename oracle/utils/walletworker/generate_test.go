package walletworker_test

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"oracle/utils/walletworker"
	"testing"
)

const privateKeyPreinit = "0xf54ca099a480e75a417a676855aed602f559d27f6f461f3754667b0b8af11ba6"
const publicKeyPreinit = "0x0438500622d7c0366e362e0abe96c2b724e8c8361fb60e2d16c22c41c109aa58a46af6a38cc36d71f75020c78560a910d35c6e39311176dee20f5f26e44eb74882"

func TestGeneratePrivate(t *testing.T) {
	privateKey, stringPrivateKey, err := walletworker.GeneratePrivate()
	if err != nil {
		t.Error("error: ", err)
	}
	t.Log("Private Key: ", hexutil.Encode(crypto.FromECDSA(privateKey)))
	t.Log("String Private Key: ", stringPrivateKey)
}

func TestGeneratePublic(t *testing.T) {
	privateKey, stringPrivateKey, err := walletworker.GeneratePrivate()
	if err != nil {
		t.Error("error: ", err)
	}
	t.Log("Private Key: ", hexutil.Encode(crypto.FromECDSA(privateKey)))
	t.Log("String Private Key: ", stringPrivateKey)
	_, publicKey := walletworker.GeneratePublic(privateKey)
	t.Log(publicKey)
}

func TestGeneratePublic2(t *testing.T) {
	privateKey, err := walletworker.StringToPrivate(string([]byte(privateKeyPreinit)[2:]))
	if err != nil {
		t.Error("error: ", err)
	}
	t.Log("Private Key: ", hexutil.Encode(crypto.FromECDSA(privateKey)))
	publicKey, _ := walletworker.GeneratePublic(privateKey)
	t.Log("Public Key: ", hexutil.Encode(crypto.FromECDSAPub(publicKey)))
	_, address := walletworker.GenerateAddress(publicKey)
	t.Log("Address: ", address)
}

func TestGenerateAddress(t *testing.T) {
	ECDSApublicKey, err := crypto.UnmarshalPubkey([]byte(publicKeyPreinit)[2:])
	if err != nil {
		t.Error("error: ", err)
	}
	_, address := walletworker.GenerateAddress(ECDSApublicKey)
	t.Log("Address: ", address)
}
