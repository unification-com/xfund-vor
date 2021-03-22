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

// VORCoordinatorABI is the input ABI used to generate the binding from.
const VORCoordinatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xfund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blockHashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ChangeFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasRefundedToProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"NewServiceAgreement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"RandomnessRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"SetGasTopUpLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"providerPays\",\"type\":\"bool\"}],\"name\":\"SetProviderPaysGas\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EXPECTED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPECTED_GAS_FIRST_FULFILMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRESEED_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROOF_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_KEY_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"callbacks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"callbackContract\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"randomnessFee\",\"type\":\"uint96\"},{\"internalType\":\"bytes32\",\"name\":\"seedAndBlockNum\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"consumerPreviousFulfillment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"fulfillRandomnessRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasTopUpLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"getProviderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalGasDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_consumerSeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePaid\",\"type\":\"uint256\"}],\"name\":\"randomnessRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"_publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_providerPaysGas\",\"type\":\"bool\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"serviceAgreements\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"vOROracle\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"providerPaysGas\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasTopUpLimit\",\"type\":\"uint256\"}],\"name\":\"setGasTopUpLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_providerPays\",\"type\":\"bool\"}],\"name\":\"setProviderPaysGas\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"topUpGas\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VORCoordinator is an auto generated Go binding around an Ethereum contract.
type VORCoordinator struct {
	VORCoordinatorCaller     // Read-only binding to the contract
	VORCoordinatorTransactor // Write-only binding to the contract
	VORCoordinatorFilterer   // Log filterer for contract events
}

// VORCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VORCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VORCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VORCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VORCoordinatorSession struct {
	Contract     *VORCoordinator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VORCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VORCoordinatorCallerSession struct {
	Contract *VORCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VORCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VORCoordinatorTransactorSession struct {
	Contract     *VORCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VORCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VORCoordinatorRaw struct {
	Contract *VORCoordinator // Generic contract binding to access the raw methods on
}

// VORCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VORCoordinatorCallerRaw struct {
	Contract *VORCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// VORCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VORCoordinatorTransactorRaw struct {
	Contract *VORCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVORCoordinator creates a new instance of VORCoordinator, bound to a specific deployed contract.
func NewVORCoordinator(address common.Address, backend bind.ContractBackend) (*VORCoordinator, error) {
	contract, err := bindVORCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VORCoordinator{VORCoordinatorCaller: VORCoordinatorCaller{contract: contract}, VORCoordinatorTransactor: VORCoordinatorTransactor{contract: contract}, VORCoordinatorFilterer: VORCoordinatorFilterer{contract: contract}}, nil
}

// NewVORCoordinatorCaller creates a new read-only instance of VORCoordinator, bound to a specific deployed contract.
func NewVORCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*VORCoordinatorCaller, error) {
	contract, err := bindVORCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorCaller{contract: contract}, nil
}

// NewVORCoordinatorTransactor creates a new write-only instance of VORCoordinator, bound to a specific deployed contract.
func NewVORCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VORCoordinatorTransactor, error) {
	contract, err := bindVORCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorTransactor{contract: contract}, nil
}

// NewVORCoordinatorFilterer creates a new log filterer instance of VORCoordinator, bound to a specific deployed contract.
func NewVORCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VORCoordinatorFilterer, error) {
	contract, err := bindVORCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorFilterer{contract: contract}, nil
}

// bindVORCoordinator binds a generic wrapper to an already deployed contract.
func bindVORCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VORCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VORCoordinator *VORCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VORCoordinator.Contract.VORCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VORCoordinator *VORCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORCoordinator.Contract.VORCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VORCoordinator *VORCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VORCoordinator.Contract.VORCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VORCoordinator *VORCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VORCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VORCoordinator *VORCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VORCoordinator *VORCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VORCoordinator.Contract.contract.Transact(opts, method, params...)
}

// EXPECTEDGAS is a free data retrieval call binding the contract method 0x610dfa53.
//
// Solidity: function EXPECTED_GAS() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) EXPECTEDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "EXPECTED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPECTEDGAS is a free data retrieval call binding the contract method 0x610dfa53.
//
// Solidity: function EXPECTED_GAS() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) EXPECTEDGAS() (*big.Int, error) {
	return _VORCoordinator.Contract.EXPECTEDGAS(&_VORCoordinator.CallOpts)
}

// EXPECTEDGAS is a free data retrieval call binding the contract method 0x610dfa53.
//
// Solidity: function EXPECTED_GAS() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) EXPECTEDGAS() (*big.Int, error) {
	return _VORCoordinator.Contract.EXPECTEDGAS(&_VORCoordinator.CallOpts)
}

// EXPECTEDGASFIRSTFULFILMENT is a free data retrieval call binding the contract method 0xa5a89f8e.
//
// Solidity: function EXPECTED_GAS_FIRST_FULFILMENT() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) EXPECTEDGASFIRSTFULFILMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "EXPECTED_GAS_FIRST_FULFILMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPECTEDGASFIRSTFULFILMENT is a free data retrieval call binding the contract method 0xa5a89f8e.
//
// Solidity: function EXPECTED_GAS_FIRST_FULFILMENT() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) EXPECTEDGASFIRSTFULFILMENT() (*big.Int, error) {
	return _VORCoordinator.Contract.EXPECTEDGASFIRSTFULFILMENT(&_VORCoordinator.CallOpts)
}

// EXPECTEDGASFIRSTFULFILMENT is a free data retrieval call binding the contract method 0xa5a89f8e.
//
// Solidity: function EXPECTED_GAS_FIRST_FULFILMENT() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) EXPECTEDGASFIRSTFULFILMENT() (*big.Int, error) {
	return _VORCoordinator.Contract.EXPECTEDGASFIRSTFULFILMENT(&_VORCoordinator.CallOpts)
}

