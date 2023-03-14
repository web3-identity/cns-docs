// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// AddrResolverABI is the input ABI used to generate the binding from.
const AddrResolverABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVersion\",\"type\":\"uint64\"}],\"name\":\"VersionChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"clearRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"recordVersions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AddrResolver is an auto generated Go binding around an Conflux contract.
type AddrResolver struct {
	AddrResolverCaller     // Read-only binding to the contract
	AddrResolverTransactor // Write-only binding to the contract
	AddrResolverFilterer   // Log filterer for contract events
}

// AddrResolverCaller is an auto generated read-only Go binding around an Conflux contract.
type AddrResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrResolverBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type AddrResolverBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrResolverTransactor is an auto generated write-only Go binding around an Conflux contract.
type AddrResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrResolverBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type AddrResolverBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrResolverFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type AddrResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrResolverSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type AddrResolverSession struct {
	Contract     *AddrResolver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddrResolverCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type AddrResolverCallerSession struct {
	Contract *AddrResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AddrResolverTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type AddrResolverTransactorSession struct {
	Contract     *AddrResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AddrResolverRaw is an auto generated low-level Go binding around an Conflux contract.
type AddrResolverRaw struct {
	Contract *AddrResolver // Generic contract binding to access the raw methods on
}

// AddrResolverCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type AddrResolverCallerRaw struct {
	Contract *AddrResolverCaller // Generic read-only contract binding to access the raw methods on
}

// AddrResolverTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type AddrResolverTransactorRaw struct {
	Contract *AddrResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddrResolver creates a new instance of AddrResolver, bound to a specific deployed contract.
func NewAddrResolver(address types.Address, backend bind.ContractBackend) (*AddrResolver, error) {
	contract, err := bindAddrResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddrResolver{AddrResolverCaller: AddrResolverCaller{contract: contract}, AddrResolverTransactor: AddrResolverTransactor{contract: contract}, AddrResolverFilterer: AddrResolverFilterer{contract: contract}}, nil
}

// NewAddrResolverCaller creates a new read-only instance of AddrResolver, bound to a specific deployed contract.
func NewAddrResolverCaller(address types.Address, caller bind.ContractCaller) (*AddrResolverCaller, error) {
	contract, err := bindAddrResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddrResolverCaller{contract: contract}, nil
}

// NewAddrResolverTransactor creates a new write-only instance of AddrResolver, bound to a specific deployed contract.
func NewAddrResolverTransactor(address types.Address, transactor bind.ContractTransactor) (*AddrResolverTransactor, error) {
	contract, err := bindAddrResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddrResolverTransactor{contract: contract}, nil
}

// NewAddrResolverFilterer creates a new log filterer instance of AddrResolver, bound to a specific deployed contract.
func NewAddrResolverFilterer(address types.Address, filterer bind.ContractFilterer) (*AddrResolverFilterer, error) {
	contract, err := bindAddrResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddrResolverFilterer{contract: contract}, nil
}

// NewAddrResolverCaller creates a new read-only instance of AddrResolver, bound to a specific deployed contract.
func NewAddrResolverBulkCaller(address types.Address, caller bind.ContractCaller) (*AddrResolverBulkCaller, error) {
	contract, err := bindAddrResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddrResolverBulkCaller{contract: contract}, nil
}

// NewAddrResolverBulkTransactor creates a new write-only instance of AddrResolver, bound to a specific deployed contract.
func NewAddrResolverBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*AddrResolverBulkTransactor, error) {
	contract, err := bindAddrResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddrResolverBulkTransactor{contract: contract}, nil
}

// bindAddrResolver binds a generic wrapper to an already deployed contract.
func bindAddrResolver(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrResolver *AddrResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddrResolver.Contract.AddrResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrResolver *AddrResolverRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.AddrResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrResolver *AddrResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.AddrResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrResolver *AddrResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddrResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrResolver *AddrResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrResolver *AddrResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddrResolver *AddrResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	__err := _AddrResolver.contract.Call(opts, &out, "addr", node)

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddrResolver *AddrResolverBulkCaller) Addr(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, node [32]byte) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _AddrResolver.contract.GenRequest(opts, "addr", node)

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _AddrResolver.contract.DecodeOutput(&out, rawOut, "addr")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddrResolver *AddrResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _AddrResolver.Contract.Addr(&_AddrResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_AddrResolver *AddrResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _AddrResolver.Contract.Addr(&_AddrResolver.CallOpts, node)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_AddrResolver *AddrResolverCaller) Addr0(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	__err := _AddrResolver.contract.Call(opts, &out, "addr0", node, coinType)

	if __err != nil {
		return *new([]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, __err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_AddrResolver *AddrResolverBulkCaller) Addr0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, node [32]byte, coinType *big.Int) (*[]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _AddrResolver.contract.GenRequest(opts, "addr0", node, coinType)

	out0 := new([]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _AddrResolver.contract.DecodeOutput(&out, rawOut, "addr0")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_AddrResolver *AddrResolverSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _AddrResolver.Contract.Addr0(&_AddrResolver.CallOpts, node, coinType)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_AddrResolver *AddrResolverCallerSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _AddrResolver.Contract.Addr0(&_AddrResolver.CallOpts, node, coinType)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_AddrResolver *AddrResolverCaller) RecordVersions(opts *bind.CallOpts, arg0 [32]byte) (uint64, error) {
	var out []interface{}
	__err := _AddrResolver.contract.Call(opts, &out, "recordVersions", arg0)

	if __err != nil {
		return *new(uint64), __err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, __err

}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_AddrResolver *AddrResolverBulkCaller) RecordVersions(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 [32]byte) (*uint64, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _AddrResolver.contract.GenRequest(opts, "recordVersions", arg0)

	out0 := new(uint64)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _AddrResolver.contract.DecodeOutput(&out, rawOut, "recordVersions")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(uint64)).(*uint64)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_AddrResolver *AddrResolverSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _AddrResolver.Contract.RecordVersions(&_AddrResolver.CallOpts, arg0)
}

// RecordVersions is a free data retrieval call binding the contract method 0xd700ff33.
//
// Solidity: function recordVersions(bytes32 ) view returns(uint64)
func (_AddrResolver *AddrResolverCallerSession) RecordVersions(arg0 [32]byte) (uint64, error) {
	return _AddrResolver.Contract.RecordVersions(&_AddrResolver.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_AddrResolver *AddrResolverCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	__err := _AddrResolver.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_AddrResolver *AddrResolverBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceID [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _AddrResolver.contract.GenRequest(opts, "supportsInterface", interfaceID)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _AddrResolver.contract.DecodeOutput(&out, rawOut, "supportsInterface")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_AddrResolver *AddrResolverSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _AddrResolver.Contract.SupportsInterface(&_AddrResolver.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_AddrResolver *AddrResolverCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _AddrResolver.Contract.SupportsInterface(&_AddrResolver.CallOpts, interfaceID)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_AddrResolver *AddrResolverTransactor) ClearRecords(opts *bind.TransactOpts, node [32]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.contract.Transact(opts, "clearRecords", node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_AddrResolver *AddrResolverBulkTransactor) ClearRecords(opts *bind.TransactOpts, node [32]byte) types.UnsignedTransaction {
	return _AddrResolver.contract.GenUnsignedTransaction(opts, "clearRecords", node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_AddrResolver *AddrResolverSession) ClearRecords(node [32]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.ClearRecords(&_AddrResolver.TransactOpts, node)
}

// ClearRecords is a paid mutator transaction binding the contract method 0x3603d758.
//
// Solidity: function clearRecords(bytes32 node) returns()
func (_AddrResolver *AddrResolverTransactorSession) ClearRecords(node [32]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.ClearRecords(&_AddrResolver.TransactOpts, node)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_AddrResolver *AddrResolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.contract.Transact(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_AddrResolver *AddrResolverBulkTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) types.UnsignedTransaction {
	return _AddrResolver.contract.GenUnsignedTransaction(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_AddrResolver *AddrResolverSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.SetAddr(&_AddrResolver.TransactOpts, node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_AddrResolver *AddrResolverTransactorSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.SetAddr(&_AddrResolver.TransactOpts, node, coinType, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_AddrResolver *AddrResolverTransactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.contract.Transact(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_AddrResolver *AddrResolverBulkTransactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) types.UnsignedTransaction {
	return _AddrResolver.contract.GenUnsignedTransaction(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_AddrResolver *AddrResolverSession) SetAddr0(node [32]byte, a common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.SetAddr0(&_AddrResolver.TransactOpts, node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_AddrResolver *AddrResolverTransactorSession) SetAddr0(node [32]byte, a common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _AddrResolver.Contract.SetAddr0(&_AddrResolver.TransactOpts, node, a)
}

// AddrResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the AddrResolver contract.
type AddrResolverAddrChangedIterator struct {
	Event *AddrResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *AddrResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddrResolverAddrChanged)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddrResolverAddrChanged)
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
func (it *AddrResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddrResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddrResolverAddrChanged represents a AddrChanged event raised by the AddrResolver contract.
type AddrResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// AddrResolverAddrChangedOrChainReorg represents a AddrChanged subscription event raised by the AddrResolver contract.
type AddrResolverAddrChangedOrChainReorg struct {
	Event      *AddrResolverAddrChanged
	ChainReorg *types.ChainReorg
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_AddrResolver *AddrResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*AddrResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, err := _AddrResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &AddrResolverAddrChangedIterator{contract: _AddrResolver.contract, event: "AddrChanged", logs: logs}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_AddrResolver *AddrResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *AddrResolverAddrChangedOrChainReorg, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AddrResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddrResolverAddrChangedOrChainReorg)
				event.Event = new(AddrResolverAddrChanged)

				if log.ChainReorg == nil {
					if err := _AddrResolver.contract.UnpackLog(event.Event, "AddrChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_AddrResolver *AddrResolverFilterer) ParseAddrChanged(log types.Log) (*AddrResolverAddrChanged, error) {
	event := new(AddrResolverAddrChanged)
	if err := _AddrResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddrResolverAddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the AddrResolver contract.
type AddrResolverAddressChangedIterator struct {
	Event *AddrResolverAddressChanged // Event containing the contract specifics and raw log

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
func (it *AddrResolverAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddrResolverAddressChanged)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddrResolverAddressChanged)
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
func (it *AddrResolverAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddrResolverAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddrResolverAddressChanged represents a AddressChanged event raised by the AddrResolver contract.
type AddrResolverAddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// AddrResolverAddressChangedOrChainReorg represents a AddressChanged subscription event raised by the AddrResolver contract.
type AddrResolverAddressChangedOrChainReorg struct {
	Event      *AddrResolverAddressChanged
	ChainReorg *types.ChainReorg
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_AddrResolver *AddrResolverFilterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*AddrResolverAddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, err := _AddrResolver.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &AddrResolverAddressChangedIterator{contract: _AddrResolver.contract, event: "AddressChanged", logs: logs}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_AddrResolver *AddrResolverFilterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *AddrResolverAddressChangedOrChainReorg, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AddrResolver.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddrResolverAddressChangedOrChainReorg)
				event.Event = new(AddrResolverAddressChanged)

				if log.ChainReorg == nil {
					if err := _AddrResolver.contract.UnpackLog(event.Event, "AddressChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseAddressChanged is a log parse operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_AddrResolver *AddrResolverFilterer) ParseAddressChanged(log types.Log) (*AddrResolverAddressChanged, error) {
	event := new(AddrResolverAddressChanged)
	if err := _AddrResolver.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddrResolverVersionChangedIterator is returned from FilterVersionChanged and is used to iterate over the raw logs and unpacked data for VersionChanged events raised by the AddrResolver contract.
type AddrResolverVersionChangedIterator struct {
	Event *AddrResolverVersionChanged // Event containing the contract specifics and raw log

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
func (it *AddrResolverVersionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddrResolverVersionChanged)
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

	if it.sub == nil {
		it.done = true
		return it.Next()
	}

	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AddrResolverVersionChanged)
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
func (it *AddrResolverVersionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddrResolverVersionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddrResolverVersionChanged represents a VersionChanged event raised by the AddrResolver contract.
type AddrResolverVersionChanged struct {
	Node       [32]byte
	NewVersion uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// AddrResolverVersionChangedOrChainReorg represents a VersionChanged subscription event raised by the AddrResolver contract.
type AddrResolverVersionChangedOrChainReorg struct {
	Event      *AddrResolverVersionChanged
	ChainReorg *types.ChainReorg
}

// FilterVersionChanged is a free log retrieval operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_AddrResolver *AddrResolverFilterer) FilterVersionChanged(opts *bind.FilterOpts, node [][32]byte) (*AddrResolverVersionChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, err := _AddrResolver.contract.FilterLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &AddrResolverVersionChangedIterator{contract: _AddrResolver.contract, event: "VersionChanged", logs: logs}, nil
}

// WatchVersionChanged is a free log subscription operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_AddrResolver *AddrResolverFilterer) WatchVersionChanged(opts *bind.WatchOpts, sink chan<- *AddrResolverVersionChangedOrChainReorg, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _AddrResolver.contract.WatchLogs(opts, "VersionChanged", nodeRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddrResolverVersionChangedOrChainReorg)
				event.Event = new(AddrResolverVersionChanged)

				if log.ChainReorg == nil {
					if err := _AddrResolver.contract.UnpackLog(event.Event, "VersionChanged", *log.Log); err != nil {
						return err
					}
					event.Event.Raw = *log.Log
				} else {
					event.ChainReorg = log.ChainReorg
				}

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

// ParseVersionChanged is a log parse operation binding the contract event 0xc6621ccb8f3f5a04bb6502154b2caf6adf5983fe76dfef1cfc9c42e3579db444.
//
// Solidity: event VersionChanged(bytes32 indexed node, uint64 newVersion)
func (_AddrResolver *AddrResolverFilterer) ParseVersionChanged(log types.Log) (*AddrResolverVersionChanged, error) {
	event := new(AddrResolverVersionChanged)
	if err := _AddrResolver.contract.UnpackLog(event, "VersionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
