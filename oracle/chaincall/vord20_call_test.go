package chaincall_test

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"path/filepath"
	"testing"
)

func TestVORD20Caller_AddXfund(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
	if err != nil {
		t.Error(err)
	}
	amount := big.NewInt(1000000000000000000)
	tx, err := MockERC20Caller.Transfer(Config.ContractCallerAddress, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

func TestVORD20Caller_RollDice(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
	if err != nil {
		t.Error(err)
	}
	tx, err := VORD20Caller.SetFee(big.NewInt(1))
	tx, err = VORD20Caller.IncreaseVORCoordinatorAllowance(big.NewInt(1000000000))
	tx, err = VORD20Caller.RollDice(big.NewInt(24), "0xF0D5BC18421fa04D0a2A2ef540ba5A9f04014BE3")
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

func TestVORD20Caller_SetFee(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.SetFee(big.NewInt(2))
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}

//{0x43475BA872641b28Ad5C840f292DA06Dcfb8Dc4d 0xf1CcaE4f01db7F5D363d38B7499e7facC65CEC54 0 20000000000 100000 [230 28 81 202 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 134 160] []}

func TestVORD20Caller_KeyHash2(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
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
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
	if err != nil {
		t.Error(err)
	}

	tx, err := VORD20Caller.House("0xF0D5BC18421fa04D0a2A2ef540ba5A9f04014BE3")
	if err != nil {
		assert.Equal(t, "VM Exception while processing transaction: revert Roll in progress", err.Error())
	}
	t.Log(tx)
}

func TestVORD20Caller_Fee(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
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
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
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
