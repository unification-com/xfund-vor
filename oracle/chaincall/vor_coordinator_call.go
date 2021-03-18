package chaincall

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"oracle/contracts/vor_coordinator"
	//"oracle/service/secp256k1"
	secp256k1_2 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type VORCoordinatorCaller struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_coordinator.VORCoordinator
	transactOpts    *bind.TransactOpts

	publicProvingKey [2]*big.Int
}

func pair(x, y *big.Int) [2]*big.Int { return [2]*big.Int{x, y} }

func NewVORCoordinatorCaller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte, oraclePublicKey []byte, oracleAddress []byte) (*VORCoordinatorCaller, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.BytesToAddress([]byte(contractStringAddress))
	instance, err := vor_coordinator.NewVORCoordinator(contractAddress, client)
	if err != nil {
		return nil, err
	}
	oraclePrivateKeyECDSA, err := crypto.HexToECDSA(string(oraclePrivateKey[2:]))
	if err != nil {
		return nil, err
	}
	transactOpts, err := bind.NewKeyedTransactorWithChainID(oraclePrivateKeyECDSA, chainID)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), common.BytesToAddress(oracleAddress))
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(1)
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = uint64(1000000) // in units
	transactOpts.Context = context.Background()

	return &VORCoordinatorCaller{
		client:           client,
		contractAddress:  contractAddress,
		instance:         instance,
		transactOpts:     transactOpts,
		publicProvingKey: pair(secp256k1_2.DecompressPubkey(oraclePublicKey)),
	}, err
}

func (d *VORCoordinatorCaller) GetTotalGasDeposits(bindOpts bind.CallOpts) (*big.Int, error) {
	return d.instance.GetTotalGasDeposits(&bindOpts)
}

func (d *VORCoordinatorCaller) GetGasTopUpLimit(bindOpts bind.CallOpts) (*big.Int, error) {
	return d.instance.GetGasTopUpLimit(&bindOpts)
}

//func (d *VORCoordinatorCaller) GetProviderAddress(bindOpts bind.CallOpts, keyHash string) (common.Address, error) {
//	return d.instance.GetProviderAddress(&bindOpts, [32]byte(keyHash))
//}

func (d *VORCoordinatorCaller) Withdraw(recipientAddress string, amount big.Int) (*types.Transaction, error) {
	recipientAddr := common.HexToAddress(recipientAddress)
	return d.instance.Withdraw(d.transactOpts, recipientAddr, &amount)
}

func (d *VORCoordinatorCaller) RegisterProvingKey(amount *big.Int, oracleAddress string, providerPaysGas bool) (*types.Transaction, error) {
	oracleAddr := common.HexToAddress(oracleAddress)
	transaction, err := d.instance.RegisterProvingKey(d.transactOpts, amount, oracleAddr, d.publicProvingKey, providerPaysGas)
	return transaction, err
}

func (d *VORCoordinatorCaller) ChangeFee(fee *big.Int) (*types.Transaction, error) {
	return d.instance.ChangeFee(d.transactOpts, d.publicProvingKey, fee)
}

func (d *VORCoordinatorCaller) SetProviderPaysGas(bindOpts bind.TransactOpts, providerPaysFee bool) (*types.Transaction, error) {
	return d.instance.SetProviderPaysGas(&bindOpts, d.publicProvingKey, providerPaysFee)
}

func (d *VORCoordinatorCaller) FulfillRandomnessRequest(bindOpts bind.TransactOpts, proof []byte) (*types.Transaction, error) {
	return d.instance.FulfillRandomnessRequest(&bindOpts, proof)
}