// PRESEEDOFFSET is a free data retrieval call binding the contract method 0xb415f4f5.
//
// Solidity: function PRESEED_OFFSET() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) PRESEEDOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "PRESEED_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRESEEDOFFSET is a free data retrieval call binding the contract method 0xb415f4f5.
//
// Solidity: function PRESEED_OFFSET() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) PRESEEDOFFSET() (*big.Int, error) {
	return _VORCoordinator.Contract.PRESEEDOFFSET(&_VORCoordinator.CallOpts)
}

// PRESEEDOFFSET is a free data retrieval call binding the contract method 0xb415f4f5.
//
// Solidity: function PRESEED_OFFSET() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) PRESEEDOFFSET() (*big.Int, error) {
	return _VORCoordinator.Contract.PRESEEDOFFSET(&_VORCoordinator.CallOpts)
}

// PROOFLENGTH is a free data retrieval call binding the contract method 0xe911439c.
//
// Solidity: function PROOF_LENGTH() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) PROOFLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "PROOF_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROOFLENGTH is a free data retrieval call binding the contract method 0xe911439c.
//
// Solidity: function PROOF_LENGTH() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) PROOFLENGTH() (*big.Int, error) {
	return _VORCoordinator.Contract.PROOFLENGTH(&_VORCoordinator.CallOpts)
}

// PROOFLENGTH is a free data retrieval call binding the contract method 0xe911439c.
//
// Solidity: function PROOF_LENGTH() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) PROOFLENGTH() (*big.Int, error) {
	return _VORCoordinator.Contract.PROOFLENGTH(&_VORCoordinator.CallOpts)
}

// PUBLICKEYOFFSET is a free data retrieval call binding the contract method 0x8aa7927b.
//
// Solidity: function PUBLIC_KEY_OFFSET() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) PUBLICKEYOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "PUBLIC_KEY_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBLICKEYOFFSET is a free data retrieval call binding the contract method 0x8aa7927b.
//
// Solidity: function PUBLIC_KEY_OFFSET() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) PUBLICKEYOFFSET() (*big.Int, error) {
	return _VORCoordinator.Contract.PUBLICKEYOFFSET(&_VORCoordinator.CallOpts)
}

// PUBLICKEYOFFSET is a free data retrieval call binding the contract method 0x8aa7927b.
//
// Solidity: function PUBLIC_KEY_OFFSET() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) PUBLICKEYOFFSET() (*big.Int, error) {
	return _VORCoordinator.Contract.PUBLICKEYOFFSET(&_VORCoordinator.CallOpts)
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) view returns(address callbackContract, uint96 randomnessFee, bytes32 seedAndBlockNum)
func (_VORCoordinator *VORCoordinatorCaller) Callbacks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	SeedAndBlockNum  [32]byte
}, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "callbacks", arg0)

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
func (_VORCoordinator *VORCoordinatorSession) Callbacks(arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	SeedAndBlockNum  [32]byte
}, error) {
	return _VORCoordinator.Contract.Callbacks(&_VORCoordinator.CallOpts, arg0)
}

// Callbacks is a free data retrieval call binding the contract method 0x21f36509.
//
// Solidity: function callbacks(bytes32 ) view returns(address callbackContract, uint96 randomnessFee, bytes32 seedAndBlockNum)
func (_VORCoordinator *VORCoordinatorCallerSession) Callbacks(arg0 [32]byte) (struct {
	CallbackContract common.Address
	RandomnessFee    *big.Int
	SeedAndBlockNum  [32]byte
}, error) {
	return _VORCoordinator.Contract.Callbacks(&_VORCoordinator.CallOpts, arg0)
}

// ConsumerPreviousFulfillment is a free data retrieval call binding the contract method 0xa0fb3bd4.
//
// Solidity: function consumerPreviousFulfillment(address ) view returns(bool)
func (_VORCoordinator *VORCoordinatorCaller) ConsumerPreviousFulfillment(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "consumerPreviousFulfillment", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConsumerPreviousFulfillment is a free data retrieval call binding the contract method 0xa0fb3bd4.
//
// Solidity: function consumerPreviousFulfillment(address ) view returns(bool)
func (_VORCoordinator *VORCoordinatorSession) ConsumerPreviousFulfillment(arg0 common.Address) (bool, error) {
	return _VORCoordinator.Contract.ConsumerPreviousFulfillment(&_VORCoordinator.CallOpts, arg0)
}

// ConsumerPreviousFulfillment is a free data retrieval call binding the contract method 0xa0fb3bd4.
//
// Solidity: function consumerPreviousFulfillment(address ) view returns(bool)
func (_VORCoordinator *VORCoordinatorCallerSession) ConsumerPreviousFulfillment(arg0 common.Address) (bool, error) {
	return _VORCoordinator.Contract.ConsumerPreviousFulfillment(&_VORCoordinator.CallOpts, arg0)
}

// GetGasTopUpLimit is a free data retrieval call binding the contract method 0x8ac2a84d.
//
// Solidity: function getGasTopUpLimit() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) GetGasTopUpLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "getGasTopUpLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasTopUpLimit is a free data retrieval call binding the contract method 0x8ac2a84d.
//
// Solidity: function getGasTopUpLimit() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) GetGasTopUpLimit() (*big.Int, error) {
	return _VORCoordinator.Contract.GetGasTopUpLimit(&_VORCoordinator.CallOpts)
}

