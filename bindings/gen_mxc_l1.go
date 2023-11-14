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

// MxcDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type MxcDataBlockMetadata struct {
	Id                uint64
	Timestamp         uint64
	L1Height          uint64
	L1Hash            [32]byte
	MixHash           [32]byte
	TxListHash        [32]byte
	TxListByteStart   *big.Int
	TxListByteEnd     *big.Int
	GasLimit          uint32
	Beneficiary       common.Address
	Treasury          common.Address
	DepositsProcessed []MxcDataEthDeposit
	BaseFee           *big.Int
	BlockReward       *big.Int
}

// MxcDataConfig is an auto generated low-level Go binding around an user-defined struct.
type MxcDataConfig struct {
	ChainId                   *big.Int
	MaxNumProposedBlocks      *big.Int
	RingBufferSize            *big.Int
	MaxVerificationsPerTx     *big.Int
	BlockMaxGasLimit          uint64
	MaxTransactionsPerBlock   uint64
	MaxBytesPerTxList         uint64
	TxListCacheExpiry         *big.Int
	ProofCooldownPeriod       *big.Int
	SystemProofCooldownPeriod *big.Int
	RealProofSkipSize         *big.Int
	EthDepositGas             *big.Int
	EthDepositMaxFee          *big.Int
	MinEthDepositsPerBlock    uint64
	MaxEthDepositsPerBlock    uint64
	MaxEthDepositAmount       *big.Int
	MinEthDepositAmount       *big.Int
	RelaySignalRoot           bool
}

// MxcDataEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type MxcDataEthDeposit struct {
	Recipient common.Address
	Amount    *big.Int
	Id        uint64
}

// MxcDataForkChoice is an auto generated low-level Go binding around an user-defined struct.
type MxcDataForkChoice struct {
	Key        [32]byte
	BlockHash  [32]byte
	SignalRoot [32]byte
	ProvenAt   uint64
	Prover     common.Address
	GasUsed    uint32
}

// MxcDataStateVariables is an auto generated low-level Go binding around an user-defined struct.
type MxcDataStateVariables struct {
	BlockFee                uint64
	AccBlockFees            uint64
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NumBlocks               uint64
	ProofTimeIssued         uint64
	ProofTimeTarget         uint64
	LastVerifiedBlockId     uint64
	AccProposedAt           uint64
	NextEthDepositToProcess uint64
	NumEthDeposits          uint64
}

// MxcL1ClientMetaData contains all meta data concerning the MxcL1Client contract.
var MxcL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_DEPOSIT_REQUIREMENT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ELECTION_INVALID_PROPOSER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ELECTION_SPEED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ELECTION_SPEED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_UNLOCK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockProposeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structMxcData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockProvenReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structMxcData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"ProofParamsChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_proposedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumProposedBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ringBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txListCacheExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"systemProofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"realProofSkipSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"minEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"}],\"internalType\":\"structMxcData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"}],\"internalType\":\"structMxcData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMxcTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proofTime\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getVerifierName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_initBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"estimateGas\",\"type\":\"uint256\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"}],\"internalType\":\"structMxcData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newProofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newProofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"newAdjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"setProofParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStakeMxcTokenBalances\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"prevBaseFee\",\"type\":\"uint48\"},{\"internalType\":\"uint64\",\"name\":\"proposerElectionTimeoutOffset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proveMetaReward\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MxcL1ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use MxcL1ClientMetaData.ABI instead.
var MxcL1ClientABI = MxcL1ClientMetaData.ABI

// MxcL1Client is an auto generated Go binding around an Ethereum contract.
type MxcL1Client struct {
	MxcL1ClientCaller     // Read-only binding to the contract
	MxcL1ClientTransactor // Write-only binding to the contract
	MxcL1ClientFilterer   // Log filterer for contract events
}

// MxcL1ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type MxcL1ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcL1ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MxcL1ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcL1ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MxcL1ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MxcL1ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MxcL1ClientSession struct {
	Contract     *MxcL1Client      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MxcL1ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MxcL1ClientCallerSession struct {
	Contract *MxcL1ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MxcL1ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MxcL1ClientTransactorSession struct {
	Contract     *MxcL1ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MxcL1ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type MxcL1ClientRaw struct {
	Contract *MxcL1Client // Generic contract binding to access the raw methods on
}

// MxcL1ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MxcL1ClientCallerRaw struct {
	Contract *MxcL1ClientCaller // Generic read-only contract binding to access the raw methods on
}

// MxcL1ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MxcL1ClientTransactorRaw struct {
	Contract *MxcL1ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMxcL1Client creates a new instance of MxcL1Client, bound to a specific deployed contract.
func NewMxcL1Client(address common.Address, backend bind.ContractBackend) (*MxcL1Client, error) {
	contract, err := bindMxcL1Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MxcL1Client{MxcL1ClientCaller: MxcL1ClientCaller{contract: contract}, MxcL1ClientTransactor: MxcL1ClientTransactor{contract: contract}, MxcL1ClientFilterer: MxcL1ClientFilterer{contract: contract}}, nil
}

// NewMxcL1ClientCaller creates a new read-only instance of MxcL1Client, bound to a specific deployed contract.
func NewMxcL1ClientCaller(address common.Address, caller bind.ContractCaller) (*MxcL1ClientCaller, error) {
	contract, err := bindMxcL1Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientCaller{contract: contract}, nil
}

// NewMxcL1ClientTransactor creates a new write-only instance of MxcL1Client, bound to a specific deployed contract.
func NewMxcL1ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*MxcL1ClientTransactor, error) {
	contract, err := bindMxcL1Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientTransactor{contract: contract}, nil
}

// NewMxcL1ClientFilterer creates a new log filterer instance of MxcL1Client, bound to a specific deployed contract.
func NewMxcL1ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*MxcL1ClientFilterer, error) {
	contract, err := bindMxcL1Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientFilterer{contract: contract}, nil
}

