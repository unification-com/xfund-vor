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
	"oracle/contracts/mock_erc20"
	"oracle/utils/walletworker"
)

type MockERC20Caller struct {
	contractAddress  common.Address
	client           *ethclient.Client
	instance         *mock_erc20.MockERC20
	transactOpts     *bind.TransactOpts
	callOpts         *bind.CallOpts
	chainID          *big.Int
	oraclePrivateKey string
	oraclePublicKey  string
	oracleAddress    string
}

func NewMockERC20Caller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*MockERC20Caller, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(contractStringAddress)
	instance, err := mock_erc20.NewMockERC20(contractAddress, client)
	if err != nil {
		return nil, err
	}
	oraclePrivateKeyECDSA, err := crypto.HexToECDSA(string(oraclePrivateKey[2:]))
	if err != nil {
		return nil, err
	}

	oraclePublicKey := oraclePrivateKeyECDSA.Public()

	ECDSAoraclePublicKey, err := crypto.UnmarshalPubkey(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey)))
	if err != nil || ECDSAoraclePublicKey == nil {
		log.Print(err)
		log.Print(ECDSAoraclePublicKey)
		return nil, err
	}

	oracleAddressObj, oracleAddress := walletworker.GenerateAddress(ECDSAoraclePublicKey)

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
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = uint64(100000) // in units
	//transactOpts.Value = big.NewInt(1000000000000000000)
	//transactOpts.Value.Mul(transactOpts.Value, big.NewInt(10))
	transactOpts.Context = context.Background()
	transactOpts.From = oracleAddressObj

	return &MockERC20Caller{
		client:           client,
		contractAddress:  contractAddress,
		instance:         instance,
		chainID:          chainID,
		transactOpts:     transactOpts,
		callOpts:         &bind.CallOpts{From: oracleAddressObj},
		oraclePrivateKey: string(oraclePrivateKey),
		oraclePublicKey:  hexutil.Encode(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey))),
		oracleAddress:    oracleAddress,
	}, err
}

func (d *MockERC20Caller) RenewTransactOpts() (err error) {
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

func (d *MockERC20Caller) Transfer(recipientAddress string, amount *big.Int) (*types.Transaction, error) {
	defer d.RenewTransactOpts()
	return d.instance.Transfer(d.transactOpts, common.HexToAddress(recipientAddress), amount)
}
