// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vord_20

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

// Vord20ABI is the input ABI used to generate the binding from.
const Vord20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vorCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xfund\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"DiceLanded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roller\",\"type\":\"address\"}],\"name\":\"DiceRolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseVorCoordinatorAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"topUpGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"withDrawGasTopUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawXFUND\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userProvidedSeed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roller\",\"type\":\"address\"}],\"name\":\"rollDice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"house\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"setKeyHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Vord20 is an auto generated Go binding around an Ethereum contract.
type Vord20 struct {
	Vord20Caller     // Read-only binding to the contract
	Vord20Transactor // Write-only binding to the contract
	Vord20Filterer   // Log filterer for contract events
}

// Vord20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Vord20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vord20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Vord20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vord20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Vord20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Vord20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Vord20Session struct {
	Contract     *Vord20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Vord20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Vord20CallerSession struct {
	Contract *Vord20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Vord20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Vord20TransactorSession struct {
	Contract     *Vord20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Vord20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Vord20Raw struct {
	Contract *Vord20 // Generic contract binding to access the raw methods on
}

// Vord20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Vord20CallerRaw struct {
	Contract *Vord20Caller // Generic read-only contract binding to access the raw methods on
}

// Vord20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Vord20TransactorRaw struct {
	Contract *Vord20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVord20 creates a new instance of Vord20, bound to a specific deployed contract.
func NewVord20(address common.Address, backend bind.ContractBackend) (*Vord20, error) {
	contract, err := bindVord20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vord20{Vord20Caller: Vord20Caller{contract: contract}, Vord20Transactor: Vord20Transactor{contract: contract}, Vord20Filterer: Vord20Filterer{contract: contract}}, nil
}

// NewVord20Caller creates a new read-only instance of Vord20, bound to a specific deployed contract.
func NewVord20Caller(address common.Address, caller bind.ContractCaller) (*Vord20Caller, error) {
	contract, err := bindVord20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Vord20Caller{contract: contract}, nil
}

// NewVord20Transactor creates a new write-only instance of Vord20, bound to a specific deployed contract.
func NewVord20Transactor(address common.Address, transactor bind.ContractTransactor) (*Vord20Transactor, error) {
	contract, err := bindVord20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Vord20Transactor{contract: contract}, nil
}

// NewVord20Filterer creates a new log filterer instance of Vord20, bound to a specific deployed contract.
func NewVord20Filterer(address common.Address, filterer bind.ContractFilterer) (*Vord20Filterer, error) {
	contract, err := bindVord20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Vord20Filterer{contract: contract}, nil
}

// bindVord20 binds a generic wrapper to an already deployed contract.
func bindVord20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Vord20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vord20 *Vord20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vord20.Contract.Vord20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vord20 *Vord20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vord20.Contract.Vord20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vord20 *Vord20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vord20.Contract.Vord20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vord20 *Vord20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vord20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vord20 *Vord20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vord20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vord20 *Vord20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vord20.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Vord20 *Vord20Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vord20.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Vord20 *Vord20Session) Fee() (*big.Int, error) {
	return _Vord20.Contract.Fee(&_Vord20.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Vord20 *Vord20CallerSession) Fee() (*big.Int, error) {
	return _Vord20.Contract.Fee(&_Vord20.CallOpts)
}

// House is a free data retrieval call binding the contract method 0xb1cad5e3.
//
// Solidity: function house(address player) view returns(string)
func (_Vord20 *Vord20Caller) House(opts *bind.CallOpts, player common.Address) (string, error) {
	var out []interface{}
	err := _Vord20.contract.Call(opts, &out, "house", player)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// House is a free data retrieval call binding the contract method 0xb1cad5e3.
//
// Solidity: function house(address player) view returns(string)
func (_Vord20 *Vord20Session) House(player common.Address) (string, error) {
	return _Vord20.Contract.House(&_Vord20.CallOpts, player)
}

// House is a free data retrieval call binding the contract method 0xb1cad5e3.
//
// Solidity: function house(address player) view returns(string)
func (_Vord20 *Vord20CallerSession) House(player common.Address) (string, error) {
	return _Vord20.Contract.House(&_Vord20.CallOpts, player)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_Vord20 *Vord20Caller) KeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vord20.contract.Call(opts, &out, "keyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_Vord20 *Vord20Session) KeyHash() ([32]byte, error) {
	return _Vord20.Contract.KeyHash(&_Vord20.CallOpts)
}

// KeyHash is a free data retrieval call binding the contract method 0x61728f39.
//
// Solidity: function keyHash() view returns(bytes32)
func (_Vord20 *Vord20CallerSession) KeyHash() ([32]byte, error) {
	return _Vord20.Contract.KeyHash(&_Vord20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vord20 *Vord20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vord20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vord20 *Vord20Session) Owner() (common.Address, error) {
	return _Vord20.Contract.Owner(&_Vord20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vord20 *Vord20CallerSession) Owner() (common.Address, error) {
	return _Vord20.Contract.Owner(&_Vord20.CallOpts)
}

// IncreaseVorCoordinatorAllowance is a paid mutator transaction binding the contract method 0xf4b981e0.
//
// Solidity: function increaseVorCoordinatorAllowance(uint256 _amount) returns()
func (_Vord20 *Vord20Transactor) IncreaseVorCoordinatorAllowance(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "increaseVorCoordinatorAllowance", _amount)
}

// IncreaseVorCoordinatorAllowance is a paid mutator transaction binding the contract method 0xf4b981e0.
//
// Solidity: function increaseVorCoordinatorAllowance(uint256 _amount) returns()
func (_Vord20 *Vord20Session) IncreaseVorCoordinatorAllowance(_amount *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.IncreaseVorCoordinatorAllowance(&_Vord20.TransactOpts, _amount)
}

// IncreaseVorCoordinatorAllowance is a paid mutator transaction binding the contract method 0xf4b981e0.
//
// Solidity: function increaseVorCoordinatorAllowance(uint256 _amount) returns()
func (_Vord20 *Vord20TransactorSession) IncreaseVorCoordinatorAllowance(_amount *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.IncreaseVorCoordinatorAllowance(&_Vord20.TransactOpts, _amount)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Vord20 *Vord20Transactor) RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "rawFulfillRandomness", requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Vord20 *Vord20Session) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.RawFulfillRandomness(&_Vord20.TransactOpts, requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Vord20 *Vord20TransactorSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.RawFulfillRandomness(&_Vord20.TransactOpts, requestId, randomness)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vord20 *Vord20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vord20 *Vord20Session) RenounceOwnership() (*types.Transaction, error) {
	return _Vord20.Contract.RenounceOwnership(&_Vord20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vord20 *Vord20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vord20.Contract.RenounceOwnership(&_Vord20.TransactOpts)
}

// RollDice is a paid mutator transaction binding the contract method 0x5e816740.
//
// Solidity: function rollDice(uint256 userProvidedSeed, address roller) returns(bytes32 requestId)
func (_Vord20 *Vord20Transactor) RollDice(opts *bind.TransactOpts, userProvidedSeed *big.Int, roller common.Address) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "rollDice", userProvidedSeed, roller)
}

// RollDice is a paid mutator transaction binding the contract method 0x5e816740.
//
// Solidity: function rollDice(uint256 userProvidedSeed, address roller) returns(bytes32 requestId)
func (_Vord20 *Vord20Session) RollDice(userProvidedSeed *big.Int, roller common.Address) (*types.Transaction, error) {
	return _Vord20.Contract.RollDice(&_Vord20.TransactOpts, userProvidedSeed, roller)
}

// RollDice is a paid mutator transaction binding the contract method 0x5e816740.
//
// Solidity: function rollDice(uint256 userProvidedSeed, address roller) returns(bytes32 requestId)
func (_Vord20 *Vord20TransactorSession) RollDice(userProvidedSeed *big.Int, roller common.Address) (*types.Transaction, error) {
	return _Vord20.Contract.RollDice(&_Vord20.TransactOpts, userProvidedSeed, roller)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_Vord20 *Vord20Transactor) SetFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "setFee", fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_Vord20 *Vord20Session) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.SetFee(&_Vord20.TransactOpts, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 fee) returns()
