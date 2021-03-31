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

	tx, err := VORD20Caller.RollDice(big.NewInt(26))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

func TestVORD20Caller_SetFee(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.SetFee(big.NewInt(1000000000000000000))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

func TestVORD20Caller_KeyHash2(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.KeyHash()
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

	tx, err := VORD20Caller.Fee()
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

func TestVORD20Caller_Owner(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.Owner()
	if err != nil {
		t.Error(err)
	}
	t.Log(tx.String())
}

//func TestVORD20Caller_SetKeyHash(t *testing.T) {
//	err := Init(os.Args[len(os.Args)-1])
//	if err != nil {
//		t.Error(err)
//	}
//
//
//	copy(hex.DecodeString(""), [32])
//	tx, err := VORD20Caller.SetKeyHash()
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(tx)
//}
