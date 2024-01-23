package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"net/http"
)

// ABI arguments marshaling components.
var (
	blockMetadataInputComponents = []abi.ArgumentMarshaling{
		{
			Name: "txListHash",
			Type: "bytes32",
		},
		{
			Name: "beneficiary",
			Type: "address",
		},
		{
			Name: "gasLimit",
			Type: "uint32",
		},
		{
			Name: "txListByteStart",
			Type: "uint24",
		},
		{
			Name: "txListByteEnd",
			Type: "uint24",
		},
		{
			Name: "cacheTxListInfo",
			Type: "uint8",
		},
	}
	blockMetadataComponents = []abi.ArgumentMarshaling{
		{
			Name: "id",
			Type: "uint64",
		},
		{
			Name: "timestamp",
			Type: "uint64",
		},
		{
			Name: "l1Height",
			Type: "uint64",
		},
		{
			Name: "l1Hash",
			Type: "bytes32",
		},
		{
			Name: "mixHash",
			Type: "bytes32",
		},
		{
			Name: "txListHash",
			Type: "bytes32",
		},
		{
			Name: "txListByteStart",
			Type: "uint24",
		},
		{
			Name: "txListByteEnd",
			Type: "uint24",
		},
		{
			Name: "gasLimit",
			Type: "uint32",
		},
		{
			Name: "beneficiary",
			Type: "address",
		},
		{
			Name: "treasury",
			Type: "address",
		},
		{
			Name: "depositsProcessed",
			Type: "tuple[]",
			Components: []abi.ArgumentMarshaling{
				{
					Name: "recipient",
					Type: "address",
				},
				{
					Name: "amount",
					Type: "uint96",
				},
				{
					Name: "id",
					Type: "uint64",
				},
			},
		},
	}
	evidenceComponents = []abi.ArgumentMarshaling{
		{
			Name: "metaHash",
			Type: "bytes32",
		},
		{
			Name: "parentHash",
			Type: "bytes32",
		},
		{
			Name: "blockHash",
			Type: "bytes32",
		},
		{
			Name: "signalRoot",
			Type: "bytes32",
		},
		{
			Name: "graffiti",
			Type: "bytes32",
		},
		{
			Name: "prover",
			Type: "address",
		},
		{
			Name: "parentGasUsed",
			Type: "uint32",
		},
		{
			Name: "gasUsed",
			Type: "uint32",
		},
		{
			Name: "verifierId",
			Type: "uint16",
		},
		{
			Name: "proof",
			Type: "bytes",
		},
	}
)

var (
	// BlockMetadataInput
	blockMetadataInputType, _ = abi.NewType("tuple", "TaikoData.BlockMetadataInput", blockMetadataInputComponents)
	blockMetadataInputArgs    = abi.Arguments{{Name: "BlockMetadataInput", Type: blockMetadataInputType}}
	// BlockMetadata
	blockMetadataType, _ = abi.NewType("tuple", "TaikoData.BlockMetadata", blockMetadataComponents)
	blockMetadataArgs    = abi.Arguments{{Name: "BlockMetadata", Type: blockMetadataType}}
	// Evidence
	EvidenceType, _ = abi.NewType("tuple", "TaikoData.BlockEvidence", evidenceComponents)
	EvidenceArgs    = abi.Arguments{{Name: "Evidence", Type: EvidenceType}}
)

var MxcL1V2ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_DEPOSIT_REQUIREMENT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ELECTION_INVALID_PROPOSER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ELECTION_SPEED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ELECTION_SPEED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_UNLOCK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockProposeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structMxcData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockProvenReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structMxcData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"ProofParamsChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_proposedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumProposedBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ringBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txListCacheExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"systemProofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"realProofSkipSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"minEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"}],\"internalType\":\"structMxcData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"}],\"internalType\":\"structMxcData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMxcTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proofTime\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getVerifierName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_initBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"estimateGas\",\"type\":\"uint256\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"}],\"internalType\":\"structMxcData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newProofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newProofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"newAdjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"setProofParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStakeMxcTokenBalances\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"prevBaseFee\",\"type\":\"uint48\"},{\"internalType\":\"uint64\",\"name\":\"proposerElectionTimeoutOffset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proveMetaReward\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"}

