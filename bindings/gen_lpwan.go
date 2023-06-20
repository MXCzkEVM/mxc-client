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

// LPWANRelaySyncStatus is an auto generated low-level Go binding around an user-defined struct.
type LPWANRelaySyncStatus struct {
	ProposedRewardEventHeight *big.Int
	ProvedRewardEventHeight   *big.Int
}

// LPWANRewardData is an auto generated low-level Go binding around an user-defined struct.
type LPWANRewardData struct {
	ProposedReward     *big.Int
	ProvenReward       *big.Int
	ProposedCostReward *big.Int
	ProvenCostReward   *big.Int
}

// LPWANRewardEvent is an auto generated low-level Go binding around an user-defined struct.
type LPWANRewardEvent struct {
	RewardHeight *big.Int
	Account      common.Address
	Amount       *big.Int
	Cost         *big.Int
}

// LPWANClientMetaData contains all meta data concerning the LPWANClient contract.
var LPWANClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"RELAY_INVALID_HEIGHT\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnExcessToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"burn\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"ControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"__Controllable_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnExcessToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"burn\",\"type\":\"bool\"}],\"name\":\"claimAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"burn\",\"type\":\"bool\"}],\"name\":\"claimProposedCostReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"burn\",\"type\":\"bool\"}],\"name\":\"claimProposedReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"burn\",\"type\":\"bool\"}],\"name\":\"claimProvenCostReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"burn\",\"type\":\"bool\"}],\"name\":\"claimProvenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelaySyncStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ProposedRewardEventHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ProvedRewardEventHeight\",\"type\":\"uint256\"}],\"internalType\":\"structLPWAN.RelaySyncStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRewardData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provenReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposedCostReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provenCostReward\",\"type\":\"uint256\"}],\"internalType\":\"structLPWAN.RewardData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRewardData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provenReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposedCostReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provenCostReward\",\"type\":\"uint256\"}],\"internalType\":\"structLPWAN.RewardData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"MEP1004Address_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_SNCode\",\"type\":\"string\"}],\"name\":\"mintMEP1004Stations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MEP1002TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_MEP1004TokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_item\",\"type\":\"string\"}],\"name\":\"submitLocationProofs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"internalType\":\"structLPWAN.RewardEvent[]\",\"name\":\"rewardEvents\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"setting\",\"type\":\"bool\"}],\"name\":\"syncProposedRewardEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"internalType\":\"structLPWAN.RewardEvent[]\",\"name\":\"rewardEvents\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"setting\",\"type\":\"bool\"}],\"name\":\"syncProvenRewardEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LPWANClientABI is the input ABI used to generate the binding from.
// Deprecated: Use LPWANClientMetaData.ABI instead.
var LPWANClientABI = LPWANClientMetaData.ABI

// LPWANClient is an auto generated Go binding around an Ethereum contract.
type LPWANClient struct {
	LPWANClientCaller     // Read-only binding to the contract
	LPWANClientTransactor // Write-only binding to the contract
	LPWANClientFilterer   // Log filterer for contract events
}

// LPWANClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type LPWANClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPWANClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LPWANClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPWANClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LPWANClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPWANClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LPWANClientSession struct {
	Contract     *LPWANClient      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LPWANClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LPWANClientCallerSession struct {
	Contract *LPWANClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LPWANClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LPWANClientTransactorSession struct {
	Contract     *LPWANClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LPWANClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type LPWANClientRaw struct {
	Contract *LPWANClient // Generic contract binding to access the raw methods on
}

// LPWANClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LPWANClientCallerRaw struct {
	Contract *LPWANClientCaller // Generic read-only contract binding to access the raw methods on
}

// LPWANClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LPWANClientTransactorRaw struct {
	Contract *LPWANClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLPWANClient creates a new instance of LPWANClient, bound to a specific deployed contract.
func NewLPWANClient(address common.Address, backend bind.ContractBackend) (*LPWANClient, error) {
	contract, err := bindLPWANClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LPWANClient{LPWANClientCaller: LPWANClientCaller{contract: contract}, LPWANClientTransactor: LPWANClientTransactor{contract: contract}, LPWANClientFilterer: LPWANClientFilterer{contract: contract}}, nil
}

// NewLPWANClientCaller creates a new read-only instance of LPWANClient, bound to a specific deployed contract.
func NewLPWANClientCaller(address common.Address, caller bind.ContractCaller) (*LPWANClientCaller, error) {
	contract, err := bindLPWANClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LPWANClientCaller{contract: contract}, nil
}

// NewLPWANClientTransactor creates a new write-only instance of LPWANClient, bound to a specific deployed contract.
func NewLPWANClientTransactor(address common.Address, transactor bind.ContractTransactor) (*LPWANClientTransactor, error) {
	contract, err := bindLPWANClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LPWANClientTransactor{contract: contract}, nil
}

// NewLPWANClientFilterer creates a new log filterer instance of LPWANClient, bound to a specific deployed contract.
func NewLPWANClientFilterer(address common.Address, filterer bind.ContractFilterer) (*LPWANClientFilterer, error) {
	contract, err := bindLPWANClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LPWANClientFilterer{contract: contract}, nil
}

// bindLPWANClient binds a generic wrapper to an already deployed contract.
func bindLPWANClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LPWANClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LPWANClient *LPWANClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LPWANClient.Contract.LPWANClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LPWANClient *LPWANClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPWANClient.Contract.LPWANClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LPWANClient *LPWANClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LPWANClient.Contract.LPWANClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LPWANClient *LPWANClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LPWANClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LPWANClient *LPWANClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPWANClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LPWANClient *LPWANClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LPWANClient.Contract.contract.Transact(opts, method, params...)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_LPWANClient *LPWANClientCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LPWANClient.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_LPWANClient *LPWANClientSession) Controllers(arg0 common.Address) (bool, error) {
	return _LPWANClient.Contract.Controllers(&_LPWANClient.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_LPWANClient *LPWANClientCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _LPWANClient.Contract.Controllers(&_LPWANClient.CallOpts, arg0)
}

// GetRelaySyncStatus is a free data retrieval call binding the contract method 0xe3fb4590.
//
// Solidity: function getRelaySyncStatus() view returns((uint256,uint256))
func (_LPWANClient *LPWANClientCaller) GetRelaySyncStatus(opts *bind.CallOpts) (LPWANRelaySyncStatus, error) {
	var out []interface{}
	err := _LPWANClient.contract.Call(opts, &out, "getRelaySyncStatus")

	if err != nil {
		return *new(LPWANRelaySyncStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(LPWANRelaySyncStatus)).(*LPWANRelaySyncStatus)

	return out0, err

}

// GetRelaySyncStatus is a free data retrieval call binding the contract method 0xe3fb4590.
//
// Solidity: function getRelaySyncStatus() view returns((uint256,uint256))
func (_LPWANClient *LPWANClientSession) GetRelaySyncStatus() (LPWANRelaySyncStatus, error) {
	return _LPWANClient.Contract.GetRelaySyncStatus(&_LPWANClient.CallOpts)
}

// GetRelaySyncStatus is a free data retrieval call binding the contract method 0xe3fb4590.
//
// Solidity: function getRelaySyncStatus() view returns((uint256,uint256))
func (_LPWANClient *LPWANClientCallerSession) GetRelaySyncStatus() (LPWANRelaySyncStatus, error) {
	return _LPWANClient.Contract.GetRelaySyncStatus(&_LPWANClient.CallOpts)
}

// GetRewardData is a free data retrieval call binding the contract method 0x451831e4.
//
// Solidity: function getRewardData(address account) view returns((uint256,uint256,uint256,uint256))
func (_LPWANClient *LPWANClientCaller) GetRewardData(opts *bind.CallOpts, account common.Address) (LPWANRewardData, error) {
	var out []interface{}
	err := _LPWANClient.contract.Call(opts, &out, "getRewardData", account)

	if err != nil {
		return *new(LPWANRewardData), err
	}

	out0 := *abi.ConvertType(out[0], new(LPWANRewardData)).(*LPWANRewardData)

	return out0, err

}

// GetRewardData is a free data retrieval call binding the contract method 0x451831e4.
//
// Solidity: function getRewardData(address account) view returns((uint256,uint256,uint256,uint256))
func (_LPWANClient *LPWANClientSession) GetRewardData(account common.Address) (LPWANRewardData, error) {
	return _LPWANClient.Contract.GetRewardData(&_LPWANClient.CallOpts, account)
}

// GetRewardData is a free data retrieval call binding the contract method 0x451831e4.
//
// Solidity: function getRewardData(address account) view returns((uint256,uint256,uint256,uint256))
func (_LPWANClient *LPWANClientCallerSession) GetRewardData(account common.Address) (LPWANRewardData, error) {
	return _LPWANClient.Contract.GetRewardData(&_LPWANClient.CallOpts, account)
}

// GetTotalRewardData is a free data retrieval call binding the contract method 0x6fa4e42d.
//
// Solidity: function getTotalRewardData() view returns((uint256,uint256,uint256,uint256))
func (_LPWANClient *LPWANClientCaller) GetTotalRewardData(opts *bind.CallOpts) (LPWANRewardData, error) {
	var out []interface{}
	err := _LPWANClient.contract.Call(opts, &out, "getTotalRewardData")

	if err != nil {
		return *new(LPWANRewardData), err
	}

	out0 := *abi.ConvertType(out[0], new(LPWANRewardData)).(*LPWANRewardData)

	return out0, err

}

// GetTotalRewardData is a free data retrieval call binding the contract method 0x6fa4e42d.
//
// Solidity: function getTotalRewardData() view returns((uint256,uint256,uint256,uint256))
func (_LPWANClient *LPWANClientSession) GetTotalRewardData() (LPWANRewardData, error) {
	return _LPWANClient.Contract.GetTotalRewardData(&_LPWANClient.CallOpts)
}

// GetTotalRewardData is a free data retrieval call binding the contract method 0x6fa4e42d.
//
// Solidity: function getTotalRewardData() view returns((uint256,uint256,uint256,uint256))
func (_LPWANClient *LPWANClientCallerSession) GetTotalRewardData() (LPWANRewardData, error) {
	return _LPWANClient.Contract.GetTotalRewardData(&_LPWANClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPWANClient *LPWANClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPWANClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPWANClient *LPWANClientSession) Owner() (common.Address, error) {
	return _LPWANClient.Contract.Owner(&_LPWANClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPWANClient *LPWANClientCallerSession) Owner() (common.Address, error) {
	return _LPWANClient.Contract.Owner(&_LPWANClient.CallOpts)
}

// ControllableInit is a paid mutator transaction binding the contract method 0x5d79343d.
//
// Solidity: function __Controllable_init() returns()
func (_LPWANClient *LPWANClientTransactor) ControllableInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "__Controllable_init")
}

// ControllableInit is a paid mutator transaction binding the contract method 0x5d79343d.
//
// Solidity: function __Controllable_init() returns()
func (_LPWANClient *LPWANClientSession) ControllableInit() (*types.Transaction, error) {
	return _LPWANClient.Contract.ControllableInit(&_LPWANClient.TransactOpts)
}

// ControllableInit is a paid mutator transaction binding the contract method 0x5d79343d.
//
// Solidity: function __Controllable_init() returns()
func (_LPWANClient *LPWANClientTransactorSession) ControllableInit() (*types.Transaction, error) {
	return _LPWANClient.Contract.ControllableInit(&_LPWANClient.TransactOpts)
}

// BurnExcessToken is a paid mutator transaction binding the contract method 0xf4a6842a.
//
// Solidity: function burnExcessToken(uint256 amount) returns()
func (_LPWANClient *LPWANClientTransactor) BurnExcessToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "burnExcessToken", amount)
}

