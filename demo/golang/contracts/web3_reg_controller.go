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

// IPriceOraclePrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceOraclePrice struct {
	Base    *big.Int
	Premium *big.Int
}

// Web3RegisterControllerABI is the input ABI used to generate the binding from.
const Web3RegisterControllerABI = "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrarImplementation\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractIFiatPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"contractReverseRegistrar\",\"name\":\"_reverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractICNameWrapper\",\"name\":\"_nameWrapper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooLow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseRegistrarImplementation\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractIFiatPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"contractReverseRegistrar\",\"name\":\"_reverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractICNameWrapper\",\"name\":\"_nameWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"labelStatus\",\"outputs\":[{\"internalType\":\"enumWeb3RegistrarController.LabelStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"wrapperExpiry\",\"type\":\"uint64\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameWhitelist\",\"outputs\":[{\"internalType\":\"contractINameWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameWrapper\",\"outputs\":[{\"internalType\":\"contractICNameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIFiatPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"wrapperExpiry\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"fuses\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"wrapperExpiry\",\"type\":\"uint64\"}],\"name\":\"registerWithFiat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"wrapperExpiry\",\"type\":\"uint64\"}],\"name\":\"renewWithFiat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPriceInFiat\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverseRegistrar\",\"outputs\":[{\"internalType\":\"contractReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"}],\"name\":\"setCommitmentAge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quota\",\"type\":\"uint256\"}],\"name\":\"setLabel45Quota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINameWhitelist\",\"name\":\"_nameWhitelist\",\"type\":\"address\"}],\"name\":\"setNameWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFiatPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_len\",\"type\":\"uint256\"}],\"name\":\"setValidLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Web3RegisterController is an auto generated Go binding around an Conflux contract.
type Web3RegisterController struct {
	Web3RegisterControllerCaller     // Read-only binding to the contract
	Web3RegisterControllerTransactor // Write-only binding to the contract
	Web3RegisterControllerFilterer   // Log filterer for contract events
}

// Web3RegisterControllerCaller is an auto generated read-only Go binding around an Conflux contract.
type Web3RegisterControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3RegisterControllerBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type Web3RegisterControllerBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3RegisterControllerTransactor is an auto generated write-only Go binding around an Conflux contract.
type Web3RegisterControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3RegisterControllerBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type Web3RegisterControllerBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3RegisterControllerFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type Web3RegisterControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3RegisterControllerSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type Web3RegisterControllerSession struct {
	Contract     *Web3RegisterController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Web3RegisterControllerCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type Web3RegisterControllerCallerSession struct {
	Contract *Web3RegisterControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// Web3RegisterControllerTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type Web3RegisterControllerTransactorSession struct {
	Contract     *Web3RegisterControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Web3RegisterControllerRaw is an auto generated low-level Go binding around an Conflux contract.
type Web3RegisterControllerRaw struct {
	Contract *Web3RegisterController // Generic contract binding to access the raw methods on
}

// Web3RegisterControllerCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type Web3RegisterControllerCallerRaw struct {
	Contract *Web3RegisterControllerCaller // Generic read-only contract binding to access the raw methods on
}

// Web3RegisterControllerTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type Web3RegisterControllerTransactorRaw struct {
	Contract *Web3RegisterControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWeb3RegisterController creates a new instance of Web3RegisterController, bound to a specific deployed contract.
func NewWeb3RegisterController(address types.Address, backend bind.ContractBackend) (*Web3RegisterController, error) {
	contract, err := bindWeb3RegisterController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterController{Web3RegisterControllerCaller: Web3RegisterControllerCaller{contract: contract}, Web3RegisterControllerTransactor: Web3RegisterControllerTransactor{contract: contract}, Web3RegisterControllerFilterer: Web3RegisterControllerFilterer{contract: contract}}, nil
}

// NewWeb3RegisterControllerCaller creates a new read-only instance of Web3RegisterController, bound to a specific deployed contract.
func NewWeb3RegisterControllerCaller(address types.Address, caller bind.ContractCaller) (*Web3RegisterControllerCaller, error) {
	contract, err := bindWeb3RegisterController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerCaller{contract: contract}, nil
}

// NewWeb3RegisterControllerTransactor creates a new write-only instance of Web3RegisterController, bound to a specific deployed contract.
func NewWeb3RegisterControllerTransactor(address types.Address, transactor bind.ContractTransactor) (*Web3RegisterControllerTransactor, error) {
	contract, err := bindWeb3RegisterController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerTransactor{contract: contract}, nil
}

// NewWeb3RegisterControllerFilterer creates a new log filterer instance of Web3RegisterController, bound to a specific deployed contract.
func NewWeb3RegisterControllerFilterer(address types.Address, filterer bind.ContractFilterer) (*Web3RegisterControllerFilterer, error) {
	contract, err := bindWeb3RegisterController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerFilterer{contract: contract}, nil
}

// NewWeb3RegisterControllerCaller creates a new read-only instance of Web3RegisterController, bound to a specific deployed contract.
func NewWeb3RegisterControllerBulkCaller(address types.Address, caller bind.ContractCaller) (*Web3RegisterControllerBulkCaller, error) {
	contract, err := bindWeb3RegisterController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerBulkCaller{contract: contract}, nil
}

// NewWeb3RegisterControllerBulkTransactor creates a new write-only instance of Web3RegisterController, bound to a specific deployed contract.
func NewWeb3RegisterControllerBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*Web3RegisterControllerBulkTransactor, error) {
	contract, err := bindWeb3RegisterController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerBulkTransactor{contract: contract}, nil
}

