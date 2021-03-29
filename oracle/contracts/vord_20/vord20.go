// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vord20

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

// VORD20ABI is the input ABI used to generate the binding from.
const VORD20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vorCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xfund\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"DiceLanded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roller\",\"type\":\"address\"}],\"name\":\"DiceRolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"house\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userProvidedSeed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roller\",\"type\":\"address\"}],\"name\":\"rollDice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"setKeyHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawXFUND\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// VORD20 is an auto generated Go binding around an Ethereum contract.
type VORD20 struct {
	VORD20Caller     // Read-only binding to the contract
	VORD20Transactor // Write-only binding to the contract
	VORD20Filterer   // Log filterer for contract events
}

// VORD20Caller is an auto generated read-only Go binding around an Ethereum contract.
type VORD20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORD20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VORD20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORD20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VORD20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORD20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VORD20Session struct {
	Contract     *VORD20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VORD20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VORD20CallerSession struct {
	Contract *VORD20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VORD20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VORD20TransactorSession struct {
	Contract     *VORD20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VORD20Raw is an auto generated low-level Go binding around an Ethereum contract.
type VORD20Raw struct {
	Contract *VORD20 // Generic contract binding to access the raw methods on
}

// VORD20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VORD20CallerRaw struct {
	Contract *VORD20Caller // Generic read-only contract binding to access the raw methods on
}

// VORD20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VORD20TransactorRaw struct {
	Contract *VORD20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVORD20 creates a new instance of VORD20, bound to a specific deployed contract.
func NewVORD20(address common.Address, backend bind.ContractBackend) (*VORD20, error) {
	contract, err := bindVORD20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VORD20{VORD20Caller: VORD20Caller{contract: contract}, VORD20Transactor: VORD20Transactor{contract: contract}, VORD20Filterer: VORD20Filterer{contract: contract}}, nil
}

// NewVORD20Caller creates a new read-only instance of VORD20, bound to a specific deployed contract.
func NewVORD20Caller(address common.Address, caller bind.ContractCaller) (*VORD20Caller, error) {
	contract, err := bindVORD20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VORD20Caller{contract: contract}, nil
}

// NewVORD20Transactor creates a new write-only instance of VORD20, bound to a specific deployed contract.
func NewVORD20Transactor(address common.Address, transactor bind.ContractTransactor) (*VORD20Transactor, error) {
	contract, err := bindVORD20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VORD20Transactor{contract: contract}, nil
}

// NewVORD20Filterer creates a new log filterer instance of VORD20, bound to a specific deployed contract.
func NewVORD20Filterer(address common.Address, filterer bind.ContractFilterer) (*VORD20Filterer, error) {
	contract, err := bindVORD20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VORD20Filterer{contract: contract}, nil
}

// bindVORD20 binds a generic wrapper to an already deployed contract.
func bindVORD20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VORD20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VORD20 *VORD20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VORD20.Contract.VORD20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VORD20 *VORD20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORD20.Contract.VORD20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VORD20 *VORD20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VORD20.Contract.VORD20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VORD20 *VORD20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VORD20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VORD20 *VORD20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORD20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VORD20 *VORD20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VORD20.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VORD20 *VORD20Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORD20.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VORD20 *VORD20Session) Fee() (*big.Int, error) {
	return _VORD20.Contract.Fee(&_VORD20.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VORD20 *VORD20CallerSession) Fee() (*big.Int, error) {
	return _VORD20.Contract.Fee(&_VORD20.CallOpts)
}

// House is a free data retrieval call binding the contract method 0xb1cad5e3.
//
// Solidity: function house(address player) view returns(string)
func (_VORD20 *VORD20Caller) House(opts *bind.CallOpts, player common.Address) (string, error) {
	var out []interface{}
	err := _VORD20.contract.Call(opts, &out, "house", player)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// House is a free data retrieval call binding the contract method 0xb1cad5e3.
//
// Solidity: function house(address player) view returns(string)
func (_VORD20 *VORD20Session) House(player common.Address) (string, error) {
	return _VORD20.Contract.House(&_VORD20.CallOpts, player)
}

// House is a free data retrieval call binding the contract method 0xb1cad5e3.
//
// Solidity: function house(address player) view returns(string)
func (_VORD20 *VORD20CallerSession) House(player common.Address) (string, error) {
	return _VORD20.Contract.House(&_VORD20.CallOpts, player)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_VORD20 *VORD20Caller) KeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VORD20.contract.Call(opts, &out, "keyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_VORD20 *VORD20Session) KeyHash() ([32]byte, error) {
	return _VORD20.Contract.KeyHash(&_VORD20.CallOpts)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_VORD20 *VORD20CallerSession) KeyHash() ([32]byte, error) {
	return _VORD20.Contract.KeyHash(&_VORD20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORD20 *VORD20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VORD20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORD20 *VORD20Session) Owner() (common.Address, error) {
	return _VORD20.Contract.Owner(&_VORD20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORD20 *VORD20CallerSession) Owner() (common.Address, error) {
	return _VORD20.Contract.Owner(&_VORD20.CallOpts)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_VORD20 *VORD20Transactor) RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "rawFulfillRandomness", requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_VORD20 *VORD20Session) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.RawFulfillRandomness(&_VORD20.TransactOpts, requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_VORD20 *VORD20TransactorSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.RawFulfillRandomness(&_VORD20.TransactOpts, requestId, randomness)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORD20 *VORD20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORD20 *VORD20Session) RenounceOwnership() (*types.Transaction, error) {
	return _VORD20.Contract.RenounceOwnership(&_VORD20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORD20 *VORD20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VORD20.Contract.RenounceOwnership(&_VORD20.TransactOpts)
}

// RollDice is a paid mutator transaction binding the contract method 0x5e816740.
//
// Solidity: function rollDice(uint256 userProvidedSeed, address roller) returns(bytes32 requestId)
func (_VORD20 *VORD20Transactor) RollDice(opts *bind.TransactOpts, userProvidedSeed *big.Int, roller common.Address) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "rollDice", userProvidedSeed, roller)
}

// RollDice is a paid mutator transaction binding the contract method 0x5e816740.
//
// Solidity: function rollDice(uint256 userProvidedSeed, address roller) returns(bytes32 requestId)
func (_VORD20 *VORD20Session) RollDice(userProvidedSeed *big.Int, roller common.Address) (*types.Transaction, error) {
	return _VORD20.Contract.RollDice(&_VORD20.TransactOpts, userProvidedSeed, roller)
}

// RollDice is a paid mutator transaction binding the contract method 0x5e816740.
//
// Solidity: function rollDice(uint256 userProvidedSeed, address roller) returns(bytes32 requestId)
func (_VORD20 *VORD20TransactorSession) RollDice(userProvidedSeed *big.Int, roller common.Address) (*types.Transaction, error) {
	return _VORD20.Contract.RollDice(&_VORD20.TransactOpts, userProvidedSeed, roller)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_VORD20 *VORD20Transactor) SetFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "setFee", fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_VORD20 *VORD20Session) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.SetFee(&_VORD20.TransactOpts, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_VORD20 *VORD20TransactorSession) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.SetFee(&_VORD20.TransactOpts, fee)
}

// SetKeyHash is a paid mutator transaction binding the contract method 0x98544710.
//
// Solidity: function setKeyHash(bytes32 keyHash) returns()
func (_VORD20 *VORD20Transactor) SetKeyHash(opts *bind.TransactOpts, keyHash [32]byte) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "setKeyHash", keyHash)
}