// BurnExcessToken is a paid mutator transaction binding the contract method 0xf4a6842a.
//
// Solidity: function burnExcessToken(uint256 amount) returns()
func (_LPWANClient *LPWANClientSession) BurnExcessToken(amount *big.Int) (*types.Transaction, error) {
	return _LPWANClient.Contract.BurnExcessToken(&_LPWANClient.TransactOpts, amount)
}

// BurnExcessToken is a paid mutator transaction binding the contract method 0xf4a6842a.
//
// Solidity: function burnExcessToken(uint256 amount) returns()
func (_LPWANClient *LPWANClientTransactorSession) BurnExcessToken(amount *big.Int) (*types.Transaction, error) {
	return _LPWANClient.Contract.BurnExcessToken(&_LPWANClient.TransactOpts, amount)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x41ddcee4.
//
// Solidity: function claimAllReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactor) ClaimAllReward(opts *bind.TransactOpts, burn bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "claimAllReward", burn)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x41ddcee4.
//
// Solidity: function claimAllReward(bool burn) returns()
func (_LPWANClient *LPWANClientSession) ClaimAllReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimAllReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x41ddcee4.
//
// Solidity: function claimAllReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactorSession) ClaimAllReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimAllReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProposedCostReward is a paid mutator transaction binding the contract method 0x1368243a.
//
// Solidity: function claimProposedCostReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactor) ClaimProposedCostReward(opts *bind.TransactOpts, burn bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "claimProposedCostReward", burn)
}