// bindWeb3RegisterController binds a generic wrapper to an already deployed contract.
func bindWeb3RegisterController(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Web3RegisterControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3RegisterController *Web3RegisterControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3RegisterController.Contract.Web3RegisterControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3RegisterController *Web3RegisterControllerRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Web3RegisterControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3RegisterController *Web3RegisterControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Web3RegisterControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3RegisterController *Web3RegisterControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3RegisterController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3RegisterController *Web3RegisterControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3RegisterController *Web3RegisterControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "ADMIN_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) ADMINROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "ADMIN_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "ADMIN_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerSession) ADMINROLE() ([32]byte, error) {
	return _Web3RegisterController.Contract.ADMINROLE(&_Web3RegisterController.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) ADMINROLE() ([32]byte, error) {
	return _Web3RegisterController.Contract.ADMINROLE(&_Web3RegisterController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) DEFAULTADMINROLE(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "DEFAULT_ADMIN_ROLE")

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "DEFAULT_ADMIN_ROLE")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Web3RegisterController.Contract.DEFAULTADMINROLE(&_Web3RegisterController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Web3RegisterController.Contract.DEFAULTADMINROLE(&_Web3RegisterController.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) MINREGISTRATIONDURATION(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "MIN_REGISTRATION_DURATION")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "MIN_REGISTRATION_DURATION")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _Web3RegisterController.Contract.MINREGISTRATIONDURATION(&_Web3RegisterController.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _Web3RegisterController.Contract.MINREGISTRATIONDURATION(&_Web3RegisterController.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "available", name)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) Available(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, name string) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "available", name)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "available")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerSession) Available(name string) (bool, error) {
	return _Web3RegisterController.Contract.Available(&_Web3RegisterController.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) Available(name string) (bool, error) {
	return _Web3RegisterController.Contract.Available(&_Web3RegisterController.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "commitments", arg0)

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) Commitments(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, arg0 [32]byte) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "commitments", arg0)

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "commitments")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _Web3RegisterController.Contract.Commitments(&_Web3RegisterController.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _Web3RegisterController.Contract.Commitments(&_Web3RegisterController.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "getRoleAdmin", role)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) GetRoleAdmin(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "getRoleAdmin", role)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "getRoleAdmin")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Web3RegisterController.Contract.GetRoleAdmin(&_Web3RegisterController.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Web3RegisterController.Contract.GetRoleAdmin(&_Web3RegisterController.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "hasRole", role, account)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) HasRole(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, role [32]byte, account common.Address) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "hasRole", role, account)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "hasRole")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Web3RegisterController.Contract.HasRole(&_Web3RegisterController.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Web3RegisterController.Contract.HasRole(&_Web3RegisterController.CallOpts, role, account)
}

// LabelStatus is a free data retrieval call binding the contract method 0xb2f100f0.
//
// Solidity: function labelStatus(string _label) view returns(uint8)
func (_Web3RegisterController *Web3RegisterControllerCaller) LabelStatus(opts *bind.CallOpts, _label string) (uint8, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "labelStatus", _label)

	if __err != nil {
		return *new(uint8), __err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, __err

}

// LabelStatus is a free data retrieval call binding the contract method 0xb2f100f0.
//
// Solidity: function labelStatus(string _label) view returns(uint8)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) LabelStatus(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, _label string) (*uint8, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "labelStatus", _label)

	out0 := new(uint8)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "labelStatus")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(uint8)).(*uint8)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// LabelStatus is a free data retrieval call binding the contract method 0xb2f100f0.
//
// Solidity: function labelStatus(string _label) view returns(uint8)
func (_Web3RegisterController *Web3RegisterControllerSession) LabelStatus(_label string) (uint8, error) {
	return _Web3RegisterController.Contract.LabelStatus(&_Web3RegisterController.CallOpts, _label)
}

