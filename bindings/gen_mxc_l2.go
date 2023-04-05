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

// MXCL2ClientMetaData contains all meta data concerning the MXCL2Client contract.
var MXCL2ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ERR_INVALID_HINT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERR_INVALID_TX_IDX\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERR_PARAMS_NOT_DEFAULTS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERR_VERIFICAITON_FAILURE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_CHAIN_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_GAS_PRICE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_INVALID_SENDER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2_PUBLIC_INPUT_HASH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"}],\"name\":\"BlockInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"}],\"name\":\"HeaderSynced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l1Height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"}],\"name\":\"anchor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockHashHistory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitConfirmations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"anchorTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slotSmoothingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBurnBips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposerDepositPctg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBaseMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofTimeMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"rewardMultiplierPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feeGracePeriodPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feeMaxPeriodPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimeCap\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeCap\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"emptyBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"proposeBlockDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"bootstrapDiscountHalvingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enableTokenomics\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enablePublicInputsCheck\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableAnchorValidation\",\"type\":\"bool\"}],\"internalType\":\"structMXCData.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"},{\"internalType\":\"enumLibInvalidTxList.Hint\",\"name\":\"hint\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"txIdx\",\"type\":\"uint256\"}],\"name\":\"invalidateBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSyncedL1Height\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MXCL2ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use MXCL2ClientMetaData.ABI instead.
var MXCL2ClientABI = MXCL2ClientMetaData.ABI

// MXCL2Client is an auto generated Go binding around an Ethereum contract.
type MXCL2Client struct {
	MXCL2ClientCaller     // Read-only binding to the contract
	MXCL2ClientTransactor // Write-only binding to the contract
	MXCL2ClientFilterer   // Log filterer for contract events
}

// MXCL2ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type MXCL2ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MXCL2ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MXCL2ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MXCL2ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MXCL2ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MXCL2ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MXCL2ClientSession struct {
	Contract     *MXCL2Client      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MXCL2ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MXCL2ClientCallerSession struct {
	Contract *MXCL2ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MXCL2ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MXCL2ClientTransactorSession struct {
	Contract     *MXCL2ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MXCL2ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type MXCL2ClientRaw struct {
	Contract *MXCL2Client // Generic contract binding to access the raw methods on
}

// MXCL2ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MXCL2ClientCallerRaw struct {
	Contract *MXCL2ClientCaller // Generic read-only contract binding to access the raw methods on
}

// MXCL2ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MXCL2ClientTransactorRaw struct {
	Contract *MXCL2ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMXCL2Client creates a new instance of MXCL2Client, bound to a specific deployed contract.
func NewMXCL2Client(address common.Address, backend bind.ContractBackend) (*MXCL2Client, error) {
	contract, err := bindMXCL2Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MXCL2Client{MXCL2ClientCaller: MXCL2ClientCaller{contract: contract}, MXCL2ClientTransactor: MXCL2ClientTransactor{contract: contract}, MXCL2ClientFilterer: MXCL2ClientFilterer{contract: contract}}, nil
}

// NewMXCL2ClientCaller creates a new read-only instance of MXCL2Client, bound to a specific deployed contract.
func NewMXCL2ClientCaller(address common.Address, caller bind.ContractCaller) (*MXCL2ClientCaller, error) {
	contract, err := bindMXCL2Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MXCL2ClientCaller{contract: contract}, nil
}

// NewMXCL2ClientTransactor creates a new write-only instance of MXCL2Client, bound to a specific deployed contract.
func NewMXCL2ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*MXCL2ClientTransactor, error) {
	contract, err := bindMXCL2Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MXCL2ClientTransactor{contract: contract}, nil
}

// NewMXCL2ClientFilterer creates a new log filterer instance of MXCL2Client, bound to a specific deployed contract.
func NewMXCL2ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*MXCL2ClientFilterer, error) {
	contract, err := bindMXCL2Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MXCL2ClientFilterer{contract: contract}, nil
}

