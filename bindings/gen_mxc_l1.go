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

// LibUtilsStateVariables is an auto generated low-level Go binding around an user-defined struct.
type LibUtilsStateVariables struct {
	FeeBase              *big.Int
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
}

// MXCDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type MXCDataBlockMetadata struct {
	Id           *big.Int
	L1Height     *big.Int
	L1Hash       [32]byte
	Beneficiary  common.Address
	TxListHash   [32]byte
	MixHash      [32]byte
	ExtraData    []byte
	GasLimit     uint64
	Timestamp    uint64
	CommitHeight uint64
	CommitSlot   uint64
}

// MXCDataConfig is an auto generated low-level Go binding around an user-defined struct.
type MXCDataConfig struct {
	ChainId                        *big.Int
	MaxNumBlocks                   *big.Int
	BlockHashHistory               *big.Int
	MaxVerificationsPerTx          *big.Int
	CommitConfirmations            *big.Int
	BlockMaxGasLimit               *big.Int
	MaxTransactionsPerBlock        *big.Int
	MaxBytesPerTxList              *big.Int
	MinTxGasLimit                  *big.Int
	AnchorTxGasLimit               *big.Int
	SlotSmoothingFactor            *big.Int
	RewardBurnBips                 *big.Int
	ProposerDepositPctg            *big.Int
	FeeBaseMAF                     *big.Int
	BlockTimeMAF                   *big.Int
	ProofTimeMAF                   *big.Int
	RewardMultiplierPctg           uint64
	FeeGracePeriodPctg             uint64
	FeeMaxPeriodPctg               uint64
	BlockTimeCap                   uint64
	ProofTimeCap                   uint64
	BlockInterval                  uint64
	EmptyBlockInterval             uint64
	ProposeBlockDeposit            *big.Int
	BootstrapDiscountHalvingPeriod uint64
	EnableTokenomics               bool
	EnablePublicInputsCheck        bool
	EnableAnchorValidation         bool
}

// MXCDataForkChoice is an auto generated low-level Go binding around an user-defined struct.
type MXCDataForkChoice struct {
	BlockHash [32]byte
	Prover    common.Address
	ProvenAt  uint64
}

// MXCDataProposedBlock is an auto generated low-level Go binding around an user-defined struct.
type MXCDataProposedBlock struct {
	MetaHash   [32]byte
	Deposit    *big.Int
	Proposer   common.Address
	ProposedAt uint64
}

// MXCL1ClientMetaData contains all meta data concerning the MXCL1Client contract.
var MXCL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_0_FEE_BASE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_0_FEE_BASE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_CALLDATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_DEST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_GAS_LIMIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_ADDR\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_DATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_LOGS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_STATUS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_RECEIPT_TOPICS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_SIG_R\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_SIG_S\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_TX_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ANCHOR_TYPE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_NUMBER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_NUMBER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_CANNOT_BE_FIRST_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_COMMITTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_CONFLICT_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_CONTRACT_NOT_ALLOWED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_DUP_PROVERS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_EXTRA_DATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_GAS_LIMIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INPUT_SIZE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_METADATA_FIELD\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_META_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_COMMITTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_ORACLE_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_PROOF_LENGTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SOLO_PROPOSER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_STAKE_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ZKP\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"commitSlot\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"BlockCommitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commitHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commitSlot\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structMXCData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"}],\"name\":\"HeaderSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"commitSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"commitBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockHashHistory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitConfirmations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"anchorTxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slotSmoothingFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBurnBips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposerDepositPctg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBaseMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofTimeMAF\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"rewardMultiplierPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feeGracePeriodPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"feeMaxPeriodPctg\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimeCap\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeCap\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"emptyBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"proposeBlockDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"bootstrapDiscountHalvingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enableTokenomics\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enablePublicInputsCheck\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableAnchorValidation\",\"type\":\"bool\"}],\"internalType\":\"structMXCData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"}],\"internalType\":\"structMXCData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getProposedBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"proposedAt\",\"type\":\"uint64\"}],\"internalType\":\"structMXCData.ProposedBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRewardBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgBlockTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgProofTime\",\"type\":\"uint64\"}],\"internalType\":\"structLibUtils.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getSyncedHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_feeBase\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commitSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"isCommitValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proposeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"proveBlockInvalid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"signWithGoldenTouch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reservedA1\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reservedA2\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nextBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgBlockTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__avgGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"latestVerifiedId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avgProofTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"__reservedC1\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MXCL1ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use MXCL1ClientMetaData.ABI instead.
var MXCL1ClientABI = MXCL1ClientMetaData.ABI

// MXCL1Client is an auto generated Go binding around an Ethereum contract.
type MXCL1Client struct {
	MXCL1ClientCaller     // Read-only binding to the contract
	MXCL1ClientTransactor // Write-only binding to the contract
	MXCL1ClientFilterer   // Log filterer for contract events
}

