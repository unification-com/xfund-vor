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

// VorRandomnessRequestMockABI is the input ABI used to generate the binding from.
const VorRandomnessRequestMockABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"name\":\"RandomnessRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_consumerSeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePaid\",\"type\":\"uint256\"}],\"name\":\"randomnessRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VorRandomnessRequestMock is an auto generated Go binding around an Ethereum contract.
type VorRandomnessRequestMock struct {
	VorRandomnessRequestMockCaller     // Read-only binding to the contract
	VorRandomnessRequestMockTransactor // Write-only binding to the contract
	VorRandomnessRequestMockFilterer   // Log filterer for contract events
}

// VorRandomnessRequestMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type VorRandomnessRequestMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VorRandomnessRequestMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VorRandomnessRequestMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VorRandomnessRequestMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VorRandomnessRequestMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VorRandomnessRequestMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VorRandomnessRequestMockSession struct {
	Contract     *VorRandomnessRequestMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VorRandomnessRequestMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VorRandomnessRequestMockCallerSession struct {
	Contract *VorRandomnessRequestMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// VorRandomnessRequestMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VorRandomnessRequestMockTransactorSession struct {
	Contract     *VorRandomnessRequestMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// VorRandomnessRequestMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type VorRandomnessRequestMockRaw struct {
	Contract *VorRandomnessRequestMock // Generic contract binding to access the raw methods on
}

// VorRandomnessRequestMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VorRandomnessRequestMockCallerRaw struct {
	Contract *VorRandomnessRequestMockCaller // Generic read-only contract binding to access the raw methods on
}

// VorRandomnessRequestMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VorRandomnessRequestMockTransactorRaw struct {
	Contract *VorRandomnessRequestMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVorRandomnessRequestMock creates a new instance of VorRandomnessRequestMock, bound to a specific deployed contract.
func NewVorRandomnessRequestMock(address common.Address, backend bind.ContractBackend) (*VorRandomnessRequestMock, error) {
	contract, err := bindVorRandomnessRequestMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VorRandomnessRequestMock{VorRandomnessRequestMockCaller: VorRandomnessRequestMockCaller{contract: contract}, VorRandomnessRequestMockTransactor: VorRandomnessRequestMockTransactor{contract: contract}, VorRandomnessRequestMockFilterer: VorRandomnessRequestMockFilterer{contract: contract}}, nil
}

// NewVorRandomnessRequestMockCaller creates a new read-only instance of VorRandomnessRequestMock, bound to a specific deployed contract.
func NewVorRandomnessRequestMockCaller(address common.Address, caller bind.ContractCaller) (*VorRandomnessRequestMockCaller, error) {
	contract, err := bindVorRandomnessRequestMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VorRandomnessRequestMockCaller{contract: contract}, nil
}

// NewVorRandomnessRequestMockTransactor creates a new write-only instance of VorRandomnessRequestMock, bound to a specific deployed contract.
func NewVorRandomnessRequestMockTransactor(address common.Address, transactor bind.ContractTransactor) (*VorRandomnessRequestMockTransactor, error) {
	contract, err := bindVorRandomnessRequestMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VorRandomnessRequestMockTransactor{contract: contract}, nil
}

// NewVorRandomnessRequestMockFilterer creates a new log filterer instance of VorRandomnessRequestMock, bound to a specific deployed contract.
func NewVorRandomnessRequestMockFilterer(address common.Address, filterer bind.ContractFilterer) (*VorRandomnessRequestMockFilterer, error) {
	contract, err := bindVorRandomnessRequestMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VorRandomnessRequestMockFilterer{contract: contract}, nil
}

