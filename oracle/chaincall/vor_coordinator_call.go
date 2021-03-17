package chaincall

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"oracle/contracts/vor_coordinator"
	"oracle/service/secp256k1"
)

type VORCoordinatorCaller struct {
	contractAddress  common.Address
	client           *ethclient.Client
	instance         *vor_coordinator.Token
	publicProvingKey [2]*big.Int
}

func pair(x, y *big.Int) [2]*big.Int { return [2]*big.Int{x, y} }

func NewVORCoordinatorCaller(contractStringAddress string, ethHostAddress string, oraclePublicKey []byte) (*VORCoordinatorCaller, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.BytesToAddress([]byte(contractStringAddress))
	instance, err := vor_coordinator.NewToken(contractAddress, client)
	if err != nil {
		return nil, err
	}
	kyberOraclePublicKey, err := secp256k1.LongUnmarshal(oraclePublicKey)
	return &VORCoordinatorCaller{
		client:           client,
		contractAddress:  contractAddress,
		instance:         instance,
		publicProvingKey: pair(secp256k1.Coordinates(kyberOraclePublicKey)),
	}, err
}

func (d *VORCoordinatorCaller) GetTotalGasDeposits(bindOpts bind.CallOpts) (*big.Int, error) {
	return d.instance.GetTotalGasDeposits(&bindOpts)
}

func (d *VORCoordinatorCaller) Withdraw(bindOpts bind.TransactOpts, recipientAddress string, amount big.Int) (*types.Transaction, error) {
	recipientAddr := common.HexToAddress(recipientAddress)
	return d.instance.Withdraw(&bindOpts, recipientAddr, &amount)
}

func (d *VORCoordinatorCaller) RegisterProvingKey(bindOpts bind.TransactOpts, amount big.Int, oracleAddress string, publicProvingKey [2]*big.Int, providerPaysGas bool) (*types.Transaction, error) {
	oracleAddr := common.HexToAddress(oracleAddress)
	return d.instance.RegisterProvingKey(&bindOpts, &amount, oracleAddr, publicProvingKey, providerPaysGas)
}

func (d *VORCoordinatorCaller) ChangeFee(bindOpts bind.TransactOpts, publicProvingKey [2]*big.Int, fee big.Int) (*types.Transaction, error) {
	return d.instance.ChangeFee(&bindOpts, publicProvingKey, &fee)
}

func (d *VORCoordinatorCaller) SetProviderPaysGas(bindOpts bind.TransactOpts, publicProvingKey [2]*big.Int, providerPaysFee bool) (*types.Transaction, error) {
	return d.instance.SetProviderPaysGas(&bindOpts, publicProvingKey, providerPaysFee)
}

func (d *VORCoordinatorCaller) FulfillRandomnessRequest(bindOpts bind.TransactOpts, publicProvingKey [2]*big.Int, proof []byte) (*types.Transaction, error) {
	return d.instance.FulfillRandomnessRequest(&bindOpts, proof)
}