// GetGasTopUpLimit is a free data retrieval call binding the contract method 0x8ac2a84d.
//
// Solidity: function getGasTopUpLimit() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) GetGasTopUpLimit() (*big.Int, error) {
	return _VORCoordinator.Contract.GetGasTopUpLimit(&_VORCoordinator.CallOpts)
}

// GetProviderAddress is a free data retrieval call binding the contract method 0x9845fb9c.
//
// Solidity: function getProviderAddress(bytes32 _keyHash) view returns(address)
func (_VORCoordinator *VORCoordinatorCaller) GetProviderAddress(opts *bind.CallOpts, _keyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "getProviderAddress", _keyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProviderAddress is a free data retrieval call binding the contract method 0x9845fb9c.
//
// Solidity: function getProviderAddress(bytes32 _keyHash) view returns(address)
func (_VORCoordinator *VORCoordinatorSession) GetProviderAddress(_keyHash [32]byte) (common.Address, error) {
	return _VORCoordinator.Contract.GetProviderAddress(&_VORCoordinator.CallOpts, _keyHash)
}

// GetProviderAddress is a free data retrieval call binding the contract method 0x9845fb9c.
//
// Solidity: function getProviderAddress(bytes32 _keyHash) view returns(address)
func (_VORCoordinator *VORCoordinatorCallerSession) GetProviderAddress(_keyHash [32]byte) (common.Address, error) {
	return _VORCoordinator.Contract.GetProviderAddress(&_VORCoordinator.CallOpts, _keyHash)
}

// GetTotalGasDeposits is a free data retrieval call binding the contract method 0x6b7f0a15.
//
// Solidity: function getTotalGasDeposits() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) GetTotalGasDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "getTotalGasDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalGasDeposits is a free data retrieval call binding the contract method 0x6b7f0a15.
//
// Solidity: function getTotalGasDeposits() view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) GetTotalGasDeposits() (*big.Int, error) {
	return _VORCoordinator.Contract.GetTotalGasDeposits(&_VORCoordinator.CallOpts)
}

// GetTotalGasDeposits is a free data retrieval call binding the contract method 0x6b7f0a15.
//
// Solidity: function getTotalGasDeposits() view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) GetTotalGasDeposits() (*big.Int, error) {
	return _VORCoordinator.Contract.GetTotalGasDeposits(&_VORCoordinator.CallOpts)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) pure returns(bytes32)
func (_VORCoordinator *VORCoordinatorCaller) HashOfKey(opts *bind.CallOpts, _publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "hashOfKey", _publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) pure returns(bytes32)
func (_VORCoordinator *VORCoordinatorSession) HashOfKey(_publicKey [2]*big.Int) ([32]byte, error) {
	return _VORCoordinator.Contract.HashOfKey(&_VORCoordinator.CallOpts, _publicKey)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] _publicKey) pure returns(bytes32)
