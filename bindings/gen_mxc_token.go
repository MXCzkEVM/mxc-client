// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ERC20VotesUpgradeableCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ERC20VotesUpgradeableCheckpoint struct {
	FromBlock uint32
	Votes     *big.Int
}

// MxcTokenClientMetaData contains all meta data concerning the MxcTokenClient contract.
var MxcTokenClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MXC_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MXC_INVALID_PREMINT_PARAMS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MXC_MINT_DISALLOWED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"votes\",\"type\":\"uint224\"}],\"internalType\":\"structERC20VotesUpgradeable.Checkpoint\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_premintRecipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_premintAmounts\",\"type\":\"uint256[]\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldMxcToken_\",\"type\":\"address\"}],\"name\":\"setMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"syncL2Burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MxcTokenClientABI is the input ABI used to generate the binding from.
// Deprecated: Use MxcTokenClientMetaData.ABI instead.
var MxcTokenClientABI = MxcTokenClientMetaData.ABI

// MxcTokenClient is an auto generated Go binding around an Ethereum contract.
type MxcTokenClient struct {
	MxcTokenClientCaller     // Read-only binding to the contract
	MxcTokenClientTransactor // Write-only binding to the contract
	MxcTokenClientFilterer   // Log filterer for contract events
}

// MxcTokenClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type MxcTokenClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcTokenClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MxcTokenClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcTokenClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MxcTokenClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcTokenClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MxcTokenClientSession struct {
	Contract     *MxcTokenClient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MxcTokenClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MxcTokenClientCallerSession struct {
	Contract *MxcTokenClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MxcTokenClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MxcTokenClientTransactorSession struct {
	Contract     *MxcTokenClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MxcTokenClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type MxcTokenClientRaw struct {
	Contract *MxcTokenClient // Generic contract binding to access the raw methods on
}

// MxcTokenClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MxcTokenClientCallerRaw struct {
	Contract *MxcTokenClientCaller // Generic read-only contract binding to access the raw methods on
}

// MxcTokenClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MxcTokenClientTransactorRaw struct {
	Contract *MxcTokenClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMxcTokenClient creates a new instance of MxcTokenClient, bound to a specific deployed contract.
func NewMxcTokenClient(address common.Address, backend bind.ContractBackend) (*MxcTokenClient, error) {
	contract, err := bindMxcTokenClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClient{MxcTokenClientCaller: MxcTokenClientCaller{contract: contract}, MxcTokenClientTransactor: MxcTokenClientTransactor{contract: contract}, MxcTokenClientFilterer: MxcTokenClientFilterer{contract: contract}}, nil
}

// NewMxcTokenClientCaller creates a new read-only instance of MxcTokenClient, bound to a specific deployed contract.
func NewMxcTokenClientCaller(address common.Address, caller bind.ContractCaller) (*MxcTokenClientCaller, error) {
	contract, err := bindMxcTokenClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientCaller{contract: contract}, nil
}

// NewMxcTokenClientTransactor creates a new write-only instance of MxcTokenClient, bound to a specific deployed contract.
func NewMxcTokenClientTransactor(address common.Address, transactor bind.ContractTransactor) (*MxcTokenClientTransactor, error) {
	contract, err := bindMxcTokenClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientTransactor{contract: contract}, nil
}

// NewMxcTokenClientFilterer creates a new log filterer instance of MxcTokenClient, bound to a specific deployed contract.
func NewMxcTokenClientFilterer(address common.Address, filterer bind.ContractFilterer) (*MxcTokenClientFilterer, error) {
	contract, err := bindMxcTokenClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientFilterer{contract: contract}, nil
}

// bindMxcTokenClient binds a generic wrapper to an already deployed contract.
func bindMxcTokenClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MxcTokenClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MxcTokenClient *MxcTokenClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MxcTokenClient.Contract.MxcTokenClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MxcTokenClient *MxcTokenClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.MxcTokenClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MxcTokenClient *MxcTokenClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.MxcTokenClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MxcTokenClient *MxcTokenClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MxcTokenClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MxcTokenClient *MxcTokenClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MxcTokenClient *MxcTokenClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MxcTokenClient *MxcTokenClientCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MxcTokenClient *MxcTokenClientSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MxcTokenClient.Contract.DOMAINSEPARATOR(&_MxcTokenClient.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MxcTokenClient *MxcTokenClientCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MxcTokenClient.Contract.DOMAINSEPARATOR(&_MxcTokenClient.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcTokenClient *MxcTokenClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcTokenClient *MxcTokenClientSession) AddressManager() (common.Address, error) {
	return _MxcTokenClient.Contract.AddressManager(&_MxcTokenClient.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcTokenClient *MxcTokenClientCallerSession) AddressManager() (common.Address, error) {
	return _MxcTokenClient.Contract.AddressManager(&_MxcTokenClient.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.Allowance(&_MxcTokenClient.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.Allowance(&_MxcTokenClient.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.BalanceOf(&_MxcTokenClient.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.BalanceOf(&_MxcTokenClient.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.BalanceOfAt(&_MxcTokenClient.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.BalanceOfAt(&_MxcTokenClient.CallOpts, account, snapshotId)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_MxcTokenClient *MxcTokenClientCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (ERC20VotesUpgradeableCheckpoint, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(ERC20VotesUpgradeableCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20VotesUpgradeableCheckpoint)).(*ERC20VotesUpgradeableCheckpoint)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_MxcTokenClient *MxcTokenClientSession) Checkpoints(account common.Address, pos uint32) (ERC20VotesUpgradeableCheckpoint, error) {
	return _MxcTokenClient.Contract.Checkpoints(&_MxcTokenClient.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_MxcTokenClient *MxcTokenClientCallerSession) Checkpoints(account common.Address, pos uint32) (ERC20VotesUpgradeableCheckpoint, error) {
	return _MxcTokenClient.Contract.Checkpoints(&_MxcTokenClient.CallOpts, account, pos)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MxcTokenClient *MxcTokenClientCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MxcTokenClient *MxcTokenClientSession) Decimals() (uint8, error) {
	return _MxcTokenClient.Contract.Decimals(&_MxcTokenClient.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MxcTokenClient *MxcTokenClientCallerSession) Decimals() (uint8, error) {
	return _MxcTokenClient.Contract.Decimals(&_MxcTokenClient.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_MxcTokenClient *MxcTokenClientCaller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_MxcTokenClient *MxcTokenClientSession) Delegates(account common.Address) (common.Address, error) {
	return _MxcTokenClient.Contract.Delegates(&_MxcTokenClient.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_MxcTokenClient *MxcTokenClientCallerSession) Delegates(account common.Address) (common.Address, error) {
	return _MxcTokenClient.Contract.Delegates(&_MxcTokenClient.CallOpts, account)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) GetPastTotalSupply(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "getPastTotalSupply", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.GetPastTotalSupply(&_MxcTokenClient.CallOpts, blockNumber)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.GetPastTotalSupply(&_MxcTokenClient.CallOpts, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "getPastVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.GetPastVotes(&_MxcTokenClient.CallOpts, account, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.GetPastVotes(&_MxcTokenClient.CallOpts, account, blockNumber)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) GetVotes(account common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.GetVotes(&_MxcTokenClient.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.GetVotes(&_MxcTokenClient.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MxcTokenClient *MxcTokenClientCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MxcTokenClient *MxcTokenClientSession) Name() (string, error) {
	return _MxcTokenClient.Contract.Name(&_MxcTokenClient.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MxcTokenClient *MxcTokenClientCallerSession) Name() (string, error) {
	return _MxcTokenClient.Contract.Name(&_MxcTokenClient.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.Nonces(&_MxcTokenClient.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MxcTokenClient.Contract.Nonces(&_MxcTokenClient.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_MxcTokenClient *MxcTokenClientCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_MxcTokenClient *MxcTokenClientSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _MxcTokenClient.Contract.NumCheckpoints(&_MxcTokenClient.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_MxcTokenClient *MxcTokenClientCallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _MxcTokenClient.Contract.NumCheckpoints(&_MxcTokenClient.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcTokenClient *MxcTokenClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcTokenClient *MxcTokenClientSession) Owner() (common.Address, error) {
	return _MxcTokenClient.Contract.Owner(&_MxcTokenClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcTokenClient *MxcTokenClientCallerSession) Owner() (common.Address, error) {
	return _MxcTokenClient.Contract.Owner(&_MxcTokenClient.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MxcTokenClient *MxcTokenClientCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MxcTokenClient *MxcTokenClientSession) Paused() (bool, error) {
	return _MxcTokenClient.Contract.Paused(&_MxcTokenClient.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MxcTokenClient *MxcTokenClientCallerSession) Paused() (bool, error) {
	return _MxcTokenClient.Contract.Paused(&_MxcTokenClient.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcTokenClient *MxcTokenClientCaller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcTokenClient *MxcTokenClientSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcTokenClient.Contract.Resolve(&_MxcTokenClient.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcTokenClient *MxcTokenClientCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcTokenClient.Contract.Resolve(&_MxcTokenClient.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcTokenClient *MxcTokenClientCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcTokenClient *MxcTokenClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcTokenClient.Contract.Resolve0(&_MxcTokenClient.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcTokenClient *MxcTokenClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcTokenClient.Contract.Resolve0(&_MxcTokenClient.CallOpts, name, allowZeroAddress)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MxcTokenClient *MxcTokenClientCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MxcTokenClient *MxcTokenClientSession) Symbol() (string, error) {
	return _MxcTokenClient.Contract.Symbol(&_MxcTokenClient.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MxcTokenClient *MxcTokenClientCallerSession) Symbol() (string, error) {
	return _MxcTokenClient.Contract.Symbol(&_MxcTokenClient.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) TotalSupply() (*big.Int, error) {
	return _MxcTokenClient.Contract.TotalSupply(&_MxcTokenClient.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) TotalSupply() (*big.Int, error) {
	return _MxcTokenClient.Contract.TotalSupply(&_MxcTokenClient.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCaller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MxcTokenClient.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.TotalSupplyAt(&_MxcTokenClient.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_MxcTokenClient *MxcTokenClientCallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _MxcTokenClient.Contract.TotalSupplyAt(&_MxcTokenClient.CallOpts, snapshotId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Approve(&_MxcTokenClient.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Approve(&_MxcTokenClient.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Burn(&_MxcTokenClient.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Burn(&_MxcTokenClient.TransactOpts, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Burn0(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "burn0", from, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientSession) Burn0(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Burn0(&_MxcTokenClient.TransactOpts, from, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Burn0(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Burn0(&_MxcTokenClient.TransactOpts, from, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.BurnFrom(&_MxcTokenClient.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.BurnFrom(&_MxcTokenClient.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MxcTokenClient *MxcTokenClientSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.DecreaseAllowance(&_MxcTokenClient.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.DecreaseAllowance(&_MxcTokenClient.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_MxcTokenClient *MxcTokenClientSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Delegate(&_MxcTokenClient.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Delegate(&_MxcTokenClient.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_MxcTokenClient *MxcTokenClientSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.DelegateBySig(&_MxcTokenClient.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.DelegateBySig(&_MxcTokenClient.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(address to, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Exchange(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "exchange", to, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(address to, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientSession) Exchange(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Exchange(&_MxcTokenClient.TransactOpts, to, amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x045d0389.
//
// Solidity: function exchange(address to, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Exchange(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Exchange(&_MxcTokenClient.TransactOpts, to, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MxcTokenClient *MxcTokenClientSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.IncreaseAllowance(&_MxcTokenClient.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.IncreaseAllowance(&_MxcTokenClient.TransactOpts, spender, addedValue)
}

// Init is a paid mutator transaction binding the contract method 0xc3a57123.
//
// Solidity: function init(address _addressManager, string _name, string _symbol, address[] _premintRecipients, uint256[] _premintAmounts) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _name string, _symbol string, _premintRecipients []common.Address, _premintAmounts []*big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "init", _addressManager, _name, _symbol, _premintRecipients, _premintAmounts)
}

// Init is a paid mutator transaction binding the contract method 0xc3a57123.
//
// Solidity: function init(address _addressManager, string _name, string _symbol, address[] _premintRecipients, uint256[] _premintAmounts) returns()
func (_MxcTokenClient *MxcTokenClientSession) Init(_addressManager common.Address, _name string, _symbol string, _premintRecipients []common.Address, _premintAmounts []*big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Init(&_MxcTokenClient.TransactOpts, _addressManager, _name, _symbol, _premintRecipients, _premintAmounts)
}

// Init is a paid mutator transaction binding the contract method 0xc3a57123.
//
// Solidity: function init(address _addressManager, string _name, string _symbol, address[] _premintRecipients, uint256[] _premintAmounts) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Init(_addressManager common.Address, _name string, _symbol string, _premintRecipients []common.Address, _premintAmounts []*big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Init(&_MxcTokenClient.TransactOpts, _addressManager, _name, _symbol, _premintRecipients, _premintAmounts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Mint(&_MxcTokenClient.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Mint(&_MxcTokenClient.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MxcTokenClient *MxcTokenClientSession) Pause() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Pause(&_MxcTokenClient.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Pause() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Pause(&_MxcTokenClient.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MxcTokenClient *MxcTokenClientSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Permit(&_MxcTokenClient.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Permit(&_MxcTokenClient.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcTokenClient *MxcTokenClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcTokenClient *MxcTokenClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.RenounceOwnership(&_MxcTokenClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.RenounceOwnership(&_MxcTokenClient.TransactOpts)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) SetAddressManager(opts *bind.TransactOpts, newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "setAddressManager", newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcTokenClient *MxcTokenClientSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.SetAddressManager(&_MxcTokenClient.TransactOpts, newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.SetAddressManager(&_MxcTokenClient.TransactOpts, newAddressManager)
}

// SetMxcToken is a paid mutator transaction binding the contract method 0x59e302f8.
//
// Solidity: function setMxcToken(address oldMxcToken_) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) SetMxcToken(opts *bind.TransactOpts, oldMxcToken_ common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "setMxcToken", oldMxcToken_)
}

// SetMxcToken is a paid mutator transaction binding the contract method 0x59e302f8.
//
// Solidity: function setMxcToken(address oldMxcToken_) returns()
func (_MxcTokenClient *MxcTokenClientSession) SetMxcToken(oldMxcToken_ common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.SetMxcToken(&_MxcTokenClient.TransactOpts, oldMxcToken_)
}

// SetMxcToken is a paid mutator transaction binding the contract method 0x59e302f8.
//
// Solidity: function setMxcToken(address oldMxcToken_) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) SetMxcToken(oldMxcToken_ common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.SetMxcToken(&_MxcTokenClient.TransactOpts, oldMxcToken_)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_MxcTokenClient *MxcTokenClientSession) Snapshot() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Snapshot(&_MxcTokenClient.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Snapshot() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Snapshot(&_MxcTokenClient.TransactOpts)
}

// SyncL2Burn is a paid mutator transaction binding the contract method 0x62451ba2.
//
// Solidity: function syncL2Burn(uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) SyncL2Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "syncL2Burn", amount)
}

// SyncL2Burn is a paid mutator transaction binding the contract method 0x62451ba2.
//
// Solidity: function syncL2Burn(uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientSession) SyncL2Burn(amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.SyncL2Burn(&_MxcTokenClient.TransactOpts, amount)
}

// SyncL2Burn is a paid mutator transaction binding the contract method 0x62451ba2.
//
// Solidity: function syncL2Burn(uint256 amount) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) SyncL2Burn(amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.SyncL2Burn(&_MxcTokenClient.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Transfer(&_MxcTokenClient.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Transfer(&_MxcTokenClient.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.TransferFrom(&_MxcTokenClient.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MxcTokenClient *MxcTokenClientTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.TransferFrom(&_MxcTokenClient.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcTokenClient *MxcTokenClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcTokenClient *MxcTokenClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.TransferOwnership(&_MxcTokenClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MxcTokenClient.Contract.TransferOwnership(&_MxcTokenClient.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MxcTokenClient *MxcTokenClientTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcTokenClient.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MxcTokenClient *MxcTokenClientSession) Unpause() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Unpause(&_MxcTokenClient.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MxcTokenClient *MxcTokenClientTransactorSession) Unpause() (*types.Transaction, error) {
	return _MxcTokenClient.Contract.Unpause(&_MxcTokenClient.TransactOpts)
}

// MxcTokenClientAddressManagerChangedIterator is returned from FilterAddressManagerChanged and is used to iterate over the raw logs and unpacked data for AddressManagerChanged events raised by the MxcTokenClient contract.
type MxcTokenClientAddressManagerChangedIterator struct {
	Event *MxcTokenClientAddressManagerChanged // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientAddressManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientAddressManagerChanged)
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
		it.Event = new(MxcTokenClientAddressManagerChanged)
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
func (it *MxcTokenClientAddressManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientAddressManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientAddressManagerChanged represents a AddressManagerChanged event raised by the MxcTokenClient contract.
type MxcTokenClientAddressManagerChanged struct {
	AddressManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressManagerChanged is a free log retrieval operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterAddressManagerChanged(opts *bind.FilterOpts) (*MxcTokenClientAddressManagerChangedIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientAddressManagerChangedIterator{contract: _MxcTokenClient.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *MxcTokenClientAddressManagerChanged) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientAddressManagerChanged)
				if err := _MxcTokenClient.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
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

// ParseAddressManagerChanged is a log parse operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseAddressManagerChanged(log types.Log) (*MxcTokenClientAddressManagerChanged, error) {
	event := new(MxcTokenClientAddressManagerChanged)
	if err := _MxcTokenClient.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MxcTokenClient contract.
type MxcTokenClientApprovalIterator struct {
	Event *MxcTokenClientApproval // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientApproval)
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
		it.Event = new(MxcTokenClientApproval)
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
func (it *MxcTokenClientApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientApproval represents a Approval event raised by the MxcTokenClient contract.
type MxcTokenClientApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MxcTokenClientApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientApprovalIterator{contract: _MxcTokenClient.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MxcTokenClientApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientApproval)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseApproval(log types.Log) (*MxcTokenClientApproval, error) {
	event := new(MxcTokenClientApproval)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the MxcTokenClient contract.
type MxcTokenClientBurnIterator struct {
	Event *MxcTokenClientBurn // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientBurn)
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
		it.Event = new(MxcTokenClientBurn)
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
func (it *MxcTokenClientBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientBurn represents a Burn event raised by the MxcTokenClient contract.
type MxcTokenClientBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address account, uint256 amount)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterBurn(opts *bind.FilterOpts) (*MxcTokenClientBurnIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientBurnIterator{contract: _MxcTokenClient.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address account, uint256 amount)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *MxcTokenClientBurn) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientBurn)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address account, uint256 amount)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseBurn(log types.Log) (*MxcTokenClientBurn, error) {
	event := new(MxcTokenClientBurn)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the MxcTokenClient contract.
type MxcTokenClientDelegateChangedIterator struct {
	Event *MxcTokenClientDelegateChanged // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientDelegateChanged)
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
		it.Event = new(MxcTokenClientDelegateChanged)
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
func (it *MxcTokenClientDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientDelegateChanged represents a DelegateChanged event raised by the MxcTokenClient contract.
type MxcTokenClientDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*MxcTokenClientDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientDelegateChangedIterator{contract: _MxcTokenClient.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *MxcTokenClientDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientDelegateChanged)
				if err := _MxcTokenClient.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseDelegateChanged(log types.Log) (*MxcTokenClientDelegateChanged, error) {
	event := new(MxcTokenClientDelegateChanged)
	if err := _MxcTokenClient.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the MxcTokenClient contract.
type MxcTokenClientDelegateVotesChangedIterator struct {
	Event *MxcTokenClientDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientDelegateVotesChanged)
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
		it.Event = new(MxcTokenClientDelegateVotesChanged)
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
func (it *MxcTokenClientDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientDelegateVotesChanged represents a DelegateVotesChanged event raised by the MxcTokenClient contract.
type MxcTokenClientDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*MxcTokenClientDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientDelegateVotesChangedIterator{contract: _MxcTokenClient.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *MxcTokenClientDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientDelegateVotesChanged)
				if err := _MxcTokenClient.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseDelegateVotesChanged(log types.Log) (*MxcTokenClientDelegateVotesChanged, error) {
	event := new(MxcTokenClientDelegateVotesChanged)
	if err := _MxcTokenClient.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MxcTokenClient contract.
type MxcTokenClientInitializedIterator struct {
	Event *MxcTokenClientInitialized // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientInitialized)
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
		it.Event = new(MxcTokenClientInitialized)
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
func (it *MxcTokenClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientInitialized represents a Initialized event raised by the MxcTokenClient contract.
type MxcTokenClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*MxcTokenClientInitializedIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientInitializedIterator{contract: _MxcTokenClient.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MxcTokenClientInitialized) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientInitialized)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseInitialized(log types.Log) (*MxcTokenClientInitialized, error) {
	event := new(MxcTokenClientInitialized)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MxcTokenClient contract.
type MxcTokenClientMintIterator struct {
	Event *MxcTokenClientMint // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientMint)
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
		it.Event = new(MxcTokenClientMint)
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
func (it *MxcTokenClientMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientMint represents a Mint event raised by the MxcTokenClient contract.
type MxcTokenClientMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address account, uint256 amount)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterMint(opts *bind.FilterOpts) (*MxcTokenClientMintIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientMintIterator{contract: _MxcTokenClient.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address account, uint256 amount)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MxcTokenClientMint) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientMint)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address account, uint256 amount)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseMint(log types.Log) (*MxcTokenClientMint, error) {
	event := new(MxcTokenClientMint)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MxcTokenClient contract.
type MxcTokenClientOwnershipTransferredIterator struct {
	Event *MxcTokenClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientOwnershipTransferred)
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
		it.Event = new(MxcTokenClientOwnershipTransferred)
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
func (it *MxcTokenClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientOwnershipTransferred represents a OwnershipTransferred event raised by the MxcTokenClient contract.
type MxcTokenClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MxcTokenClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientOwnershipTransferredIterator{contract: _MxcTokenClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MxcTokenClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientOwnershipTransferred)
				if err := _MxcTokenClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MxcTokenClient *MxcTokenClientFilterer) ParseOwnershipTransferred(log types.Log) (*MxcTokenClientOwnershipTransferred, error) {
	event := new(MxcTokenClientOwnershipTransferred)
	if err := _MxcTokenClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MxcTokenClient contract.
type MxcTokenClientPausedIterator struct {
	Event *MxcTokenClientPaused // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientPaused)
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
		it.Event = new(MxcTokenClientPaused)
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
func (it *MxcTokenClientPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientPaused represents a Paused event raised by the MxcTokenClient contract.
type MxcTokenClientPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterPaused(opts *bind.FilterOpts) (*MxcTokenClientPausedIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientPausedIterator{contract: _MxcTokenClient.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MxcTokenClientPaused) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientPaused)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MxcTokenClient *MxcTokenClientFilterer) ParsePaused(log types.Log) (*MxcTokenClientPaused, error) {
	event := new(MxcTokenClientPaused)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the MxcTokenClient contract.
type MxcTokenClientSnapshotIterator struct {
	Event *MxcTokenClientSnapshot // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientSnapshot)
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
		it.Event = new(MxcTokenClientSnapshot)
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
func (it *MxcTokenClientSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientSnapshot represents a Snapshot event raised by the MxcTokenClient contract.
type MxcTokenClientSnapshot struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterSnapshot(opts *bind.FilterOpts) (*MxcTokenClientSnapshotIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientSnapshotIterator{contract: _MxcTokenClient.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *MxcTokenClientSnapshot) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientSnapshot)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x8030e83b04d87bef53480e26263266d6ca66863aa8506aca6f2559d18aa1cb67.
//
// Solidity: event Snapshot(uint256 id)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseSnapshot(log types.Log) (*MxcTokenClientSnapshot, error) {
	event := new(MxcTokenClientSnapshot)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MxcTokenClient contract.
type MxcTokenClientTransferIterator struct {
	Event *MxcTokenClientTransfer // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientTransfer)
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
		it.Event = new(MxcTokenClientTransfer)
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
func (it *MxcTokenClientTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientTransfer represents a Transfer event raised by the MxcTokenClient contract.
type MxcTokenClientTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MxcTokenClientTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientTransferIterator{contract: _MxcTokenClient.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MxcTokenClientTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientTransfer)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseTransfer(log types.Log) (*MxcTokenClientTransfer, error) {
	event := new(MxcTokenClientTransfer)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcTokenClientUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MxcTokenClient contract.
type MxcTokenClientUnpausedIterator struct {
	Event *MxcTokenClientUnpaused // Event containing the contract specifics and raw log

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
func (it *MxcTokenClientUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcTokenClientUnpaused)
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
		it.Event = new(MxcTokenClientUnpaused)
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
func (it *MxcTokenClientUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcTokenClientUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcTokenClientUnpaused represents a Unpaused event raised by the MxcTokenClient contract.
type MxcTokenClientUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MxcTokenClient *MxcTokenClientFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MxcTokenClientUnpausedIterator, error) {

	logs, sub, err := _MxcTokenClient.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MxcTokenClientUnpausedIterator{contract: _MxcTokenClient.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MxcTokenClient *MxcTokenClientFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MxcTokenClientUnpaused) (event.Subscription, error) {

	logs, sub, err := _MxcTokenClient.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcTokenClientUnpaused)
				if err := _MxcTokenClient.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MxcTokenClient *MxcTokenClientFilterer) ParseUnpaused(log types.Log) (*MxcTokenClientUnpaused, error) {
	event := new(MxcTokenClientUnpaused)
	if err := _MxcTokenClient.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