// MXCL1ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type MXCL1ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MXCL1ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MXCL1ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MXCL1ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MXCL1ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MXCL1ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MXCL1ClientSession struct {
	Contract     *MXCL1Client      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MXCL1ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MXCL1ClientCallerSession struct {
	Contract *MXCL1ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MXCL1ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MXCL1ClientTransactorSession struct {
	Contract     *MXCL1ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MXCL1ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type MXCL1ClientRaw struct {
	Contract *MXCL1Client // Generic contract binding to access the raw methods on
}

// MXCL1ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MXCL1ClientCallerRaw struct {
	Contract *MXCL1ClientCaller // Generic read-only contract binding to access the raw methods on
}

// MXCL1ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MXCL1ClientTransactorRaw struct {
	Contract *MXCL1ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMXCL1Client creates a new instance of MXCL1Client, bound to a specific deployed contract.
func NewMXCL1Client(address common.Address, backend bind.ContractBackend) (*MXCL1Client, error) {
	contract, err := bindMXCL1Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MXCL1Client{MXCL1ClientCaller: MXCL1ClientCaller{contract: contract}, MXCL1ClientTransactor: MXCL1ClientTransactor{contract: contract}, MXCL1ClientFilterer: MXCL1ClientFilterer{contract: contract}}, nil
}

// NewMXCL1ClientCaller creates a new read-only instance of MXCL1Client, bound to a specific deployed contract.
func NewMXCL1ClientCaller(address common.Address, caller bind.ContractCaller) (*MXCL1ClientCaller, error) {
	contract, err := bindMXCL1Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientCaller{contract: contract}, nil
}

// NewMXCL1ClientTransactor creates a new write-only instance of MXCL1Client, bound to a specific deployed contract.
func NewMXCL1ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*MXCL1ClientTransactor, error) {
	contract, err := bindMXCL1Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientTransactor{contract: contract}, nil
}

// NewMXCL1ClientFilterer creates a new log filterer instance of MXCL1Client, bound to a specific deployed contract.
func NewMXCL1ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*MXCL1ClientFilterer, error) {
	contract, err := bindMXCL1Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientFilterer{contract: contract}, nil
}

// bindMXCL1Client binds a generic wrapper to an already deployed contract.
func bindMXCL1Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MXCL1ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MXCL1Client *MXCL1ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MXCL1Client.Contract.MXCL1ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MXCL1Client *MXCL1ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL1Client.Contract.MXCL1ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MXCL1Client *MXCL1ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MXCL1Client.Contract.MXCL1ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MXCL1Client *MXCL1ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MXCL1Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MXCL1Client *MXCL1ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL1Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MXCL1Client *MXCL1ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MXCL1Client.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MXCL1Client *MXCL1ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MXCL1Client *MXCL1ClientSession) AddressManager() (common.Address, error) {
	return _MXCL1Client.Contract.AddressManager(&_MXCL1Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MXCL1Client *MXCL1ClientCallerSession) AddressManager() (common.Address, error) {
	return _MXCL1Client.Contract.AddressManager(&_MXCL1Client.CallOpts)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint256)
func (_MXCL1Client *MXCL1ClientCaller) GetBlockFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getBlockFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint256)
func (_MXCL1Client *MXCL1ClientSession) GetBlockFee() (*big.Int, error) {
	return _MXCL1Client.Contract.GetBlockFee(&_MXCL1Client.CallOpts)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint256)
func (_MXCL1Client *MXCL1ClientCallerSession) GetBlockFee() (*big.Int, error) {
	return _MXCL1Client.Contract.GetBlockFee(&_MXCL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint64,bool,bool,bool))
func (_MXCL1Client *MXCL1ClientCaller) GetConfig(opts *bind.CallOpts) (MXCDataConfig, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(MXCDataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MXCDataConfig)).(*MXCDataConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint64,bool,bool,bool))
func (_MXCL1Client *MXCL1ClientSession) GetConfig() (MXCDataConfig, error) {
	return _MXCL1Client.Contract.GetConfig(&_MXCL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint64,bool,bool,bool))
func (_MXCL1Client *MXCL1ClientCallerSession) GetConfig() (MXCDataConfig, error) {
	return _MXCL1Client.Contract.GetConfig(&_MXCL1Client.CallOpts)
}

// GetForkChoice is a free data retrieval call binding the contract method 0xe00ea1e1.
//
// Solidity: function getForkChoice(uint256 id, bytes32 parentHash) view returns((bytes32,address,uint64))
func (_MXCL1Client *MXCL1ClientCaller) GetForkChoice(opts *bind.CallOpts, id *big.Int, parentHash [32]byte) (MXCDataForkChoice, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getForkChoice", id, parentHash)

	if err != nil {
		return *new(MXCDataForkChoice), err
	}

	out0 := *abi.ConvertType(out[0], new(MXCDataForkChoice)).(*MXCDataForkChoice)

	return out0, err

}

// GetForkChoice is a free data retrieval call binding the contract method 0xe00ea1e1.
//
// Solidity: function getForkChoice(uint256 id, bytes32 parentHash) view returns((bytes32,address,uint64))
func (_MXCL1Client *MXCL1ClientSession) GetForkChoice(id *big.Int, parentHash [32]byte) (MXCDataForkChoice, error) {
	return _MXCL1Client.Contract.GetForkChoice(&_MXCL1Client.CallOpts, id, parentHash)
}