// LabelStatus is a free data retrieval call binding the contract method 0xb2f100f0.
//
// Solidity: function labelStatus(string _label) view returns(uint8)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) LabelStatus(_label string) (uint8, error) {
	return _Web3RegisterController.Contract.LabelStatus(&_Web3RegisterController.CallOpts, _label)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x1f6e854a.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCaller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) ([32]byte, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "makeCommitment", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// MakeCommitment is a free data retrieval call binding the contract method 0x1f6e854a.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) MakeCommitment(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "makeCommitment", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "makeCommitment")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MakeCommitment is a free data retrieval call binding the contract method 0x1f6e854a.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) ([32]byte, error) {
	return _Web3RegisterController.Contract.MakeCommitment(&_Web3RegisterController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x1f6e854a.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) ([32]byte, error) {
	return _Web3RegisterController.Contract.MakeCommitment(&_Web3RegisterController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)
}

// MakeCommitment0 is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCaller) MakeCommitment0(opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "makeCommitment0", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)

	if __err != nil {
		return *new([32]byte), __err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, __err

}

// MakeCommitment0 is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) MakeCommitment0(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*[32]byte, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "makeCommitment0", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)

	out0 := new([32]byte)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "makeCommitment0")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MakeCommitment0 is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerSession) MakeCommitment0(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _Web3RegisterController.Contract.MakeCommitment0(&_Web3RegisterController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MakeCommitment0 is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) MakeCommitment0(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _Web3RegisterController.Contract.MakeCommitment0(&_Web3RegisterController.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCaller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "maxCommitmentAge")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) MaxCommitmentAge(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "maxCommitmentAge")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "maxCommitmentAge")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerSession) MaxCommitmentAge() (*big.Int, error) {
	return _Web3RegisterController.Contract.MaxCommitmentAge(&_Web3RegisterController.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _Web3RegisterController.Contract.MaxCommitmentAge(&_Web3RegisterController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCaller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "minCommitmentAge")

	if __err != nil {
		return *new(*big.Int), __err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, __err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) MinCommitmentAge(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (**big.Int, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "minCommitmentAge")

	out0 := new(*big.Int)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "minCommitmentAge")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerSession) MinCommitmentAge() (*big.Int, error) {
	return _Web3RegisterController.Contract.MinCommitmentAge(&_Web3RegisterController.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) MinCommitmentAge() (*big.Int, error) {
	return _Web3RegisterController.Contract.MinCommitmentAge(&_Web3RegisterController.CallOpts)
}

// NameWhitelist is a free data retrieval call binding the contract method 0x1e803098.
//
// Solidity: function nameWhitelist() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCaller) NameWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "nameWhitelist")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// NameWhitelist is a free data retrieval call binding the contract method 0x1e803098.
//
// Solidity: function nameWhitelist() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) NameWhitelist(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "nameWhitelist")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "nameWhitelist")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// NameWhitelist is a free data retrieval call binding the contract method 0x1e803098.
//
// Solidity: function nameWhitelist() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerSession) NameWhitelist() (common.Address, error) {
	return _Web3RegisterController.Contract.NameWhitelist(&_Web3RegisterController.CallOpts)
}

// NameWhitelist is a free data retrieval call binding the contract method 0x1e803098.
//
// Solidity: function nameWhitelist() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) NameWhitelist() (common.Address, error) {
	return _Web3RegisterController.Contract.NameWhitelist(&_Web3RegisterController.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCaller) NameWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "nameWrapper")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) NameWrapper(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "nameWrapper")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "nameWrapper")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerSession) NameWrapper() (common.Address, error) {
	return _Web3RegisterController.Contract.NameWrapper(&_Web3RegisterController.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) NameWrapper() (common.Address, error) {
	return _Web3RegisterController.Contract.NameWrapper(&_Web3RegisterController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "owner")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) Owner(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "owner")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "owner")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerSession) Owner() (common.Address, error) {
	return _Web3RegisterController.Contract.Owner(&_Web3RegisterController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) Owner() (common.Address, error) {
	return _Web3RegisterController.Contract.Owner(&_Web3RegisterController.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCaller) Prices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "prices")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) Prices(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "prices")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "prices")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerSession) Prices() (common.Address, error) {
	return _Web3RegisterController.Contract.Prices(&_Web3RegisterController.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) Prices() (common.Address, error) {
	return _Web3RegisterController.Contract.Prices(&_Web3RegisterController.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerCaller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (IPriceOraclePrice, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "rentPrice", name, duration)

	if __err != nil {
		return *new(IPriceOraclePrice), __err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

	return out0, __err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) RentPrice(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, name string, duration *big.Int) (*IPriceOraclePrice, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "rentPrice", name, duration)

	out0 := new(IPriceOraclePrice)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "rentPrice")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _Web3RegisterController.Contract.RentPrice(&_Web3RegisterController.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _Web3RegisterController.Contract.RentPrice(&_Web3RegisterController.CallOpts, name, duration)
}

// RentPriceInFiat is a free data retrieval call binding the contract method 0x9da4f2e1.
//
// Solidity: function rentPriceInFiat(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerCaller) RentPriceInFiat(opts *bind.CallOpts, name string, duration *big.Int) (IPriceOraclePrice, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "rentPriceInFiat", name, duration)

	if __err != nil {
		return *new(IPriceOraclePrice), __err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

	return out0, __err

}

// RentPriceInFiat is a free data retrieval call binding the contract method 0x9da4f2e1.
//
// Solidity: function rentPriceInFiat(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) RentPriceInFiat(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, name string, duration *big.Int) (*IPriceOraclePrice, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "rentPriceInFiat", name, duration)

	out0 := new(IPriceOraclePrice)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "rentPriceInFiat")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// RentPriceInFiat is a free data retrieval call binding the contract method 0x9da4f2e1.
//
// Solidity: function rentPriceInFiat(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerSession) RentPriceInFiat(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _Web3RegisterController.Contract.RentPriceInFiat(&_Web3RegisterController.CallOpts, name, duration)
}

