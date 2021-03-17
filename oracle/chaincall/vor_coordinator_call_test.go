package chaincall_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"oracle/chaincall"
	"testing"
)

const (
	VORCoordinatorAddress = "0xad22a7bf841F553a6feB4AA1A3EbBb7aEbEdf740"
	OraclePubkey          = []byte("")
)

func TestVORCoordinatorCaller_GetTotalGasDeposits(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorAddress, "http://127.0.0.1:7546", OraclePubkey)
	if err != nil {
		t.Error(err)
	}
	GasDeposits, err := VORCoordinator.GetTotalGasDeposits(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
}

func TestVORCoordinatorCaller_Withdraw(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorAddress, "http://127.0.0.1:57810", OraclePubkey)
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.Withdraw(bind.TransactOpts{}, "0x04FBC34DCf60c88e701a8B3161154451e33Eef75", *big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_ChangeFee(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorAddress, "http://127.0.0.1:57810")
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.ChangeFee(bind.TransactOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}
