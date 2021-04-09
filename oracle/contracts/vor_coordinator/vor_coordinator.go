// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vor_coordinator

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

// VorCoordinatorABI is the input ABI used to generate the binding from.
const VorCoordinatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xfund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blockHashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ChangeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewServiceAgreement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"RandomnessRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequestFulfilled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRESEED_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROOF_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_KEY_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"callbacks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"randomnessFee\",\"type\":\"uint96\"},{\"internalType\":\"bytes32\",\"name\":\"seedAndBlockNum\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"serviceAgreements\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"vOROracle\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"getProviderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"_publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_consumerSeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePaid\",\"type\":\"uint256\"}],\"name\":\"randomnessRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"fulfillRandomnessRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VorCoordinator is an auto generated Go binding around an Ethereum contract.
type VorCoordinator struct {
	VorCoordinatorCaller     // Read-only binding to the contract
	VorCoordinatorTransactor // Write-only binding to the contract
	VorCoordinatorFilterer   // Log filterer for contract events
}

// VorCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VorCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VorCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VorCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VorCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VorCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VorCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VorCoordinatorSession struct {
	Contract     *VorCoordinator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VorCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VorCoordinatorCallerSession struct {
	Contract *VorCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VorCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VorCoordinatorTransactorSession struct {
	Contract     *VorCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VorCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VorCoordinatorRaw struct {
	Contract *VorCoordinator // Generic contract binding to access the raw methods on
}

// VorCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VorCoordinatorCallerRaw struct {
	Contract *VorCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// VorCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VorCoordinatorTransactorRaw struct {
	Contract *VorCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVorCoordinator creates a new instance of VorCoordinator, bound to a specific deployed contract.
func NewVorCoordinator(address common.Address, backend bind.ContractBackend) (*VorCoordinator, error) {
	contract, err := bindVorCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VorCoordinator{VorCoordinatorCaller: VorCoordinatorCaller{contract: contract}, VorCoordinatorTransactor: VorCoordinatorTransactor{contract: contract}, VorCoordinatorFilterer: VorCoordinatorFilterer{contract: contract}}, nil
}

// NewVorCoordinatorCaller creates a new read-only instance of VorCoordinator, bound to a specific deployed contract.
func NewVorCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*VorCoordinatorCaller, error) {
	contract, err := bindVorCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorCaller{contract: contract}, nil
}

// NewVorCoordinatorTransactor creates a new write-only instance of VorCoordinator, bound to a specific deployed contract.
func NewVorCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VorCoordinatorTransactor, error) {
	contract, err := bindVorCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorTransactor{contract: contract}, nil
}

// NewVorCoordinatorFilterer creates a new log filterer instance of VorCoordinator, bound to a specific deployed contract.
func NewVorCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VorCoordinatorFilterer, error) {
	contract, err := bindVorCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorFilterer{contract: contract}, nil
}

// bindVorCoordinator binds a generic wrapper to an already deployed contract.
func bindVorCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VorCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VorCoordinator *VorCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VorCoordinator.Contract.VorCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VorCoordinator *VorCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VorCoordinator.Contract.VorCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VorCoordinator *VorCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VorCoordinator.Contract.VorCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VorCoordinator *VorCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VorCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VorCoordinator *VorCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VorCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VorCoordinator *VorCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VorCoordinator.Contract.contract.Transact(opts, method, params...)
}

// PRESEEDOFFSET is a free data retrieval call binding the contract method 0xb415f4f5.
//
// Solidity: function PRESEED_OFFSET() view returns(uint256)
func (_VorCoordinator *VorCoordinatorCaller) PRESEEDOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "PRESEED_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRESEEDOFFSET is a free data retrieval call binding the contract method 0xb415f4f5.
//
// Solidity: function PRESEED_OFFSET() view returns(uint256)
func (_VorCoordinator *VorCoordinatorSession) PRESEEDOFFSET() (*big.Int, error) {
	return _VorCoordinator.Contract.PRESEEDOFFSET(&_VorCoordinator.CallOpts)
}

// PRESEEDOFFSET is a free data retrieval call binding the contract method 0xb415f4f5.
//
// Solidity: function PRESEED_OFFSET() view returns(uint256)
func (_VorCoordinator *VorCoordinatorCallerSession) PRESEEDOFFSET() (*big.Int, error) {
	return _VorCoordinator.Contract.PRESEEDOFFSET(&_VorCoordinator.CallOpts)
}

