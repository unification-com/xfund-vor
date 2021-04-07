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
	"oracle/config"
	"oracle/contracts/vor_randomness_request_mock"
	"oracle/utils/walletworker"
)

type VORRandomnessRequestMockCaller struct {
	contractAddress common.Address
	client          *ethclient.Client
	instance        *vor_randomness_request_mock.VorRandomnessRequestMock
	transactOpts    *bind.TransactOpts
	callOpts        *bind.CallOpts

	publicProvingKey [2]*big.Int
	oraclePrivateKey string
	oraclePublicKey  string
	oracleAddress    string
}

func NewVORRandomnessRequestMockCaller(contractStringAddress string, ethHostAddress string, chainID *big.Int, oraclePrivateKey []byte) (*VORRandomnessRequestMockCaller, error) {
	client, err := ethclient.Dial(ethHostAddress)
	if err != nil {
		return nil, err
	}
	fmt.Println("contractStringAddress: ", contractStringAddress)
	contractAddress := common.HexToAddress(contractStringAddress)
	instance, err := vor_randomness_request_mock.NewVorRandomnessRequestMock(contractAddress, client)
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

	_, oracleAddress := walletworker.GenerateAddress(ECDSAoraclePublicKey)
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
	transactOpts.GasLimit = uint64(config.Conf.LimitGasPrice) // in units
	transactOpts.Context = context.Background()

	return &VORRandomnessRequestMockCaller{
		client:           client,
		contractAddress:  contractAddress,
		instance:         instance,
		transactOpts:     transactOpts,
		callOpts:         &bind.CallOpts{},
		publicProvingKey: [2]*big.Int{ECDSAoraclePublicKey.X, ECDSAoraclePublicKey.Y},
		oraclePrivateKey: string(oraclePrivateKey),
		oraclePublicKey:  hexutil.Encode(crypto.FromECDSAPub(oraclePublicKey.(*ecdsa.PublicKey))),
		oracleAddress:    oracleAddress,
	}, err
}

func (d *VORRandomnessRequestMockCaller) RandomnessRequest(keyHash [32]byte, consumerSeed *big.Int, feePaid *big.Int) (*types.Transaction, error) {
	return d.instance.RandomnessRequest(d.transactOpts, keyHash, consumerSeed, feePaid)
}