// GetForkChoice is a free data retrieval call binding the contract method 0xe00ea1e1.
//
// Solidity: function getForkChoice(uint256 id, bytes32 parentHash) view returns((bytes32,address,uint64))
func (_MXCL1Client *MXCL1ClientCallerSession) GetForkChoice(id *big.Int, parentHash [32]byte) (MXCDataForkChoice, error) {
	return _MXCL1Client.Contract.GetForkChoice(&_MXCL1Client.CallOpts, id, parentHash)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_MXCL1Client *MXCL1ClientCaller) GetLatestSyncedHeader(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getLatestSyncedHeader")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_MXCL1Client *MXCL1ClientSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _MXCL1Client.Contract.GetLatestSyncedHeader(&_MXCL1Client.CallOpts)
}

// GetLatestSyncedHeader is a free data retrieval call binding the contract method 0x5155ce9f.
//
// Solidity: function getLatestSyncedHeader() view returns(bytes32)
func (_MXCL1Client *MXCL1ClientCallerSession) GetLatestSyncedHeader() ([32]byte, error) {
	return _MXCL1Client.Contract.GetLatestSyncedHeader(&_MXCL1Client.CallOpts)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint256 reward)
func (_MXCL1Client *MXCL1ClientCaller) GetProofReward(opts *bind.CallOpts, provenAt uint64, proposedAt uint64) (*big.Int, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getProofReward", provenAt, proposedAt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint256 reward)
func (_MXCL1Client *MXCL1ClientSession) GetProofReward(provenAt uint64, proposedAt uint64) (*big.Int, error) {
	return _MXCL1Client.Contract.GetProofReward(&_MXCL1Client.CallOpts, provenAt, proposedAt)
}

// GetProofReward is a free data retrieval call binding the contract method 0x4ee56f9e.
//
// Solidity: function getProofReward(uint64 provenAt, uint64 proposedAt) view returns(uint256 reward)
func (_MXCL1Client *MXCL1ClientCallerSession) GetProofReward(provenAt uint64, proposedAt uint64) (*big.Int, error) {
	return _MXCL1Client.Contract.GetProofReward(&_MXCL1Client.CallOpts, provenAt, proposedAt)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32,uint256,address,uint64))
