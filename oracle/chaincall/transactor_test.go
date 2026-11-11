package chaincall

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"oracle/config"
)

// fakeNode is a JSON-RPC endpoint with just enough of eth_ to send transactions: its pending nonce rises with
// every transaction it accepts, and it answers the nonce slowly, so unserialised senders would read the same
// one.
type fakeNode struct {
	mu      sync.Mutex
	nonce   uint64
	nonces  []uint64
	baseFee *big.Int // nil: a pre-London chain
}

func (n *fakeNode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     json.RawMessage   `json:"id"`
		Method string            `json:"method"`
		Params []json.RawMessage `json:"params"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var result interface{}
	switch req.Method {
	case "eth_getTransactionCount":
		time.Sleep(5 * time.Millisecond)
		n.mu.Lock()
		result = hexutil.Uint64(n.nonce)
		n.mu.Unlock()
	case "eth_gasPrice":
		result = (*hexutil.Big)(gwei(1))
	case "eth_maxPriorityFeePerGas":
		result = (*hexutil.Big)(gwei(2))
	case "eth_getBlockByNumber":
		result = &types.Header{Number: big.NewInt(1), Difficulty: big.NewInt(0), BaseFee: n.baseFee}
	case "eth_sendRawTransaction":
		var raw hexutil.Bytes
		_ = json.Unmarshal(req.Params[0], &raw)
		tx := new(types.Transaction)
		if err := tx.UnmarshalBinary(raw); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		n.mu.Lock()
		n.nonces = append(n.nonces, tx.Nonce())
		if tx.Nonce() >= n.nonce {
			n.nonce = tx.Nonce() + 1
		}
		n.mu.Unlock()
		result = tx.Hash()
	default:
		http.Error(w, "unexpected method "+req.Method, http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": result})
}

// The poller and the HTTP API send from the same account at once: each send must get its own nonce.
func TestTransactorSendsSerialise(t *testing.T) {
	node := &fakeNode{}
	server := httptest.NewServer(node)
	defer server.Close()

	tr, err := connect(server.URL, []byte("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"), big.NewInt(1337))
	require.NoError(t, err)

	const senders = 10
	var wg sync.WaitGroup
	errs := make(chan error, senders)
	for range senders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := tr.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
				if opts.GasLimit != 0 {
					return nil, fmt.Errorf("gas limit %d: it must be estimated", opts.GasLimit)
				}
				tx, err := opts.Signer(opts.From, types.NewTransaction(opts.Nonce.Uint64(), common.Address{1}, big.NewInt(0), 21000, opts.GasPrice, nil))
				if err != nil {
					return nil, err
				}
				return tx, tr.client.SendTransaction(opts.Context, tx)
			})
			errs <- err
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		require.NoError(t, err)
	}

	sort.Slice(node.nonces, func(i, j int) bool { return node.nonces[i] < node.nonces[j] })
	want := make([]uint64, senders)
	for i := range want {
		want[i] = uint64(i)
	}
	assert.Equal(t, want, node.nonces, "every transaction needs its own nonce")
}

func gwei(n int64) *big.Int { return new(big.Int).Mul(big.NewInt(n), big.NewInt(params.GWei)) }

// Where the chain has a base fee, transactions are EIP-1559: tip as suggested, fee cap twice the base fee plus
// the tip, both within max_gas_price; elsewhere, or with legacy_gas, a legacy gas price.
func TestTransactorFees(t *testing.T) {
	maxGasPrice, legacyGas := config.Conf.MaxGasPrice, config.Conf.LegacyGas
	t.Cleanup(func() { config.Conf.MaxGasPrice, config.Conf.LegacyGas = maxGasPrice, legacyGas })

	cases := []struct {
		name            string
		baseFee         *big.Int
		maxGwei         int64
		legacy          bool
		price, tip, cap *big.Int
	}{
		{name: "london", baseFee: gwei(10), maxGwei: 150, tip: gwei(2), cap: gwei(22)},
		{name: "fee cap capped", baseFee: gwei(10), maxGwei: 15, tip: gwei(2), cap: gwei(15)},
		{name: "tip within the cap", baseFee: gwei(10), maxGwei: 1, tip: gwei(1), cap: gwei(1)},
		{name: "pre-london", maxGwei: 150, price: gwei(1)},
		{name: "legacy_gas", baseFee: gwei(10), maxGwei: 150, legacy: true, price: gwei(1)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			config.Conf.MaxGasPrice, config.Conf.LegacyGas = c.maxGwei, c.legacy
			server := httptest.NewServer(&fakeNode{baseFee: c.baseFee})
			defer server.Close()
			tr, err := connect(server.URL, []byte("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"), big.NewInt(1337))
			require.NoError(t, err)

			_, err = tr.send(func(opts *bind.TransactOpts) (*types.Transaction, error) {
				assert.Equal(t, c.price, opts.GasPrice, "gas price")
				assert.Equal(t, c.tip, opts.GasTipCap, "tip")
				assert.Equal(t, c.cap, opts.GasFeeCap, "fee cap")
				return nil, nil
			})
			require.NoError(t, err)
		})
	}
}

func TestCapToMaxGwei(t *testing.T) {
	assert.Equal(t, gwei(150), capToMaxGwei(gwei(200), 150))
	assert.Equal(t, gwei(100), capToMaxGwei(gwei(100), 150))
	assert.Equal(t, gwei(200), capToMaxGwei(gwei(200), 0), "0 means uncapped")
}

func TestConnectRejectsABadKey(t *testing.T) {
	_, err := connect("http://127.0.0.1:1", []byte("0xnotakey"), big.NewInt(1))
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "notakey", "errors must not echo the key")
}