// bindMXCL2Client binds a generic wrapper to an already deployed contract.
func bindMXCL2Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MXCL2ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MXCL2Client *MXCL2ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MXCL2Client.Contract.MXCL2ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MXCL2Client *MXCL2ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL2Client.Contract.MXCL2ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MXCL2Client *MXCL2ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MXCL2Client.Contract.MXCL2ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MXCL2Client *MXCL2ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MXCL2Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MXCL2Client *MXCL2ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL2Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MXCL2Client *MXCL2ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MXCL2Client.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MXCL2Client *MXCL2ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MXCL2Client *MXCL2ClientSession) AddressManager() (common.Address, error) {
	return _MXCL2Client.Contract.AddressManager(&_MXCL2Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MXCL2Client *MXCL2ClientCallerSession) AddressManager() (common.Address, error) {
	return _MXCL2Client.Contract.AddressManager(&_MXCL2Client.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_MXCL2Client *MXCL2ClientCaller) GetBlockHash(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "getBlockHash", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_MXCL2Client *MXCL2ClientSession) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _MXCL2Client.Contract.GetBlockHash(&_MXCL2Client.CallOpts, number)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 number) view returns(bytes32)
func (_MXCL2Client *MXCL2ClientCallerSession) GetBlockHash(number *big.Int) ([32]byte, error) {
	return _MXCL2Client.Contract.GetBlockHash(&_MXCL2Client.CallOpts, number)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint64,bool,bool,bool) config)
func (_MXCL2Client *MXCL2ClientCaller) GetConfig(opts *bind.CallOpts) (MXCDataConfig, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(MXCDataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MXCDataConfig)).(*MXCDataConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint64,bool,bool,bool) config)
func (_MXCL2Client *MXCL2ClientSession) GetConfig() (MXCDataConfig, error) {
	return _MXCL2Client.Contract.GetConfig(&_MXCL2Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint64,bool,bool,bool) config)
func (_MXCL2Client *MXCL2ClientCallerSession) GetConfig() (MXCDataConfig, error) {
	return _MXCL2Client.Contract.GetConfig(&_MXCL2Client.CallOpts)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_MXCL2Client *MXCL2ClientCaller) GetLatestSyncedHeader(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "getLatestSyncedHeader")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_MXCL2Client *MXCL2ClientSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _MXCL2Client.Contract.GetLatestSyncedHeader(&_MXCL2Client.CallOpts)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_MXCL2Client *MXCL2ClientCallerSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _MXCL2Client.Contract.GetLatestSyncedHeader(&_MXCL2Client.CallOpts)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_MXCL2Client *MXCL2ClientCaller) GetSyncedHeader(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "getSyncedHeader", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_MXCL2Client *MXCL2ClientSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _MXCL2Client.Contract.GetSyncedHeader(&_MXCL2Client.CallOpts, number)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_MXCL2Client *MXCL2ClientCallerSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _MXCL2Client.Contract.GetSyncedHeader(&_MXCL2Client.CallOpts, number)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint256)
func (_MXCL2Client *MXCL2ClientCaller) LatestSyncedL1Height(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "latestSyncedL1Height")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint256)
func (_MXCL2Client *MXCL2ClientSession) LatestSyncedL1Height() (*big.Int, error) {
	return _MXCL2Client.Contract.LatestSyncedL1Height(&_MXCL2Client.CallOpts)
}

