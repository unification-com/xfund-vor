package chaincall_test

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"
)

func TestMockERC20Caller_Transfer(t *testing.T) {
	dir, _ := os.Getwd()
	configPath := filepath.Join(dir, "..", "test_data", "erc20_transfer_test_config.json")
	err := Init(configPath, "fq516b1boc8vrm7nasnb8fy7u5rb6zhh")
	if err != nil {
		t.Error(err)
	}
	amount := big.NewInt(100000)
	amount.Mul(amount, big.NewInt(10))
	tx, err := MockERC20Caller.Transfer(Config.ContractCallerAddress, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log(tx)
}
