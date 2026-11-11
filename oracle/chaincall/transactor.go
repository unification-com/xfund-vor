package chaincall

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"oracle/config"
	"oracle/utils"
)

// rpcTimeout bounds every JSON-RPC call over HTTP, so a hung endpoint can't stall the poller (or an API
// request) indefinitely.
const rpcTimeout = 30 * time.Second

// Dial connects to an Ethereum JSON-RPC endpoint. Over HTTP(S) every call is bounded by rpcTimeout; other
// transports only bound the connection itself.
func Dial(url string) (*ethclient.Client, error) {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		client, err := rpc.DialHTTPWithClient(url, &http.Client{Timeout: rpcTimeout})
		if err != nil {
			return nil, err
		}
		return ethclient.NewClient(client), nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	return ethclient.DialContext(ctx, url)
}

// transactor sends transactions from one account. Each transaction gets fresh options (the nonce read from
// the chain just before sending, current gas fees, gas estimated) built and sent while holding mu, so
// concurrent senders, the poller and the HTTP API, never share options or race for a nonce.
type transactor struct {
	client  *ethclient.Client
	ctx     context.Context
	key     *ecdsa.PrivateKey
	address common.Address
	chainID *big.Int
	mu      sync.Mutex
}

// connect dials ethHost and returns a transactor for the hex private key (with or without 0x).
func connect(ethHost string, privateKey []byte, chainID *big.Int) (*transactor, error) {
	key, err := crypto.HexToECDSA(utils.RemoveHexPrefix(string(privateKey)))
	if err != nil {
		return nil, errors.New("not a valid private key")
	}
	client, err := Dial(ethHost)
	if err != nil {
		return nil, err
	}
	return &transactor{
		client:  client,
		ctx:     context.Background(),
		key:     key,
		address: crypto.PubkeyToAddress(key.PublicKey),
		chainID: chainID,
	}, nil
}

// send builds fresh options and passes them to tx, which sends one transaction with them.
func (t *transactor) send(tx func(opts *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	opts, err := bind.NewKeyedTransactorWithChainID(t.key, t.chainID)
	if err != nil {
		return nil, err
	}
	nonce, err := t.client.PendingNonceAt(t.ctx, t.address)
	if err != nil {
		return nil, fmt.Errorf("getting the pending nonce: %w", err)
	}
	fees, err := t.suggestFees()
	if err != nil {
		return nil, err
	}
	opts.Context = t.ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	fees.apply(opts)
	// GasLimit 0 makes go-ethereum estimate gas, which fails for a transaction that would revert: a fixed limit
	// skipped that, so a doomed fulfilment (a request paying below the fee) was broadcast, reverted on-chain and
	// was re-sent forever (the 2026-06 fee-revert runaway).
	opts.GasLimit = 0
	return tx(opts)
}

// gasFees prices one transaction: a legacy gas price, or EIP-1559 tip and fee caps, the others nil.
type gasFees struct {
	gasPrice, gasTipCap, gasFeeCap *big.Int
}

func (f gasFees) apply(opts *bind.TransactOpts) {
	opts.GasPrice, opts.GasTipCap, opts.GasFeeCap = f.gasPrice, f.gasTipCap, f.gasFeeCap
}

// feeCapBaseFeeMultiplier: the base fee rises at most 12.5% a block, so twice it lasts ~6 full blocks before
// an EIP-1559 fee cap would bind.
const feeCapBaseFeeMultiplier = 2

// suggestFees prices a transaction with EIP-1559 fees where the chain has a base fee (unless legacy_gas is set),
// else with a legacy gas price; max_gas_price caps either.
func (t *transactor) suggestFees() (gasFees, error) {
	maxGwei := config.Conf.MaxGasPrice
	if !config.Conf.LegacyGas {
		head, err := t.client.HeaderByNumber(t.ctx, nil)
		if err != nil {
			return gasFees{}, fmt.Errorf("getting the head block: %w", err)
		}
		if head.BaseFee != nil {
			tip, err := t.client.SuggestGasTipCap(t.ctx)
			if err != nil {
				return gasFees{}, fmt.Errorf("getting the priority fee: %w", err)
			}
			feeCap := new(big.Int).Mul(head.BaseFee, big.NewInt(feeCapBaseFeeMultiplier))
			return dynamicFees(tip, feeCap.Add(feeCap, tip), maxGwei), nil
		}
	}
	gasPrice, err := t.client.SuggestGasPrice(t.ctx)
	if err != nil {
		return gasFees{}, fmt.Errorf("getting the gas price: %w", err)
	}
	return gasFees{gasPrice: capToMaxGwei(gasPrice, maxGwei)}, nil
}

// dynamicFees caps the fee cap at maxGwei and keeps the tip within it.
func dynamicFees(tip, feeCap *big.Int, maxGwei int64) gasFees {
	feeCap = capToMaxGwei(feeCap, maxGwei)
	if tip.Cmp(feeCap) > 0 {
		tip = new(big.Int).Set(feeCap)
	}
	return gasFees{gasTipCap: tip, gasFeeCap: feeCap}
}

// callOpts are the options for a read-only contract call from this account.
func (t *transactor) callOpts() *bind.CallOpts {
	return &bind.CallOpts{From: t.address, Context: t.ctx}
}

// capToMaxGwei limits gasPrice (wei) to maxGwei; maxGwei <= 0 means uncapped.
func capToMaxGwei(gasPrice *big.Int, maxGwei int64) *big.Int {
	if maxGwei > 0 {
		if maxWei := new(big.Int).Mul(big.NewInt(maxGwei), big.NewInt(params.GWei)); gasPrice.Cmp(maxWei) > 0 {
			return maxWei
		}
	}
	return gasPrice
}
