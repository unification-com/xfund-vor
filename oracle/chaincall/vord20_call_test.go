package chaincall_test

import (
	"math/big"
	"os"
	"testing"
)

func TestVORD20Caller_RollDice(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.RollDice(big.NewInt(23))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

func TestVORD20Caller_Fee(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.SetFee(big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}