// ClaimProposedCostReward is a paid mutator transaction binding the contract method 0x1368243a.
//
// Solidity: function claimProposedCostReward(bool burn) returns()
func (_LPWANClient *LPWANClientSession) ClaimProposedCostReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProposedCostReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProposedCostReward is a paid mutator transaction binding the contract method 0x1368243a.
//
// Solidity: function claimProposedCostReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactorSession) ClaimProposedCostReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProposedCostReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProposedReward is a paid mutator transaction binding the contract method 0x5abf3150.
//
// Solidity: function claimProposedReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactor) ClaimProposedReward(opts *bind.TransactOpts, burn bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "claimProposedReward", burn)
}

// ClaimProposedReward is a paid mutator transaction binding the contract method 0x5abf3150.
//
// Solidity: function claimProposedReward(bool burn) returns()
func (_LPWANClient *LPWANClientSession) ClaimProposedReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProposedReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProposedReward is a paid mutator transaction binding the contract method 0x5abf3150.
//
// Solidity: function claimProposedReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactorSession) ClaimProposedReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProposedReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProvenCostReward is a paid mutator transaction binding the contract method 0x694a61f1.
//
// Solidity: function claimProvenCostReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactor) ClaimProvenCostReward(opts *bind.TransactOpts, burn bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "claimProvenCostReward", burn)
}

// ClaimProvenCostReward is a paid mutator transaction binding the contract method 0x694a61f1.
//
// Solidity: function claimProvenCostReward(bool burn) returns()
func (_LPWANClient *LPWANClientSession) ClaimProvenCostReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProvenCostReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProvenCostReward is a paid mutator transaction binding the contract method 0x694a61f1.
//
// Solidity: function claimProvenCostReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactorSession) ClaimProvenCostReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProvenCostReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProvenReward is a paid mutator transaction binding the contract method 0xf5a0545e.
//
// Solidity: function claimProvenReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactor) ClaimProvenReward(opts *bind.TransactOpts, burn bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "claimProvenReward", burn)
}

// ClaimProvenReward is a paid mutator transaction binding the contract method 0xf5a0545e.
//
// Solidity: function claimProvenReward(bool burn) returns()
func (_LPWANClient *LPWANClientSession) ClaimProvenReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProvenReward(&_LPWANClient.TransactOpts, burn)
}