// bindMxcL1Client binds a generic wrapper to an already deployed contract.
func bindMxcL1Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MxcL1ClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MxcL1Client *MxcL1ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MxcL1Client.Contract.MxcL1ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MxcL1Client *MxcL1ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL1Client.Contract.MxcL1ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MxcL1Client *MxcL1ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MxcL1Client.Contract.MxcL1ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MxcL1Client *MxcL1ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MxcL1Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MxcL1Client *MxcL1ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL1Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MxcL1Client *MxcL1ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MxcL1Client.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcL1Client *MxcL1ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcL1Client *MxcL1ClientSession) AddressManager() (common.Address, error) {
	return _MxcL1Client.Contract.AddressManager(&_MxcL1Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_MxcL1Client *MxcL1ClientCallerSession) AddressManager() (common.Address, error) {
	return _MxcL1Client.Contract.AddressManager(&_MxcL1Client.CallOpts)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_MxcL1Client *MxcL1ClientCaller) GetBlock(opts *bind.CallOpts, blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getBlock", blockId)

	outstruct := new(struct {
		MetaHash   [32]byte
		Proposer   common.Address
		ProposedAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetaHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ProposedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_MxcL1Client *MxcL1ClientSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _MxcL1Client.Contract.GetBlock(&_MxcL1Client.CallOpts, blockId)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 blockId) view returns(bytes32 _metaHash, address _proposer, uint64 _proposedAt)
func (_MxcL1Client *MxcL1ClientCallerSession) GetBlock(blockId *big.Int) (struct {
	MetaHash   [32]byte
	Proposer   common.Address
	ProposedAt uint64
}, error) {
	return _MxcL1Client.Contract.GetBlock(&_MxcL1Client.CallOpts, blockId)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_MxcL1Client *MxcL1ClientCaller) GetBlockFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getBlockFee")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_MxcL1Client *MxcL1ClientSession) GetBlockFee() (uint64, error) {
	return _MxcL1Client.Contract.GetBlockFee(&_MxcL1Client.CallOpts)
}

// GetBlockFee is a free data retrieval call binding the contract method 0x7baf0bc7.
//
// Solidity: function getBlockFee() view returns(uint64)
func (_MxcL1Client *MxcL1ClientCallerSession) GetBlockFee() (uint64, error) {
	return _MxcL1Client.Contract.GetBlockFee(&_MxcL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,bool))
func (_MxcL1Client *MxcL1ClientCaller) GetConfig(opts *bind.CallOpts) (MxcDataConfig, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(MxcDataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MxcDataConfig)).(*MxcDataConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,bool))
func (_MxcL1Client *MxcL1ClientSession) GetConfig() (MxcDataConfig, error) {
	return _MxcL1Client.Contract.GetConfig(&_MxcL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() pure returns((uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint96,uint96,bool))
func (_MxcL1Client *MxcL1ClientCallerSession) GetConfig() (MxcDataConfig, error) {
	return _MxcL1Client.Contract.GetConfig(&_MxcL1Client.CallOpts)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_MxcL1Client *MxcL1ClientCaller) GetCrossChainBlockHash(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getCrossChainBlockHash", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_MxcL1Client *MxcL1ClientSession) GetCrossChainBlockHash(blockId *big.Int) ([32]byte, error) {
	return _MxcL1Client.Contract.GetCrossChainBlockHash(&_MxcL1Client.CallOpts, blockId)
}

// GetCrossChainBlockHash is a free data retrieval call binding the contract method 0xbacb386d.
//
// Solidity: function getCrossChainBlockHash(uint256 blockId) view returns(bytes32)
func (_MxcL1Client *MxcL1ClientCallerSession) GetCrossChainBlockHash(blockId *big.Int) ([32]byte, error) {
	return _MxcL1Client.Contract.GetCrossChainBlockHash(&_MxcL1Client.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_MxcL1Client *MxcL1ClientCaller) GetCrossChainSignalRoot(opts *bind.CallOpts, blockId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getCrossChainSignalRoot", blockId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_MxcL1Client *MxcL1ClientSession) GetCrossChainSignalRoot(blockId *big.Int) ([32]byte, error) {
	return _MxcL1Client.Contract.GetCrossChainSignalRoot(&_MxcL1Client.CallOpts, blockId)
}

// GetCrossChainSignalRoot is a free data retrieval call binding the contract method 0xb8914ce4.
//
// Solidity: function getCrossChainSignalRoot(uint256 blockId) view returns(bytes32)
func (_MxcL1Client *MxcL1ClientCallerSession) GetCrossChainSignalRoot(blockId *big.Int) ([32]byte, error) {
	return _MxcL1Client.Contract.GetCrossChainSignalRoot(&_MxcL1Client.CallOpts, blockId)
}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_MxcL1Client *MxcL1ClientCaller) GetForkChoice(opts *bind.CallOpts, blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (MxcDataForkChoice, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getForkChoice", blockId, parentHash, parentGasUsed)

	if err != nil {
		return *new(MxcDataForkChoice), err
	}

	out0 := *abi.ConvertType(out[0], new(MxcDataForkChoice)).(*MxcDataForkChoice)

	return out0, err

}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_MxcL1Client *MxcL1ClientSession) GetForkChoice(blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (MxcDataForkChoice, error) {
	return _MxcL1Client.Contract.GetForkChoice(&_MxcL1Client.CallOpts, blockId, parentHash, parentGasUsed)
}

// GetForkChoice is a free data retrieval call binding the contract method 0x7163e0ed.
//
// Solidity: function getForkChoice(uint256 blockId, bytes32 parentHash, uint32 parentGasUsed) view returns((bytes32,bytes32,bytes32,uint64,address,uint32))
func (_MxcL1Client *MxcL1ClientCallerSession) GetForkChoice(blockId *big.Int, parentHash [32]byte, parentGasUsed uint32) (MxcDataForkChoice, error) {
	return _MxcL1Client.Contract.GetForkChoice(&_MxcL1Client.CallOpts, blockId, parentHash, parentGasUsed)
}

// GetMxcTokenBalance is a free data retrieval call binding the contract method 0xc21f0775.
//
// Solidity: function getMxcTokenBalance(address addr) view returns(uint256)
func (_MxcL1Client *MxcL1ClientCaller) GetMxcTokenBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getMxcTokenBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMxcTokenBalance is a free data retrieval call binding the contract method 0xc21f0775.
//
// Solidity: function getMxcTokenBalance(address addr) view returns(uint256)
func (_MxcL1Client *MxcL1ClientSession) GetMxcTokenBalance(addr common.Address) (*big.Int, error) {
	return _MxcL1Client.Contract.GetMxcTokenBalance(&_MxcL1Client.CallOpts, addr)
}

// GetMxcTokenBalance is a free data retrieval call binding the contract method 0xc21f0775.
//
// Solidity: function getMxcTokenBalance(address addr) view returns(uint256)
func (_MxcL1Client *MxcL1ClientCallerSession) GetMxcTokenBalance(addr common.Address) (*big.Int, error) {
	return _MxcL1Client.Contract.GetMxcTokenBalance(&_MxcL1Client.CallOpts, addr)
}

// GetProofReward is a free data retrieval call binding the contract method 0x55f7259e.
//
// Solidity: function getProofReward(uint64 proofTime) view returns(uint256)
func (_MxcL1Client *MxcL1ClientCaller) GetProofReward(opts *bind.CallOpts, proofTime uint64) (*big.Int, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getProofReward", proofTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProofReward is a free data retrieval call binding the contract method 0x55f7259e.
//
// Solidity: function getProofReward(uint64 proofTime) view returns(uint256)
func (_MxcL1Client *MxcL1ClientSession) GetProofReward(proofTime uint64) (*big.Int, error) {
	return _MxcL1Client.Contract.GetProofReward(&_MxcL1Client.CallOpts, proofTime)
}

// GetProofReward is a free data retrieval call binding the contract method 0x55f7259e.
//
// Solidity: function getProofReward(uint64 proofTime) view returns(uint256)
func (_MxcL1Client *MxcL1ClientCallerSession) GetProofReward(proofTime uint64) (*big.Int, error) {
	return _MxcL1Client.Contract.GetProofReward(&_MxcL1Client.CallOpts, proofTime)
}

// GetProposeReward is a free data retrieval call binding the contract method 0xe527c8b7.
//
// Solidity: function getProposeReward() view returns(uint256)
func (_MxcL1Client *MxcL1ClientCaller) GetProposeReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getProposeReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposeReward is a free data retrieval call binding the contract method 0xe527c8b7.
//
// Solidity: function getProposeReward() view returns(uint256)
func (_MxcL1Client *MxcL1ClientSession) GetProposeReward() (*big.Int, error) {
	return _MxcL1Client.Contract.GetProposeReward(&_MxcL1Client.CallOpts)
}

// GetProposeReward is a free data retrieval call binding the contract method 0xe527c8b7.
//
// Solidity: function getProposeReward() view returns(uint256)
func (_MxcL1Client *MxcL1ClientCallerSession) GetProposeReward() (*big.Int, error) {
	return _MxcL1Client.Contract.GetProposeReward(&_MxcL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_MxcL1Client *MxcL1ClientCaller) GetStateVariables(opts *bind.CallOpts) (MxcDataStateVariables, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getStateVariables")

	if err != nil {
		return *new(MxcDataStateVariables), err
	}

	out0 := *abi.ConvertType(out[0], new(MxcDataStateVariables)).(*MxcDataStateVariables)

	return out0, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_MxcL1Client *MxcL1ClientSession) GetStateVariables() (MxcDataStateVariables, error) {
	return _MxcL1Client.Contract.GetStateVariables(&_MxcL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_MxcL1Client *MxcL1ClientCallerSession) GetStateVariables() (MxcDataStateVariables, error) {
	return _MxcL1Client.Contract.GetStateVariables(&_MxcL1Client.CallOpts)
}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_MxcL1Client *MxcL1ClientCaller) GetVerifierName(opts *bind.CallOpts, id uint16) ([32]byte, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "getVerifierName", id)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_MxcL1Client *MxcL1ClientSession) GetVerifierName(id uint16) ([32]byte, error) {
	return _MxcL1Client.Contract.GetVerifierName(&_MxcL1Client.CallOpts, id)
}

// GetVerifierName is a free data retrieval call binding the contract method 0x0372303d.
//
// Solidity: function getVerifierName(uint16 id) pure returns(bytes32)
func (_MxcL1Client *MxcL1ClientCallerSession) GetVerifierName(id uint16) ([32]byte, error) {
	return _MxcL1Client.Contract.GetVerifierName(&_MxcL1Client.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcL1Client *MxcL1ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcL1Client *MxcL1ClientSession) Owner() (common.Address, error) {
	return _MxcL1Client.Contract.Owner(&_MxcL1Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MxcL1Client *MxcL1ClientCallerSession) Owner() (common.Address, error) {
	return _MxcL1Client.Contract.Owner(&_MxcL1Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL1Client *MxcL1ClientCaller) Resolve(opts *bind.CallOpts, chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL1Client *MxcL1ClientSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL1Client.Contract.Resolve(&_MxcL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x6c6563f6.
//
// Solidity: function resolve(uint256 chainId, bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL1Client *MxcL1ClientCallerSession) Resolve(chainId *big.Int, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL1Client.Contract.Resolve(&_MxcL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL1Client *MxcL1ClientCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL1Client *MxcL1ClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL1Client.Contract.Resolve0(&_MxcL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address)
func (_MxcL1Client *MxcL1ClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _MxcL1Client.Contract.Resolve0(&_MxcL1Client.CallOpts, name, allowZeroAddress)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 totalStakeMxcTokenBalances, uint64 genesisHeight, uint64 genesisTimestamp, uint16 adjustmentQuotient, uint48 prevBaseFee, uint64 proposerElectionTimeoutOffset, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 proveMetaReward, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 proofTimeTarget)
func (_MxcL1Client *MxcL1ClientCaller) State(opts *bind.CallOpts) (struct {
	TotalStakeMxcTokenBalances    *big.Int
	GenesisHeight                 uint64
	GenesisTimestamp              uint64
	AdjustmentQuotient            uint16
	PrevBaseFee                   *big.Int
	ProposerElectionTimeoutOffset uint64
	AccProposedAt                 uint64
	AccBlockFees                  uint64
	NumBlocks                     uint64
	ProveMetaReward               uint64
	BlockFee                      uint64
	ProofTimeIssued               uint64
	LastVerifiedBlockId           uint64
	ProofTimeTarget               uint64
}, error) {
	var out []interface{}
	err := _MxcL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		TotalStakeMxcTokenBalances    *big.Int
		GenesisHeight                 uint64
		GenesisTimestamp              uint64
		AdjustmentQuotient            uint16
		PrevBaseFee                   *big.Int
		ProposerElectionTimeoutOffset uint64
		AccProposedAt                 uint64
		AccBlockFees                  uint64
		NumBlocks                     uint64
		ProveMetaReward               uint64
		BlockFee                      uint64
		ProofTimeIssued               uint64
		LastVerifiedBlockId           uint64
		ProofTimeTarget               uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalStakeMxcTokenBalances = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GenesisHeight = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.GenesisTimestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.AdjustmentQuotient = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.PrevBaseFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ProposerElectionTimeoutOffset = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.AccProposedAt = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.AccBlockFees = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.NumBlocks = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.ProveMetaReward = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.BlockFee = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.ProofTimeIssued = *abi.ConvertType(out[11], new(uint64)).(*uint64)
	outstruct.LastVerifiedBlockId = *abi.ConvertType(out[12], new(uint64)).(*uint64)
	outstruct.ProofTimeTarget = *abi.ConvertType(out[13], new(uint64)).(*uint64)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 totalStakeMxcTokenBalances, uint64 genesisHeight, uint64 genesisTimestamp, uint16 adjustmentQuotient, uint48 prevBaseFee, uint64 proposerElectionTimeoutOffset, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 proveMetaReward, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 proofTimeTarget)
func (_MxcL1Client *MxcL1ClientSession) State() (struct {
	TotalStakeMxcTokenBalances    *big.Int
	GenesisHeight                 uint64
	GenesisTimestamp              uint64
	AdjustmentQuotient            uint16
	PrevBaseFee                   *big.Int
	ProposerElectionTimeoutOffset uint64
	AccProposedAt                 uint64
	AccBlockFees                  uint64
	NumBlocks                     uint64
	ProveMetaReward               uint64
	BlockFee                      uint64
	ProofTimeIssued               uint64
	LastVerifiedBlockId           uint64
	ProofTimeTarget               uint64
}, error) {
	return _MxcL1Client.Contract.State(&_MxcL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint256 totalStakeMxcTokenBalances, uint64 genesisHeight, uint64 genesisTimestamp, uint16 adjustmentQuotient, uint48 prevBaseFee, uint64 proposerElectionTimeoutOffset, uint64 accProposedAt, uint64 accBlockFees, uint64 numBlocks, uint64 proveMetaReward, uint64 blockFee, uint64 proofTimeIssued, uint64 lastVerifiedBlockId, uint64 proofTimeTarget)
func (_MxcL1Client *MxcL1ClientCallerSession) State() (struct {
	TotalStakeMxcTokenBalances    *big.Int
	GenesisHeight                 uint64
	GenesisTimestamp              uint64
	AdjustmentQuotient            uint16
	PrevBaseFee                   *big.Int
	ProposerElectionTimeoutOffset uint64
	AccProposedAt                 uint64
	AccBlockFees                  uint64
	NumBlocks                     uint64
	ProveMetaReward               uint64
	BlockFee                      uint64
	ProofTimeIssued               uint64
	LastVerifiedBlockId           uint64
	ProofTimeTarget               uint64
}, error) {
	return _MxcL1Client.Contract.State(&_MxcL1Client.CallOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_MxcL1Client *MxcL1ClientTransactor) DepositEtherToL2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "depositEtherToL2")
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_MxcL1Client *MxcL1ClientSession) DepositEtherToL2() (*types.Transaction, error) {
	return _MxcL1Client.Contract.DepositEtherToL2(&_MxcL1Client.TransactOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0xa22f7670.
//
// Solidity: function depositEtherToL2() payable returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) DepositEtherToL2() (*types.Transaction, error) {
	return _MxcL1Client.Contract.DepositEtherToL2(&_MxcL1Client.TransactOpts)
}

// DepositMxcToken is a paid mutator transaction binding the contract method 0x5c9c0a14.
//
// Solidity: function depositMxcToken(uint256 amount) returns()
func (_MxcL1Client *MxcL1ClientTransactor) DepositMxcToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "depositMxcToken", amount)
}

// DepositMxcToken is a paid mutator transaction binding the contract method 0x5c9c0a14.
//
// Solidity: function depositMxcToken(uint256 amount) returns()
func (_MxcL1Client *MxcL1ClientSession) DepositMxcToken(amount *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.DepositMxcToken(&_MxcL1Client.TransactOpts, amount)
}

// DepositMxcToken is a paid mutator transaction binding the contract method 0x5c9c0a14.
//
// Solidity: function depositMxcToken(uint256 amount) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) DepositMxcToken(amount *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.DepositMxcToken(&_MxcL1Client.TransactOpts, amount)
}

// Init is a paid mutator transaction binding the contract method 0xc7a64b19.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeTarget, uint64 _initProofTimeIssued, uint16 _adjustmentQuotient) returns()
func (_MxcL1Client *MxcL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeTarget uint64, _initProofTimeIssued uint64, _adjustmentQuotient uint16) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeTarget, _initProofTimeIssued, _adjustmentQuotient)
}

// Init is a paid mutator transaction binding the contract method 0xc7a64b19.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeTarget, uint64 _initProofTimeIssued, uint16 _adjustmentQuotient) returns()
func (_MxcL1Client *MxcL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeTarget uint64, _initProofTimeIssued uint64, _adjustmentQuotient uint16) (*types.Transaction, error) {
	return _MxcL1Client.Contract.Init(&_MxcL1Client.TransactOpts, _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeTarget, _initProofTimeIssued, _adjustmentQuotient)
}

// Init is a paid mutator transaction binding the contract method 0xc7a64b19.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash, uint64 _initBlockFee, uint64 _initProofTimeTarget, uint64 _initProofTimeIssued, uint16 _adjustmentQuotient) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte, _initBlockFee uint64, _initProofTimeTarget uint64, _initProofTimeIssued uint64, _adjustmentQuotient uint16) (*types.Transaction, error) {
	return _MxcL1Client.Contract.Init(&_MxcL1Client.TransactOpts, _addressManager, _genesisBlockHash, _initBlockFee, _initProofTimeTarget, _initProofTimeIssued, _adjustmentQuotient)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x68bd2e7a.
//
// Solidity: function proposeBlock(bytes input, bytes txList, uint256 estimateGas) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[],uint256,uint256) meta)
func (_MxcL1Client *MxcL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, input []byte, txList []byte, estimateGas *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "proposeBlock", input, txList, estimateGas)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x68bd2e7a.
//
// Solidity: function proposeBlock(bytes input, bytes txList, uint256 estimateGas) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[],uint256,uint256) meta)
func (_MxcL1Client *MxcL1ClientSession) ProposeBlock(input []byte, txList []byte, estimateGas *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.ProposeBlock(&_MxcL1Client.TransactOpts, input, txList, estimateGas)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0x68bd2e7a.
//
// Solidity: function proposeBlock(bytes input, bytes txList, uint256 estimateGas) returns((uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[],uint256,uint256) meta)
func (_MxcL1Client *MxcL1ClientTransactorSession) ProposeBlock(input []byte, txList []byte, estimateGas *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.ProposeBlock(&_MxcL1Client.TransactOpts, input, txList, estimateGas)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_MxcL1Client *MxcL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "proveBlock", blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_MxcL1Client *MxcL1ClientSession) ProveBlock(blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _MxcL1Client.Contract.ProveBlock(&_MxcL1Client.TransactOpts, blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0xf3840f60.
//
// Solidity: function proveBlock(uint256 blockId, bytes input) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) ProveBlock(blockId *big.Int, input []byte) (*types.Transaction, error) {
	return _MxcL1Client.Contract.ProveBlock(&_MxcL1Client.TransactOpts, blockId, input)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcL1Client *MxcL1ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcL1Client *MxcL1ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _MxcL1Client.Contract.RenounceOwnership(&_MxcL1Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MxcL1Client.Contract.RenounceOwnership(&_MxcL1Client.TransactOpts)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcL1Client *MxcL1ClientTransactor) SetAddressManager(opts *bind.TransactOpts, newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "setAddressManager", newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcL1Client *MxcL1ClientSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcL1Client.Contract.SetAddressManager(&_MxcL1Client.TransactOpts, newAddressManager)
}

// SetAddressManager is a paid mutator transaction binding the contract method 0x0652b57a.
//
// Solidity: function setAddressManager(address newAddressManager) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) SetAddressManager(newAddressManager common.Address) (*types.Transaction, error) {
	return _MxcL1Client.Contract.SetAddressManager(&_MxcL1Client.TransactOpts, newAddressManager)
}

// SetProofParams is a paid mutator transaction binding the contract method 0xae46a347.
//
// Solidity: function setProofParams(uint64 newProofTimeTarget, uint64 newProofTimeIssued, uint64 newBlockFee, uint16 newAdjustmentQuotient) returns()
func (_MxcL1Client *MxcL1ClientTransactor) SetProofParams(opts *bind.TransactOpts, newProofTimeTarget uint64, newProofTimeIssued uint64, newBlockFee uint64, newAdjustmentQuotient uint16) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "setProofParams", newProofTimeTarget, newProofTimeIssued, newBlockFee, newAdjustmentQuotient)
}

// SetProofParams is a paid mutator transaction binding the contract method 0xae46a347.
//
// Solidity: function setProofParams(uint64 newProofTimeTarget, uint64 newProofTimeIssued, uint64 newBlockFee, uint16 newAdjustmentQuotient) returns()
func (_MxcL1Client *MxcL1ClientSession) SetProofParams(newProofTimeTarget uint64, newProofTimeIssued uint64, newBlockFee uint64, newAdjustmentQuotient uint16) (*types.Transaction, error) {
	return _MxcL1Client.Contract.SetProofParams(&_MxcL1Client.TransactOpts, newProofTimeTarget, newProofTimeIssued, newBlockFee, newAdjustmentQuotient)
}

// SetProofParams is a paid mutator transaction binding the contract method 0xae46a347.
//
// Solidity: function setProofParams(uint64 newProofTimeTarget, uint64 newProofTimeIssued, uint64 newBlockFee, uint16 newAdjustmentQuotient) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) SetProofParams(newProofTimeTarget uint64, newProofTimeIssued uint64, newBlockFee uint64, newAdjustmentQuotient uint16) (*types.Transaction, error) {
	return _MxcL1Client.Contract.SetProofParams(&_MxcL1Client.TransactOpts, newProofTimeTarget, newProofTimeIssued, newBlockFee, newAdjustmentQuotient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcL1Client *MxcL1ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcL1Client *MxcL1ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MxcL1Client.Contract.TransferOwnership(&_MxcL1Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MxcL1Client.Contract.TransferOwnership(&_MxcL1Client.TransactOpts, newOwner)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_MxcL1Client *MxcL1ClientTransactor) VerifyBlocks(opts *bind.TransactOpts, maxBlocks *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "verifyBlocks", maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_MxcL1Client *MxcL1ClientSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.VerifyBlocks(&_MxcL1Client.TransactOpts, maxBlocks)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x2fb5ae0a.
//
// Solidity: function verifyBlocks(uint256 maxBlocks) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) VerifyBlocks(maxBlocks *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.VerifyBlocks(&_MxcL1Client.TransactOpts, maxBlocks)
}

// WithdrawMxcToken is a paid mutator transaction binding the contract method 0x51e15268.
//
// Solidity: function withdrawMxcToken(uint256 amount) returns()
func (_MxcL1Client *MxcL1ClientTransactor) WithdrawMxcToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.contract.Transact(opts, "withdrawMxcToken", amount)
}

// WithdrawMxcToken is a paid mutator transaction binding the contract method 0x51e15268.
//
// Solidity: function withdrawMxcToken(uint256 amount) returns()
func (_MxcL1Client *MxcL1ClientSession) WithdrawMxcToken(amount *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.WithdrawMxcToken(&_MxcL1Client.TransactOpts, amount)
}

// WithdrawMxcToken is a paid mutator transaction binding the contract method 0x51e15268.
//
// Solidity: function withdrawMxcToken(uint256 amount) returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) WithdrawMxcToken(amount *big.Int) (*types.Transaction, error) {
	return _MxcL1Client.Contract.WithdrawMxcToken(&_MxcL1Client.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MxcL1Client *MxcL1ClientTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MxcL1Client.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MxcL1Client *MxcL1ClientSession) Receive() (*types.Transaction, error) {
	return _MxcL1Client.Contract.Receive(&_MxcL1Client.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MxcL1Client *MxcL1ClientTransactorSession) Receive() (*types.Transaction, error) {
	return _MxcL1Client.Contract.Receive(&_MxcL1Client.TransactOpts)
}

// MxcL1ClientAddressManagerChangedIterator is returned from FilterAddressManagerChanged and is used to iterate over the raw logs and unpacked data for AddressManagerChanged events raised by the MxcL1Client contract.
type MxcL1ClientAddressManagerChangedIterator struct {
	Event *MxcL1ClientAddressManagerChanged // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientAddressManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientAddressManagerChanged)
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
		it.Event = new(MxcL1ClientAddressManagerChanged)
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
func (it *MxcL1ClientAddressManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientAddressManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientAddressManagerChanged represents a AddressManagerChanged event raised by the MxcL1Client contract.
type MxcL1ClientAddressManagerChanged struct {
	AddressManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressManagerChanged is a free log retrieval operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcL1Client *MxcL1ClientFilterer) FilterAddressManagerChanged(opts *bind.FilterOpts) (*MxcL1ClientAddressManagerChangedIterator, error) {

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientAddressManagerChangedIterator{contract: _MxcL1Client.contract, event: "AddressManagerChanged", logs: logs, sub: sub}, nil
}

// WatchAddressManagerChanged is a free log subscription operation binding the contract event 0x399ded90cb5ed8d89ef7e76ff4af65c373f06d3bf5d7eef55f4228e7b702a18b.
//
// Solidity: event AddressManagerChanged(address addressManager)
func (_MxcL1Client *MxcL1ClientFilterer) WatchAddressManagerChanged(opts *bind.WatchOpts, sink chan<- *MxcL1ClientAddressManagerChanged) (event.Subscription, error) {

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "AddressManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientAddressManagerChanged)
				if err := _MxcL1Client.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
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
func (_MxcL1Client *MxcL1ClientFilterer) ParseAddressManagerChanged(log types.Log) (*MxcL1ClientAddressManagerChanged, error) {
	event := new(MxcL1ClientAddressManagerChanged)
	if err := _MxcL1Client.contract.UnpackLog(event, "AddressManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientBlockProposeRewardIterator is returned from FilterBlockProposeReward and is used to iterate over the raw logs and unpacked data for BlockProposeReward events raised by the MxcL1Client contract.
type MxcL1ClientBlockProposeRewardIterator struct {
	Event *MxcL1ClientBlockProposeReward // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientBlockProposeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientBlockProposeReward)
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
		it.Event = new(MxcL1ClientBlockProposeReward)
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
func (it *MxcL1ClientBlockProposeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientBlockProposeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientBlockProposeReward represents a BlockProposeReward event raised by the MxcL1Client contract.
type MxcL1ClientBlockProposeReward struct {
	Id       *big.Int
	Proposer common.Address
	Reward   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlockProposeReward is a free log retrieval operation binding the contract event 0x6c37877020064aabb3c00d8dff4cf6b32ec9dfc400d86b943af79c46492ca791.
//
// Solidity: event BlockProposeReward(uint256 indexed id, address proposer, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) FilterBlockProposeReward(opts *bind.FilterOpts, id []*big.Int) (*MxcL1ClientBlockProposeRewardIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "BlockProposeReward", idRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientBlockProposeRewardIterator{contract: _MxcL1Client.contract, event: "BlockProposeReward", logs: logs, sub: sub}, nil
}

// WatchBlockProposeReward is a free log subscription operation binding the contract event 0x6c37877020064aabb3c00d8dff4cf6b32ec9dfc400d86b943af79c46492ca791.
//
// Solidity: event BlockProposeReward(uint256 indexed id, address proposer, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) WatchBlockProposeReward(opts *bind.WatchOpts, sink chan<- *MxcL1ClientBlockProposeReward, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "BlockProposeReward", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientBlockProposeReward)
				if err := _MxcL1Client.contract.UnpackLog(event, "BlockProposeReward", log); err != nil {
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

// ParseBlockProposeReward is a log parse operation binding the contract event 0x6c37877020064aabb3c00d8dff4cf6b32ec9dfc400d86b943af79c46492ca791.
//
// Solidity: event BlockProposeReward(uint256 indexed id, address proposer, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) ParseBlockProposeReward(log types.Log) (*MxcL1ClientBlockProposeReward, error) {
	event := new(MxcL1ClientBlockProposeReward)
	if err := _MxcL1Client.contract.UnpackLog(event, "BlockProposeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientBlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the MxcL1Client contract.
type MxcL1ClientBlockProposedIterator struct {
	Event *MxcL1ClientBlockProposed // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientBlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientBlockProposed)
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
		it.Event = new(MxcL1ClientBlockProposed)
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
func (it *MxcL1ClientBlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientBlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientBlockProposed represents a BlockProposed event raised by the MxcL1Client contract.
type MxcL1ClientBlockProposed struct {
	Id       *big.Int
	Meta     MxcDataBlockMetadata
	BlockFee uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0xb250038c85825d1ba0b724295351e6c1fb3e629867d2510e4b6938546abfecf8.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[],uint256,uint256) meta, uint64 blockFee)
func (_MxcL1Client *MxcL1ClientFilterer) FilterBlockProposed(opts *bind.FilterOpts, id []*big.Int) (*MxcL1ClientBlockProposedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientBlockProposedIterator{contract: _MxcL1Client.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0xb250038c85825d1ba0b724295351e6c1fb3e629867d2510e4b6938546abfecf8.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[],uint256,uint256) meta, uint64 blockFee)
func (_MxcL1Client *MxcL1ClientFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *MxcL1ClientBlockProposed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "BlockProposed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientBlockProposed)
				if err := _MxcL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
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

// ParseBlockProposed is a log parse operation binding the contract event 0xb250038c85825d1ba0b724295351e6c1fb3e629867d2510e4b6938546abfecf8.
//
// Solidity: event BlockProposed(uint256 indexed id, (uint64,uint64,uint64,bytes32,bytes32,bytes32,uint24,uint24,uint32,address,address,(address,uint96,uint64)[],uint256,uint256) meta, uint64 blockFee)
func (_MxcL1Client *MxcL1ClientFilterer) ParseBlockProposed(log types.Log) (*MxcL1ClientBlockProposed, error) {
	event := new(MxcL1ClientBlockProposed)
	if err := _MxcL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientBlockProvenIterator is returned from FilterBlockProven and is used to iterate over the raw logs and unpacked data for BlockProven events raised by the MxcL1Client contract.
type MxcL1ClientBlockProvenIterator struct {
	Event *MxcL1ClientBlockProven // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientBlockProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientBlockProven)
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
		it.Event = new(MxcL1ClientBlockProven)
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
func (it *MxcL1ClientBlockProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientBlockProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientBlockProven represents a BlockProven event raised by the MxcL1Client contract.
type MxcL1ClientBlockProven struct {
	Id            *big.Int
	ParentHash    [32]byte
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Prover        common.Address
	ParentGasUsed uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockProven is a free log retrieval operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_MxcL1Client *MxcL1ClientFilterer) FilterBlockProven(opts *bind.FilterOpts, id []*big.Int) (*MxcL1ClientBlockProvenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientBlockProvenIterator{contract: _MxcL1Client.contract, event: "BlockProven", logs: logs, sub: sub}, nil
}

// WatchBlockProven is a free log subscription operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_MxcL1Client *MxcL1ClientFilterer) WatchBlockProven(opts *bind.WatchOpts, sink chan<- *MxcL1ClientBlockProven, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "BlockProven", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientBlockProven)
				if err := _MxcL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
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

// ParseBlockProven is a log parse operation binding the contract event 0x2295930c498c7b1f60143439a63dd1d24bbb730f08ff6ed383b490ba2c1cafa4.
//
// Solidity: event BlockProven(uint256 indexed id, bytes32 parentHash, bytes32 blockHash, bytes32 signalRoot, address prover, uint32 parentGasUsed)
func (_MxcL1Client *MxcL1ClientFilterer) ParseBlockProven(log types.Log) (*MxcL1ClientBlockProven, error) {
	event := new(MxcL1ClientBlockProven)
	if err := _MxcL1Client.contract.UnpackLog(event, "BlockProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientBlockProvenRewardIterator is returned from FilterBlockProvenReward and is used to iterate over the raw logs and unpacked data for BlockProvenReward events raised by the MxcL1Client contract.
type MxcL1ClientBlockProvenRewardIterator struct {
	Event *MxcL1ClientBlockProvenReward // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientBlockProvenRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientBlockProvenReward)
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
		it.Event = new(MxcL1ClientBlockProvenReward)
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
func (it *MxcL1ClientBlockProvenRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientBlockProvenRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientBlockProvenReward represents a BlockProvenReward event raised by the MxcL1Client contract.
type MxcL1ClientBlockProvenReward struct {
	Id     *big.Int
	Prover common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlockProvenReward is a free log retrieval operation binding the contract event 0x3f660dbecdc997b020e22e0c005183d9b72e405db24f96490dc1484344242989.
//
// Solidity: event BlockProvenReward(uint256 indexed id, address prover, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) FilterBlockProvenReward(opts *bind.FilterOpts, id []*big.Int) (*MxcL1ClientBlockProvenRewardIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "BlockProvenReward", idRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientBlockProvenRewardIterator{contract: _MxcL1Client.contract, event: "BlockProvenReward", logs: logs, sub: sub}, nil
}

// WatchBlockProvenReward is a free log subscription operation binding the contract event 0x3f660dbecdc997b020e22e0c005183d9b72e405db24f96490dc1484344242989.
//
// Solidity: event BlockProvenReward(uint256 indexed id, address prover, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) WatchBlockProvenReward(opts *bind.WatchOpts, sink chan<- *MxcL1ClientBlockProvenReward, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "BlockProvenReward", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientBlockProvenReward)
				if err := _MxcL1Client.contract.UnpackLog(event, "BlockProvenReward", log); err != nil {
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

// ParseBlockProvenReward is a log parse operation binding the contract event 0x3f660dbecdc997b020e22e0c005183d9b72e405db24f96490dc1484344242989.
//
// Solidity: event BlockProvenReward(uint256 indexed id, address prover, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) ParseBlockProvenReward(log types.Log) (*MxcL1ClientBlockProvenReward, error) {
	event := new(MxcL1ClientBlockProvenReward)
	if err := _MxcL1Client.contract.UnpackLog(event, "BlockProvenReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientBlockVerifiedIterator is returned from FilterBlockVerified and is used to iterate over the raw logs and unpacked data for BlockVerified events raised by the MxcL1Client contract.
type MxcL1ClientBlockVerifiedIterator struct {
	Event *MxcL1ClientBlockVerified // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientBlockVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientBlockVerified)
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
		it.Event = new(MxcL1ClientBlockVerified)
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
func (it *MxcL1ClientBlockVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientBlockVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientBlockVerified represents a BlockVerified event raised by the MxcL1Client contract.
type MxcL1ClientBlockVerified struct {
	Id        *big.Int
	BlockHash [32]byte
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0x7e98ef90898cf74532e8f918c3b89c5ce86c0b10a0d9bf3d0526af7fa7b099da.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) FilterBlockVerified(opts *bind.FilterOpts, id []*big.Int) (*MxcL1ClientBlockVerifiedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientBlockVerifiedIterator{contract: _MxcL1Client.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0x7e98ef90898cf74532e8f918c3b89c5ce86c0b10a0d9bf3d0526af7fa7b099da.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *MxcL1ClientBlockVerified, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "BlockVerified", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientBlockVerified)
				if err := _MxcL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
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

// ParseBlockVerified is a log parse operation binding the contract event 0x7e98ef90898cf74532e8f918c3b89c5ce86c0b10a0d9bf3d0526af7fa7b099da.
//
// Solidity: event BlockVerified(uint256 indexed id, bytes32 blockHash, uint256 reward)
func (_MxcL1Client *MxcL1ClientFilterer) ParseBlockVerified(log types.Log) (*MxcL1ClientBlockVerified, error) {
	event := new(MxcL1ClientBlockVerified)
	if err := _MxcL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientCrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the MxcL1Client contract.
type MxcL1ClientCrossChainSyncedIterator struct {
	Event *MxcL1ClientCrossChainSynced // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientCrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientCrossChainSynced)
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
		it.Event = new(MxcL1ClientCrossChainSynced)
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
func (it *MxcL1ClientCrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientCrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientCrossChainSynced represents a CrossChainSynced event raised by the MxcL1Client contract.
type MxcL1ClientCrossChainSynced struct {
	SrcHeight  *big.Int
	BlockHash  [32]byte
	SignalRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_MxcL1Client *MxcL1ClientFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, srcHeight []*big.Int) (*MxcL1ClientCrossChainSyncedIterator, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientCrossChainSyncedIterator{contract: _MxcL1Client.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_MxcL1Client *MxcL1ClientFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *MxcL1ClientCrossChainSynced, srcHeight []*big.Int) (event.Subscription, error) {

	var srcHeightRule []interface{}
	for _, srcHeightItem := range srcHeight {
		srcHeightRule = append(srcHeightRule, srcHeightItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "CrossChainSynced", srcHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientCrossChainSynced)
				if err := _MxcL1Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
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

// ParseCrossChainSynced is a log parse operation binding the contract event 0x7528bbd1cef0e5d13408706892a51ee8ef82bbf33d4ec0c37216f8beba71205b.
//
// Solidity: event CrossChainSynced(uint256 indexed srcHeight, bytes32 blockHash, bytes32 signalRoot)
func (_MxcL1Client *MxcL1ClientFilterer) ParseCrossChainSynced(log types.Log) (*MxcL1ClientCrossChainSynced, error) {
	event := new(MxcL1ClientCrossChainSynced)
	if err := _MxcL1Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientEthDepositedIterator is returned from FilterEthDeposited and is used to iterate over the raw logs and unpacked data for EthDeposited events raised by the MxcL1Client contract.
type MxcL1ClientEthDepositedIterator struct {
	Event *MxcL1ClientEthDeposited // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientEthDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientEthDeposited)
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
		it.Event = new(MxcL1ClientEthDeposited)
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
func (it *MxcL1ClientEthDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientEthDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientEthDeposited represents a EthDeposited event raised by the MxcL1Client contract.
type MxcL1ClientEthDeposited struct {
	Deposit MxcDataEthDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEthDeposited is a free log retrieval operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_MxcL1Client *MxcL1ClientFilterer) FilterEthDeposited(opts *bind.FilterOpts) (*MxcL1ClientEthDepositedIterator, error) {

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientEthDepositedIterator{contract: _MxcL1Client.contract, event: "EthDeposited", logs: logs, sub: sub}, nil
}

// WatchEthDeposited is a free log subscription operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_MxcL1Client *MxcL1ClientFilterer) WatchEthDeposited(opts *bind.WatchOpts, sink chan<- *MxcL1ClientEthDeposited) (event.Subscription, error) {

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientEthDeposited)
				if err := _MxcL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
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

// ParseEthDeposited is a log parse operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_MxcL1Client *MxcL1ClientFilterer) ParseEthDeposited(log types.Log) (*MxcL1ClientEthDeposited, error) {
	event := new(MxcL1ClientEthDeposited)
	if err := _MxcL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MxcL1Client contract.
type MxcL1ClientInitializedIterator struct {
	Event *MxcL1ClientInitialized // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientInitialized)
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
		it.Event = new(MxcL1ClientInitialized)
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
func (it *MxcL1ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientInitialized represents a Initialized event raised by the MxcL1Client contract.
type MxcL1ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcL1Client *MxcL1ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*MxcL1ClientInitializedIterator, error) {

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientInitializedIterator{contract: _MxcL1Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MxcL1Client *MxcL1ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MxcL1ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientInitialized)
				if err := _MxcL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MxcL1Client *MxcL1ClientFilterer) ParseInitialized(log types.Log) (*MxcL1ClientInitialized, error) {
	event := new(MxcL1ClientInitialized)
	if err := _MxcL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MxcL1Client contract.
type MxcL1ClientOwnershipTransferredIterator struct {
	Event *MxcL1ClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientOwnershipTransferred)
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
		it.Event = new(MxcL1ClientOwnershipTransferred)
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
func (it *MxcL1ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientOwnershipTransferred represents a OwnershipTransferred event raised by the MxcL1Client contract.
type MxcL1ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MxcL1Client *MxcL1ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MxcL1ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientOwnershipTransferredIterator{contract: _MxcL1Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MxcL1Client *MxcL1ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MxcL1ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientOwnershipTransferred)
				if err := _MxcL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MxcL1Client *MxcL1ClientFilterer) ParseOwnershipTransferred(log types.Log) (*MxcL1ClientOwnershipTransferred, error) {
	event := new(MxcL1ClientOwnershipTransferred)
	if err := _MxcL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MxcL1ClientProofParamsChangedIterator is returned from FilterProofParamsChanged and is used to iterate over the raw logs and unpacked data for ProofParamsChanged events raised by the MxcL1Client contract.
type MxcL1ClientProofParamsChangedIterator struct {
	Event *MxcL1ClientProofParamsChanged // Event containing the contract specifics and raw log

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
func (it *MxcL1ClientProofParamsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MxcL1ClientProofParamsChanged)
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
		it.Event = new(MxcL1ClientProofParamsChanged)
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
func (it *MxcL1ClientProofParamsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MxcL1ClientProofParamsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MxcL1ClientProofParamsChanged represents a ProofParamsChanged event raised by the MxcL1Client contract.
type MxcL1ClientProofParamsChanged struct {
	ProofTimeTarget    uint64
	ProofTimeIssued    uint64
	BlockFee           uint64
	AdjustmentQuotient uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProofParamsChanged is a free log retrieval operation binding the contract event 0x565e5aa69c99d81e441dd3bb8535d888585683743f3c6a3bf49e5e1b227bd8f9.
//
// Solidity: event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
func (_MxcL1Client *MxcL1ClientFilterer) FilterProofParamsChanged(opts *bind.FilterOpts) (*MxcL1ClientProofParamsChangedIterator, error) {

	logs, sub, err := _MxcL1Client.contract.FilterLogs(opts, "ProofParamsChanged")
	if err != nil {
		return nil, err
	}
	return &MxcL1ClientProofParamsChangedIterator{contract: _MxcL1Client.contract, event: "ProofParamsChanged", logs: logs, sub: sub}, nil
}

// WatchProofParamsChanged is a free log subscription operation binding the contract event 0x565e5aa69c99d81e441dd3bb8535d888585683743f3c6a3bf49e5e1b227bd8f9.
//
// Solidity: event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
func (_MxcL1Client *MxcL1ClientFilterer) WatchProofParamsChanged(opts *bind.WatchOpts, sink chan<- *MxcL1ClientProofParamsChanged) (event.Subscription, error) {

	logs, sub, err := _MxcL1Client.contract.WatchLogs(opts, "ProofParamsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MxcL1ClientProofParamsChanged)
				if err := _MxcL1Client.contract.UnpackLog(event, "ProofParamsChanged", log); err != nil {
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

// ParseProofParamsChanged is a log parse operation binding the contract event 0x565e5aa69c99d81e441dd3bb8535d888585683743f3c6a3bf49e5e1b227bd8f9.
//
// Solidity: event ProofParamsChanged(uint64 proofTimeTarget, uint64 proofTimeIssued, uint64 blockFee, uint16 adjustmentQuotient)
func (_MxcL1Client *MxcL1ClientFilterer) ParseProofParamsChanged(log types.Log) (*MxcL1ClientProofParamsChanged, error) {
	event := new(MxcL1ClientProofParamsChanged)
	if err := _MxcL1Client.contract.UnpackLog(event, "ProofParamsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