func (_MXCL1Client *MXCL1ClientCaller) GetProposedBlock(opts *bind.CallOpts, id *big.Int) (MXCDataProposedBlock, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getProposedBlock", id)

	if err != nil {
		return *new(MXCDataProposedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(MXCDataProposedBlock)).(*MXCDataProposedBlock)

	return out0, err

}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32,uint256,address,uint64))
func (_MXCL1Client *MXCL1ClientSession) GetProposedBlock(id *big.Int) (MXCDataProposedBlock, error) {
	return _MXCL1Client.Contract.GetProposedBlock(&_MXCL1Client.CallOpts, id)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0x8972b10c.
//
// Solidity: function getProposedBlock(uint256 id) view returns((bytes32,uint256,address,uint64))
func (_MXCL1Client *MXCL1ClientCallerSession) GetProposedBlock(id *big.Int) (MXCDataProposedBlock, error) {
	return _MXCL1Client.Contract.GetProposedBlock(&_MXCL1Client.CallOpts, id)
}

// GetRewardBalance is a free data retrieval call binding the contract method 0xd5a849e9.
//
// Solidity: function getRewardBalance(address addr) view returns(uint256)
func (_MXCL1Client *MXCL1ClientCaller) GetRewardBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getRewardBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardBalance is a free data retrieval call binding the contract method 0xd5a849e9.
//
// Solidity: function getRewardBalance(address addr) view returns(uint256)
func (_MXCL1Client *MXCL1ClientSession) GetRewardBalance(addr common.Address) (*big.Int, error) {
	return _MXCL1Client.Contract.GetRewardBalance(&_MXCL1Client.CallOpts, addr)
}

// GetRewardBalance is a free data retrieval call binding the contract method 0xd5a849e9.
//
// Solidity: function getRewardBalance(address addr) view returns(uint256)
func (_MXCL1Client *MXCL1ClientCallerSession) GetRewardBalance(addr common.Address) (*big.Int, error) {
	return _MXCL1Client.Contract.GetRewardBalance(&_MXCL1Client.CallOpts, addr)
}

// GetStakeAmount is a free data retrieval call binding the contract method 0x722580b6.
//
// Solidity: function getStakeAmount() view returns(uint256)
func (_MXCL1Client *MXCL1ClientCaller) GetStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAmount is a free data retrieval call binding the contract method 0x722580b6.
//
// Solidity: function getStakeAmount() view returns(uint256)
func (_MXCL1Client *MXCL1ClientSession) GetStakeAmount() (*big.Int, error) {
	return _MXCL1Client.Contract.GetStakeAmount(&_MXCL1Client.CallOpts)
}

// GetStakeAmount is a free data retrieval call binding the contract method 0x722580b6.
//
// Solidity: function getStakeAmount() view returns(uint256)
func (_MXCL1Client *MXCL1ClientCallerSession) GetStakeAmount() (*big.Int, error) {
	return _MXCL1Client.Contract.GetStakeAmount(&_MXCL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_MXCL1Client *MXCL1ClientCaller) GetStateVariables(opts *bind.CallOpts) (LibUtilsStateVariables, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getStateVariables")

	if err != nil {
		return *new(LibUtilsStateVariables), err
	}

	out0 := *abi.ConvertType(out[0], new(LibUtilsStateVariables)).(*LibUtilsStateVariables)

	return out0, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_MXCL1Client *MXCL1ClientSession) GetStateVariables() (LibUtilsStateVariables, error) {
	return _MXCL1Client.Contract.GetStateVariables(&_MXCL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint256,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_MXCL1Client *MXCL1ClientCallerSession) GetStateVariables() (LibUtilsStateVariables, error) {
	return _MXCL1Client.Contract.GetStateVariables(&_MXCL1Client.CallOpts)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_MXCL1Client *MXCL1ClientCaller) GetSyncedHeader(opts *bind.CallOpts, number *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "getSyncedHeader", number)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_MXCL1Client *MXCL1ClientSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _MXCL1Client.Contract.GetSyncedHeader(&_MXCL1Client.CallOpts, number)
}

// GetSyncedHeader is a free data retrieval call binding the contract method 0x25bf86f2.
//
// Solidity: function getSyncedHeader(uint256 number) view returns(bytes32)
func (_MXCL1Client *MXCL1ClientCallerSession) GetSyncedHeader(number *big.Int) ([32]byte, error) {
	return _MXCL1Client.Contract.GetSyncedHeader(&_MXCL1Client.CallOpts, number)
}

// IsCommitValid is a free data retrieval call binding the contract method 0x340d9599.
//
// Solidity: function isCommitValid(uint256 commitSlot, uint256 commitHeight, bytes32 commitHash) view returns(bool)
func (_MXCL1Client *MXCL1ClientCaller) IsCommitValid(opts *bind.CallOpts, commitSlot *big.Int, commitHeight *big.Int, commitHash [32]byte) (bool, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "isCommitValid", commitSlot, commitHeight, commitHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommitValid is a free data retrieval call binding the contract method 0x340d9599.
//
// Solidity: function isCommitValid(uint256 commitSlot, uint256 commitHeight, bytes32 commitHash) view returns(bool)
func (_MXCL1Client *MXCL1ClientSession) IsCommitValid(commitSlot *big.Int, commitHeight *big.Int, commitHash [32]byte) (bool, error) {
	return _MXCL1Client.Contract.IsCommitValid(&_MXCL1Client.CallOpts, commitSlot, commitHeight, commitHash)
}

// IsCommitValid is a free data retrieval call binding the contract method 0x340d9599.
//
// Solidity: function isCommitValid(uint256 commitSlot, uint256 commitHeight, bytes32 commitHash) view returns(bool)
func (_MXCL1Client *MXCL1ClientCallerSession) IsCommitValid(commitSlot *big.Int, commitHeight *big.Int, commitHash [32]byte) (bool, error) {
	return _MXCL1Client.Contract.IsCommitValid(&_MXCL1Client.CallOpts, commitSlot, commitHeight, commitHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MXCL1Client *MXCL1ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MXCL1Client *MXCL1ClientSession) Owner() (common.Address, error) {
	return _MXCL1Client.Contract.Owner(&_MXCL1Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MXCL1Client *MXCL1ClientCallerSession) Owner() (common.Address, error) {
	return _MXCL1Client.Contract.Owner(&_MXCL1Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_MXCL1Client *MXCL1ClientCaller) Resolve(opts *bind.CallOpts, name string, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "resolve", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_MXCL1Client *MXCL1ClientSession) Resolve(name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL1Client.Contract.Resolve(&_MXCL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x0ca4dffd.
//
// Solidity: function resolve(string name, bool allowZeroAddress) view returns(address)
func (_MXCL1Client *MXCL1ClientCallerSession) Resolve(name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL1Client.Contract.Resolve(&_MXCL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_MXCL1Client *MXCL1ClientCaller) Resolve0(opts *bind.CallOpts, chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "resolve0", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_MXCL1Client *MXCL1ClientSession) Resolve0(chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL1Client.Contract.Resolve0(&_MXCL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0x1be2bfa7.
//
// Solidity: function resolve(uint256 chainId, string name, bool allowZeroAddress) view returns(address)
func (_MXCL1Client *MXCL1ClientCallerSession) Resolve0(chainId *big.Int, name string, allowZeroAddress bool) (common.Address, error) {
	return _MXCL1Client.Contract.Resolve0(&_MXCL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// SignWithGoldenTouch is a free data retrieval call binding the contract method 0xdadec12a.
//
// Solidity: function signWithGoldenTouch(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_MXCL1Client *MXCL1ClientCaller) SignWithGoldenTouch(opts *bind.CallOpts, hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "signWithGoldenTouch", hash, k)

	outstruct := new(struct {
		V uint8
		R *big.Int
		S *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SignWithGoldenTouch is a free data retrieval call binding the contract method 0xdadec12a.
//
// Solidity: function signWithGoldenTouch(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_MXCL1Client *MXCL1ClientSession) SignWithGoldenTouch(hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _MXCL1Client.Contract.SignWithGoldenTouch(&_MXCL1Client.CallOpts, hash, k)
}

// SignWithGoldenTouch is a free data retrieval call binding the contract method 0xdadec12a.
//
// Solidity: function signWithGoldenTouch(bytes32 hash, uint8 k) view returns(uint8 v, uint256 r, uint256 s)
func (_MXCL1Client *MXCL1ClientCallerSession) SignWithGoldenTouch(hash [32]byte, k uint8) (struct {
	V uint8
	R *big.Int
	S *big.Int
}, error) {
	return _MXCL1Client.Contract.SignWithGoldenTouch(&_MXCL1Client.CallOpts, hash, k)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reservedA1, uint64 __reservedA2, uint256 feeBase, uint64 nextBlockId, uint64 lastProposedAt, uint64 avgBlockTime, uint64 __avgGasLimit, uint64 latestVerifiedHeight, uint64 latestVerifiedId, uint64 avgProofTime, uint64 __reservedC1)
func (_MXCL1Client *MXCL1ClientCaller) State(opts *bind.CallOpts) (struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	ReservedA1           uint64
	ReservedA2           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	AvgGasLimit          uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
	ReservedC1           uint64
}, error) {
	var out []interface{}
	err := _MXCL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		GenesisHeight        uint64
		GenesisTimestamp     uint64
		ReservedA1           uint64
		ReservedA2           uint64
		FeeBase              *big.Int
		NextBlockId          uint64
		LastProposedAt       uint64
		AvgBlockTime         uint64
		AvgGasLimit          uint64
		LatestVerifiedHeight uint64
		LatestVerifiedId     uint64
		AvgProofTime         uint64
		ReservedC1           uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GenesisHeight = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.GenesisTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ReservedA1 = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ReservedA2 = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.FeeBase = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NextBlockId = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.LastProposedAt = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.AvgBlockTime = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.AvgGasLimit = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.LatestVerifiedHeight = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.LatestVerifiedId = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.AvgProofTime = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	outstruct.ReservedC1 = *abi.ConvertType(out[12], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reservedA1, uint64 __reservedA2, uint256 feeBase, uint64 nextBlockId, uint64 lastProposedAt, uint64 avgBlockTime, uint64 __avgGasLimit, uint64 latestVerifiedHeight, uint64 latestVerifiedId, uint64 avgProofTime, uint64 __reservedC1)
func (_MXCL1Client *MXCL1ClientSession) State() (struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	ReservedA1           uint64
	ReservedA2           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	AvgGasLimit          uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
	ReservedC1           uint64
}, error) {
	return _MXCL1Client.Contract.State(&_MXCL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint64 genesisHeight, uint64 genesisTimestamp, uint64 __reservedA1, uint64 __reservedA2, uint256 feeBase, uint64 nextBlockId, uint64 lastProposedAt, uint64 avgBlockTime, uint64 __avgGasLimit, uint64 latestVerifiedHeight, uint64 latestVerifiedId, uint64 avgProofTime, uint64 __reservedC1)
func (_MXCL1Client *MXCL1ClientCallerSession) State() (struct {
	GenesisHeight        uint64
	GenesisTimestamp     uint64
	ReservedA1           uint64
	ReservedA2           uint64
	FeeBase              *big.Int
	NextBlockId          uint64
	LastProposedAt       uint64
	AvgBlockTime         uint64
	AvgGasLimit          uint64
	LatestVerifiedHeight uint64
	LatestVerifiedId     uint64
	AvgProofTime         uint64
	ReservedC1           uint64
}, error) {
	return _MXCL1Client.Contract.State(&_MXCL1Client.CallOpts)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x7e7a262c.
//
// Solidity: function commitBlock(uint64 commitSlot, bytes32 commitHash) returns()
func (_MXCL1Client *MXCL1ClientTransactor) CommitBlock(opts *bind.TransactOpts, commitSlot uint64, commitHash [32]byte) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "commitBlock", commitSlot, commitHash)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x7e7a262c.
//
// Solidity: function commitBlock(uint64 commitSlot, bytes32 commitHash) returns()
func (_MXCL1Client *MXCL1ClientSession) CommitBlock(commitSlot uint64, commitHash [32]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.CommitBlock(&_MXCL1Client.TransactOpts, commitSlot, commitHash)
}

// CommitBlock is a paid mutator transaction binding the contract method 0x7e7a262c.
//
// Solidity: function commitBlock(uint64 commitSlot, bytes32 commitHash) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) CommitBlock(commitSlot uint64, commitHash [32]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.CommitBlock(&_MXCL1Client.TransactOpts, commitSlot, commitHash)
}

// Init is a paid mutator transaction binding the contract method 0x9c5e9f06.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint256 _feeBase) returns()
func (_MXCL1Client *MXCL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte, _feeBase *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash, _feeBase)
}

// Init is a paid mutator transaction binding the contract method 0x9c5e9f06.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint256 _feeBase) returns()
func (_MXCL1Client *MXCL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _feeBase *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.Contract.Init(&_MXCL1Client.TransactOpts, _addressManager, _genesisBlockHash, _feeBase)
}

// Init is a paid mutator transaction binding the contract method 0x9c5e9f06.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint256 _feeBase) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _feeBase *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.Contract.Init(&_MXCL1Client.TransactOpts, _addressManager, _genesisBlockHash, _feeBase)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "proposeBlock", inputs)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientSession) ProposeBlock(inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.ProposeBlock(&_MXCL1Client.TransactOpts, inputs)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xa043dbdf.
//
// Solidity: function proposeBlock(bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) ProposeBlock(inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.ProposeBlock(&_MXCL1Client.TransactOpts, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockId, bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "proveBlock", blockId, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockId, bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientSession) ProveBlock(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.ProveBlock(&_MXCL1Client.TransactOpts, blockId, inputs)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x8ed7b3be.
//
// Solidity: function proveBlock(uint256 blockId, bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) ProveBlock(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.ProveBlock(&_MXCL1Client.TransactOpts, blockId, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockId, bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientTransactor) ProveBlockInvalid(opts *bind.TransactOpts, blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "proveBlockInvalid", blockId, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockId, bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientSession) ProveBlockInvalid(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.ProveBlockInvalid(&_MXCL1Client.TransactOpts, blockId, inputs)
}

// ProveBlockInvalid is a paid mutator transaction binding the contract method 0xa279cec7.
//
// Solidity: function proveBlockInvalid(uint256 blockId, bytes[] inputs) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) ProveBlockInvalid(blockId *big.Int, inputs [][]byte) (*types.Transaction, error) {
	return _MXCL1Client.Contract.ProveBlockInvalid(&_MXCL1Client.TransactOpts, blockId, inputs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MXCL1Client *MXCL1ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MXCL1Client *MXCL1ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _MXCL1Client.Contract.RenounceOwnership(&_MXCL1Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MXCL1Client.Contract.RenounceOwnership(&_MXCL1Client.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_MXCL1Client *MXCL1ClientTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_MXCL1Client *MXCL1ClientSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.Contract.Stake(&_MXCL1Client.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.Contract.Stake(&_MXCL1Client.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MXCL1Client *MXCL1ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MXCL1Client *MXCL1ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MXCL1Client.Contract.TransferOwnership(&_MXCL1Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MXCL1Client.Contract.TransferOwnership(&_MXCL1Client.TransactOpts, newOwner)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns()
func (_MXCL1Client *MXCL1ClientTransactor) UnStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "unStake")
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns()
func (_MXCL1Client *MXCL1ClientSession) UnStake() (*types.Transaction, error) {
	return _MXCL1Client.Contract.UnStake(&_MXCL1Client.TransactOpts)
}

// UnStake is a paid mutator transaction binding the contract method 0x73cf575a.
//
// Solidity: function unStake() returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) UnStake() (*types.Transaction, error) {
	return _MXCL1Client.Contract.UnStake(&_MXCL1Client.TransactOpts)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_MXCL1Client *MXCL1ClientTransactor) VerifyBlocks(opts *bind.TransactOpts, maxBlocks *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "verifyBlocks", maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_MXCL1Client *MXCL1ClientSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.Contract.VerifyBlocks(&_MXCL1Client.TransactOpts, maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _MXCL1Client.Contract.VerifyBlocks(&_MXCL1Client.TransactOpts, maxBlocks)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_MXCL1Client *MXCL1ClientTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MXCL1Client.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_MXCL1Client *MXCL1ClientSession) WithdrawBalance() (*types.Transaction, error) {
	return _MXCL1Client.Contract.WithdrawBalance(&_MXCL1Client.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_MXCL1Client *MXCL1ClientTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _MXCL1Client.Contract.WithdrawBalance(&_MXCL1Client.TransactOpts)
}

// MXCL1ClientBlockCommittedIterator is returned from FilterBlockCommitted and is used to iterate over the raw logs and unpacked data for BlockCommitted events raised by the MXCL1Client contract.
type MXCL1ClientBlockCommittedIterator struct {
	Event *MXCL1ClientBlockCommitted // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientBlockCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientBlockCommitted)
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
		it.Event = new(MXCL1ClientBlockCommitted)
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
func (it *MXCL1ClientBlockCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientBlockCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientBlockCommitted represents a BlockCommitted event raised by the MXCL1Client contract.
type MXCL1ClientBlockCommitted struct {
	CommitSlot uint64
	CommitHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockCommitted is a free log retrieval operation binding the contract event 0x51264991e22d808f3bcbb1cbffa82b752eae327c24055259a5c455c0aa5b7136.
//
// Solidity: event BlockCommitted(uint64 commitSlot, bytes32 commitHash)
func (_MXCL1Client *MXCL1ClientFilterer) FilterBlockCommitted(opts *bind.FilterOpts) (*MXCL1ClientBlockCommittedIterator, error) {

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "BlockCommitted")
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientBlockCommittedIterator{contract: _MXCL1Client.contract, event: "BlockCommitted", logs: logs, sub: sub}, nil
}

// WatchBlockCommitted is a free log subscription operation binding the contract event 0x51264991e22d808f3bcbb1cbffa82b752eae327c24055259a5c455c0aa5b7136.
//
// Solidity: event BlockCommitted(uint64 commitSlot, bytes32 commitHash)
func (_MXCL1Client *MXCL1ClientFilterer) WatchBlockCommitted(opts *bind.WatchOpts, sink chan<- *MXCL1ClientBlockCommitted) (event.Subscription, error) {

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "BlockCommitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientBlockCommitted)
				if err := _MXCL1Client.contract.UnpackLog(event, "BlockCommitted", log); err != nil {
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

// ParseBlockCommitted is a log parse operation binding the contract event 0x51264991e22d808f3bcbb1cbffa82b752eae327c24055259a5c455c0aa5b7136.
//
// Solidity: event BlockCommitted(uint64 commitSlot, bytes32 commitHash)
func (_MXCL1Client *MXCL1ClientFilterer) ParseBlockCommitted(log types.Log) (*MXCL1ClientBlockCommitted, error) {
	event := new(MXCL1ClientBlockCommitted)
	if err := _MXCL1Client.contract.UnpackLog(event, "BlockCommitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL1ClientBlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the MXCL1Client contract.
type MXCL1ClientBlockProposedIterator struct {
	Event *MXCL1ClientBlockProposed // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientBlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientBlockProposed)
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
		it.Event = new(MXCL1ClientBlockProposed)
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
func (it *MXCL1ClientBlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientBlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientBlockProposed represents a BlockProposed event raised by the MXCL1Client contract.
type MXCL1ClientBlockProposed struct {
	Id   *big.Int
	Meta MXCDataBlockMetadata
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0x344fc5d5f80c4a29bd7e06ad7c28a0c8c9c08d682129da3a31936d5982e4f044.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,bytes32,bytes32,bytes,uint64,uint64,uint64,uint64) meta)
func (_MXCL1Client *MXCL1ClientFilterer) FilterBlockProposed(opts *bind.FilterOpts, id []*big.Int) (*MXCL1ClientBlockProposedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientBlockProposedIterator{contract: _MXCL1Client.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0x344fc5d5f80c4a29bd7e06ad7c28a0c8c9c08d682129da3a31936d5982e4f044.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,bytes32,bytes32,bytes,uint64,uint64,uint64,uint64) meta)
func (_MXCL1Client *MXCL1ClientFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *MXCL1ClientBlockProposed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientBlockProposed)
				if err := _MXCL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
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

// ParseBlockProposed is a log parse operation binding the contract event 0x344fc5d5f80c4a29bd7e06ad7c28a0c8c9c08d682129da3a31936d5982e4f044.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint256,uint256,bytes32,address,bytes32,bytes32,bytes,uint64,uint64,uint64,uint64) meta)
func (_MXCL1Client *MXCL1ClientFilterer) ParseBlockProposed(log types.Log) (*MXCL1ClientBlockProposed, error) {
	event := new(MXCL1ClientBlockProposed)
	if err := _MXCL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL1ClientBlockProvenIterator is returned from FilterBlockProven and is used to iterate over the raw logs and unpacked data for BlockProven events raised by the MXCL1Client contract.
type MXCL1ClientBlockProvenIterator struct {
	Event *MXCL1ClientBlockProven // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientBlockProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientBlockProven)
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
		it.Event = new(MXCL1ClientBlockProven)
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
func (it *MXCL1ClientBlockProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientBlockProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientBlockProven represents a BlockProven event raised by the MXCL1Client contract.
type MXCL1ClientBlockProven struct {
	Id         *big.Int
	ParentHash [32]byte
	BlockHash  [32]byte
	Prover     common.Address
	ProvenAt   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0x45848a3b2a67571e5876283456675aa3e05880e4f5a73447bd86cef5a181db38.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, address prover, uint64 provenAt)
func (_MXCL1Client *MXCL1ClientFilterer) FilterBlockProven(opts *bind.FilterOpts, id []*big.Int) (*MXCL1ClientBlockProvenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientBlockProvenIterator{contract: _MXCL1Client.contract, event: "BlockProven", logs: logs, sub: sub}, nil
}

// WatchBlockProven is a free log subscription operation binding the contract event 0x45848a3b2a67571e5876283456675aa3e05880e4f5a73447bd86cef5a181db38.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, address prover, uint64 provenAt)
func (_MXCL1Client *MXCL1ClientFilterer) WatchBlockProven(opts *bind.WatchOpts, sink chan<- *MXCL1ClientBlockProven, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientBlockProven)
				if err := _MXCL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
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

// ParseBlockProven is a log parse operation binding the contract event 0x45848a3b2a67571e5876283456675aa3e05880e4f5a73447bd86cef5a181db38.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, address prover, uint64 provenAt)
func (_MXCL1Client *MXCL1ClientFilterer) ParseBlockProven(log types.Log) (*MXCL1ClientBlockProven, error) {
	event := new(MXCL1ClientBlockProven)
	if err := _MXCL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL1ClientBlockVerifiedIterator is returned from FilterBlockVerified and is used to iterate over the raw logs and unpacked data for BlockVerified events raised by the MXCL1Client contract.
type MXCL1ClientBlockVerifiedIterator struct {
	Event *MXCL1ClientBlockVerified // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientBlockVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientBlockVerified)
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
		it.Event = new(MXCL1ClientBlockVerified)
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
func (it *MXCL1ClientBlockVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientBlockVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientBlockVerified represents a BlockVerified event raised by the MXCL1Client contract.
type MXCL1ClientBlockVerified struct {
	Id        *big.Int
	BlockHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_MXCL1Client *MXCL1ClientFilterer) FilterBlockVerified(opts *bind.FilterOpts, id []*big.Int) (*MXCL1ClientBlockVerifiedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientBlockVerifiedIterator{contract: _MXCL1Client.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_MXCL1Client *MXCL1ClientFilterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *MXCL1ClientBlockVerified, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientBlockVerified)
				if err := _MXCL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
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

// ParseBlockVerified is a log parse operation binding the contract event 0x68b82650828a9621868d09dc161400acbe189fa002e3fb7cf9dea5c2c6f928ee.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash)
func (_MXCL1Client *MXCL1ClientFilterer) ParseBlockVerified(log types.Log) (*MXCL1ClientBlockVerified, error) {
	event := new(MXCL1ClientBlockVerified)
	if err := _MXCL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL1ClientHeaderSyncedIterator is returned from FilterHeaderSynced and is used to iterate over the raw logs and unpacked data for HeaderSynced events raised by the MXCL1Client contract.
type MXCL1ClientHeaderSyncedIterator struct {
	Event *MXCL1ClientHeaderSynced // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientHeaderSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientHeaderSynced)
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
		it.Event = new(MXCL1ClientHeaderSynced)
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
func (it *MXCL1ClientHeaderSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientHeaderSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientHeaderSynced represents a HeaderSynced event raised by the MXCL1Client contract.
type MXCL1ClientHeaderSynced struct {
	SrcHeight *big.Int
	SrcHash   [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHeaderSynced is a free log retrieval operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_MXCL1Client *MXCL1ClientFilterer) FilterHeaderSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*MXCL1ClientHeaderSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "HeaderSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientHeaderSyncedIterator{contract: _MXCL1Client.contract, event: "HeaderSynced", logs: logs, sub: sub}, nil
}

// WatchHeaderSynced is a free log subscription operation binding the contract event 0x58313b60ec6c5bfc381e52f0de3ede0faac3cdffea26f7d6bcc3d09b61018691.
//
// Solidity: event HeaderSynced(uint256 indexed srcHeight, bytes32 srcHash)
func (_MXCL1Client *MXCL1ClientFilterer) WatchHeaderSynced(opts *bind.WatchOpts, sink chan<- *MXCL1ClientHeaderSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "HeaderSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientHeaderSynced)
				if err := _MXCL1Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
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
func (_MXCL1Client *MXCL1ClientFilterer) ParseHeaderSynced(log types.Log) (*MXCL1ClientHeaderSynced, error) {
	event := new(MXCL1ClientHeaderSynced)
	if err := _MXCL1Client.contract.UnpackLog(event, "HeaderSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL1ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MXCL1Client contract.
type MXCL1ClientInitializedIterator struct {
	Event *MXCL1ClientInitialized // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientInitialized)
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
		it.Event = new(MXCL1ClientInitialized)
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
func (it *MXCL1ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientInitialized represents a Initialized event raised by the MXCL1Client contract.
type MXCL1ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MXCL1Client *MXCL1ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*MXCL1ClientInitializedIterator, error) {

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientInitializedIterator{contract: _MXCL1Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MXCL1Client *MXCL1ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MXCL1ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientInitialized)
				if err := _MXCL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MXCL1Client *MXCL1ClientFilterer) ParseInitialized(log types.Log) (*MXCL1ClientInitialized, error) {
	event := new(MXCL1ClientInitialized)
	if err := _MXCL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MXCL1ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MXCL1Client contract.
type MXCL1ClientOwnershipTransferredIterator struct {
	Event *MXCL1ClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MXCL1ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MXCL1ClientOwnershipTransferred)
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
		it.Event = new(MXCL1ClientOwnershipTransferred)
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
func (it *MXCL1ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MXCL1ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MXCL1ClientOwnershipTransferred represents a OwnershipTransferred event raised by the MXCL1Client contract.
type MXCL1ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MXCL1Client *MXCL1ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MXCL1ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MXCL1Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MXCL1ClientOwnershipTransferredIterator{contract: _MXCL1Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MXCL1Client *MXCL1ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MXCL1ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MXCL1Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MXCL1ClientOwnershipTransferred)
				if err := _MXCL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MXCL1Client *MXCL1ClientFilterer) ParseOwnershipTransferred(log types.Log) (*MXCL1ClientOwnershipTransferred, error) {
	event := new(MXCL1ClientOwnershipTransferred)
	if err := _MXCL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
