package chaincall_test

import (
	"math/big"
	"os"
	"testing"
)

func TestMockERC20Caller_Transfer(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}
	amount := big.NewInt(1000000000000000000)
	amount.Mul(amount, big.NewInt(10))
	tx, err := MockERC20Caller.Transfer("0xf56C666822fd97d71604BCF6aBe7BF062120fd15", amount)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}