// ClaimProvenReward is a paid mutator transaction binding the contract method 0xf5a0545e.
//
// Solidity: function claimProvenReward(bool burn) returns()
func (_LPWANClient *LPWANClientTransactorSession) ClaimProvenReward(burn bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.ClaimProvenReward(&_LPWANClient.TransactOpts, burn)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address MEP1004Address_) returns()
func (_LPWANClient *LPWANClientTransactor) Initialize(opts *bind.TransactOpts, MEP1004Address_ common.Address) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "initialize", MEP1004Address_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address MEP1004Address_) returns()
func (_LPWANClient *LPWANClientSession) Initialize(MEP1004Address_ common.Address) (*types.Transaction, error) {
	return _LPWANClient.Contract.Initialize(&_LPWANClient.TransactOpts, MEP1004Address_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address MEP1004Address_) returns()
func (_LPWANClient *LPWANClientTransactorSession) Initialize(MEP1004Address_ common.Address) (*types.Transaction, error) {
	return _LPWANClient.Contract.Initialize(&_LPWANClient.TransactOpts, MEP1004Address_)
}

// MintMEP1004Stations is a paid mutator transaction binding the contract method 0xf2b9ec40.
//
// Solidity: function mintMEP1004Stations(address _to, string _SNCode) returns()
func (_LPWANClient *LPWANClientTransactor) MintMEP1004Stations(opts *bind.TransactOpts, _to common.Address, _SNCode string) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "mintMEP1004Stations", _to, _SNCode)
}

// MintMEP1004Stations is a paid mutator transaction binding the contract method 0xf2b9ec40.
//
// Solidity: function mintMEP1004Stations(address _to, string _SNCode) returns()
func (_LPWANClient *LPWANClientSession) MintMEP1004Stations(_to common.Address, _SNCode string) (*types.Transaction, error) {
	return _LPWANClient.Contract.MintMEP1004Stations(&_LPWANClient.TransactOpts, _to, _SNCode)
}

// MintMEP1004Stations is a paid mutator transaction binding the contract method 0xf2b9ec40.
//
// Solidity: function mintMEP1004Stations(address _to, string _SNCode) returns()
func (_LPWANClient *LPWANClientTransactorSession) MintMEP1004Stations(_to common.Address, _SNCode string) (*types.Transaction, error) {
	return _LPWANClient.Contract.MintMEP1004Stations(&_LPWANClient.TransactOpts, _to, _SNCode)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPWANClient *LPWANClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPWANClient *LPWANClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _LPWANClient.Contract.RenounceOwnership(&_LPWANClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPWANClient *LPWANClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LPWANClient.Contract.RenounceOwnership(&_LPWANClient.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_LPWANClient *LPWANClientTransactor) SetController(opts *bind.TransactOpts, controller common.Address, enabled bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "setController", controller, enabled)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_LPWANClient *LPWANClientSession) SetController(controller common.Address, enabled bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.SetController(&_LPWANClient.TransactOpts, controller, enabled)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool enabled) returns()
func (_LPWANClient *LPWANClientTransactorSession) SetController(controller common.Address, enabled bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.SetController(&_LPWANClient.TransactOpts, controller, enabled)
}

// SubmitLocationProofs is a paid mutator transaction binding the contract method 0xcdceeba0.
//
// Solidity: function submitLocationProofs(uint256 _MEP1002TokenId, uint256[] _MEP1004TokenIds, string _item) returns()
func (_LPWANClient *LPWANClientTransactor) SubmitLocationProofs(opts *bind.TransactOpts, _MEP1002TokenId *big.Int, _MEP1004TokenIds []*big.Int, _item string) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "submitLocationProofs", _MEP1002TokenId, _MEP1004TokenIds, _item)
}

// SubmitLocationProofs is a paid mutator transaction binding the contract method 0xcdceeba0.
//
// Solidity: function submitLocationProofs(uint256 _MEP1002TokenId, uint256[] _MEP1004TokenIds, string _item) returns()
func (_LPWANClient *LPWANClientSession) SubmitLocationProofs(_MEP1002TokenId *big.Int, _MEP1004TokenIds []*big.Int, _item string) (*types.Transaction, error) {
	return _LPWANClient.Contract.SubmitLocationProofs(&_LPWANClient.TransactOpts, _MEP1002TokenId, _MEP1004TokenIds, _item)
}

// SubmitLocationProofs is a paid mutator transaction binding the contract method 0xcdceeba0.
//
// Solidity: function submitLocationProofs(uint256 _MEP1002TokenId, uint256[] _MEP1004TokenIds, string _item) returns()
func (_LPWANClient *LPWANClientTransactorSession) SubmitLocationProofs(_MEP1002TokenId *big.Int, _MEP1004TokenIds []*big.Int, _item string) (*types.Transaction, error) {
	return _LPWANClient.Contract.SubmitLocationProofs(&_LPWANClient.TransactOpts, _MEP1002TokenId, _MEP1004TokenIds, _item)
}