// PROOFLENGTH is a free data retrieval call binding the contract method 0xe911439c.
//
// Solidity: function PROOF_LENGTH() view returns(uint256)
func (_VorCoordinator *VorCoordinatorCaller) PROOFLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "PROOF_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROOFLENGTH is a free data retrieval call binding the contract method 0xe911439c.
//
// Solidity: function PROOF_LENGTH() view returns(uint256)
func (_VorCoordinator *VorCoordinatorSession) PROOFLENGTH() (*big.Int, error) {
	return _VorCoordinator.Contract.PROOFLENGTH(&_VorCoordinator.CallOpts)
}

// PROOFLENGTH is a free data retrieval call binding the contract method 0xe911439c.
//
// Solidity: function PROOF_LENGTH() view returns(uint256)
func (_VorCoordinator *VorCoordinatorCallerSession) PROOFLENGTH() (*big.Int, error) {
	return _VorCoordinator.Contract.PROOFLENGTH(&_VorCoordinator.CallOpts)
}

// PUBLICKEYOFFSET is a free data retrieval call binding the contract method 0x8aa7927b.
//
// Solidity: function PUBLIC_KEY_OFFSET() view returns(uint256)
func (_VorCoordinator *VorCoordinatorCaller) PUBLICKEYOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "PUBLIC_KEY_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBLICKEYOFFSET is a free data retrieval call binding the contract method 0x8aa7927b.
//
// Solidity: function PUBLIC_KEY_OFFSET() view returns(uint256)
func (_VorCoordinator *VorCoordinatorSession) PUBLICKEYOFFSET() (*big.Int, error) {
	return _VorCoordinator.Contract.PUBLICKEYOFFSET(&_VorCoordinator.CallOpts)
}

// PUBLICKEYOFFSET is a free data retrieval call binding the contract method 0x8aa7927b.
//
// Solidity: function PUBLIC_KEY_OFFSET() view returns(uint256)
func (_VorCoordinator *VorCoordinatorCallerSession) PUBLICKEYOFFSET() (*big.Int, error) {
	return _VorCoordinator.Contract.PUBLICKEYOFFSET(&_VorCoordinator.CallOpts)
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) view returns(address callbackContract, uint96 randomnessFee, bytes32 seedAndBlockNum)
func (_VorCoordinator *VorCoordinatorCaller) Callbacks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	SeedAndBlockNum  [32]byte
}, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "callbacks", arg0)

	outstruct := new(struct {
		CallbackContract common.Address
		RandomnessFee    *big.Int
		SeedAndBlockNum  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CallbackContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RandomnessFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SeedAndBlockNum = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) view returns(address callbackContract, uint96 randomnessFee, bytes32 seedAndBlockNum)
func (_VorCoordinator *VorCoordinatorSession) Callbacks(arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	SeedAndBlockNum  [32]byte
}, error) {
	return _VorCoordinator.Contract.Callbacks(&_VorCoordinator.CallOpts, arg0)
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) view returns(address callbackContract, uint96 randomnessFee, bytes32 seedAndBlockNum)
func (_VorCoordinator *VorCoordinatorCallerSession) Callbacks(arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	SeedAndBlockNum  [32]byte
}, error) {
	return _VorCoordinator.Contract.Callbacks(&_VorCoordinator.CallOpts, arg0)
}