// RentPriceInFiat is a free data retrieval call binding the contract method 0x9da4f2e1.
//
// Solidity: function rentPriceInFiat(string name, uint256 duration) view returns((uint256,uint256) price)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) RentPriceInFiat(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _Web3RegisterController.Contract.RentPriceInFiat(&_Web3RegisterController.CallOpts, name, duration)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCaller) ReverseRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "reverseRegistrar")

	if __err != nil {
		return *new(common.Address), __err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, __err

}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) ReverseRegistrar(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) (*common.Address, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "reverseRegistrar")

	out0 := new(common.Address)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "reverseRegistrar")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerSession) ReverseRegistrar() (common.Address, error) {
	return _Web3RegisterController.Contract.ReverseRegistrar(&_Web3RegisterController.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) ReverseRegistrar() (common.Address, error) {
	return _Web3RegisterController.Contract.ReverseRegistrar(&_Web3RegisterController.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) SupportsInterface(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, interfaceID [4]byte) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "supportsInterface", interfaceID)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "supportsInterface")
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
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Web3RegisterController *Web3RegisterControllerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Web3RegisterController.Contract.SupportsInterface(&_Web3RegisterController.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Web3RegisterController.Contract.SupportsInterface(&_Web3RegisterController.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	__err := _Web3RegisterController.contract.Call(opts, &out, "valid", name)

	if __err != nil {
		return *new(bool), __err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, __err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerBulkCaller) Valid(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, name string) (*bool, *error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _Web3RegisterController.contract.GenRequest(opts, "valid", name)

	out0 := new(bool)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _Web3RegisterController.contract.DecodeOutput(&out, rawOut, "valid")
		if err != nil {
			return err
		}

		*out0 = *abi.ConvertType(out[0], new(bool)).(*bool)

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return out0, __err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerSession) Valid(name string) (bool, error) {
	return _Web3RegisterController.Contract.Valid(&_Web3RegisterController.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) view returns(bool)
func (_Web3RegisterController *Web3RegisterControllerCallerSession) Valid(name string) (bool, error) {
	return _Web3RegisterController.Contract.Valid(&_Web3RegisterController.CallOpts, name)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) AddAdmin(addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.AddAdmin(&_Web3RegisterController.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) AddAdmin(addr common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.AddAdmin(&_Web3RegisterController.TransactOpts, addr)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) Commit(opts *bind.TransactOpts, commitment [32]byte) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) Commit(commitment [32]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Commit(&_Web3RegisterController.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) Commit(commitment [32]byte) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Commit(&_Web3RegisterController.TransactOpts, commitment)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.GrantRole(&_Web3RegisterController.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.GrantRole(&_Web3RegisterController.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e785737.
//
// Solidity: function initialize(address _base, address _prices, uint256 _minCommitmentAge, uint256 _maxCommitmentAge, address _reverseRegistrar, address _nameWrapper, address _admin) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) Initialize(opts *bind.TransactOpts, _base common.Address, _prices common.Address, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int, _reverseRegistrar common.Address, _nameWrapper common.Address, _admin common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "initialize", _base, _prices, _minCommitmentAge, _maxCommitmentAge, _reverseRegistrar, _nameWrapper, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e785737.
//
// Solidity: function initialize(address _base, address _prices, uint256 _minCommitmentAge, uint256 _maxCommitmentAge, address _reverseRegistrar, address _nameWrapper, address _admin) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) Initialize(opts *bind.TransactOpts, _base common.Address, _prices common.Address, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int, _reverseRegistrar common.Address, _nameWrapper common.Address, _admin common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "initialize", _base, _prices, _minCommitmentAge, _maxCommitmentAge, _reverseRegistrar, _nameWrapper, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e785737.
//
// Solidity: function initialize(address _base, address _prices, uint256 _minCommitmentAge, uint256 _maxCommitmentAge, address _reverseRegistrar, address _nameWrapper, address _admin) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) Initialize(_base common.Address, _prices common.Address, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int, _reverseRegistrar common.Address, _nameWrapper common.Address, _admin common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Initialize(&_Web3RegisterController.TransactOpts, _base, _prices, _minCommitmentAge, _maxCommitmentAge, _reverseRegistrar, _nameWrapper, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x3e785737.
//
// Solidity: function initialize(address _base, address _prices, uint256 _minCommitmentAge, uint256 _maxCommitmentAge, address _reverseRegistrar, address _nameWrapper, address _admin) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) Initialize(_base common.Address, _prices common.Address, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int, _reverseRegistrar common.Address, _nameWrapper common.Address, _admin common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Initialize(&_Web3RegisterController.TransactOpts, _base, _prices, _minCommitmentAge, _maxCommitmentAge, _reverseRegistrar, _nameWrapper, _admin)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RecoverFunds(&_Web3RegisterController.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RecoverFunds(&_Web3RegisterController.TransactOpts, _token, _to, _amount)
}

// Register is a paid mutator transaction binding the contract method 0x4e73c666.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) payable returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "register", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)
}

// Register is a paid mutator transaction binding the contract method 0x4e73c666.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) payable returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "register", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)
}

// Register is a paid mutator transaction binding the contract method 0x4e73c666.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) payable returns()
func (_Web3RegisterController *Web3RegisterControllerSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Register(&_Web3RegisterController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)
}

// Register is a paid mutator transaction binding the contract method 0x4e73c666.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses, uint64 wrapperExpiry) payable returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Register(&_Web3RegisterController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses, wrapperExpiry)
}

// Register0 is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) Register0(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "register0", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register0 is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) Register0(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "register0", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register0 is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_Web3RegisterController *Web3RegisterControllerSession) Register0(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Register0(&_Web3RegisterController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register0 is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) Register0(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Register0(&_Web3RegisterController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// RegisterWithFiat is a paid mutator transaction binding the contract method 0xa72e1cb2.
//
// Solidity: function registerWithFiat(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) RegisterWithFiat(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, fuses uint16, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "registerWithFiat", name, owner, duration, secret, resolver, data, reverseRecord, fuses, wrapperExpiry)
}

// RegisterWithFiat is a paid mutator transaction binding the contract method 0xa72e1cb2.
//
// Solidity: function registerWithFiat(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) RegisterWithFiat(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, fuses uint16, wrapperExpiry uint64) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "registerWithFiat", name, owner, duration, secret, resolver, data, reverseRecord, fuses, wrapperExpiry)
}

