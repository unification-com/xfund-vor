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
	"log"
	"math/big"
	"oracle/contracts/vord_20"
	"oracle/utils/walletworker"
)

type VORD20Caller struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vord_20.VORD20
	transactOpts    *bind.TransactOpts
	callOpts        *bind.CallOpts

	oraclePrivateKey string
	oraclePublicKey  string
	oracleAddress    string
}

func NewVORD20Caller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*VORD20Caller, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	fmt.Println("contractStringAddress: ", contractStringAddress)
	contractAddress := common.HexToAddress(contractStringAddress)
	instance, err := vord_20.NewVORD20(contractAddress, client)
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
	if err != nil || ECDSAoraclePublicKey == nil {
		log.Print(err)
		log.Print(ECDSAoraclePublicKey)
		return nil, err
	}

	oracleAddressObj, oracleAddress := walletworker.GenerateAddress(ECDSAoraclePublicKey)
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
	//transactOpts.Value = big.NewInt(1000000000000000000)
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = uint64(100000) // in units
	transactOpts.Context = context.Background()
	transactOpts.From = oracleAddressObj

	return &VORD20Caller{
		client:           client,
		contractAddress:  contractAddress,
		instance:         instance,
		transactOpts:     transactOpts,
		callOpts:         &bind.CallOpts{From: oracleAddressObj},
		oraclePrivateKey: string(oraclePrivateKey),
		oraclePublicKey:  hexutil.Encode(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey))),
		oracleAddress:    oracleAddress,
	}, err
}

func (d *VORD20Caller) RenewTransactOpts() (err error) {
	fmt.Println("RenewTransactOpts")
	gasPrice, err := d.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	nonce, err := d.client.PendingNonceAt(context.Background(), common.HexToAddress(d.oracleAddress))
	if err != nil {
		return
	}
	d.transactOpts.Nonce = big.NewInt(int64(nonce))
	d.transactOpts.GasPrice = gasPrice
	d.transactOpts.GasLimit = uint64(100000) // in units

	return
}

func (d *VORD20Caller) RollDice(seed *big.Int) (*types.Transaction, error) {
	defer d.RenewTransactOpts()
	return d.instance.RollDice(d.transactOpts, seed, common.HexToAddress(d.oracleAddress))
}

func (d VORD20Caller) TopUpGas(amount *big.Int) (*types.Transaction, error) {
	defer d.RenewTransactOpts()
	d.transactOpts.Value = amount
	return d.instance.TopUpGas(d.transactOpts, amount)
}

func (d *VORD20Caller) SetFee(fee *big.Int) (*types.Transaction, error) {
	defer d.RenewTransactOpts()
	return d.instance.SetFee(d.transactOpts, fee)
}

func (d *VORD20Caller) Owner() (common.Address, error) {
	defer d.RenewTransactOpts()
	return d.instance.Owner(d.callOpts)
}

func (d *VORD20Caller) SetKeyHash(keyHash [32]byte) (*types.Transaction, error) {
	defer d.RenewTransactOpts()
	return d.instance.SetKeyHash(d.transactOpts, keyHash)
}

func (d *VORD20Caller) KeyHash() ([32]byte, error) {
	defer d.RenewTransactOpts()
	return d.instance.KeyHash(d.callOpts)
}

func (d *VORD20Caller) Fee() (*big.Int, error) {
	defer d.RenewTransactOpts()
	return d.instance.Fee(d.callOpts)
}

func (d *VORD20Caller) Transfer() {
	defer d.RenewTransactOpts()
}
