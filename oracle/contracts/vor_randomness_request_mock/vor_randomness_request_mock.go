// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vor_randomness_request_mock

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

// VORRandomnessRequestMockABI is the input ABI used to generate the binding from.
const VORRandomnessRequestMockABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"RandomnessRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_consumerSeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePaid\",\"type\":\"uint256\"}],\"name\":\"randomnessRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VORRandomnessRequestMock is an auto generated Go binding around an Ethereum contract.
type VORRandomnessRequestMock struct {
	VORRandomnessRequestMockCaller     // Read-only binding to the contract
	VORRandomnessRequestMockTransactor // Write-only binding to the contract
	VORRandomnessRequestMockFilterer   // Log filterer for contract events
}

// VORRandomnessRequestMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type VORRandomnessRequestMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORRandomnessRequestMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VORRandomnessRequestMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORRandomnessRequestMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VORRandomnessRequestMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VORRandomnessRequestMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VORRandomnessRequestMockSession struct {
	Contract     *VORRandomnessRequestMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VORRandomnessRequestMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VORRandomnessRequestMockCallerSession struct {
	Contract *VORRandomnessRequestMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// VORRandomnessRequestMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VORRandomnessRequestMockTransactorSession struct {
	Contract     *VORRandomnessRequestMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// VORRandomnessRequestMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type VORRandomnessRequestMockRaw struct {
	Contract *VORRandomnessRequestMock // Generic contract binding to access the raw methods on
}

// VORRandomnessRequestMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VORRandomnessRequestMockCallerRaw struct {
	Contract *VORRandomnessRequestMockCaller // Generic read-only contract binding to access the raw methods on
}

// VORRandomnessRequestMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VORRandomnessRequestMockTransactorRaw struct {
	Contract *VORRandomnessRequestMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVORRandomnessRequestMock creates a new instance of VORRandomnessRequestMock, bound to a specific deployed contract.
func NewVORRandomnessRequestMock(address common.Address, backend bind.ContractBackend) (*VORRandomnessRequestMock, error) {
	contract, err := bindVORRandomnessRequestMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMock{VORRandomnessRequestMockCaller: VORRandomnessRequestMockCaller{contract: contract}, VORRandomnessRequestMockTransactor: VORRandomnessRequestMockTransactor{contract: contract}, VORRandomnessRequestMockFilterer: VORRandomnessRequestMockFilterer{contract: contract}}, nil
}

// NewVORRandomnessRequestMockCaller creates a new read-only instance of VORRandomnessRequestMock, bound to a specific deployed contract.
func NewVORRandomnessRequestMockCaller(address common.Address, caller bind.ContractCaller) (*VORRandomnessRequestMockCaller, error) {
	contract, err := bindVORRandomnessRequestMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockCaller{contract: contract}, nil
}

// NewVORRandomnessRequestMockTransactor creates a new write-only instance of VORRandomnessRequestMock, bound to a specific deployed contract.
func NewVORRandomnessRequestMockTransactor(address common.Address, transactor bind.ContractTransactor) (*VORRandomnessRequestMockTransactor, error) {
	contract, err := bindVORRandomnessRequestMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockTransactor{contract: contract}, nil
}

// NewVORRandomnessRequestMockFilterer creates a new log filterer instance of VORRandomnessRequestMock, bound to a specific deployed contract.
func NewVORRandomnessRequestMockFilterer(address common.Address, filterer bind.ContractFilterer) (*VORRandomnessRequestMockFilterer, error) {
	contract, err := bindVORRandomnessRequestMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockFilterer{contract: contract}, nil
}