// RegisterWithFiat is a paid mutator transaction binding the contract method 0xa72e1cb2.
//
// Solidity: function registerWithFiat(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) RegisterWithFiat(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, fuses uint16, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RegisterWithFiat(&_Web3RegisterController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, fuses, wrapperExpiry)
}

// RegisterWithFiat is a paid mutator transaction binding the contract method 0xa72e1cb2.
//
// Solidity: function registerWithFiat(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) RegisterWithFiat(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, fuses uint16, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RegisterWithFiat(&_Web3RegisterController.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, fuses, wrapperExpiry)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Web3RegisterController *Web3RegisterControllerSession) Renew(name string, duration *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Renew(&_Web3RegisterController.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) Renew(name string, duration *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Renew(&_Web3RegisterController.TransactOpts, name, duration)
}

// RenewWithFiat is a paid mutator transaction binding the contract method 0x1ea61fc4.
//
// Solidity: function renewWithFiat(string name, uint256 duration, uint32 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) RenewWithFiat(opts *bind.TransactOpts, name string, duration *big.Int, fuses uint32, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "renewWithFiat", name, duration, fuses, wrapperExpiry)
}

// RenewWithFiat is a paid mutator transaction binding the contract method 0x1ea61fc4.
//
// Solidity: function renewWithFiat(string name, uint256 duration, uint32 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) RenewWithFiat(opts *bind.TransactOpts, name string, duration *big.Int, fuses uint32, wrapperExpiry uint64) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "renewWithFiat", name, duration, fuses, wrapperExpiry)
}

// RenewWithFiat is a paid mutator transaction binding the contract method 0x1ea61fc4.
//
// Solidity: function renewWithFiat(string name, uint256 duration, uint32 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) RenewWithFiat(name string, duration *big.Int, fuses uint32, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RenewWithFiat(&_Web3RegisterController.TransactOpts, name, duration, fuses, wrapperExpiry)
}

