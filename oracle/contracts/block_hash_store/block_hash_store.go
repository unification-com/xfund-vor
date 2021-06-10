// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package block_hash_store

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlockHashStoreABI is the input ABI used to generate the binding from.
const BlockHashStoreABI = "[{\"inputs\":[],\"name\":\"storeEarliest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getBlockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"storeVerifyHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BlockHashStore is an auto generated Go binding around an Ethereum contract.
type BlockHashStore struct {
	BlockHashStoreCaller     // Read-only binding to the contract
	BlockHashStoreTransactor // Write-only binding to the contract
	BlockHashStoreFilterer   // Log filterer for contract events
}

// BlockHashStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockHashStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockHashStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockHashStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockHashStoreSession struct {
	Contract     *BlockHashStore   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockHashStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockHashStoreCallerSession struct {
	Contract *BlockHashStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BlockHashStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockHashStoreTransactorSession struct {
	Contract     *BlockHashStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BlockHashStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockHashStoreRaw struct {
	Contract *BlockHashStore // Generic contract binding to access the raw methods on
}

// BlockHashStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockHashStoreCallerRaw struct {
	Contract *BlockHashStoreCaller // Generic read-only contract binding to access the raw methods on
}

// BlockHashStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockHashStoreTransactorRaw struct {
	Contract *BlockHashStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockHashStore creates a new instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStore(address common.Address, backend bind.ContractBackend) (*BlockHashStore, error) {
	contract, err := bindBlockHashStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockHashStore{BlockHashStoreCaller: BlockHashStoreCaller{contract: contract}, BlockHashStoreTransactor: BlockHashStoreTransactor{contract: contract}, BlockHashStoreFilterer: BlockHashStoreFilterer{contract: contract}}, nil
}

// NewBlockHashStoreCaller creates a new read-only instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStoreCaller(address common.Address, caller bind.ContractCaller) (*BlockHashStoreCaller, error) {
	contract, err := bindBlockHashStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashStoreCaller{contract: contract}, nil
}

// NewBlockHashStoreTransactor creates a new write-only instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockHashStoreTransactor, error) {
	contract, err := bindBlockHashStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashStoreTransactor{contract: contract}, nil
}

// NewBlockHashStoreFilterer creates a new log filterer instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockHashStoreFilterer, error) {
	contract, err := bindBlockHashStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockHashStoreFilterer{contract: contract}, nil
}

// bindBlockHashStore binds a generic wrapper to an already deployed contract.
func bindBlockHashStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockHashStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashStore *BlockHashStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashStore.Contract.BlockHashStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashStore *BlockHashStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashStore.Contract.BlockHashStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashStore *BlockHashStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashStore.Contract.BlockHashStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashStore *BlockHashStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashStore *BlockHashStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashStore *BlockHashStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashStore.Contract.contract.Transact(opts, method, params...)
}

// GetBlockhash is a free data retrieval call binding the contract method 0xe9413d38.
//
// Solidity: function getBlockhash(uint256 n) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreCaller) GetBlockhash(opts *bind.CallOpts, n *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashStore.contract.Call(opts, &out, "getBlockhash", n)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockhash is a free data retrieval call binding the contract method 0xe9413d38.
//
// Solidity: function getBlockhash(uint256 n) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreSession) GetBlockhash(n *big.Int) ([32]byte, error) {
	return _BlockHashStore.Contract.GetBlockhash(&_BlockHashStore.CallOpts, n)
}

// GetBlockhash is a free data retrieval call binding the contract method 0xe9413d38.
//
// Solidity: function getBlockhash(uint256 n) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreCallerSession) GetBlockhash(n *big.Int) ([32]byte, error) {
	return _BlockHashStore.Contract.GetBlockhash(&_BlockHashStore.CallOpts, n)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 n) returns()
func (_BlockHashStore *BlockHashStoreTransactor) Store(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _BlockHashStore.contract.Transact(opts, "store", n)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 n) returns()
func (_BlockHashStore *BlockHashStoreSession) Store(n *big.Int) (*types.Transaction, error) {
	return _BlockHashStore.Contract.Store(&_BlockHashStore.TransactOpts, n)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 n) returns()
func (_BlockHashStore *BlockHashStoreTransactorSession) Store(n *big.Int) (*types.Transaction, error) {
	return _BlockHashStore.Contract.Store(&_BlockHashStore.TransactOpts, n)
}

// StoreEarliest is a paid mutator transaction binding the contract method 0x83b6d6b7.
//
// Solidity: function storeEarliest() returns()
func (_BlockHashStore *BlockHashStoreTransactor) StoreEarliest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashStore.contract.Transact(opts, "storeEarliest")
}

// StoreEarliest is a paid mutator transaction binding the contract method 0x83b6d6b7.
//
// Solidity: function storeEarliest() returns()
func (_BlockHashStore *BlockHashStoreSession) StoreEarliest() (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreEarliest(&_BlockHashStore.TransactOpts)
}

// StoreEarliest is a paid mutator transaction binding the contract method 0x83b6d6b7.
//
// Solidity: function storeEarliest() returns()
func (_BlockHashStore *BlockHashStoreTransactorSession) StoreEarliest() (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreEarliest(&_BlockHashStore.TransactOpts)
}

// StoreVerifyHeader is a paid mutator transaction binding the contract method 0xfadff0e1.
//
// Solidity: function storeVerifyHeader(uint256 n, bytes header) returns()
func (_BlockHashStore *BlockHashStoreTransactor) StoreVerifyHeader(opts *bind.TransactOpts, n *big.Int, header []byte) (*types.Transaction, error) {
	return _BlockHashStore.contract.Transact(opts, "storeVerifyHeader", n, header)
}

// StoreVerifyHeader is a paid mutator transaction binding the contract method 0xfadff0e1.
//
// Solidity: function storeVerifyHeader(uint256 n, bytes header) returns()
func (_BlockHashStore *BlockHashStoreSession) StoreVerifyHeader(n *big.Int, header []byte) (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreVerifyHeader(&_BlockHashStore.TransactOpts, n, header)
}

// StoreVerifyHeader is a paid mutator transaction binding the contract method 0xfadff0e1.
//
// Solidity: function storeVerifyHeader(uint256 n, bytes header) returns()
func (_BlockHashStore *BlockHashStoreTransactorSession) StoreVerifyHeader(n *big.Int, header []byte) (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreVerifyHeader(&_BlockHashStore.TransactOpts, n, header)
}