// bindVORRandomnessRequestMock binds a generic wrapper to an already deployed contract.
func bindVORRandomnessRequestMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VORRandomnessRequestMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VORRandomnessRequestMock *VORRandomnessRequestMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VORRandomnessRequestMock.Contract.VORRandomnessRequestMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VORRandomnessRequestMock *VORRandomnessRequestMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.VORRandomnessRequestMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VORRandomnessRequestMock *VORRandomnessRequestMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.VORRandomnessRequestMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VORRandomnessRequestMock *VORRandomnessRequestMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VORRandomnessRequestMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORRandomnessRequestMock *VORRandomnessRequestMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VORRandomnessRequestMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORRandomnessRequestMock *VORRandomnessRequestMockSession) Owner() (common.Address, error) {
	return _VORRandomnessRequestMock.Contract.Owner(&_VORRandomnessRequestMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VORRandomnessRequestMock *VORRandomnessRequestMockCallerSession) Owner() (common.Address, error) {
	return _VORRandomnessRequestMock.Contract.Owner(&_VORRandomnessRequestMock.CallOpts)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactor) RandomnessRequest(opts *bind.TransactOpts, _keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.contract.Transact(opts, "randomnessRequest", _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.RandomnessRequest(&_VORRandomnessRequestMock.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactorSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.RandomnessRequest(&_VORRandomnessRequestMock.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockSession) RenounceOwnership() (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.RenounceOwnership(&_VORRandomnessRequestMock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.RenounceOwnership(&_VORRandomnessRequestMock.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.TransferOwnership(&_VORRandomnessRequestMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VORRandomnessRequestMock *VORRandomnessRequestMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VORRandomnessRequestMock.Contract.TransferOwnership(&_VORRandomnessRequestMock.TransactOpts, newOwner)
}

// VORRandomnessRequestMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VORRandomnessRequestMock contract.
type VORRandomnessRequestMockOwnershipTransferredIterator struct {
	Event *VORRandomnessRequestMockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VORRandomnessRequestMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORRandomnessRequestMockOwnershipTransferred)
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
		it.Event = new(VORRandomnessRequestMockOwnershipTransferred)
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
func (it *VORRandomnessRequestMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORRandomnessRequestMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORRandomnessRequestMockOwnershipTransferred represents a OwnershipTransferred event raised by the VORRandomnessRequestMock contract.
type VORRandomnessRequestMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORRandomnessRequestMock *VORRandomnessRequestMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VORRandomnessRequestMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VORRandomnessRequestMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockOwnershipTransferredIterator{contract: _VORRandomnessRequestMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VORRandomnessRequestMock *VORRandomnessRequestMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VORRandomnessRequestMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VORRandomnessRequestMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORRandomnessRequestMockOwnershipTransferred)
				if err := _VORRandomnessRequestMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VORRandomnessRequestMock *VORRandomnessRequestMockFilterer) ParseOwnershipTransferred(log types.Log) (*VORRandomnessRequestMockOwnershipTransferred, error) {
	event := new(VORRandomnessRequestMockOwnershipTransferred)
	if err := _VORRandomnessRequestMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VORRandomnessRequestMockRandomnessRequestIterator is returned from FilterRandomnessRequest and is used to iterate over the raw logs and unpacked data for RandomnessRequest events raised by the VORRandomnessRequestMock contract.
type VORRandomnessRequestMockRandomnessRequestIterator struct {
	Event *VORRandomnessRequestMockRandomnessRequest // Event containing the contract specifics and raw log

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
func (it *VORRandomnessRequestMockRandomnessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VORRandomnessRequestMockRandomnessRequest)
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
		it.Event = new(VORRandomnessRequestMockRandomnessRequest)
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
func (it *VORRandomnessRequestMockRandomnessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VORRandomnessRequestMockRandomnessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VORRandomnessRequestMockRandomnessRequest represents a RandomnessRequest event raised by the VORRandomnessRequestMock contract.
type VORRandomnessRequestMockRandomnessRequest struct {
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
func (_VORRandomnessRequestMock *VORRandomnessRequestMockFilterer) FilterRandomnessRequest(opts *bind.FilterOpts) (*VORRandomnessRequestMockRandomnessRequestIterator, error) {

	logs, sub, err := _VORRandomnessRequestMock.contract.FilterLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return &VORRandomnessRequestMockRandomnessRequestIterator{contract: _VORRandomnessRequestMock.contract, event: "RandomnessRequest", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequest is a free log subscription operation binding the contract event 0xebb37373bb11123e38f964627878b02c247f92f3913df7cf3f270b5222c8d2be.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)
func (_VORRandomnessRequestMock *VORRandomnessRequestMockFilterer) WatchRandomnessRequest(opts *bind.WatchOpts, sink chan<- *VORRandomnessRequestMockRandomnessRequest) (event.Subscription, error) {

	logs, sub, err := _VORRandomnessRequestMock.contract.WatchLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VORRandomnessRequestMockRandomnessRequest)
				if err := _VORRandomnessRequestMock.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
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
func (_VORRandomnessRequestMock *VORRandomnessRequestMockFilterer) ParseRandomnessRequest(log types.Log) (*VORRandomnessRequestMockRandomnessRequest, error) {
	event := new(VORRandomnessRequestMockRandomnessRequest)
	if err := _VORRandomnessRequestMock.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