// RenewWithFiat is a paid mutator transaction binding the contract method 0x1ea61fc4.
//
// Solidity: function renewWithFiat(string name, uint256 duration, uint32 fuses, uint64 wrapperExpiry) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) RenewWithFiat(name string, duration *big.Int, fuses uint32, wrapperExpiry uint64) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RenewWithFiat(&_Web3RegisterController.TransactOpts, name, duration, fuses, wrapperExpiry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) RenounceOwnership(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3RegisterController *Web3RegisterControllerSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RenounceOwnership(&_Web3RegisterController.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) RenounceOwnership() (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RenounceOwnership(&_Web3RegisterController.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RenounceRole(&_Web3RegisterController.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RenounceRole(&_Web3RegisterController.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RevokeRole(&_Web3RegisterController.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.RevokeRole(&_Web3RegisterController.TransactOpts, role, account)
}

// SetCommitmentAge is a paid mutator transaction binding the contract method 0xe4214b46.
//
// Solidity: function setCommitmentAge(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) SetCommitmentAge(opts *bind.TransactOpts, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "setCommitmentAge", _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAge is a paid mutator transaction binding the contract method 0xe4214b46.
//
// Solidity: function setCommitmentAge(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) SetCommitmentAge(opts *bind.TransactOpts, _minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "setCommitmentAge", _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAge is a paid mutator transaction binding the contract method 0xe4214b46.
//
// Solidity: function setCommitmentAge(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) SetCommitmentAge(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetCommitmentAge(&_Web3RegisterController.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetCommitmentAge is a paid mutator transaction binding the contract method 0xe4214b46.
//
// Solidity: function setCommitmentAge(uint256 _minCommitmentAge, uint256 _maxCommitmentAge) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) SetCommitmentAge(_minCommitmentAge *big.Int, _maxCommitmentAge *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetCommitmentAge(&_Web3RegisterController.TransactOpts, _minCommitmentAge, _maxCommitmentAge)
}

// SetLabel45Quota is a paid mutator transaction binding the contract method 0xa97bf04b.
//
// Solidity: function setLabel45Quota(uint256 _quota) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) SetLabel45Quota(opts *bind.TransactOpts, _quota *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "setLabel45Quota", _quota)
}

// SetLabel45Quota is a paid mutator transaction binding the contract method 0xa97bf04b.
//
// Solidity: function setLabel45Quota(uint256 _quota) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) SetLabel45Quota(opts *bind.TransactOpts, _quota *big.Int) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "setLabel45Quota", _quota)
}

// SetLabel45Quota is a paid mutator transaction binding the contract method 0xa97bf04b.
//
// Solidity: function setLabel45Quota(uint256 _quota) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) SetLabel45Quota(_quota *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetLabel45Quota(&_Web3RegisterController.TransactOpts, _quota)
}

// SetLabel45Quota is a paid mutator transaction binding the contract method 0xa97bf04b.
//
// Solidity: function setLabel45Quota(uint256 _quota) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) SetLabel45Quota(_quota *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetLabel45Quota(&_Web3RegisterController.TransactOpts, _quota)
}

// SetNameWhitelist is a paid mutator transaction binding the contract method 0xe221496f.
//
// Solidity: function setNameWhitelist(address _nameWhitelist) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) SetNameWhitelist(opts *bind.TransactOpts, _nameWhitelist common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "setNameWhitelist", _nameWhitelist)
}

// SetNameWhitelist is a paid mutator transaction binding the contract method 0xe221496f.
//
// Solidity: function setNameWhitelist(address _nameWhitelist) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) SetNameWhitelist(opts *bind.TransactOpts, _nameWhitelist common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "setNameWhitelist", _nameWhitelist)
}

// SetNameWhitelist is a paid mutator transaction binding the contract method 0xe221496f.
//
// Solidity: function setNameWhitelist(address _nameWhitelist) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) SetNameWhitelist(_nameWhitelist common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetNameWhitelist(&_Web3RegisterController.TransactOpts, _nameWhitelist)
}

// SetNameWhitelist is a paid mutator transaction binding the contract method 0xe221496f.
//
// Solidity: function setNameWhitelist(address _nameWhitelist) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) SetNameWhitelist(_nameWhitelist common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetNameWhitelist(&_Web3RegisterController.TransactOpts, _nameWhitelist)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) SetPriceOracle(opts *bind.TransactOpts, _prices common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "setPriceOracle", _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) SetPriceOracle(opts *bind.TransactOpts, _prices common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "setPriceOracle", _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) SetPriceOracle(_prices common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetPriceOracle(&_Web3RegisterController.TransactOpts, _prices)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _prices) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) SetPriceOracle(_prices common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetPriceOracle(&_Web3RegisterController.TransactOpts, _prices)
}

// SetValidLen is a paid mutator transaction binding the contract method 0x602ece21.
//
// Solidity: function setValidLen(uint256 _len) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) SetValidLen(opts *bind.TransactOpts, _len *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "setValidLen", _len)
}

// SetValidLen is a paid mutator transaction binding the contract method 0x602ece21.
//
// Solidity: function setValidLen(uint256 _len) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) SetValidLen(opts *bind.TransactOpts, _len *big.Int) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "setValidLen", _len)
}