func (_Vord20 *Vord20TransactorSession) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.SetFee(&_Vord20.TransactOpts, fee)
}

// SetKeyHash is a paid mutator transaction binding the contract method 0x98544710.
//
// Solidity: function setKeyHash(bytes32 keyHash) returns()
func (_Vord20 *Vord20Transactor) SetKeyHash(opts *bind.TransactOpts, keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "setKeyHash", keyHash)
}

// SetKeyHash is a paid mutator transaction binding the contract method 0x98544710.
//
// Solidity: function setKeyHash(bytes32 keyHash) returns()
func (_Vord20 *Vord20Session) SetKeyHash(keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.Contract.SetKeyHash(&_Vord20.TransactOpts, keyHash)
}

// SetKeyHash is a paid mutator transaction binding the contract method 0x98544710.
//
// Solidity: function setKeyHash(bytes32 keyHash) returns()
func (_Vord20 *Vord20TransactorSession) SetKeyHash(keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.Contract.SetKeyHash(&_Vord20.TransactOpts, keyHash)
}

// TopUpGas is a paid mutator transaction binding the contract method 0x3064e1cd.
//
// Solidity: function topUpGas(bytes32 _keyHash) payable returns()
func (_Vord20 *Vord20Transactor) TopUpGas(opts *bind.TransactOpts, _keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "topUpGas", _keyHash)
}

