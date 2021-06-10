package chaincall

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"oracle/config"
	"oracle/contracts/block_hash_store"
	"oracle/contracts/vor_coordinator"
	"oracle/utils"
	"oracle/utils/walletworker"
)

type VORCoordinatorCaller struct {
	vorCoordinatorContractAddress common.Address
	blockHashStoreContractAddress common.Address
	client                        *ethclient.Client
	vorCoordinatorInstance        *vor_coordinator.VorCoordinator
	blockHashStoreInstance        *block_hash_store.BlockHashStore
	transactOpts                  *bind.TransactOpts
	callOpts                      *bind.CallOpts

	context          context.Context
	publicProvingKey [2]*big.Int
	oraclePrivateKey string
	oraclePublicKey  string
	oracleAddress    string
}

func NewVORCoordinatorCaller(vorCoordinatorStringAddress string,
	blockHashStoreStringAddress string,
	ethHostAddress string,
	chainID *big.Int,
	oraclePrivateKey []byte) (*VORCoordinatorCaller, error) {
	client, err := ethclient.Dial(ethHostAddress)
	ctx := context.Background()
	if err != nil {
		return nil, err
	}

	vorCoordinatorContractAddress := common.HexToAddress(vorCoordinatorStringAddress)
	vorCoordinatorInstance, err := vor_coordinator.NewVorCoordinator(vorCoordinatorContractAddress, client)
	if err != nil {
		return nil, err
	}

	blockHashContractAddress := common.HexToAddress(blockHashStoreStringAddress)
	blockHashStoreInstance, err := block_hash_store.NewBlockHashStore(blockHashContractAddress, client)

	if err != nil {
		return nil, err
	}

	oraclePrivateKeyECDSA, err := crypto.HexToECDSA(utils.RemoveHexPrefix(string(oraclePrivateKey)))
	if err != nil {
		return nil, err
	}

	oraclePublicKey := oraclePrivateKeyECDSA.Public()
	//log.Print("Public Key: ", hexutil.Encode(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey))))

	ECDSAoraclePublicKey, err := crypto.UnmarshalPubkey(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey)))
	if err != nil || ECDSAoraclePublicKey == nil {
		log.Print(err)
		log.Print(ECDSAoraclePublicKey)
		return nil, err
	}
	_, oracleAddress := walletworker.GenerateAddress(ECDSAoraclePublicKey)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(oraclePrivateKeyECDSA, chainID)
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(oracleAddress))
	if err != nil {
		return nil, err
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(0)

	transactOpts.GasPrice = nil
	transactOpts.GasLimit = uint64(config.Conf.GasLimit) // in units
	transactOpts.Context = ctx

	callOpts := &bind.CallOpts{From: common.HexToAddress(oracleAddress), Context: ctx}

	return &VORCoordinatorCaller{
		client:                        client,
		vorCoordinatorContractAddress: vorCoordinatorContractAddress,
		vorCoordinatorInstance:        vorCoordinatorInstance,
		blockHashStoreContractAddress: blockHashContractAddress,
		blockHashStoreInstance:        blockHashStoreInstance,
		transactOpts:                  transactOpts,
		callOpts:                      callOpts,
		context:                       ctx,
		publicProvingKey:              [2]*big.Int{ECDSAoraclePublicKey.X, ECDSAoraclePublicKey.Y},
		oraclePrivateKey:              string(oraclePrivateKey),
		oraclePublicKey:               hexutil.Encode(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey))),
		oracleAddress:                 oracleAddress,
	}, err
}

func (d *VORCoordinatorCaller) RenewTransactOpts() (err error) {
	nonce, err := d.client.PendingNonceAt(d.context, common.HexToAddress(d.oracleAddress))
	if err != nil {
		return err
	}

	d.transactOpts.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := d.client.SuggestGasPrice(d.context)
	if err != nil {
		return
	}
	d.transactOpts.GasPrice = gasPrice

	if config.Conf.MaxGasPrice > 0 {

		maxGasPrice := big.NewInt(0).Mul(big.NewInt(config.Conf.MaxGasPrice), big.NewInt(params.GWei))
		if gasPrice.Cmp(maxGasPrice) > 0 {
			d.transactOpts.GasPrice = maxGasPrice
		}
	}

	return
}