// SyncProposedRewardEvent is a paid mutator transaction binding the contract method 0x23196cf2.
//
// Solidity: function syncProposedRewardEvent((uint256,address,uint256,uint256)[] rewardEvents, bool setting) returns()
func (_LPWANClient *LPWANClientTransactor) SyncProposedRewardEvent(opts *bind.TransactOpts, rewardEvents []LPWANRewardEvent, setting bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "syncProposedRewardEvent", rewardEvents, setting)
}

// SyncProposedRewardEvent is a paid mutator transaction binding the contract method 0x23196cf2.
//
// Solidity: function syncProposedRewardEvent((uint256,address,uint256,uint256)[] rewardEvents, bool setting) returns()
func (_LPWANClient *LPWANClientSession) SyncProposedRewardEvent(rewardEvents []LPWANRewardEvent, setting bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.SyncProposedRewardEvent(&_LPWANClient.TransactOpts, rewardEvents, setting)
}

// SyncProposedRewardEvent is a paid mutator transaction binding the contract method 0x23196cf2.
//
// Solidity: function syncProposedRewardEvent((uint256,address,uint256,uint256)[] rewardEvents, bool setting) returns()
func (_LPWANClient *LPWANClientTransactorSession) SyncProposedRewardEvent(rewardEvents []LPWANRewardEvent, setting bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.SyncProposedRewardEvent(&_LPWANClient.TransactOpts, rewardEvents, setting)
}

// SyncProvenRewardEvent is a paid mutator transaction binding the contract method 0xc6ea2882.
//
// Solidity: function syncProvenRewardEvent((uint256,address,uint256,uint256)[] rewardEvents, bool setting) returns()
func (_LPWANClient *LPWANClientTransactor) SyncProvenRewardEvent(opts *bind.TransactOpts, rewardEvents []LPWANRewardEvent, setting bool) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "syncProvenRewardEvent", rewardEvents, setting)
}

// SyncProvenRewardEvent is a paid mutator transaction binding the contract method 0xc6ea2882.
//
// Solidity: function syncProvenRewardEvent((uint256,address,uint256,uint256)[] rewardEvents, bool setting) returns()
func (_LPWANClient *LPWANClientSession) SyncProvenRewardEvent(rewardEvents []LPWANRewardEvent, setting bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.SyncProvenRewardEvent(&_LPWANClient.TransactOpts, rewardEvents, setting)
}