// GetProviderAddress is a free data retrieval call binding the contract method 0x9845fb9c.
//
// Solidity: function getProviderAddress(bytes32 _keyHash) view returns(address)
func (_VorCoordinator *VorCoordinatorCaller) GetProviderAddress(opts *bind.CallOpts, _keyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "getProviderAddress", _keyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProviderAddress is a free data retrieval call binding the contract method 0x9845fb9c.
//
// Solidity: function getProviderAddress(bytes32 _keyHash) view returns(address)
func (_VorCoordinator *VorCoordinatorSession) GetProviderAddress(_keyHash [32]byte) (common.Address, error) {
	return _VorCoordinator.Contract.GetProviderAddress(&_VorCoordinator.CallOpts, _keyHash)
}

// GetProviderAddress is a free data retrieval call binding the contract method 0x9845fb9c.
//
// Solidity: function getProviderAddress(bytes32 _keyHash) view returns(address)
func (_VorCoordinator *VorCoordinatorCallerSession) GetProviderAddress(_keyHash [32]byte) (common.Address, error) {
	return _VorCoordinator.Contract.GetProviderAddress(&_VorCoordinator.CallOpts, _keyHash)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) pure returns(bytes32)
func (_VorCoordinator *VorCoordinatorCaller) HashOfKey(opts *bind.CallOpts, _publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "hashOfKey", _publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) pure returns(bytes32)
func (_VorCoordinator *VorCoordinatorSession) HashOfKey(_publicKey [2]*big.Int) ([32]byte, error) {
	return _VorCoordinator.Contract.HashOfKey(&_VorCoordinator.CallOpts, _publicKey)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) pure returns(bytes32)
func (_VorCoordinator *VorCoordinatorCallerSession) HashOfKey(_publicKey [2]*big.Int) ([32]byte, error) {
	return _VorCoordinator.Contract.HashOfKey(&_VorCoordinator.CallOpts, _publicKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VorCoordinator *VorCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VorCoordinator *VorCoordinatorSession) Owner() (common.Address, error) {
	return _VorCoordinator.Contract.Owner(&_VorCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VorCoordinator *VorCoordinatorCallerSession) Owner() (common.Address, error) {
	return _VorCoordinator.Contract.Owner(&_VorCoordinator.CallOpts)
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) view returns(address vOROracle, uint96 fee)
func (_VorCoordinator *VorCoordinatorCaller) ServiceAgreements(opts *bind.CallOpts, arg0 [32]byte) (struct {
	VOROracle common.Address
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "serviceAgreements", arg0)

	outstruct := new(struct {
		VOROracle common.Address
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VOROracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) view returns(address vOROracle, uint96 fee)
func (_VorCoordinator *VorCoordinatorSession) ServiceAgreements(arg0 [32]byte) (struct {
	VOROracle common.Address
	Fee       *big.Int
}, error) {
	return _VorCoordinator.Contract.ServiceAgreements(&_VorCoordinator.CallOpts, arg0)
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) view returns(address vOROracle, uint96 fee)
func (_VorCoordinator *VorCoordinatorCallerSession) ServiceAgreements(arg0 [32]byte) (struct {
	VOROracle common.Address
	Fee       *big.Int
}, error) {
	return _VorCoordinator.Contract.ServiceAgreements(&_VorCoordinator.CallOpts, arg0)
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) view returns(uint256)
func (_VorCoordinator *VorCoordinatorCaller) WithdrawableTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VorCoordinator.contract.Call(opts, &out, "withdrawableTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) view returns(uint256)
func (_VorCoordinator *VorCoordinatorSession) WithdrawableTokens(arg0 common.Address) (*big.Int, error) {
	return _VorCoordinator.Contract.WithdrawableTokens(&_VorCoordinator.CallOpts, arg0)
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) view returns(uint256)
func (_VorCoordinator *VorCoordinatorCallerSession) WithdrawableTokens(arg0 common.Address) (*big.Int, error) {
	return _VorCoordinator.Contract.WithdrawableTokens(&_VorCoordinator.CallOpts, arg0)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x81f1e046.
//
// Solidity: function changeFee(uint256[2] _publicProvingKey, uint256 _fee) returns()
func (_VorCoordinator *VorCoordinatorTransactor) ChangeFee(opts *bind.TransactOpts, _publicProvingKey [2]*big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "changeFee", _publicProvingKey, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x81f1e046.
//
// Solidity: function changeFee(uint256[2] _publicProvingKey, uint256 _fee) returns()
func (_VorCoordinator *VorCoordinatorSession) ChangeFee(_publicProvingKey [2]*big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.ChangeFee(&_VorCoordinator.TransactOpts, _publicProvingKey, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x81f1e046.
//
// Solidity: function changeFee(uint256[2] _publicProvingKey, uint256 _fee) returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) ChangeFee(_publicProvingKey [2]*big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.ChangeFee(&_VorCoordinator.TransactOpts, _publicProvingKey, _fee)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns()
func (_VorCoordinator *VorCoordinatorTransactor) FulfillRandomnessRequest(opts *bind.TransactOpts, _proof []byte) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "fulfillRandomnessRequest", _proof)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns()
func (_VorCoordinator *VorCoordinatorSession) FulfillRandomnessRequest(_proof []byte) (*types.Transaction, error) {
	return _VorCoordinator.Contract.FulfillRandomnessRequest(&_VorCoordinator.TransactOpts, _proof)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) FulfillRandomnessRequest(_proof []byte) (*types.Transaction, error) {
	return _VorCoordinator.Contract.FulfillRandomnessRequest(&_VorCoordinator.TransactOpts, _proof)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VorCoordinator *VorCoordinatorTransactor) RandomnessRequest(opts *bind.TransactOpts, _keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "randomnessRequest", _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VorCoordinator *VorCoordinatorSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.RandomnessRequest(&_VorCoordinator.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.RandomnessRequest(&_VorCoordinator.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xbbf70d69.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey) returns()
func (_VorCoordinator *VorCoordinatorTransactor) RegisterProvingKey(opts *bind.TransactOpts, _fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "registerProvingKey", _fee, _oracle, _publicProvingKey)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xbbf70d69.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey) returns()
func (_VorCoordinator *VorCoordinatorSession) RegisterProvingKey(_fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.RegisterProvingKey(&_VorCoordinator.TransactOpts, _fee, _oracle, _publicProvingKey)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xbbf70d69.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey) returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) RegisterProvingKey(_fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.RegisterProvingKey(&_VorCoordinator.TransactOpts, _fee, _oracle, _publicProvingKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VorCoordinator *VorCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VorCoordinator *VorCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VorCoordinator.Contract.RenounceOwnership(&_VorCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VorCoordinator.Contract.RenounceOwnership(&_VorCoordinator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VorCoordinator *VorCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VorCoordinator *VorCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VorCoordinator.Contract.TransferOwnership(&_VorCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VorCoordinator.Contract.TransferOwnership(&_VorCoordinator.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VorCoordinator *VorCoordinatorTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.contract.Transact(opts, "withdraw", _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VorCoordinator *VorCoordinatorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.Withdraw(&_VorCoordinator.TransactOpts, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VorCoordinator *VorCoordinatorTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VorCoordinator.Contract.Withdraw(&_VorCoordinator.TransactOpts, _recipient, _amount)
}

// VorCoordinatorChangeFeeIterator is returned from FilterChangeFee and is used to iterate over the raw logs and unpacked data for ChangeFee events raised by the VorCoordinator contract.
type VorCoordinatorChangeFeeIterator struct {
	Event *VorCoordinatorChangeFee // Event containing the contract specifics and raw log

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
func (it *VorCoordinatorChangeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorCoordinatorChangeFee)
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
		it.Event = new(VorCoordinatorChangeFee)
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
func (it *VorCoordinatorChangeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorCoordinatorChangeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorCoordinatorChangeFee represents a ChangeFee event raised by the VorCoordinator contract.
type VorCoordinatorChangeFee struct {
	KeyHash [32]byte
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangeFee is a free log retrieval operation binding the contract event 0xfc6ca2918cb95fbfc07e4567da6387d5e4209b4ced636bbe450c09895d447526.
//
// Solidity: event ChangeFee(bytes32 keyHash, uint256 fee)
func (_VorCoordinator *VorCoordinatorFilterer) FilterChangeFee(opts *bind.FilterOpts) (*VorCoordinatorChangeFeeIterator, error) {

	logs, sub, err := _VorCoordinator.contract.FilterLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorChangeFeeIterator{contract: _VorCoordinator.contract, event: "ChangeFee", logs: logs, sub: sub}, nil
}

// WatchChangeFee is a free log subscription operation binding the contract event 0xfc6ca2918cb95fbfc07e4567da6387d5e4209b4ced636bbe450c09895d447526.
//
// Solidity: event ChangeFee(bytes32 keyHash, uint256 fee)
func (_VorCoordinator *VorCoordinatorFilterer) WatchChangeFee(opts *bind.WatchOpts, sink chan<- *VorCoordinatorChangeFee) (event.Subscription, error) {

	logs, sub, err := _VorCoordinator.contract.WatchLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorCoordinatorChangeFee)
				if err := _VorCoordinator.contract.UnpackLog(event, "ChangeFee", log); err != nil {
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

// ParseChangeFee is a log parse operation binding the contract event 0xfc6ca2918cb95fbfc07e4567da6387d5e4209b4ced636bbe450c09895d447526.
//
// Solidity: event ChangeFee(bytes32 keyHash, uint256 fee)
func (_VorCoordinator *VorCoordinatorFilterer) ParseChangeFee(log types.Log) (*VorCoordinatorChangeFee, error) {
	event := new(VorCoordinatorChangeFee)
	if err := _VorCoordinator.contract.UnpackLog(event, "ChangeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VorCoordinatorNewServiceAgreementIterator is returned from FilterNewServiceAgreement and is used to iterate over the raw logs and unpacked data for NewServiceAgreement events raised by the VorCoordinator contract.
type VorCoordinatorNewServiceAgreementIterator struct {
	Event *VorCoordinatorNewServiceAgreement // Event containing the contract specifics and raw log

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
func (it *VorCoordinatorNewServiceAgreementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorCoordinatorNewServiceAgreement)
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
		it.Event = new(VorCoordinatorNewServiceAgreement)
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
func (it *VorCoordinatorNewServiceAgreementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorCoordinatorNewServiceAgreementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorCoordinatorNewServiceAgreement represents a NewServiceAgreement event raised by the VorCoordinator contract.
type VorCoordinatorNewServiceAgreement struct {
	KeyHash [32]byte
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewServiceAgreement is a free log retrieval operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VorCoordinator *VorCoordinatorFilterer) FilterNewServiceAgreement(opts *bind.FilterOpts) (*VorCoordinatorNewServiceAgreementIterator, error) {

	logs, sub, err := _VorCoordinator.contract.FilterLogs(opts, "NewServiceAgreement")
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorNewServiceAgreementIterator{contract: _VorCoordinator.contract, event: "NewServiceAgreement", logs: logs, sub: sub}, nil
}

// WatchNewServiceAgreement is a free log subscription operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VorCoordinator *VorCoordinatorFilterer) WatchNewServiceAgreement(opts *bind.WatchOpts, sink chan<- *VorCoordinatorNewServiceAgreement) (event.Subscription, error) {

	logs, sub, err := _VorCoordinator.contract.WatchLogs(opts, "NewServiceAgreement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorCoordinatorNewServiceAgreement)
				if err := _VorCoordinator.contract.UnpackLog(event, "NewServiceAgreement", log); err != nil {
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

// ParseNewServiceAgreement is a log parse operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VorCoordinator *VorCoordinatorFilterer) ParseNewServiceAgreement(log types.Log) (*VorCoordinatorNewServiceAgreement, error) {
	event := new(VorCoordinatorNewServiceAgreement)
	if err := _VorCoordinator.contract.UnpackLog(event, "NewServiceAgreement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VorCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VorCoordinator contract.
type VorCoordinatorOwnershipTransferredIterator struct {
	Event *VorCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VorCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorCoordinatorOwnershipTransferred)
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
		it.Event = new(VorCoordinatorOwnershipTransferred)
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
func (it *VorCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the VorCoordinator contract.
type VorCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VorCoordinator *VorCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VorCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VorCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorOwnershipTransferredIterator{contract: _VorCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VorCoordinator *VorCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VorCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VorCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorCoordinatorOwnershipTransferred)
				if err := _VorCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VorCoordinator *VorCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*VorCoordinatorOwnershipTransferred, error) {
	event := new(VorCoordinatorOwnershipTransferred)
	if err := _VorCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VorCoordinatorRandomnessRequestIterator is returned from FilterRandomnessRequest and is used to iterate over the raw logs and unpacked data for RandomnessRequest events raised by the VorCoordinator contract.
type VorCoordinatorRandomnessRequestIterator struct {
	Event *VorCoordinatorRandomnessRequest // Event containing the contract specifics and raw log

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
func (it *VorCoordinatorRandomnessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorCoordinatorRandomnessRequest)
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
		it.Event = new(VorCoordinatorRandomnessRequest)
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
func (it *VorCoordinatorRandomnessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorCoordinatorRandomnessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorCoordinatorRandomnessRequest represents a RandomnessRequest event raised by the VorCoordinator contract.
type VorCoordinatorRandomnessRequest struct {
	KeyHash   [32]byte
	Seed      *big.Int
	Sender    common.Address
	Fee       *big.Int
	RequestID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessRequest is a free log retrieval operation binding the contract event 0xebb37373bb11123e38f964627878b02c247f92f3913df7cf3f270b5222c8d2be.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)
func (_VorCoordinator *VorCoordinatorFilterer) FilterRandomnessRequest(opts *bind.FilterOpts) (*VorCoordinatorRandomnessRequestIterator, error) {

	logs, sub, err := _VorCoordinator.contract.FilterLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorRandomnessRequestIterator{contract: _VorCoordinator.contract, event: "RandomnessRequest", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequest is a free log subscription operation binding the contract event 0xebb37373bb11123e38f964627878b02c247f92f3913df7cf3f270b5222c8d2be.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)
func (_VorCoordinator *VorCoordinatorFilterer) WatchRandomnessRequest(opts *bind.WatchOpts, sink chan<- *VorCoordinatorRandomnessRequest) (event.Subscription, error) {

	logs, sub, err := _VorCoordinator.contract.WatchLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorCoordinatorRandomnessRequest)
				if err := _VorCoordinator.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
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

// ParseRandomnessRequest is a log parse operation binding the contract event 0xebb37373bb11123e38f964627878b02c247f92f3913df7cf3f270b5222c8d2be.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)
func (_VorCoordinator *VorCoordinatorFilterer) ParseRandomnessRequest(log types.Log) (*VorCoordinatorRandomnessRequest, error) {
	event := new(VorCoordinatorRandomnessRequest)
	if err := _VorCoordinator.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VorCoordinatorRandomnessRequestFulfilledIterator is returned from FilterRandomnessRequestFulfilled and is used to iterate over the raw logs and unpacked data for RandomnessRequestFulfilled events raised by the VorCoordinator contract.
type VorCoordinatorRandomnessRequestFulfilledIterator struct {
	Event *VorCoordinatorRandomnessRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *VorCoordinatorRandomnessRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorCoordinatorRandomnessRequestFulfilled)
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
		it.Event = new(VorCoordinatorRandomnessRequestFulfilled)
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
func (it *VorCoordinatorRandomnessRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorCoordinatorRandomnessRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorCoordinatorRandomnessRequestFulfilled represents a RandomnessRequestFulfilled event raised by the VorCoordinator contract.
type VorCoordinatorRandomnessRequestFulfilled struct {
	RequestId [32]byte
	Output    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessRequestFulfilled is a free log retrieval operation binding the contract event 0xa2e7a402243ebda4a69ceeb3dfb682943b7a9b3ac66d6eefa8db65894009611c.
//
// Solidity: event RandomnessRequestFulfilled(bytes32 requestId, uint256 output)
func (_VorCoordinator *VorCoordinatorFilterer) FilterRandomnessRequestFulfilled(opts *bind.FilterOpts) (*VorCoordinatorRandomnessRequestFulfilledIterator, error) {

	logs, sub, err := _VorCoordinator.contract.FilterLogs(opts, "RandomnessRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &VorCoordinatorRandomnessRequestFulfilledIterator{contract: _VorCoordinator.contract, event: "RandomnessRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequestFulfilled is a free log subscription operation binding the contract event 0xa2e7a402243ebda4a69ceeb3dfb682943b7a9b3ac66d6eefa8db65894009611c.
//
// Solidity: event RandomnessRequestFulfilled(bytes32 requestId, uint256 output)
func (_VorCoordinator *VorCoordinatorFilterer) WatchRandomnessRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VorCoordinatorRandomnessRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _VorCoordinator.contract.WatchLogs(opts, "RandomnessRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorCoordinatorRandomnessRequestFulfilled)
				if err := _VorCoordinator.contract.UnpackLog(event, "RandomnessRequestFulfilled", log); err != nil {
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

// ParseRandomnessRequestFulfilled is a log parse operation binding the contract event 0xa2e7a402243ebda4a69ceeb3dfb682943b7a9b3ac66d6eefa8db65894009611c.
//
// Solidity: event RandomnessRequestFulfilled(bytes32 requestId, uint256 output)
func (_VorCoordinator *VorCoordinatorFilterer) ParseRandomnessRequestFulfilled(log types.Log) (*VorCoordinatorRandomnessRequestFulfilled, error) {
	event := new(VorCoordinatorRandomnessRequestFulfilled)
	if err := _VorCoordinator.contract.UnpackLog(event, "RandomnessRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