var MxcL1V1ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ALREADY_PROVEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_BLOCK_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_DEPOSIT_REQUIREMENT\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"L1_EVIDENCE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_FORK_CHOICE_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INSUFFICIENT_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_CONFIG\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_EVIDENCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_METADATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_INVALID_PROOF_OVERWRITE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_SPECIAL_PROVER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_NOT_UNLOCK\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_ORACLE_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SAME_PROOF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_DISABLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_SYSTEM_PROVER_PROHIBITED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TOO_MANY_BLOCKS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_NOT_EXIST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1_TX_LIST_RANGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_DENIED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RESOLVER_INVALID_ADDR\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"RESOLVER_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addressManager\",\"type\":\"address\"}],\"name\":\"AddressManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockProposeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structMxcData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"}],\"name\":\"BlockProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"BlockProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockProvenReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"BlockVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"}],\"name\":\"CrossChainSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structMxcData.EthDeposit\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"EthDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"ProofParamsChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEtherToL2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_metaHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_proposedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumProposedBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ringBufferSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVerificationsPerTx\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockMaxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTransactionsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBytesPerTxList\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"txListCacheExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"systemProofCooldownPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"realProofSkipSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxEthDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"minEthDepositAmount\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"relaySignalRoot\",\"type\":\"bool\"}],\"internalType\":\"structMxcData.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"getCrossChainSignalRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"parentGasUsed\",\"type\":\"uint32\"}],\"name\":\"getForkChoice\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"provenAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasUsed\",\"type\":\"uint32\"}],\"internalType\":\"structMxcData.ForkChoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMxcTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proofTime\",\"type\":\"uint64\"}],\"name\":\"getProofReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateVariables\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numEthDeposits\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.StateVariables\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getVerifierName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_initBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_initProofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_adjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"txList\",\"type\":\"bytes\"}],\"name\":\"proposeBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"l1Height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"l1Hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txListHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"txListByteStart\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"txListByteEnd\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"internalType\":\"structMxcData.EthDeposit[]\",\"name\":\"depositsProcessed\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"}],\"internalType\":\"structMxcData.BlockMetadata\",\"name\":\"meta\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"proveBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"allowZeroAddress\",\"type\":\"bool\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddressManager\",\"type\":\"address\"}],\"name\":\"setAddressManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newProofTimeTarget\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newProofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newBlockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"newAdjustmentQuotient\",\"type\":\"uint16\"}],\"name\":\"setProofParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"genesisHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"genesisTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"adjustmentQuotient\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"prevBaseFee\",\"type\":\"uint48\"},{\"internalType\":\"uint64\",\"name\":\"__reserved72\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accProposedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"accBlockFees\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeIssued\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proofTimeTarget\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBlocks\",\"type\":\"uint256\"}],\"name\":\"verifyBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMxcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Contract ABIs.
var (
	MxcL1ABIV3  *abi.ABI
	MxcL1ABIV2  *abi.ABI
	MxcL1ABIV1  *abi.ABI
	MxcL2ABI    *abi.ABI
	MxcTokenABI *abi.ABI
	LPWANABI    *abi.ABI
)

func init() {
	var err error

	if MxcL1ABIV3, err = bindings.MxcL1ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get MxcL1V3 ABI error", "error", err)
	}

	if MxcL1ABIV2, err = MxcL1V2ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get MxcL1V2 ABI error", "error", err)
	}

	if MxcL1ABIV1, err = MxcL1V1ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get MxcL1V1 ABI error", "error", err)
	}

	if MxcL2ABI, err = bindings.MxcL2ClientMetaData.GetAbi(); err != nil {
		log.Crit("Get MxcL2 ABI error", "error", err)
	}

	if MxcTokenABI, err = bindings.MxcTokenClientMetaData.GetAbi(); err != nil {
		log.Crit("Get MxcToken ABI error", "error", err)
	}

	if LPWANABI, err = bindings.LPWANClientMetaData.GetAbi(); err != nil {
		log.Crit("Get LPWAN ABI error", "error", err)
	}
}

// EncodeBlockMetadataInput performs the solidity `abi.encode` for the given blockMetadataInput.
func EncodeBlockMetadataInput(meta *MxcL1BlockMetadataInput) ([]byte, error) {
	b, err := blockMetadataInputArgs.Pack(meta)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode block metadata input, %w", err)
	}
	return b, nil
}