func (_VORCoordinator *VORCoordinatorCallerSession) HashOfKey(_publicKey [2]*big.Int) ([32]byte, error) {
	return _VORCoordinator.Contract.HashOfKey(&_VORCoordinator.CallOpts, _publicKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORCoordinator *VORCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORCoordinator *VORCoordinatorSession) Owner() (common.Address, error) {
	return _VORCoordinator.Contract.Owner(&_VORCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORCoordinator *VORCoordinatorCallerSession) Owner() (common.Address, error) {
	return _VORCoordinator.Contract.Owner(&_VORCoordinator.CallOpts)
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) view returns(address vOROracle, uint96 fee, bool providerPaysGas)
func (_VORCoordinator *VORCoordinatorCaller) ServiceAgreements(opts *bind.CallOpts, arg0 [32]byte) (struct {
	VOROracle       common.Address
	Fee             *big.Int
	ProviderPaysGas bool
}, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "serviceAgreements", arg0)

	outstruct := new(struct {
		VOROracle       common.Address
		Fee             *big.Int
		ProviderPaysGas bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VOROracle = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProviderPaysGas = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) view returns(address vOROracle, uint96 fee, bool providerPaysGas)
func (_VORCoordinator *VORCoordinatorSession) ServiceAgreements(arg0 [32]byte) (struct {
	VOROracle       common.Address
	Fee             *big.Int
	ProviderPaysGas bool
}, error) {
	return _VORCoordinator.Contract.ServiceAgreements(&_VORCoordinator.CallOpts, arg0)
}

// ServiceAgreements is a free data retrieval call binding the contract method 0x75d35070.
//
// Solidity: function serviceAgreements(bytes32 ) view returns(address vOROracle, uint96 fee, bool providerPaysGas)
func (_VORCoordinator *VORCoordinatorCallerSession) ServiceAgreements(arg0 [32]byte) (struct {
	VOROracle       common.Address
	Fee             *big.Int
	ProviderPaysGas bool
}, error) {
	return _VORCoordinator.Contract.ServiceAgreements(&_VORCoordinator.CallOpts, arg0)
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) view returns(uint256)
func (_VORCoordinator *VORCoordinatorCaller) WithdrawableTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VORCoordinator.contract.Call(opts, &out, "withdrawableTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) view returns(uint256)
func (_VORCoordinator *VORCoordinatorSession) WithdrawableTokens(arg0 common.Address) (*big.Int, error) {
	return _VORCoordinator.Contract.WithdrawableTokens(&_VORCoordinator.CallOpts, arg0)
}

// WithdrawableTokens is a free data retrieval call binding the contract method 0x006f6ad0.
//
// Solidity: function withdrawableTokens(address ) view returns(uint256)
func (_VORCoordinator *VORCoordinatorCallerSession) WithdrawableTokens(arg0 common.Address) (*big.Int, error) {
	return _VORCoordinator.Contract.WithdrawableTokens(&_VORCoordinator.CallOpts, arg0)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x81f1e046.
//
// Solidity: function changeFee(uint256[2] _publicProvingKey, uint256 _fee) returns()
func (_VORCoordinator *VORCoordinatorTransactor) ChangeFee(opts *bind.TransactOpts, _publicProvingKey [2]*big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "changeFee", _publicProvingKey, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x81f1e046.
//
// Solidity: function changeFee(uint256[2] _publicProvingKey, uint256 _fee) returns()
func (_VORCoordinator *VORCoordinatorSession) ChangeFee(_publicProvingKey [2]*big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.ChangeFee(&_VORCoordinator.TransactOpts, _publicProvingKey, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x81f1e046.
//
// Solidity: function changeFee(uint256[2] _publicProvingKey, uint256 _fee) returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) ChangeFee(_publicProvingKey [2]*big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.ChangeFee(&_VORCoordinator.TransactOpts, _publicProvingKey, _fee)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns()
func (_VORCoordinator *VORCoordinatorTransactor) FulfillRandomnessRequest(opts *bind.TransactOpts, _proof []byte) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "fulfillRandomnessRequest", _proof)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns()
func (_VORCoordinator *VORCoordinatorSession) FulfillRandomnessRequest(_proof []byte) (*types.Transaction, error) {
	return _VORCoordinator.Contract.FulfillRandomnessRequest(&_VORCoordinator.TransactOpts, _proof)
}

// FulfillRandomnessRequest is a paid mutator transaction binding the contract method 0x5e1c1059.
//
// Solidity: function fulfillRandomnessRequest(bytes _proof) returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) FulfillRandomnessRequest(_proof []byte) (*types.Transaction, error) {
	return _VORCoordinator.Contract.FulfillRandomnessRequest(&_VORCoordinator.TransactOpts, _proof)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VORCoordinator *VORCoordinatorTransactor) RandomnessRequest(opts *bind.TransactOpts, _keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "randomnessRequest", _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VORCoordinator *VORCoordinatorSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.RandomnessRequest(&_VORCoordinator.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.RandomnessRequest(&_VORCoordinator.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xb59b4feb.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey, bool _providerPaysGas) returns()
func (_VORCoordinator *VORCoordinatorTransactor) RegisterProvingKey(opts *bind.TransactOpts, _fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int, _providerPaysGas bool) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "registerProvingKey", _fee, _oracle, _publicProvingKey, _providerPaysGas)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xb59b4feb.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey, bool _providerPaysGas) returns()
func (_VORCoordinator *VORCoordinatorSession) RegisterProvingKey(_fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int, _providerPaysGas bool) (*types.Transaction, error) {
	return _VORCoordinator.Contract.RegisterProvingKey(&_VORCoordinator.TransactOpts, _fee, _oracle, _publicProvingKey, _providerPaysGas)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0xb59b4feb.
//
// Solidity: function registerProvingKey(uint256 _fee, address _oracle, uint256[2] _publicProvingKey, bool _providerPaysGas) returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) RegisterProvingKey(_fee *big.Int, _oracle common.Address, _publicProvingKey [2]*big.Int, _providerPaysGas bool) (*types.Transaction, error) {
	return _VORCoordinator.Contract.RegisterProvingKey(&_VORCoordinator.TransactOpts, _fee, _oracle, _publicProvingKey, _providerPaysGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORCoordinator *VORCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORCoordinator *VORCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VORCoordinator.Contract.RenounceOwnership(&_VORCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VORCoordinator.Contract.RenounceOwnership(&_VORCoordinator.TransactOpts)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _gasTopUpLimit) returns(bool success)
func (_VORCoordinator *VORCoordinatorTransactor) SetGasTopUpLimit(opts *bind.TransactOpts, _gasTopUpLimit *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "setGasTopUpLimit", _gasTopUpLimit)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _gasTopUpLimit) returns(bool success)
func (_VORCoordinator *VORCoordinatorSession) SetGasTopUpLimit(_gasTopUpLimit *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.SetGasTopUpLimit(&_VORCoordinator.TransactOpts, _gasTopUpLimit)
}

// SetGasTopUpLimit is a paid mutator transaction binding the contract method 0x0f3a85d8.
//
// Solidity: function setGasTopUpLimit(uint256 _gasTopUpLimit) returns(bool success)
func (_VORCoordinator *VORCoordinatorTransactorSession) SetGasTopUpLimit(_gasTopUpLimit *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.SetGasTopUpLimit(&_VORCoordinator.TransactOpts, _gasTopUpLimit)
}

// SetProviderPaysGas is a paid mutator transaction binding the contract method 0x7c4f46dd.
//
// Solidity: function setProviderPaysGas(uint256[2] _publicProvingKey, bool _providerPays) returns(bool success)
func (_VORCoordinator *VORCoordinatorTransactor) SetProviderPaysGas(opts *bind.TransactOpts, _publicProvingKey [2]*big.Int, _providerPays bool) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "setProviderPaysGas", _publicProvingKey, _providerPays)
}

// SetProviderPaysGas is a paid mutator transaction binding the contract method 0x7c4f46dd.
//
// Solidity: function setProviderPaysGas(uint256[2] _publicProvingKey, bool _providerPays) returns(bool success)
func (_VORCoordinator *VORCoordinatorSession) SetProviderPaysGas(_publicProvingKey [2]*big.Int, _providerPays bool) (*types.Transaction, error) {
	return _VORCoordinator.Contract.SetProviderPaysGas(&_VORCoordinator.TransactOpts, _publicProvingKey, _providerPays)
}

// SetProviderPaysGas is a paid mutator transaction binding the contract method 0x7c4f46dd.
//
// Solidity: function setProviderPaysGas(uint256[2] _publicProvingKey, bool _providerPays) returns(bool success)
func (_VORCoordinator *VORCoordinatorTransactorSession) SetProviderPaysGas(_publicProvingKey [2]*big.Int, _providerPays bool) (*types.Transaction, error) {
	return _VORCoordinator.Contract.SetProviderPaysGas(&_VORCoordinator.TransactOpts, _publicProvingKey, _providerPays)
}

// TopUpGas is a paid mutator transaction binding the contract method 0x9234bcc0.
//
// Solidity: function topUpGas(address _provider) payable returns(bool success)
func (_VORCoordinator *VORCoordinatorTransactor) TopUpGas(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "topUpGas", _provider)
}

// TopUpGas is a paid mutator transaction binding the contract method 0x9234bcc0.
//
// Solidity: function topUpGas(address _provider) payable returns(bool success)
func (_VORCoordinator *VORCoordinatorSession) TopUpGas(_provider common.Address) (*types.Transaction, error) {
	return _VORCoordinator.Contract.TopUpGas(&_VORCoordinator.TransactOpts, _provider)
}

// TopUpGas is a paid mutator transaction binding the contract method 0x9234bcc0.
//
// Solidity: function topUpGas(address _provider) payable returns(bool success)
func (_VORCoordinator *VORCoordinatorTransactorSession) TopUpGas(_provider common.Address) (*types.Transaction, error) {
	return _VORCoordinator.Contract.TopUpGas(&_VORCoordinator.TransactOpts, _provider)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORCoordinator *VORCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORCoordinator *VORCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VORCoordinator.Contract.TransferOwnership(&_VORCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VORCoordinator.Contract.TransferOwnership(&_VORCoordinator.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VORCoordinator *VORCoordinatorTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.contract.Transact(opts, "withdraw", _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VORCoordinator *VORCoordinatorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.Withdraw(&_VORCoordinator.TransactOpts, _recipient, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _recipient, uint256 _amount) returns()
func (_VORCoordinator *VORCoordinatorTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VORCoordinator.Contract.Withdraw(&_VORCoordinator.TransactOpts, _recipient, _amount)
}

// VORCoordinatorChangeFeeIterator is returned from FilterChangeFee and is used to iterate over the raw logs and unpacked data for ChangeFee events raised by the VORCoordinator contract.
type VORCoordinatorChangeFeeIterator struct {
	Event *VORCoordinatorChangeFee // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorChangeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorChangeFee)
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
		it.Event = new(VORCoordinatorChangeFee)
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
func (it *VORCoordinatorChangeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorChangeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorChangeFee represents a ChangeFee event raised by the VORCoordinator contract.
type VORCoordinatorChangeFee struct {
	KeyHash [32]byte
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangeFee is a free log retrieval operation binding the contract event 0xfc6ca2918cb95fbfc07e4567da6387d5e4209b4ced636bbe450c09895d447526.
//
// Solidity: event ChangeFee(bytes32 keyHash, uint256 fee)
func (_VORCoordinator *VORCoordinatorFilterer) FilterChangeFee(opts *bind.FilterOpts) (*VORCoordinatorChangeFeeIterator, error) {

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorChangeFeeIterator{contract: _VORCoordinator.contract, event: "ChangeFee", logs: logs, sub: sub}, nil
}

// WatchChangeFee is a free log subscription operation binding the contract event 0xfc6ca2918cb95fbfc07e4567da6387d5e4209b4ced636bbe450c09895d447526.
//
// Solidity: event ChangeFee(bytes32 keyHash, uint256 fee)
func (_VORCoordinator *VORCoordinatorFilterer) WatchChangeFee(opts *bind.WatchOpts, sink chan<- *VORCoordinatorChangeFee) (event.Subscription, error) {

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "ChangeFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorChangeFee)
				if err := _VORCoordinator.contract.UnpackLog(event, "ChangeFee", log); err != nil {
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
func (_VORCoordinator *VORCoordinatorFilterer) ParseChangeFee(log types.Log) (*VORCoordinatorChangeFee, error) {
	event := new(VORCoordinatorChangeFee)
	if err := _VORCoordinator.contract.UnpackLog(event, "ChangeFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorGasRefundedToProviderIterator is returned from FilterGasRefundedToProvider and is used to iterate over the raw logs and unpacked data for GasRefundedToProvider events raised by the VORCoordinator contract.
type VORCoordinatorGasRefundedToProviderIterator struct {
	Event *VORCoordinatorGasRefundedToProvider // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorGasRefundedToProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorGasRefundedToProvider)
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
		it.Event = new(VORCoordinatorGasRefundedToProvider)
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
func (it *VORCoordinatorGasRefundedToProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorGasRefundedToProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorGasRefundedToProvider represents a GasRefundedToProvider event raised by the VORCoordinator contract.
type VORCoordinatorGasRefundedToProvider struct {
	Consumer common.Address
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasRefundedToProvider is a free log retrieval operation binding the contract event 0xc5a3f3f707348b27fcecf895e59365d724538a5ba59bd84f5a216be628ac7624.
//
// Solidity: event GasRefundedToProvider(address indexed consumer, address indexed provider, uint256 amount)
func (_VORCoordinator *VORCoordinatorFilterer) FilterGasRefundedToProvider(opts *bind.FilterOpts, consumer []common.Address, provider []common.Address) (*VORCoordinatorGasRefundedToProviderIterator, error) {

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "GasRefundedToProvider", consumerRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorGasRefundedToProviderIterator{contract: _VORCoordinator.contract, event: "GasRefundedToProvider", logs: logs, sub: sub}, nil
}

// WatchGasRefundedToProvider is a free log subscription operation binding the contract event 0xc5a3f3f707348b27fcecf895e59365d724538a5ba59bd84f5a216be628ac7624.
//
// Solidity: event GasRefundedToProvider(address indexed consumer, address indexed provider, uint256 amount)
func (_VORCoordinator *VORCoordinatorFilterer) WatchGasRefundedToProvider(opts *bind.WatchOpts, sink chan<- *VORCoordinatorGasRefundedToProvider, consumer []common.Address, provider []common.Address) (event.Subscription, error) {

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "GasRefundedToProvider", consumerRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorGasRefundedToProvider)
				if err := _VORCoordinator.contract.UnpackLog(event, "GasRefundedToProvider", log); err != nil {
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

// ParseGasRefundedToProvider is a log parse operation binding the contract event 0xc5a3f3f707348b27fcecf895e59365d724538a5ba59bd84f5a216be628ac7624.
//
// Solidity: event GasRefundedToProvider(address indexed consumer, address indexed provider, uint256 amount)
func (_VORCoordinator *VORCoordinatorFilterer) ParseGasRefundedToProvider(log types.Log) (*VORCoordinatorGasRefundedToProvider, error) {
	event := new(VORCoordinatorGasRefundedToProvider)
	if err := _VORCoordinator.contract.UnpackLog(event, "GasRefundedToProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorGasToppedUpIterator is returned from FilterGasToppedUp and is used to iterate over the raw logs and unpacked data for GasToppedUp events raised by the VORCoordinator contract.
type VORCoordinatorGasToppedUpIterator struct {
	Event *VORCoordinatorGasToppedUp // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorGasToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorGasToppedUp)
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
		it.Event = new(VORCoordinatorGasToppedUp)
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
func (it *VORCoordinatorGasToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorGasToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorGasToppedUp represents a GasToppedUp event raised by the VORCoordinator contract.
type VORCoordinatorGasToppedUp struct {
	Consumer common.Address
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasToppedUp is a free log retrieval operation binding the contract event 0x91741e86d5a4eecb6b400c9552980cc853bc8690a11d78e496e670226d2352f7.
//
// Solidity: event GasToppedUp(address indexed consumer, address indexed provider, uint256 amount)
func (_VORCoordinator *VORCoordinatorFilterer) FilterGasToppedUp(opts *bind.FilterOpts, consumer []common.Address, provider []common.Address) (*VORCoordinatorGasToppedUpIterator, error) {

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "GasToppedUp", consumerRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorGasToppedUpIterator{contract: _VORCoordinator.contract, event: "GasToppedUp", logs: logs, sub: sub}, nil
}

// WatchGasToppedUp is a free log subscription operation binding the contract event 0x91741e86d5a4eecb6b400c9552980cc853bc8690a11d78e496e670226d2352f7.
//
// Solidity: event GasToppedUp(address indexed consumer, address indexed provider, uint256 amount)
func (_VORCoordinator *VORCoordinatorFilterer) WatchGasToppedUp(opts *bind.WatchOpts, sink chan<- *VORCoordinatorGasToppedUp, consumer []common.Address, provider []common.Address) (event.Subscription, error) {

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "GasToppedUp", consumerRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorGasToppedUp)
				if err := _VORCoordinator.contract.UnpackLog(event, "GasToppedUp", log); err != nil {
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

// ParseGasToppedUp is a log parse operation binding the contract event 0x91741e86d5a4eecb6b400c9552980cc853bc8690a11d78e496e670226d2352f7.
//
// Solidity: event GasToppedUp(address indexed consumer, address indexed provider, uint256 amount)
func (_VORCoordinator *VORCoordinatorFilterer) ParseGasToppedUp(log types.Log) (*VORCoordinatorGasToppedUp, error) {
	event := new(VORCoordinatorGasToppedUp)
	if err := _VORCoordinator.contract.UnpackLog(event, "GasToppedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorNewServiceAgreementIterator is returned from FilterNewServiceAgreement and is used to iterate over the raw logs and unpacked data for NewServiceAgreement events raised by the VORCoordinator contract.
type VORCoordinatorNewServiceAgreementIterator struct {
	Event *VORCoordinatorNewServiceAgreement // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorNewServiceAgreementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorNewServiceAgreement)
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
		it.Event = new(VORCoordinatorNewServiceAgreement)
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
func (it *VORCoordinatorNewServiceAgreementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorNewServiceAgreementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorNewServiceAgreement represents a NewServiceAgreement event raised by the VORCoordinator contract.
type VORCoordinatorNewServiceAgreement struct {
	KeyHash [32]byte
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewServiceAgreement is a free log retrieval operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VORCoordinator *VORCoordinatorFilterer) FilterNewServiceAgreement(opts *bind.FilterOpts) (*VORCoordinatorNewServiceAgreementIterator, error) {

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "NewServiceAgreement")
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorNewServiceAgreementIterator{contract: _VORCoordinator.contract, event: "NewServiceAgreement", logs: logs, sub: sub}, nil
}

// WatchNewServiceAgreement is a free log subscription operation binding the contract event 0xae189157e0628c1e62315e9179156e1ea10e90e9c15060002f7021e907dc2cfe.
//
// Solidity: event NewServiceAgreement(bytes32 keyHash, uint256 fee)
func (_VORCoordinator *VORCoordinatorFilterer) WatchNewServiceAgreement(opts *bind.WatchOpts, sink chan<- *VORCoordinatorNewServiceAgreement) (event.Subscription, error) {

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "NewServiceAgreement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorNewServiceAgreement)
				if err := _VORCoordinator.contract.UnpackLog(event, "NewServiceAgreement", log); err != nil {
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
func (_VORCoordinator *VORCoordinatorFilterer) ParseNewServiceAgreement(log types.Log) (*VORCoordinatorNewServiceAgreement, error) {
	event := new(VORCoordinatorNewServiceAgreement)
	if err := _VORCoordinator.contract.UnpackLog(event, "NewServiceAgreement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VORCoordinator contract.
type VORCoordinatorOwnershipTransferredIterator struct {
	Event *VORCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorOwnershipTransferred)
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
		it.Event = new(VORCoordinatorOwnershipTransferred)
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
func (it *VORCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the VORCoordinator contract.
type VORCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORCoordinator *VORCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VORCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorOwnershipTransferredIterator{contract: _VORCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORCoordinator *VORCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VORCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorOwnershipTransferred)
				if err := _VORCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VORCoordinator *VORCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*VORCoordinatorOwnershipTransferred, error) {
	event := new(VORCoordinatorOwnershipTransferred)
	if err := _VORCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorRandomnessRequestIterator is returned from FilterRandomnessRequest and is used to iterate over the raw logs and unpacked data for RandomnessRequest events raised by the VORCoordinator contract.
type VORCoordinatorRandomnessRequestIterator struct {
	Event *VORCoordinatorRandomnessRequest // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorRandomnessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorRandomnessRequest)
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
		it.Event = new(VORCoordinatorRandomnessRequest)
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
func (it *VORCoordinatorRandomnessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorRandomnessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorRandomnessRequest represents a RandomnessRequest event raised by the VORCoordinator contract.
type VORCoordinatorRandomnessRequest struct {
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
func (_VORCoordinator *VORCoordinatorFilterer) FilterRandomnessRequest(opts *bind.FilterOpts) (*VORCoordinatorRandomnessRequestIterator, error) {

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorRandomnessRequestIterator{contract: _VORCoordinator.contract, event: "RandomnessRequest", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequest is a free log subscription operation binding the contract event 0xebb37373bb11123e38f964627878b02c247f92f3913df7cf3f270b5222c8d2be.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)
func (_VORCoordinator *VORCoordinatorFilterer) WatchRandomnessRequest(opts *bind.WatchOpts, sink chan<- *VORCoordinatorRandomnessRequest) (event.Subscription, error) {

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorRandomnessRequest)
				if err := _VORCoordinator.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
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
func (_VORCoordinator *VORCoordinatorFilterer) ParseRandomnessRequest(log types.Log) (*VORCoordinatorRandomnessRequest, error) {
	event := new(VORCoordinatorRandomnessRequest)
	if err := _VORCoordinator.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorRandomnessRequestFulfilledIterator is returned from FilterRandomnessRequestFulfilled and is used to iterate over the raw logs and unpacked data for RandomnessRequestFulfilled events raised by the VORCoordinator contract.
type VORCoordinatorRandomnessRequestFulfilledIterator struct {
	Event *VORCoordinatorRandomnessRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorRandomnessRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorRandomnessRequestFulfilled)
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
		it.Event = new(VORCoordinatorRandomnessRequestFulfilled)
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
func (it *VORCoordinatorRandomnessRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorRandomnessRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorRandomnessRequestFulfilled represents a RandomnessRequestFulfilled event raised by the VORCoordinator contract.
type VORCoordinatorRandomnessRequestFulfilled struct {
	RequestId [32]byte
	Output    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRandomnessRequestFulfilled is a free log retrieval operation binding the contract event 0xa2e7a402243ebda4a69ceeb3dfb682943b7a9b3ac66d6eefa8db65894009611c.
//
// Solidity: event RandomnessRequestFulfilled(bytes32 requestId, uint256 output)
func (_VORCoordinator *VORCoordinatorFilterer) FilterRandomnessRequestFulfilled(opts *bind.FilterOpts) (*VORCoordinatorRandomnessRequestFulfilledIterator, error) {

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "RandomnessRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorRandomnessRequestFulfilledIterator{contract: _VORCoordinator.contract, event: "RandomnessRequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequestFulfilled is a free log subscription operation binding the contract event 0xa2e7a402243ebda4a69ceeb3dfb682943b7a9b3ac66d6eefa8db65894009611c.
//
// Solidity: event RandomnessRequestFulfilled(bytes32 requestId, uint256 output)
func (_VORCoordinator *VORCoordinatorFilterer) WatchRandomnessRequestFulfilled(opts *bind.WatchOpts, sink chan<- *VORCoordinatorRandomnessRequestFulfilled) (event.Subscription, error) {

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "RandomnessRequestFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorRandomnessRequestFulfilled)
				if err := _VORCoordinator.contract.UnpackLog(event, "RandomnessRequestFulfilled", log); err != nil {
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
func (_VORCoordinator *VORCoordinatorFilterer) ParseRandomnessRequestFulfilled(log types.Log) (*VORCoordinatorRandomnessRequestFulfilled, error) {
	event := new(VORCoordinatorRandomnessRequestFulfilled)
	if err := _VORCoordinator.contract.UnpackLog(event, "RandomnessRequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorSetGasTopUpLimitIterator is returned from FilterSetGasTopUpLimit and is used to iterate over the raw logs and unpacked data for SetGasTopUpLimit events raised by the VORCoordinator contract.
type VORCoordinatorSetGasTopUpLimitIterator struct {
	Event *VORCoordinatorSetGasTopUpLimit // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorSetGasTopUpLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorSetGasTopUpLimit)
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
		it.Event = new(VORCoordinatorSetGasTopUpLimit)
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
func (it *VORCoordinatorSetGasTopUpLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorSetGasTopUpLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorSetGasTopUpLimit represents a SetGasTopUpLimit event raised by the VORCoordinator contract.
type VORCoordinatorSetGasTopUpLimit struct {
	Sender   common.Address
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGasTopUpLimit is a free log retrieval operation binding the contract event 0x35ab6203306e69cd89de0e2e1416a00fec248268e57b99e4d2638e3b14f277f1.
//
// Solidity: event SetGasTopUpLimit(address indexed sender, uint256 oldLimit, uint256 newLimit)
func (_VORCoordinator *VORCoordinatorFilterer) FilterSetGasTopUpLimit(opts *bind.FilterOpts, sender []common.Address) (*VORCoordinatorSetGasTopUpLimitIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "SetGasTopUpLimit", senderRule)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorSetGasTopUpLimitIterator{contract: _VORCoordinator.contract, event: "SetGasTopUpLimit", logs: logs, sub: sub}, nil
}

// WatchSetGasTopUpLimit is a free log subscription operation binding the contract event 0x35ab6203306e69cd89de0e2e1416a00fec248268e57b99e4d2638e3b14f277f1.
//
// Solidity: event SetGasTopUpLimit(address indexed sender, uint256 oldLimit, uint256 newLimit)
func (_VORCoordinator *VORCoordinatorFilterer) WatchSetGasTopUpLimit(opts *bind.WatchOpts, sink chan<- *VORCoordinatorSetGasTopUpLimit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "SetGasTopUpLimit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorSetGasTopUpLimit)
				if err := _VORCoordinator.contract.UnpackLog(event, "SetGasTopUpLimit", log); err != nil {
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

// ParseSetGasTopUpLimit is a log parse operation binding the contract event 0x35ab6203306e69cd89de0e2e1416a00fec248268e57b99e4d2638e3b14f277f1.
//
// Solidity: event SetGasTopUpLimit(address indexed sender, uint256 oldLimit, uint256 newLimit)
func (_VORCoordinator *VORCoordinatorFilterer) ParseSetGasTopUpLimit(log types.Log) (*VORCoordinatorSetGasTopUpLimit, error) {
	event := new(VORCoordinatorSetGasTopUpLimit)
	if err := _VORCoordinator.contract.UnpackLog(event, "SetGasTopUpLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORCoordinatorSetProviderPaysGasIterator is returned from FilterSetProviderPaysGas and is used to iterate over the raw logs and unpacked data for SetProviderPaysGas events raised by the VORCoordinator contract.
type VORCoordinatorSetProviderPaysGasIterator struct {
	Event *VORCoordinatorSetProviderPaysGas // Event containing the contract specifics and raw log

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
func (it *VORCoordinatorSetProviderPaysGasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORCoordinatorSetProviderPaysGas)
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
		it.Event = new(VORCoordinatorSetProviderPaysGas)
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
func (it *VORCoordinatorSetProviderPaysGasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORCoordinatorSetProviderPaysGasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORCoordinatorSetProviderPaysGas represents a SetProviderPaysGas event raised by the VORCoordinator contract.
type VORCoordinatorSetProviderPaysGas struct {
	Provider     common.Address
	ProviderPays bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetProviderPaysGas is a free log retrieval operation binding the contract event 0xbd14da99da538efaf54d5467de2c08e022bd7054a36953bffe5854d6c6c4ab3f.
//
// Solidity: event SetProviderPaysGas(address indexed provider, bool providerPays)
func (_VORCoordinator *VORCoordinatorFilterer) FilterSetProviderPaysGas(opts *bind.FilterOpts, provider []common.Address) (*VORCoordinatorSetProviderPaysGasIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _VORCoordinator.contract.FilterLogs(opts, "SetProviderPaysGas", providerRule)
	if err != nil {
		return nil, err
	}
	return &VORCoordinatorSetProviderPaysGasIterator{contract: _VORCoordinator.contract, event: "SetProviderPaysGas", logs: logs, sub: sub}, nil
}

// WatchSetProviderPaysGas is a free log subscription operation binding the contract event 0xbd14da99da538efaf54d5467de2c08e022bd7054a36953bffe5854d6c6c4ab3f.
//
// Solidity: event SetProviderPaysGas(address indexed provider, bool providerPays)
func (_VORCoordinator *VORCoordinatorFilterer) WatchSetProviderPaysGas(opts *bind.WatchOpts, sink chan<- *VORCoordinatorSetProviderPaysGas, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _VORCoordinator.contract.WatchLogs(opts, "SetProviderPaysGas", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORCoordinatorSetProviderPaysGas)
				if err := _VORCoordinator.contract.UnpackLog(event, "SetProviderPaysGas", log); err != nil {
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

// ParseSetProviderPaysGas is a log parse operation binding the contract event 0xbd14da99da538efaf54d5467de2c08e022bd7054a36953bffe5854d6c6c4ab3f.
//
// Solidity: event SetProviderPaysGas(address indexed provider, bool providerPays)
func (_VORCoordinator *VORCoordinatorFilterer) ParseSetProviderPaysGas(log types.Log) (*VORCoordinatorSetProviderPaysGas, error) {
	event := new(VORCoordinatorSetProviderPaysGas)
	if err := _VORCoordinator.contract.UnpackLog(event, "SetProviderPaysGas", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