// LatestSyncedL1Height is a free data retrieval call binding the contract method 0xc7b96908.
//
// Solidity: function latestSyncedL1Height() view returns(uint256)
func (_MXCL2Client *MXCL2ClientCallerSession) LatestSyncedL1Height() (*big.Int, error) {
	return _MXCL2Client.Contract.LatestSyncedL1Height(&_MXCL2Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_MXCL2Client *MXCL2ClientCaller) Resolve(opts *bind.CallOpts, name string, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "resolve", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_MXCL2Client *MXCL2ClientSession) Resolve(name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL2Client.Contract.Resolve(&_MXCL2Client.CallOpts, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_MXCL2Client *MXCL2ClientCallerSession) Resolve(name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL2Client.Contract.Resolve(&_MXCL2Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_MXCL2Client *MXCL2ClientCaller) Resolve0(opts *bind.CallOpts, chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MXCL2Client.contract.Call(opts, &out, "resolve0", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_MXCL2Client *MXCL2ClientSession) Resolve0(chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL2Client.Contract.Resolve0(&_MXCL2Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_MXCL2Client *MXCL2ClientCallerSession) Resolve0(chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL2Client.Contract.Resolve0(&_MXCL2Client.CallOpts, chainId, name, allowZeroAddress)
}

// Anchor is a paid mutator transaction binding the contract method 0xa0ca2d08.
//
// Solidity: function anchor(uint256 l1Height, bytes32 l1Hash) returns()
func (_MXCL2Client *MXCL2ClientTransactor) Anchor(opts *bind.TransactOpts, l1Height *big.Int, l1Hash [32]byte) (*types.Transaction, error) {
	return _MXCL2Client.contract.Transact(opts, "anchor", l1Height, l1Hash)
}

// Anchor is a paid mutator transaction binding the contract method 0xa0ca2d08.
//
// Solidity: function anchor(uint256 l1Height, bytes32 l1Hash) returns()
func (_MXCL2Client *MXCL2ClientSession) Anchor(l1Height *big.Int, l1Hash [32]byte) (*types.Transaction, error) {
	return _MXCL2Client.Contract.Anchor(&_MXCL2Client.TransactOpts, l1Height, l1Hash)
}

// Anchor is a paid mutator transaction binding the contract method 0xa0ca2d08.
//
// Solidity: function anchor(uint256 l1Height, bytes32 l1Hash) returns()
func (_MXCL2Client *MXCL2ClientTransactorSession) Anchor(l1Height *big.Int, l1Hash [32]byte) (*types.Transaction, error) {
	return _MXCL2Client.Contract.Anchor(&_MXCL2Client.TransactOpts, l1Height, l1Hash)
}

// InvalidateBlock is a paid mutator transaction binding the contract method 0x975e09a0.
//
// Solidity: function invalidateBlock(bytes txList, uint8 hint, uint256 txIdx) returns()
func (_MXCL2Client *MXCL2ClientTransactor) InvalidateBlock(opts *bind.TransactOpts, txList []byte, hint uint8, txIdx *big.Int) (*types.Transaction, error) {
	return _MXCL2Client.contract.Transact(opts, "invalidateBlock", txList, hint, txIdx)
}

// InvalidateBlock is a paid mutator transaction binding the contract method 0x975e09a0.
//
// Solidity: function invalidateBlock(bytes txList, uint8 hint, uint256 txIdx) returns()
func (_MXCL2Client *MXCL2ClientSession) InvalidateBlock(txList []byte, hint uint8, txIdx *big.Int) (*types.Transaction, error) {
	return _MXCL2Client.Contract.InvalidateBlock(&_MXCL2Client.TransactOpts, txList, hint, txIdx)
}

// InvalidateBlock is a paid mutator transaction binding the contract method 0x975e09a0.
//
// Solidity: function invalidateBlock(bytes txList, uint8 hint, uint256 txIdx) returns()
func (_MXCL2Client *MXCL2ClientTransactorSession) InvalidateBlock(txList []byte, hint uint8, txIdx *big.Int) (*types.Transaction, error) {
	return _MXCL2Client.Contract.InvalidateBlock(&_MXCL2Client.TransactOpts, txList, hint, txIdx)
}

// MXCL2ClientBlockInvalidatedIterator is returned from FilterBlockInvalidated and is used to iterate over the raw logs and unpacked data for BlockInvalidated events raised by the MXCL2Client contract.
type MXCL2ClientBlockInvalidatedIterator struct {
	Event *MXCL2ClientBlockInvalidated // Event containing the contract specifics and raw log

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
func (it *MXCL2ClientBlockInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL2ClientBlockInvalidated)
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
		it.Event = new(MXCL2ClientBlockInvalidated)
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
func (it *MXCL2ClientBlockInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL2ClientBlockInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL2ClientBlockInvalidated represents a BlockInvalidated event raised by the MXCL2Client contract.
type MXCL2ClientBlockInvalidated struct {
	TxListHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockInvalidated is a free log retrieval operation binding the contract event 0x64b299ff9f8ba674288abb53380419048a4271dda03b837ecba6b40e6ddea4a2.
//
// Solidity: event BlockInvalidated(bytes32 indexed txListHash)
func (_MXCL2Client *MXCL2ClientFilterer) FilterBlockInvalidated(opts *bind.FilterOpts, txListHash [][32]byte) (*MXCL2ClientBlockInvalidatedIterator, error) {

	var txListHashRule []interface{}
	for _, txListHashItem := range txListHash {
		txListHashRule = append(txListHashRule, txListHashItem)
	}

	logs, sub, err := _MXCL2Client.contract.FilterLogs(opts, "BlockInvalidated", txListHashRule)
	if err != nil {
		return nil, err
	}
	return &MXCL2ClientBlockInvalidatedIterator{contract: _MXCL2Client.contract, event: "BlockInvalidated", logs: logs, sub: sub}, nil
}

// WatchBlockInvalidated is a free log subscription operation binding the contract event 0x64b299ff9f8ba674288abb53380419048a4271dda03b837ecba6b40e6ddea4a2.
//
// Solidity: event BlockInvalidated(bytes32 indexed txListHash)
func (_MXCL2Client *MXCL2ClientFilterer) WatchBlockInvalidated(opts *bind.WatchOpts, sink chan<- *MXCL2ClientBlockInvalidated, txListHash [][32]byte) (event.Subscription, error) {

	var txListHashRule []interface{}
	for _, txListHashItem := range txListHash {
		txListHashRule = append(txListHashRule, txListHashItem)
	}

	logs, sub, err := _MXCL2Client.contract.WatchLogs(opts, "BlockInvalidated", txListHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL2ClientBlockInvalidated)
				if err := _MXCL2Client.contract.UnpackLog(event, "BlockInvalidated", log); err != nil {
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

// ParseBlockInvalidated is a log parse operation binding the contract event 0x64b299ff9f8ba674288abb53380419048a4271dda03b837ecba6b40e6ddea4a2.
//
// Solidity: event BlockInvalidated(bytes32 indexed txListHash)
func (_MXCL2Client *MXCL2ClientFilterer) ParseBlockInvalidated(log types.Log) (*MXCL2ClientBlockInvalidated, error) {
	event := new(MXCL2ClientBlockInvalidated)
	if err := _MXCL2Client.contract.UnpackLog(event, "BlockInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL2ClientHeaderSyncedIterator is returned from FilterHeaderSynced and is used to iterate over the raw logs and unpacked data for HeaderSynced events raised by the MXCL2Client contract.
type MXCL2ClientHeaderSyncedIterator struct {
	Event *MXCL2ClientHeaderSynced // Event containing the contract specifics and raw log

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
func (it *MXCL2ClientHeaderSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL2ClientHeaderSynced)
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
		it.Event = new(MXCL2ClientHeaderSynced)
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
func (it *MXCL2ClientHeaderSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL2ClientHeaderSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL2ClientHeaderSynced represents a HeaderSynced event raised by the MXCL2Client contract.
type MXCL2ClientHeaderSynced struct {
	SrcHeight *big.Int
	SrcHash   [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHeaderSynced is a free log retrieval operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_MXCL2Client *MXCL2ClientFilterer) FilterHeaderSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*MXCL2ClientHeaderSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MXCL2Client.contract.FilterLogs(opts, "HeaderSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &MXCL2ClientHeaderSyncedIterator{contract: _MXCL2Client.contract, event: "HeaderSynced", logs: logs, sub: sub}, nil
}

// WatchHeaderSynced is a free log subscription operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_MXCL2Client *MXCL2ClientFilterer) WatchHeaderSynced(opts *bind.WatchOpts, sink chan<- *MXCL2ClientHeaderSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MXCL2Client.contract.WatchLogs(opts, "HeaderSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL2ClientHeaderSynced)
				if err := _MXCL2Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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

// ParseHeaderSynced is a log parse operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_MXCL2Client *MXCL2ClientFilterer) ParseHeaderSynced(log types.Log) (*MXCL2ClientHeaderSynced, error) {
	event := new(MXCL2ClientHeaderSynced)
	if err := _MXCL2Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