// SyncProvenRewardEvent is a paid mutator transaction binding the contract method 0xc6ea2882.
//
// Solidity: function syncProvenRewardEvent((uint256,address,uint256,uint256)[] rewardEvents, bool setting) returns()
func (_LPWANClient *LPWANClientTransactorSession) SyncProvenRewardEvent(rewardEvents []LPWANRewardEvent, setting bool) (*types.Transaction, error) {
	return _LPWANClient.Contract.SyncProvenRewardEvent(&_LPWANClient.TransactOpts, rewardEvents, setting)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPWANClient *LPWANClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LPWANClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPWANClient *LPWANClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LPWANClient.Contract.TransferOwnership(&_LPWANClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPWANClient *LPWANClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LPWANClient.Contract.TransferOwnership(&_LPWANClient.TransactOpts, newOwner)
}

// LPWANClientBurnExcessTokenIterator is returned from FilterBurnExcessToken and is used to iterate over the raw logs and unpacked data for BurnExcessToken events raised by the LPWANClient contract.
type LPWANClientBurnExcessTokenIterator struct {
	Event *LPWANClientBurnExcessToken // Event containing the contract specifics and raw log

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
func (it *LPWANClientBurnExcessTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPWANClientBurnExcessToken)
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
		it.Event = new(LPWANClientBurnExcessToken)
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
func (it *LPWANClientBurnExcessTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPWANClientBurnExcessTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPWANClientBurnExcessToken represents a BurnExcessToken event raised by the LPWANClient contract.
type LPWANClientBurnExcessToken struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnExcessToken is a free log retrieval operation binding the contract event 0x50b7fb3383de2508212f88c1dd0d32ddf0ca63749f5ff5d6c06538f9f872ae4e.
//
// Solidity: event BurnExcessToken(uint256 indexed id, uint256 indexed amount)
func (_LPWANClient *LPWANClientFilterer) FilterBurnExcessToken(opts *bind.FilterOpts, id []*big.Int, amount []*big.Int) (*LPWANClientBurnExcessTokenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LPWANClient.contract.FilterLogs(opts, "BurnExcessToken", idRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LPWANClientBurnExcessTokenIterator{contract: _LPWANClient.contract, event: "BurnExcessToken", logs: logs, sub: sub}, nil
}

// WatchBurnExcessToken is a free log subscription operation binding the contract event 0x50b7fb3383de2508212f88c1dd0d32ddf0ca63749f5ff5d6c06538f9f872ae4e.
//
// Solidity: event BurnExcessToken(uint256 indexed id, uint256 indexed amount)
func (_LPWANClient *LPWANClientFilterer) WatchBurnExcessToken(opts *bind.WatchOpts, sink chan<- *LPWANClientBurnExcessToken, id []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LPWANClient.contract.WatchLogs(opts, "BurnExcessToken", idRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPWANClientBurnExcessToken)
				if err := _LPWANClient.contract.UnpackLog(event, "BurnExcessToken", log); err != nil {
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

// ParseBurnExcessToken is a log parse operation binding the contract event 0x50b7fb3383de2508212f88c1dd0d32ddf0ca63749f5ff5d6c06538f9f872ae4e.
//
// Solidity: event BurnExcessToken(uint256 indexed id, uint256 indexed amount)
func (_LPWANClient *LPWANClientFilterer) ParseBurnExcessToken(log types.Log) (*LPWANClientBurnExcessToken, error) {
	event := new(LPWANClientBurnExcessToken)
	if err := _LPWANClient.contract.UnpackLog(event, "BurnExcessToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPWANClientClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the LPWANClient contract.
type LPWANClientClaimRewardIterator struct {
	Event *LPWANClientClaimReward // Event containing the contract specifics and raw log

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
func (it *LPWANClientClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPWANClientClaimReward)
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
		it.Event = new(LPWANClientClaimReward)
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
func (it *LPWANClientClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPWANClientClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPWANClientClaimReward represents a ClaimReward event raised by the LPWANClient contract.
type LPWANClientClaimReward struct {
	Account    common.Address
	Burn       bool
	RewardType *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0x1428870caf6bc74a66b702732942aaeff425bc9bb621cabdadbc35fb4a5804d0.
//
// Solidity: event ClaimReward(address indexed account, bool indexed burn, uint256 indexed rewardType, uint256 amount)
func (_LPWANClient *LPWANClientFilterer) FilterClaimReward(opts *bind.FilterOpts, account []common.Address, burn []bool, rewardType []*big.Int) (*LPWANClientClaimRewardIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var burnRule []interface{}
	for _, burnItem := range burn {
		burnRule = append(burnRule, burnItem)
	}
	var rewardTypeRule []interface{}
	for _, rewardTypeItem := range rewardType {
		rewardTypeRule = append(rewardTypeRule, rewardTypeItem)
	}

	logs, sub, err := _LPWANClient.contract.FilterLogs(opts, "ClaimReward", accountRule, burnRule, rewardTypeRule)
	if err != nil {
		return nil, err
	}
	return &LPWANClientClaimRewardIterator{contract: _LPWANClient.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x1428870caf6bc74a66b702732942aaeff425bc9bb621cabdadbc35fb4a5804d0.
//
// Solidity: event ClaimReward(address indexed account, bool indexed burn, uint256 indexed rewardType, uint256 amount)
func (_LPWANClient *LPWANClientFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *LPWANClientClaimReward, account []common.Address, burn []bool, rewardType []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var burnRule []interface{}
	for _, burnItem := range burn {
		burnRule = append(burnRule, burnItem)
	}
	var rewardTypeRule []interface{}
	for _, rewardTypeItem := range rewardType {
		rewardTypeRule = append(rewardTypeRule, rewardTypeItem)
	}

	logs, sub, err := _LPWANClient.contract.WatchLogs(opts, "ClaimReward", accountRule, burnRule, rewardTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPWANClientClaimReward)
				if err := _LPWANClient.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0x1428870caf6bc74a66b702732942aaeff425bc9bb621cabdadbc35fb4a5804d0.
//
// Solidity: event ClaimReward(address indexed account, bool indexed burn, uint256 indexed rewardType, uint256 amount)
func (_LPWANClient *LPWANClientFilterer) ParseClaimReward(log types.Log) (*LPWANClientClaimReward, error) {
	event := new(LPWANClientClaimReward)
	if err := _LPWANClient.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPWANClientControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the LPWANClient contract.
type LPWANClientControllerChangedIterator struct {
	Event *LPWANClientControllerChanged // Event containing the contract specifics and raw log

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
func (it *LPWANClientControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPWANClientControllerChanged)
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
		it.Event = new(LPWANClientControllerChanged)
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
func (it *LPWANClientControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPWANClientControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPWANClientControllerChanged represents a ControllerChanged event raised by the LPWANClient contract.
type LPWANClientControllerChanged struct {
	Controller common.Address
	Enabled    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_LPWANClient *LPWANClientFilterer) FilterControllerChanged(opts *bind.FilterOpts, controller []common.Address) (*LPWANClientControllerChangedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _LPWANClient.contract.FilterLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return &LPWANClientControllerChangedIterator{contract: _LPWANClient.contract, event: "ControllerChanged", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_LPWANClient *LPWANClientFilterer) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *LPWANClientControllerChanged, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _LPWANClient.contract.WatchLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPWANClientControllerChanged)
				if err := _LPWANClient.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
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

// ParseControllerChanged is a log parse operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool enabled)
func (_LPWANClient *LPWANClientFilterer) ParseControllerChanged(log types.Log) (*LPWANClientControllerChanged, error) {
	event := new(LPWANClientControllerChanged)
	if err := _LPWANClient.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPWANClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LPWANClient contract.
type LPWANClientInitializedIterator struct {
	Event *LPWANClientInitialized // Event containing the contract specifics and raw log

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
func (it *LPWANClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPWANClientInitialized)
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
		it.Event = new(LPWANClientInitialized)
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
func (it *LPWANClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPWANClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPWANClientInitialized represents a Initialized event raised by the LPWANClient contract.
type LPWANClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LPWANClient *LPWANClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*LPWANClientInitializedIterator, error) {

	logs, sub, err := _LPWANClient.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LPWANClientInitializedIterator{contract: _LPWANClient.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LPWANClient *LPWANClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LPWANClientInitialized) (event.Subscription, error) {

	logs, sub, err := _LPWANClient.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPWANClientInitialized)
				if err := _LPWANClient.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LPWANClient *LPWANClientFilterer) ParseInitialized(log types.Log) (*LPWANClientInitialized, error) {
	event := new(LPWANClientInitialized)
	if err := _LPWANClient.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPWANClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LPWANClient contract.
type LPWANClientOwnershipTransferredIterator struct {
	Event *LPWANClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LPWANClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPWANClientOwnershipTransferred)
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
		it.Event = new(LPWANClientOwnershipTransferred)
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
func (it *LPWANClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPWANClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPWANClientOwnershipTransferred represents a OwnershipTransferred event raised by the LPWANClient contract.
type LPWANClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPWANClient *LPWANClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LPWANClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPWANClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LPWANClientOwnershipTransferredIterator{contract: _LPWANClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPWANClient *LPWANClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LPWANClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPWANClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPWANClientOwnershipTransferred)
				if err := _LPWANClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LPWANClient *LPWANClientFilterer) ParseOwnershipTransferred(log types.Log) (*LPWANClientOwnershipTransferred, error) {
	event := new(LPWANClientOwnershipTransferred)
	if err := _LPWANClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