// SetValidLen is a paid mutator transaction binding the contract method 0x602ece21.
//
// Solidity: function setValidLen(uint256 _len) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) SetValidLen(_len *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetValidLen(&_Web3RegisterController.TransactOpts, _len)
}

// SetValidLen is a paid mutator transaction binding the contract method 0x602ece21.
//
// Solidity: function setValidLen(uint256 _len) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) SetValidLen(_len *big.Int) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.SetValidLen(&_Web3RegisterController.TransactOpts, _len)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3RegisterController *Web3RegisterControllerSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.TransferOwnership(&_Web3RegisterController.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.TransferOwnership(&_Web3RegisterController.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Web3RegisterController *Web3RegisterControllerTransactor) Withdraw(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Web3RegisterController *Web3RegisterControllerBulkTransactor) Withdraw(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _Web3RegisterController.contract.GenUnsignedTransaction(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Web3RegisterController *Web3RegisterControllerSession) Withdraw() (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Withdraw(&_Web3RegisterController.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Web3RegisterController *Web3RegisterControllerTransactorSession) Withdraw() (*types.UnsignedTransaction, *types.Hash, error) {
	return _Web3RegisterController.Contract.Withdraw(&_Web3RegisterController.TransactOpts)
}

// Web3RegisterControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Web3RegisterController contract.
type Web3RegisterControllerInitializedIterator struct {
	Event *Web3RegisterControllerInitialized // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerInitialized)
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
		it.Event = new(Web3RegisterControllerInitialized)
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
func (it *Web3RegisterControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerInitialized represents a Initialized event raised by the Web3RegisterController contract.
type Web3RegisterControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerInitializedOrChainReorg represents a Initialized subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerInitializedOrChainReorg struct {
	Event      *Web3RegisterControllerInitialized
	ChainReorg *types.ChainReorg
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*Web3RegisterControllerInitializedIterator, error) {

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerInitializedIterator{contract: _Web3RegisterController.contract, event: "Initialized", logs: logs}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerInitializedOrChainReorg) (event.Subscription, error) {

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerInitializedOrChainReorg)
				event.Event = new(Web3RegisterControllerInitialized)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "Initialized", *log.Log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseInitialized(log types.Log) (*Web3RegisterControllerInitialized, error) {
	event := new(Web3RegisterControllerInitialized)
	if err := _Web3RegisterController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3RegisterControllerNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Web3RegisterController contract.
type Web3RegisterControllerNameRegisteredIterator struct {
	Event *Web3RegisterControllerNameRegistered // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerNameRegistered)
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
		it.Event = new(Web3RegisterControllerNameRegistered)
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
func (it *Web3RegisterControllerNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerNameRegistered represents a NameRegistered event raised by the Web3RegisterController contract.
type Web3RegisterControllerNameRegistered struct {
	Name     string
	Label    [32]byte
	Owner    common.Address
	BaseCost *big.Int
	Premium  *big.Int
	Expires  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerNameRegisteredOrChainReorg represents a NameRegistered subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerNameRegisteredOrChainReorg struct {
	Event      *Web3RegisterControllerNameRegistered
	ChainReorg *types.ChainReorg
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*Web3RegisterControllerNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerNameRegisteredIterator{contract: _Web3RegisterController.contract, event: "NameRegistered", logs: logs}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerNameRegisteredOrChainReorg, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerNameRegisteredOrChainReorg)
				event.Event = new(Web3RegisterControllerNameRegistered)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "NameRegistered", *log.Log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseNameRegistered(log types.Log) (*Web3RegisterControllerNameRegistered, error) {
	event := new(Web3RegisterControllerNameRegistered)
	if err := _Web3RegisterController.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3RegisterControllerNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Web3RegisterController contract.
type Web3RegisterControllerNameRenewedIterator struct {
	Event *Web3RegisterControllerNameRenewed // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerNameRenewed)
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
		it.Event = new(Web3RegisterControllerNameRenewed)
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
func (it *Web3RegisterControllerNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerNameRenewed represents a NameRenewed event raised by the Web3RegisterController contract.
type Web3RegisterControllerNameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerNameRenewedOrChainReorg represents a NameRenewed subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerNameRenewedOrChainReorg struct {
	Event      *Web3RegisterControllerNameRenewed
	ChainReorg *types.ChainReorg
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*Web3RegisterControllerNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerNameRenewedIterator{contract: _Web3RegisterController.contract, event: "NameRenewed", logs: logs}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerNameRenewedOrChainReorg, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerNameRenewedOrChainReorg)
				event.Event = new(Web3RegisterControllerNameRenewed)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "NameRenewed", *log.Log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseNameRenewed(log types.Log) (*Web3RegisterControllerNameRenewed, error) {
	event := new(Web3RegisterControllerNameRenewed)
	if err := _Web3RegisterController.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3RegisterControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Web3RegisterController contract.
type Web3RegisterControllerOwnershipTransferredIterator struct {
	Event *Web3RegisterControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerOwnershipTransferred)
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
		it.Event = new(Web3RegisterControllerOwnershipTransferred)
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
func (it *Web3RegisterControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerOwnershipTransferred represents a OwnershipTransferred event raised by the Web3RegisterController contract.
type Web3RegisterControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerOwnershipTransferredOrChainReorg represents a OwnershipTransferred subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerOwnershipTransferredOrChainReorg struct {
	Event      *Web3RegisterControllerOwnershipTransferred
	ChainReorg *types.ChainReorg
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Web3RegisterControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerOwnershipTransferredIterator{contract: _Web3RegisterController.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerOwnershipTransferredOrChainReorg, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerOwnershipTransferredOrChainReorg)
				event.Event = new(Web3RegisterControllerOwnershipTransferred)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "OwnershipTransferred", *log.Log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseOwnershipTransferred(log types.Log) (*Web3RegisterControllerOwnershipTransferred, error) {
	event := new(Web3RegisterControllerOwnershipTransferred)
	if err := _Web3RegisterController.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3RegisterControllerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleAdminChangedIterator struct {
	Event *Web3RegisterControllerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerRoleAdminChanged)
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
		it.Event = new(Web3RegisterControllerRoleAdminChanged)
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
func (it *Web3RegisterControllerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerRoleAdminChanged represents a RoleAdminChanged event raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerRoleAdminChangedOrChainReorg represents a RoleAdminChanged subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleAdminChangedOrChainReorg struct {
	Event      *Web3RegisterControllerRoleAdminChanged
	ChainReorg *types.ChainReorg
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Web3RegisterControllerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerRoleAdminChangedIterator{contract: _Web3RegisterController.contract, event: "RoleAdminChanged", logs: logs}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerRoleAdminChangedOrChainReorg, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerRoleAdminChangedOrChainReorg)
				event.Event = new(Web3RegisterControllerRoleAdminChanged)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "RoleAdminChanged", *log.Log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseRoleAdminChanged(log types.Log) (*Web3RegisterControllerRoleAdminChanged, error) {
	event := new(Web3RegisterControllerRoleAdminChanged)
	if err := _Web3RegisterController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3RegisterControllerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleGrantedIterator struct {
	Event *Web3RegisterControllerRoleGranted // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerRoleGranted)
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
		it.Event = new(Web3RegisterControllerRoleGranted)
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
func (it *Web3RegisterControllerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerRoleGranted represents a RoleGranted event raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerRoleGrantedOrChainReorg represents a RoleGranted subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleGrantedOrChainReorg struct {
	Event      *Web3RegisterControllerRoleGranted
	ChainReorg *types.ChainReorg
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Web3RegisterControllerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerRoleGrantedIterator{contract: _Web3RegisterController.contract, event: "RoleGranted", logs: logs}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerRoleGrantedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerRoleGrantedOrChainReorg)
				event.Event = new(Web3RegisterControllerRoleGranted)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "RoleGranted", *log.Log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseRoleGranted(log types.Log) (*Web3RegisterControllerRoleGranted, error) {
	event := new(Web3RegisterControllerRoleGranted)
	if err := _Web3RegisterController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3RegisterControllerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleRevokedIterator struct {
	Event *Web3RegisterControllerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *Web3RegisterControllerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3RegisterControllerRoleRevoked)
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
		it.Event = new(Web3RegisterControllerRoleRevoked)
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
func (it *Web3RegisterControllerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3RegisterControllerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3RegisterControllerRoleRevoked represents a RoleRevoked event raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// Web3RegisterControllerRoleRevokedOrChainReorg represents a RoleRevoked subscription event raised by the Web3RegisterController contract.
type Web3RegisterControllerRoleRevokedOrChainReorg struct {
	Event      *Web3RegisterControllerRoleRevoked
	ChainReorg *types.ChainReorg
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Web3RegisterController *Web3RegisterControllerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Web3RegisterControllerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, err := _Web3RegisterController.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Web3RegisterControllerRoleRevokedIterator{contract: _Web3RegisterController.contract, event: "RoleRevoked", logs: logs}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Web3RegisterController *Web3RegisterControllerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Web3RegisterControllerRoleRevokedOrChainReorg, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Web3RegisterController.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}

	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3RegisterControllerRoleRevokedOrChainReorg)
				event.Event = new(Web3RegisterControllerRoleRevoked)

				if log.ChainReorg == nil {
					if err := _Web3RegisterController.contract.UnpackLog(event.Event, "RoleRevoked", *log.Log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Web3RegisterController *Web3RegisterControllerFilterer) ParseRoleRevoked(log types.Log) (*Web3RegisterControllerRoleRevoked, error) {
	event := new(Web3RegisterControllerRoleRevoked)
	if err := _Web3RegisterController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
