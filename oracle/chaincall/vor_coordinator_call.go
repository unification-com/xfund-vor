package chaincall

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"oracle/contracts/vor_coordinator"
	"oracle/walletworker"
)

type VORCoordinatorCaller struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_coordinator.VORCoordinator
	transactOpts    *bind.TransactOpts

	publicProvingKey [2]*big.Int
}

func pair(x, y *big.Int) [2]*big.Int { return [2]*big.Int{x, y} }

func NewVORCoordinatorCaller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*VORCoordinatorCaller, error) {
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

	oraclePublicKey := oraclePrivateKeyECDSA.Public()
	log.Print("Public Key: ", hexutil.Encode(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey))))

	ECDSAoraclePublicKey, err := crypto.UnmarshalPubkey(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey)))
	if err != nil || ECDSAoraclePublicKey == nil  {
		log.Print(err)
		log.Print(ECDSAoraclePublicKey)
		return nil, err
	}
	oracleAddress := walletworker.GenerateAddress(ECDSAoraclePublicKey)
	log.Print("Address: ", oracleAddress)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(oraclePrivateKeyECDSA, chainID)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(oracleAddress))
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(0)
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = uint64(100000) // in units
	transactOpts.Context = context.Background()

	return &VORCoordinatorCaller{
		client:           client,
		contractAddress:  contractAddress,
		instance:         instance,
		transactOpts:     transactOpts,
		publicProvingKey: [2]*big.Int{ECDSAoraclePublicKey.X, ECDSAoraclePublicKey.Y},
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

func (d *VORCoordinatorCaller) Withdraw(recipientAddress string, amount *big.Int) (*types.Transaction, error) {
	recipientAddr := common.HexToAddress(recipientAddress)
	return d.instance.Withdraw(d.transactOpts, recipientAddr, amount)
}

func (d *VORCoordinatorCaller) RegisterProvingKey(fee big.Int, oracleAddress string, providerPaysGas bool) (*types.Transaction, error) {
	oracleAddr := common.BytesToAddress([]byte(oracleAddress))
	log.Print(*d.transactOpts)
	log.Print(d.publicProvingKey)
	transaction, err := d.instance.RegisterProvingKey(d.transactOpts, &fee, oracleAddr, d.publicProvingKey, providerPaysGas)
	return transaction, err
}

func (d *VORCoordinatorCaller) ChangeFee(fee *big.Int) (*types.Transaction, error) {
	return d.instance.ChangeFee(d.transactOpts, d.publicProvingKey, fee)
}

func (d *VORCoordinatorCaller) SetProviderPaysGas(providerPaysFee bool) (*types.Transaction, error) {
	return d.instance.SetProviderPaysGas(d.transactOpts, d.publicProvingKey, providerPaysFee)
}

func (d *VORCoordinatorCaller) FulfillRandomnessRequest(proof []byte) (*types.Transaction, error) {
	return d.instance.FulfillRandomnessRequest(d.transactOpts, proof)
}
