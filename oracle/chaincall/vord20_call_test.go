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

func TestVORD20Caller_TopUpGas(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	keyHash, err := VORD20Caller.KeyHash()
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.TopUpGas(keyHash)
	if err != nil {
		t.Error(err)
	}

	t.Log(tx)
}

//{0x43475BA872641b28Ad5C840f292DA06Dcfb8Dc4d 0xf1CcaE4f01db7F5D363d38B7499e7facC65CEC54 0 20000000000 100000 [230 28 81 202 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 134 160] []}

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

func TestVORD20Caller_House(t *testing.T) {
	err := Init(os.Args[len(os.Args)-1])
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.House()
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