// bindVorRandomnessRequestMock binds a generic wrapper to an already deployed contract.
func bindVorRandomnessRequestMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VorRandomnessRequestMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VorRandomnessRequestMock *VorRandomnessRequestMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VorRandomnessRequestMock.Contract.VorRandomnessRequestMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VorRandomnessRequestMock *VorRandomnessRequestMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.VorRandomnessRequestMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VorRandomnessRequestMock *VorRandomnessRequestMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.VorRandomnessRequestMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VorRandomnessRequestMock *VorRandomnessRequestMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VorRandomnessRequestMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VorRandomnessRequestMock *VorRandomnessRequestMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VorRandomnessRequestMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VorRandomnessRequestMock *VorRandomnessRequestMockSession) Owner() (common.Address, error) {
	return _VorRandomnessRequestMock.Contract.Owner(&_VorRandomnessRequestMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VorRandomnessRequestMock *VorRandomnessRequestMockCallerSession) Owner() (common.Address, error) {
	return _VorRandomnessRequestMock.Contract.Owner(&_VorRandomnessRequestMock.CallOpts)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactor) RandomnessRequest(opts *bind.TransactOpts, _keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.contract.Transact(opts, "randomnessRequest", _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.RandomnessRequest(&_VorRandomnessRequestMock.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RandomnessRequest is a paid mutator transaction binding the contract method 0x264eb1cf.
//
// Solidity: function randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactorSession) RandomnessRequest(_keyHash [32]byte, _consumerSeed *big.Int, _feePaid *big.Int) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.RandomnessRequest(&_VorRandomnessRequestMock.TransactOpts, _keyHash, _consumerSeed, _feePaid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockSession) RenounceOwnership() (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.RenounceOwnership(&_VorRandomnessRequestMock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.RenounceOwnership(&_VorRandomnessRequestMock.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.TransferOwnership(&_VorRandomnessRequestMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VorRandomnessRequestMock *VorRandomnessRequestMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VorRandomnessRequestMock.Contract.TransferOwnership(&_VorRandomnessRequestMock.TransactOpts, newOwner)
}

// VorRandomnessRequestMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VorRandomnessRequestMock contract.
type VorRandomnessRequestMockOwnershipTransferredIterator struct {
	Event *VorRandomnessRequestMockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VorRandomnessRequestMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorRandomnessRequestMockOwnershipTransferred)
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
		it.Event = new(VorRandomnessRequestMockOwnershipTransferred)
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
func (it *VorRandomnessRequestMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorRandomnessRequestMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorRandomnessRequestMockOwnershipTransferred represents a OwnershipTransferred event raised by the VorRandomnessRequestMock contract.
type VorRandomnessRequestMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VorRandomnessRequestMock *VorRandomnessRequestMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VorRandomnessRequestMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VorRandomnessRequestMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VorRandomnessRequestMockOwnershipTransferredIterator{contract: _VorRandomnessRequestMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VorRandomnessRequestMock *VorRandomnessRequestMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VorRandomnessRequestMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VorRandomnessRequestMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorRandomnessRequestMockOwnershipTransferred)
				if err := _VorRandomnessRequestMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VorRandomnessRequestMock *VorRandomnessRequestMockFilterer) ParseOwnershipTransferred(log types.Log) (*VorRandomnessRequestMockOwnershipTransferred, error) {
	event := new(VorRandomnessRequestMockOwnershipTransferred)
	if err := _VorRandomnessRequestMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VorRandomnessRequestMockRandomnessRequestIterator is returned from FilterRandomnessRequest and is used to iterate over the raw logs and unpacked data for RandomnessRequest events raised by the VorRandomnessRequestMock contract.
type VorRandomnessRequestMockRandomnessRequestIterator struct {
	Event *VorRandomnessRequestMockRandomnessRequest // Event containing the contract specifics and raw log

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
func (it *VorRandomnessRequestMockRandomnessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VorRandomnessRequestMockRandomnessRequest)
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
		it.Event = new(VorRandomnessRequestMockRandomnessRequest)
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
func (it *VorRandomnessRequestMockRandomnessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VorRandomnessRequestMockRandomnessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VorRandomnessRequestMockRandomnessRequest represents a RandomnessRequest event raised by the VorRandomnessRequestMock contract.
type VorRandomnessRequestMockRandomnessRequest struct {
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
func (_VorRandomnessRequestMock *VorRandomnessRequestMockFilterer) FilterRandomnessRequest(opts *bind.FilterOpts) (*VorRandomnessRequestMockRandomnessRequestIterator, error) {

	logs, sub, err := _VorRandomnessRequestMock.contract.FilterLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return &VorRandomnessRequestMockRandomnessRequestIterator{contract: _VorRandomnessRequestMock.contract, event: "RandomnessRequest", logs: logs, sub: sub}, nil
}

// WatchRandomnessRequest is a free log subscription operation binding the contract event 0xebb37373bb11123e38f964627878b02c247f92f3913df7cf3f270b5222c8d2be.
//
// Solidity: event RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)
func (_VorRandomnessRequestMock *VorRandomnessRequestMockFilterer) WatchRandomnessRequest(opts *bind.WatchOpts, sink chan<- *VorRandomnessRequestMockRandomnessRequest) (event.Subscription, error) {

	logs, sub, err := _VorRandomnessRequestMock.contract.WatchLogs(opts, "RandomnessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VorRandomnessRequestMockRandomnessRequest)
				if err := _VorRandomnessRequestMock.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
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
func (_VorRandomnessRequestMock *VorRandomnessRequestMockFilterer) ParseRandomnessRequest(log types.Log) (*VorRandomnessRequestMockRandomnessRequest, error) {
	event := new(VorRandomnessRequestMockRandomnessRequest)
	if err := _VorRandomnessRequestMock.contract.UnpackLog(event, "RandomnessRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