// SetKeyHash is a paid mutator transaction binding the contract method 0x98544710.
//
// Solidity: function setKeyHash(bytes32 keyHash) returns()
func (_VORD20 *VORD20Session) SetKeyHash(keyHash [32]byte) (*types.Transaction, error) {
	return _VORD20.Contract.SetKeyHash(&_VORD20.TransactOpts, keyHash)
}

// SetKeyHash is a paid mutator transaction binding the contract method 0x98544710.
//
// Solidity: function setKeyHash(bytes32 keyHash) returns()
func (_VORD20 *VORD20TransactorSession) SetKeyHash(keyHash [32]byte) (*types.Transaction, error) {
	return _VORD20.Contract.SetKeyHash(&_VORD20.TransactOpts, keyHash)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_VORD20 *VORD20Transactor) TopUpGas(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "topUpGas", _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_VORD20 *VORD20Session) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.TopUpGas(&_VORD20.TransactOpts, _amount)
}

// TopUpGas is a paid mutator transaction binding the contract method 0xe61c51ca.
//
// Solidity: function topUpGas(uint256 _amount) returns()
func (_VORD20 *VORD20TransactorSession) TopUpGas(_amount *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.TopUpGas(&_VORD20.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORD20 *VORD20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORD20 *VORD20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VORD20.Contract.TransferOwnership(&_VORD20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORD20 *VORD20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VORD20.Contract.TransferOwnership(&_VORD20.TransactOpts, newOwner)
}

// WithdrawXFUND is a paid mutator transaction binding the contract method 0x2916ddb1.
//
// Solidity: function withdrawXFUND(address to, uint256 value) returns()
func (_VORD20 *VORD20Transactor) WithdrawXFUND(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VORD20.contract.Transact(opts, "withdrawXFUND", to, value)
}

// WithdrawXFUND is a paid mutator transaction binding the contract method 0x2916ddb1.
//
// Solidity: function withdrawXFUND(address to, uint256 value) returns()
func (_VORD20 *VORD20Session) WithdrawXFUND(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.WithdrawXFUND(&_VORD20.TransactOpts, to, value)
}

// WithdrawXFUND is a paid mutator transaction binding the contract method 0x2916ddb1.
//
// Solidity: function withdrawXFUND(address to, uint256 value) returns()
func (_VORD20 *VORD20TransactorSession) WithdrawXFUND(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VORD20.Contract.WithdrawXFUND(&_VORD20.TransactOpts, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VORD20 *VORD20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORD20.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VORD20 *VORD20Session) Receive() (*types.Transaction, error) {
	return _VORD20.Contract.Receive(&_VORD20.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VORD20 *VORD20TransactorSession) Receive() (*types.Transaction, error) {
	return _VORD20.Contract.Receive(&_VORD20.TransactOpts)
}

// VORD20DiceLandedIterator is returned from FilterDiceLanded and is used to iterate over the raw logs and unpacked data for DiceLanded events raised by the VORD20 contract.
type VORD20DiceLandedIterator struct {
	Event *VORD20DiceLanded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VORD20DiceLandedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORD20DiceLanded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VORD20DiceLanded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VORD20DiceLandedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORD20DiceLandedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORD20DiceLanded represents a DiceLanded event raised by the VORD20 contract.
type VORD20DiceLanded struct {
	RequestId [32]byte
	Result    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDiceLanded is a free log retrieval operation binding the contract event 0x909dd726737b7ffa4ae9d137e9edebe8a74a9c2910a4b60e8112f93ab2170837.
//
// Solidity: event DiceLanded(bytes32 indexed requestId, uint256 indexed result)
func (_VORD20 *VORD20Filterer) FilterDiceLanded(opts *bind.FilterOpts, requestId [][32]byte, result []*big.Int) (*VORD20DiceLandedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var resultRule []interface{}
	for _, resultItem := range result {
		resultRule = append(resultRule, resultItem)
	}

	logs, sub, err := _VORD20.contract.FilterLogs(opts, "DiceLanded", requestIdRule, resultRule)
	if err != nil {
		return nil, err
	}
	return &VORD20DiceLandedIterator{contract: _VORD20.contract, event: "DiceLanded", logs: logs, sub: sub}, nil
}

// WatchDiceLanded is a free log subscription operation binding the contract event 0x909dd726737b7ffa4ae9d137e9edebe8a74a9c2910a4b60e8112f93ab2170837.
//
// Solidity: event DiceLanded(bytes32 indexed requestId, uint256 indexed result)
func (_VORD20 *VORD20Filterer) WatchDiceLanded(opts *bind.WatchOpts, sink chan<- *VORD20DiceLanded, requestId [][32]byte, result []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var resultRule []interface{}
	for _, resultItem := range result {
		resultRule = append(resultRule, resultItem)
	}

	logs, sub, err := _VORD20.contract.WatchLogs(opts, "DiceLanded", requestIdRule, resultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORD20DiceLanded)
				if err := _VORD20.contract.UnpackLog(event, "DiceLanded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDiceLanded is a log parse operation binding the contract event 0x909dd726737b7ffa4ae9d137e9edebe8a74a9c2910a4b60e8112f93ab2170837.
//
// Solidity: event DiceLanded(bytes32 indexed requestId, uint256 indexed result)
func (_VORD20 *VORD20Filterer) ParseDiceLanded(log types.Log) (*VORD20DiceLanded, error) {
	event := new(VORD20DiceLanded)
	if err := _VORD20.contract.UnpackLog(event, "DiceLanded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORD20DiceRolledIterator is returned from FilterDiceRolled and is used to iterate over the raw logs and unpacked data for DiceRolled events raised by the VORD20 contract.
type VORD20DiceRolledIterator struct {
	Event *VORD20DiceRolled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VORD20DiceRolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORD20DiceRolled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VORD20DiceRolled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VORD20DiceRolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORD20DiceRolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORD20DiceRolled represents a DiceRolled event raised by the VORD20 contract.
type VORD20DiceRolled struct {
	RequestId [32]byte
	Roller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDiceRolled is a free log retrieval operation binding the contract event 0x923de4fc4aece24a78a9e4ca3009c571a742f81ac2c004a229224b0fd1883bdd.
//
// Solidity: event DiceRolled(bytes32 indexed requestId, address indexed roller)
func (_VORD20 *VORD20Filterer) FilterDiceRolled(opts *bind.FilterOpts, requestId [][32]byte, roller []common.Address) (*VORD20DiceRolledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var rollerRule []interface{}
	for _, rollerItem := range roller {
		rollerRule = append(rollerRule, rollerItem)
	}

	logs, sub, err := _VORD20.contract.FilterLogs(opts, "DiceRolled", requestIdRule, rollerRule)
	if err != nil {
		return nil, err
	}
	return &VORD20DiceRolledIterator{contract: _VORD20.contract, event: "DiceRolled", logs: logs, sub: sub}, nil
}

// WatchDiceRolled is a free log subscription operation binding the contract event 0x923de4fc4aece24a78a9e4ca3009c571a742f81ac2c004a229224b0fd1883bdd.
//
// Solidity: event DiceRolled(bytes32 indexed requestId, address indexed roller)
func (_VORD20 *VORD20Filterer) WatchDiceRolled(opts *bind.WatchOpts, sink chan<- *VORD20DiceRolled, requestId [][32]byte, roller []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var rollerRule []interface{}
	for _, rollerItem := range roller {
		rollerRule = append(rollerRule, rollerItem)
	}

	logs, sub, err := _VORD20.contract.WatchLogs(opts, "DiceRolled", requestIdRule, rollerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORD20DiceRolled)
				if err := _VORD20.contract.UnpackLog(event, "DiceRolled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDiceRolled is a log parse operation binding the contract event 0x923de4fc4aece24a78a9e4ca3009c571a742f81ac2c004a229224b0fd1883bdd.
//
// Solidity: event DiceRolled(bytes32 indexed requestId, address indexed roller)
func (_VORD20 *VORD20Filterer) ParseDiceRolled(log types.Log) (*VORD20DiceRolled, error) {
	event := new(VORD20DiceRolled)
	if err := _VORD20.contract.UnpackLog(event, "DiceRolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORD20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VORD20 contract.
type VORD20OwnershipTransferredIterator struct {
	Event *VORD20OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VORD20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORD20OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VORD20OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VORD20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORD20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORD20OwnershipTransferred represents a OwnershipTransferred event raised by the VORD20 contract.
type VORD20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORD20 *VORD20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VORD20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VORD20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VORD20OwnershipTransferredIterator{contract: _VORD20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORD20 *VORD20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VORD20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VORD20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORD20OwnershipTransferred)
				if err := _VORD20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORD20 *VORD20Filterer) ParseOwnershipTransferred(log types.Log) (*VORD20OwnershipTransferred, error) {
	event := new(VORD20OwnershipTransferred)
	if err := _VORD20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