// TopUpGas is a paid mutator transaction binding the contract method 0x3064e1cd.
//
// Solidity: function topUpGas(bytes32 _keyHash) payable returns()
func (_Vord20 *Vord20Session) TopUpGas(_keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.Contract.TopUpGas(&_Vord20.TransactOpts, _keyHash)
}

// TopUpGas is a paid mutator transaction binding the contract method 0x3064e1cd.
//
// Solidity: function topUpGas(bytes32 _keyHash) payable returns()
func (_Vord20 *Vord20TransactorSession) TopUpGas(_keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.Contract.TopUpGas(&_Vord20.TransactOpts, _keyHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vord20 *Vord20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vord20 *Vord20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vord20.Contract.TransferOwnership(&_Vord20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vord20 *Vord20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vord20.Contract.TransferOwnership(&_Vord20.TransactOpts, newOwner)
}

// WithDrawGasTopUp is a paid mutator transaction binding the contract method 0xca9c7d27.
//
// Solidity: function withDrawGasTopUp(bytes32 _keyHash) returns()
func (_Vord20 *Vord20Transactor) WithDrawGasTopUp(opts *bind.TransactOpts, _keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "withDrawGasTopUp", _keyHash)
}

// WithDrawGasTopUp is a paid mutator transaction binding the contract method 0xca9c7d27.
//
// Solidity: function withDrawGasTopUp(bytes32 _keyHash) returns()
func (_Vord20 *Vord20Session) WithDrawGasTopUp(_keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.Contract.WithDrawGasTopUp(&_Vord20.TransactOpts, _keyHash)
}

// WithDrawGasTopUp is a paid mutator transaction binding the contract method 0xca9c7d27.
//
// Solidity: function withDrawGasTopUp(bytes32 _keyHash) returns()
func (_Vord20 *Vord20TransactorSession) WithDrawGasTopUp(_keyHash [32]byte) (*types.Transaction, error) {
	return _Vord20.Contract.WithDrawGasTopUp(&_Vord20.TransactOpts, _keyHash)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns(bool success)
func (_Vord20 *Vord20Transactor) WithdrawEth(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "withdrawEth", amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns(bool success)
func (_Vord20 *Vord20Session) WithdrawEth(amount *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.WithdrawEth(&_Vord20.TransactOpts, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xc311d049.
//
// Solidity: function withdrawEth(uint256 amount) returns(bool success)
func (_Vord20 *Vord20TransactorSession) WithdrawEth(amount *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.WithdrawEth(&_Vord20.TransactOpts, amount)
}

// WithdrawXFUND is a paid mutator transaction binding the contract method 0x2916ddb1.
//
// Solidity: function withdrawXFUND(address to, uint256 value) returns()
func (_Vord20 *Vord20Transactor) WithdrawXFUND(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vord20.contract.Transact(opts, "withdrawXFUND", to, value)
}

// WithdrawXFUND is a paid mutator transaction binding the contract method 0x2916ddb1.
//
// Solidity: function withdrawXFUND(address to, uint256 value) returns()
func (_Vord20 *Vord20Session) WithdrawXFUND(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.WithdrawXFUND(&_Vord20.TransactOpts, to, value)
}

// WithdrawXFUND is a paid mutator transaction binding the contract method 0x2916ddb1.
//
// Solidity: function withdrawXFUND(address to, uint256 value) returns()
func (_Vord20 *Vord20TransactorSession) WithdrawXFUND(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Vord20.Contract.WithdrawXFUND(&_Vord20.TransactOpts, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vord20 *Vord20Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vord20.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vord20 *Vord20Session) Receive() (*types.Transaction, error) {
	return _Vord20.Contract.Receive(&_Vord20.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vord20 *Vord20TransactorSession) Receive() (*types.Transaction, error) {
	return _Vord20.Contract.Receive(&_Vord20.TransactOpts)
}

// Vord20DiceLandedIterator is returned from FilterDiceLanded and is used to iterate over the raw logs and unpacked data for DiceLanded events raised by the Vord20 contract.
type Vord20DiceLandedIterator struct {
	Event *Vord20DiceLanded // Event containing the contract specifics and raw log

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
func (it *Vord20DiceLandedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vord20DiceLanded)
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
		it.Event = new(Vord20DiceLanded)
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
func (it *Vord20DiceLandedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vord20DiceLandedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vord20DiceLanded represents a DiceLanded event raised by the Vord20 contract.
type Vord20DiceLanded struct {
	RequestId [32]byte
	Result    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDiceLanded is a free log retrieval operation binding the contract event 0x909dd726737b7ffa4ae9d137e9edebe8a74a9c2910a4b60e8112f93ab2170837.
//
// Solidity: event DiceLanded(bytes32 indexed requestId, uint256 indexed result)
func (_Vord20 *Vord20Filterer) FilterDiceLanded(opts *bind.FilterOpts, requestId [][32]byte, result []*big.Int) (*Vord20DiceLandedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var resultRule []interface{}
	for _, resultItem := range result {
		resultRule = append(resultRule, resultItem)
	}

	logs, sub, err := _Vord20.contract.FilterLogs(opts, "DiceLanded", requestIdRule, resultRule)
	if err != nil {
		return nil, err
	}
	return &Vord20DiceLandedIterator{contract: _Vord20.contract, event: "DiceLanded", logs: logs, sub: sub}, nil
}

// WatchDiceLanded is a free log subscription operation binding the contract event 0x909dd726737b7ffa4ae9d137e9edebe8a74a9c2910a4b60e8112f93ab2170837.
//
// Solidity: event DiceLanded(bytes32 indexed requestId, uint256 indexed result)
func (_Vord20 *Vord20Filterer) WatchDiceLanded(opts *bind.WatchOpts, sink chan<- *Vord20DiceLanded, requestId [][32]byte, result []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var resultRule []interface{}
	for _, resultItem := range result {
		resultRule = append(resultRule, resultItem)
	}

	logs, sub, err := _Vord20.contract.WatchLogs(opts, "DiceLanded", requestIdRule, resultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vord20DiceLanded)
				if err := _Vord20.contract.UnpackLog(event, "DiceLanded", log); err != nil {
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
func (_Vord20 *Vord20Filterer) ParseDiceLanded(log types.Log) (*Vord20DiceLanded, error) {
	event := new(Vord20DiceLanded)
	if err := _Vord20.contract.UnpackLog(event, "DiceLanded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vord20DiceRolledIterator is returned from FilterDiceRolled and is used to iterate over the raw logs and unpacked data for DiceRolled events raised by the Vord20 contract.
type Vord20DiceRolledIterator struct {
	Event *Vord20DiceRolled // Event containing the contract specifics and raw log

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
func (it *Vord20DiceRolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vord20DiceRolled)
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
		it.Event = new(Vord20DiceRolled)
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
func (it *Vord20DiceRolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vord20DiceRolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vord20DiceRolled represents a DiceRolled event raised by the Vord20 contract.
type Vord20DiceRolled struct {
	RequestId [32]byte
	Roller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDiceRolled is a free log retrieval operation binding the contract event 0x923de4fc4aece24a78a9e4ca3009c571a742f81ac2c004a229224b0fd1883bdd.
//
// Solidity: event DiceRolled(bytes32 indexed requestId, address indexed roller)
func (_Vord20 *Vord20Filterer) FilterDiceRolled(opts *bind.FilterOpts, requestId [][32]byte, roller []common.Address) (*Vord20DiceRolledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var rollerRule []interface{}
	for _, rollerItem := range roller {
		rollerRule = append(rollerRule, rollerItem)
	}

	logs, sub, err := _Vord20.contract.FilterLogs(opts, "DiceRolled", requestIdRule, rollerRule)
	if err != nil {
		return nil, err
	}
	return &Vord20DiceRolledIterator{contract: _Vord20.contract, event: "DiceRolled", logs: logs, sub: sub}, nil
}

// WatchDiceRolled is a free log subscription operation binding the contract event 0x923de4fc4aece24a78a9e4ca3009c571a742f81ac2c004a229224b0fd1883bdd.
//
// Solidity: event DiceRolled(bytes32 indexed requestId, address indexed roller)
func (_Vord20 *Vord20Filterer) WatchDiceRolled(opts *bind.WatchOpts, sink chan<- *Vord20DiceRolled, requestId [][32]byte, roller []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var rollerRule []interface{}
	for _, rollerItem := range roller {
		rollerRule = append(rollerRule, rollerItem)
	}

	logs, sub, err := _Vord20.contract.WatchLogs(opts, "DiceRolled", requestIdRule, rollerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vord20DiceRolled)
				if err := _Vord20.contract.UnpackLog(event, "DiceRolled", log); err != nil {
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
func (_Vord20 *Vord20Filterer) ParseDiceRolled(log types.Log) (*Vord20DiceRolled, error) {
	event := new(Vord20DiceRolled)
	if err := _Vord20.contract.UnpackLog(event, "DiceRolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vord20EthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the Vord20 contract.
type Vord20EthWithdrawnIterator struct {
	Event *Vord20EthWithdrawn // Event containing the contract specifics and raw log

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
func (it *Vord20EthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vord20EthWithdrawn)
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
		it.Event = new(Vord20EthWithdrawn)
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
func (it *Vord20EthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vord20EthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vord20EthWithdrawn represents a EthWithdrawn event raised by the Vord20 contract.
type Vord20EthWithdrawn struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address receiver, uint256 amount)
func (_Vord20 *Vord20Filterer) FilterEthWithdrawn(opts *bind.FilterOpts) (*Vord20EthWithdrawnIterator, error) {

	logs, sub, err := _Vord20.contract.FilterLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return &Vord20EthWithdrawnIterator{contract: _Vord20.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address receiver, uint256 amount)
func (_Vord20 *Vord20Filterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *Vord20EthWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Vord20.contract.WatchLogs(opts, "EthWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vord20EthWithdrawn)
				if err := _Vord20.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address receiver, uint256 amount)
func (_Vord20 *Vord20Filterer) ParseEthWithdrawn(log types.Log) (*Vord20EthWithdrawn, error) {
	event := new(Vord20EthWithdrawn)
	if err := _Vord20.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Vord20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vord20 contract.
type Vord20OwnershipTransferredIterator struct {
	Event *Vord20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Vord20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Vord20OwnershipTransferred)
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
		it.Event = new(Vord20OwnershipTransferred)
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
func (it *Vord20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Vord20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Vord20OwnershipTransferred represents a OwnershipTransferred event raised by the Vord20 contract.
type Vord20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vord20 *Vord20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Vord20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vord20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Vord20OwnershipTransferredIterator{contract: _Vord20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vord20 *Vord20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Vord20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vord20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Vord20OwnershipTransferred)
				if err := _Vord20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Vord20 *Vord20Filterer) ParseOwnershipTransferred(log types.Log) (*Vord20OwnershipTransferred, error) {
	event := new(Vord20OwnershipTransferred)
	if err := _Vord20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