// EncodeBlockMetadata performs the solidity `abi.encode` for the given blockMetadata.
func EncodeBlockMetadata(meta *bindings.MxcDataBlockMetadata) ([]byte, error) {
	b, err := blockMetadataArgs.Pack(meta)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode block metadata, %w", err)
	}
	return b, nil
}

// EncodeEvidence performs the solidity `abi.encode` for the given evidence.
func EncodeEvidence(e *MxcL1Evidence) ([]byte, error) {
	b, err := EvidenceArgs.Pack(e)
	if err != nil {
		return nil, fmt.Errorf("failed to abi.encode evidence, %w", err)
	}
	return b, nil
}

// EncodeCommitHash performs the solidity `abi.encodePacked` for the given
// commitHash components.
func EncodeCommitHash(beneficiary common.Address, txListHash [32]byte) []byte {
	// keccak256(abi.encodePacked(beneficiary, txListHash));
	return crypto.Keccak256(
		bytes.Join([][]byte{beneficiary.Bytes(), txListHash[:]}, nil),
	)
}

// EncodeProposeBlockInput encodes the input params for MxcL1.proposeBlock.
func EncodeProposeBlockInput(metadataInput *MxcL1BlockMetadataInput) ([]byte, error) {
	metaBytes, err := EncodeBlockMetadataInput(metadataInput)
	if err != nil {
		return nil, err
	}
	return metaBytes, nil
}

// EncodeProveBlockInput encodes the input params for MxcL1.proveBlock.
func EncodeProveBlockInput(
	evidence *MxcL1Evidence,
) ([]byte, error) {
	evidenceBytes, err := EncodeEvidence(evidence)
	if err != nil {
		return nil, err
	}

	return evidenceBytes, nil
}

// EncodeProveBlockInvalidInput encodes the input params for MxcL1.proveBlockInvalid.
func EncodeProveBlockInvalidInput(
	evidence *MxcL1Evidence,
	target *bindings.MxcDataBlockMetadata,
	receipt *types.Receipt,
) ([][]byte, error) {
	evidenceBytes, err := EncodeEvidence(evidence)
	if err != nil {
		return nil, err
	}

	metaBytes, err := EncodeBlockMetadata(target)
	if err != nil {
		return nil, err
	}

	receiptBytes, err := rlp.EncodeToBytes(receipt)
	if err != nil {
		return nil, err
	}

	return [][]byte{evidenceBytes, metaBytes, receiptBytes}, nil
}

// UnpackTxListBytes unpacks the input data of a MxcL1.proposeBlock transaction, and returns the txList bytes.
func UnpackTxListBytes(txData []byte, ipfsGateways ...string) ([]byte, error) {
	isV3 := true
	method, err := MxcL1ABIV3.MethodById(txData)
	if err != nil {
		isV3 = false
		method, err = MxcL1ABIV2.MethodById(txData)
		if err != nil {
			method, err = MxcL1ABIV1.MethodById(txData)
			if err != nil {
				return nil, err
			}
			return nil, err
		}
		return nil, err
	}

	// Only check for safety.
	if method.Name != "proposeBlock" {
		return nil, fmt.Errorf("invalid method name: %s", method.Name)
	}

	args := map[string]interface{}{}

	if err := method.Inputs.UnpackIntoMap(args, txData[4:]); err != nil {
		return nil, err
	}

	inputs, ok := args["txList"].([]byte)

	if !ok {
		return nil, errors.New("failed to get txList bytes")
	}

	if isV3 {
		if len(ipfsGateways) == 0 {
			return nil, errors.New("download ipfs failed cause empty ipfs gateways endpoint")
		}

		var lastErr error
		for _, gateway := range ipfsGateways {
			inputs = (func() []byte {
				resp, err := http.NewRequest("GET", fmt.Sprintf("%s%s", gateway, string(inputs)), nil)
				if err != nil {
					lastErr = fmt.Errorf("failed to download ipfs blob, cid: %v, gateway: %v, err: %v", string(inputs), gateway, err)
					return nil
				}
				defer resp.Body.Close()

				data, err := io.ReadAll(resp.Body)
				if err != nil {
					lastErr = fmt.Errorf("failed to parse downloaded ipfs blob, gateway: %v, err: %v", gateway, err)
					return nil
				}
				return common.FromHex(string(data))
			})()
		}

		if lastErr != nil {
			return nil, lastErr
		}

		return nil, errors.New("all IPFS gateways failed")
	}

	return inputs, nil
}
