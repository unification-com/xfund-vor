package chaincall_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"oracle/chaincall"
	"testing"
)

const (
	VORCoordinatorAddress = "0x22F043993312CB050E7F7A5C1207f68a05D3ef66"
)

var oraclePrivateKey = []byte("0xf54ca099a480e75a417a676855aed602f559d27f6f461f3754667b0b8af11ba6")
var oraclePublicKey = []byte("0x0438500622d7c0366e362e0abe96c2b724e8c8361fb60e2d16c22c41c109aa58a46af6a38cc36d71f75020c78560a910d35c6e39311176dee20f5f26e44eb74882")
var oracleAddress = []byte("0x0f4EE406Ef8cD37b655Fe0bb9F7ebADf17f47033")

func VORCoordinatorCallerTestValues() (string, string, *big.Int, []byte, []byte, []byte) {
	return VORCoordinatorAddress, "http://192.168.1.2:7545", big.NewInt(5777), oraclePrivateKey, oraclePublicKey, oracleAddress
}
func TestVORCoordinatorCaller_GetTotalGasDeposits(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	if err != nil {
		t.Error(err)
	}
	GasDeposits, err := VORCoordinator.GetTotalGasDeposits(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
}
func TestVORCoordinatorCaller_GetGasTopUpLimit(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	if err != nil {
		t.Error(err)
	}
	GasDeposits, err := VORCoordinator.GetGasTopUpLimit(bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}
	t.Log(GasDeposits)
}

func TestVORCoordinatorCaller_Withdraw(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.Withdraw("0x04FBC34DCf60c88e701a8B3161154451e33Eef75", *big.NewInt(100))
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_ChangeFee(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.ChangeFee(big.NewInt(1))
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}

func TestVORCoordinatorCaller_RegisterProvingKey(t *testing.T) {
	VORCoordinator, err := chaincall.NewVORCoordinatorCaller(VORCoordinatorCallerTestValues())
	if err != nil {
		t.Error(err)
	}
	TransactOut, err := VORCoordinator.RegisterProvingKey(big.NewInt(1), string(oracleAddress), false)
	if err != nil {
		t.Error(err)
	}
	t.Log(TransactOut)
}
