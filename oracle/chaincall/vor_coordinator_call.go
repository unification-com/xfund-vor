package chaincall

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"oracle/config"
	"oracle/contracts/block_hash_store"
	"oracle/contracts/vor_coordinator"
)

// registrationTimeout bounds the wait for the proving key's registration transaction to be mined.
const registrationTimeout = 5 * time.Minute

type VORCoordinatorCaller struct {
	*transactor
	vorCoordinatorInstance *vor_coordinator.VorCoordinator
	blockHashStoreInstance *block_hash_store.BlockHashStore
	publicProvingKey       [2]*big.Int
}

func NewVORCoordinatorCaller(chain config.ChainConfig, oraclePrivateKey []byte) (*VORCoordinatorCaller, error) {
	t, err := connect(chain, oraclePrivateKey)
	if err != nil {
		return nil, err
	}
	vorCoordinatorInstance, err := vor_coordinator.NewVorCoordinator(common.HexToAddress(chain.VORCoordinatorContractAddress), t.client)
	if err != nil {
		return nil, err
	}
	blockHashStoreInstance, err := block_hash_store.NewBlockHashStore(common.HexToAddress(chain.BlockHashStoreContractAddress), t.client)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorCaller{
		transactor:             t,
		vorCoordinatorInstance: vorCoordinatorInstance,
		blockHashStoreInstance: blockHashStoreInstance,
		publicProvingKey:       [2]*big.Int{t.key.PublicKey.X, t.key.PublicKey.Y},
	}, nil
}

func (d *VORCoordinatorCaller) HashOfKey() ([32]byte, error) {
	return d.vorCoordinatorInstance.HashOfKey(d.callOpts(), d.publicProvingKey)
}

// ProvingKeyProvider returns the provider address the coordinator holds for the oracle's proving key: the
// zero address while the key is unregistered.
func (d *VORCoordinatorCaller) ProvingKeyProvider() (common.Address, error) {
	keyHash, err := d.HashOfKey()
	if err != nil {
		return common.Address{}, err
	}
	return d.vorCoordinatorInstance.GetProviderAddress(d.callOpts(), keyHash)
}

// OracleAddress is the address the oracle sends transactions from.
func (d *VORCoordinatorCaller) OracleAddress() common.Address {
	return d.address
}

func (d *VORCoordinatorCaller) Withdraw(recipientAddress string, amount *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.vorCoordinatorInstance.Withdraw(opts, common.HexToAddress(recipientAddress), amount)
	})
}

func (d *VORCoordinatorCaller) RegisterProvingKey(fee *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.vorCoordinatorInstance.RegisterProvingKey(opts, fee, d.address, d.publicProvingKey)
	})
}

// RegisterProvingKeyAndWait registers the oracle's proving key at fee and waits for the transaction to be
// mined, failing unless it succeeded.
func (d *VORCoordinatorCaller) RegisterProvingKeyAndWait(fee *big.Int) error {
	tx, err := d.RegisterProvingKey(fee)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(d.ctx, registrationTimeout)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, d.client, tx)
	if err != nil {
		return fmt.Errorf("waiting for the registration transaction %s: %w", tx.Hash().Hex(), err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("the registration transaction %s failed", tx.Hash().Hex())
	}
	return nil
}

func (d *VORCoordinatorCaller) RandomnessRequest(keyHash [32]byte, consumerSeed *big.Int, feePaid *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.vorCoordinatorInstance.RandomnessRequest(opts, keyHash, consumerSeed, feePaid)
	})
}

func (d *VORCoordinatorCaller) ChangeFee(fee *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.vorCoordinatorInstance.ChangeFee(opts, d.publicProvingKey, fee)
	})
}

func (d *VORCoordinatorCaller) ChangeGranularFee(_consumer common.Address, fee *big.Int) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.vorCoordinatorInstance.ChangeGranularFee(opts, d.publicProvingKey, fee, _consumer)
	})
}

func (d *VORCoordinatorCaller) FulfillRandomnessRequest(proof []byte) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.vorCoordinatorInstance.FulfillRandomnessRequest(opts, proof)
	})
}

func (d *VORCoordinatorCaller) StoreBlockHash(blockNum uint64) (*types.Transaction, error) {
	return d.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return d.blockHashStoreInstance.Store(opts, new(big.Int).SetUint64(blockNum))
	})
}

func (d *VORCoordinatorCaller) QueryWithdrawableTokens() (*big.Int, error) {
	return d.vorCoordinatorInstance.WithdrawableTokens(d.callOpts(), d.address)
}

func (d *VORCoordinatorCaller) QueryFees(consumer string) (*big.Int, error) {
	keyHash, err := d.HashOfKey()
	if err != nil {
		return nil, err
	}
	if consumer == "" {
		return d.vorCoordinatorInstance.GetProviderFee(d.callOpts(), keyHash)
	}
	return d.vorCoordinatorInstance.GetProviderGranularFee(d.callOpts(), keyHash, common.HexToAddress(consumer))
}

func (d *VORCoordinatorCaller) GetOracleEthBalance() (*big.Int, error) {
	return d.client.BalanceAt(d.ctx, d.address, nil)
}

func (d *VORCoordinatorCaller) GetTx(txHashStr string) (*types.Transaction, bool, error) {
	return d.client.TransactionByHash(d.ctx, common.HexToHash(txHashStr))
}

func (d *VORCoordinatorCaller) GetTxReceipt(txHashStr string) (*types.Receipt, error) {
	return d.client.TransactionReceipt(d.ctx, common.HexToHash(txHashStr))
}

func (d *VORCoordinatorCaller) GetBlockHashFromBlockStore(blockNum uint64) (bool, common.Hash, error) {
	hash, err := d.blockHashStoreInstance.GetBlockhash(d.callOpts(), new(big.Int).SetUint64(blockNum))
	if err != nil {
		return false, common.Hash{}, err
	}
	return true, common.BytesToHash(hash[:]), nil
}