func (d *VORCoordinatorCaller) HashOfKey() ([32]byte, error) {
	return d.vorCoordinatorInstance.HashOfKey(d.callOpts, d.publicProvingKey)
}

//func (d *VORCoordinatorCallerr) HashOfKeyLocally() ([]byte, error) {
//	utils.Keccak256(d.publicProvingKey)
//	crypto.Keccak256()
//}
//func (d *VORCoordinatorCaller) GetProviderAddress(bindOpts bind.CallOpts, keyHash string) (common.Address, error) {
//	return d.vorCoordinatorInstance.GetProviderAddress(&bindOpts, [32]byte(keyHash))
//}

func (d *VORCoordinatorCaller) Withdraw(recipientAddress string, amount *big.Int) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	recipientAddr := common.HexToAddress(recipientAddress)
	return d.vorCoordinatorInstance.Withdraw(d.transactOpts, recipientAddr, amount)
}

func (d *VORCoordinatorCaller) RegisterProvingKey(fee *big.Int) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	fmt.Println(d.publicProvingKey)
	transaction, err := d.vorCoordinatorInstance.RegisterProvingKey(d.transactOpts, fee, common.HexToAddress(d.oracleAddress), d.publicProvingKey)
	return transaction, err
}

func (d *VORCoordinatorCaller) RandomnessRequest(keyHash [32]byte, consumerSeed *big.Int, feePaid *big.Int) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	transaction, err := d.vorCoordinatorInstance.RandomnessRequest(d.transactOpts, keyHash, consumerSeed, feePaid)
	return transaction, err
}

func (d *VORCoordinatorCaller) ChangeFee(fee *big.Int) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	return d.vorCoordinatorInstance.ChangeFee(d.transactOpts, d.publicProvingKey, fee)
}

func (d *VORCoordinatorCaller) ChangeGranularFee(_consumer common.Address, fee *big.Int) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	return d.vorCoordinatorInstance.ChangeGranularFee(d.transactOpts, d.publicProvingKey, fee, _consumer)
}

func (d *VORCoordinatorCaller) FulfillRandomnessRequest(proof []byte) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	return d.vorCoordinatorInstance.FulfillRandomnessRequest(d.transactOpts, proof)
}

func (d *VORCoordinatorCaller) StoreBlockHash(blockNum uint64) (*types.Transaction, error) {
	err := d.RenewTransactOpts()
	if err != nil {
		return nil, err
	}
	defer d.RenewTransactOpts()
	return d.blockHashStoreInstance.Store(d.transactOpts, big.NewInt(0).SetUint64(blockNum))
}

func (d *VORCoordinatorCaller) QueryWithdrawableTokens() (*big.Int, error) {
	return d.vorCoordinatorInstance.WithdrawableTokens(d.callOpts, common.HexToAddress(d.oracleAddress))
}

func (d *VORCoordinatorCaller) QueryFees(consumer string) (*big.Int, error) {
	keyHash, _ := d.HashOfKey()
	if consumer == "" {
		return d.vorCoordinatorInstance.GetProviderFee(d.callOpts, keyHash)
	} else {
		return d.vorCoordinatorInstance.GetProviderGranularFee(d.callOpts, keyHash, common.HexToAddress(consumer))
	}
}

func (d *VORCoordinatorCaller) GetOracleEthBalance() (*big.Int, error) {
	return d.client.BalanceAt(d.context, common.HexToAddress(d.oracleAddress), nil)
}

func (d *VORCoordinatorCaller) GetTx(txHashStr string) (*types.Transaction, bool, error) {
	txHash := common.HexToHash(txHashStr)
	return d.client.TransactionByHash(context.Background(), txHash)
}

func (d *VORCoordinatorCaller) GetTxReceipt(txHashStr string) (*types.Receipt, error) {
	txHash := common.HexToHash(txHashStr)
	return d.client.TransactionReceipt(context.Background(), txHash)
}

func (d *VORCoordinatorCaller) GetBlockHashFromBlockStore(blockNum uint64) (bool, common.Hash, error) {
	hash, err := d.blockHashStoreInstance.GetBlockhash(d.callOpts, big.NewInt(0).SetUint64(blockNum))
	h := common.BytesToHash(hash[:])

	if err != nil {
		return false, h, err
	}

	return true, h, nil
}
