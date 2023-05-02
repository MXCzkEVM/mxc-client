package submitter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/MXCzkEVM/mxc-client/bindings/encoding"
	"github.com/MXCzkEVM/mxc-client/metrics"
	"github.com/MXCzkEVM/mxc-client/pkg/rpc"
	anchorTxValidator "github.com/MXCzkEVM/mxc-client/prover/anchor_tx_validator"
	proofProducer "github.com/MXCzkEVM/mxc-client/prover/proof_producer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

var _ ProofSubmitter = (*ValidProofSubmitter)(nil)

// ValidProofSubmitter is responsible requesting zk proofs for the given valid L2
// blocks, and submitting the generated proofs to the MXCL1 smart contract.
type ValidProofSubmitter struct {
	rpc               *rpc.Client
	proofProducer     proofProducer.ProofProducer
	reusltCh          chan *proofProducer.ProofWithHeader
	anchorTxValidator *anchorTxValidator.AnchorTxValidator
	proverPrivKey     *ecdsa.PrivateKey
	proverAddress     common.Address
	submittorPrivKey  *ecdsa.PrivateKey
	mutex             *sync.Mutex
}

// NewValidProofSubmitter creates a new ValidProofSubmitter instance.
func NewValidProofSubmitter(
	rpc *rpc.Client,
	proofProducer proofProducer.ProofProducer,
	reusltCh chan *proofProducer.ProofWithHeader,
	mxcL2Address common.Address,
	proverPrivKey *ecdsa.PrivateKey,
	submittorPrivKey *ecdsa.PrivateKey,
	mutex *sync.Mutex,
) *ValidProofSubmitter {
	return &ValidProofSubmitter{
		rpc:               rpc,
		proofProducer:     proofProducer,
		reusltCh:          reusltCh,
		anchorTxValidator: anchorTxValidator.New(mxcL2Address, rpc.L2ChainID, rpc),
		proverPrivKey:     proverPrivKey,
		proverAddress:     crypto.PubkeyToAddress(proverPrivKey.PublicKey),
		submittorPrivKey:  submittorPrivKey,
		mutex:             mutex,
	}
}

// RequestProof implements the ProofSubmitter interface.
func (s *ValidProofSubmitter) RequestProof(ctx context.Context, event *bindings.MXCL1ClientBlockProposed) error {
	l1Origin, err := s.rpc.WaitL1Origin(ctx, event.Id)
	if err != nil {
		return fmt.Errorf("failed to fetch l1Origin, blockID: %d, err: %w", event.Id, err)
	}

	// This should not be reached, only check for safety.
	if l1Origin.Throwaway {
		log.Error("Get a block metadata with invalid transaction list", "l1Origin", l1Origin)
		return nil
	}

	// Get the header of the block to prove from L2 execution engine.
	header, err := s.rpc.L2.HeaderByHash(ctx, l1Origin.L2BlockHash)
	if err != nil {
		return err
	}

	// Request proof.
	opts := &proofProducer.ProofRequestOptions{
		Height:             header.Number,
		ProverAddress:      s.proverAddress,
		ProposeBlockTxHash: event.Raw.TxHash,
	}

	if err := s.proofProducer.RequestProof(ctx, opts, event.Id, &event.Meta, header, s.reusltCh); err != nil {
		return err
	}

	metrics.ProverQueuedProofCounter.Inc(1)
	metrics.ProverQueuedValidProofCounter.Inc(1)

	return nil
}

// SubmitProof implements the ProofSubmitter interface.
func (s *ValidProofSubmitter) SubmitProof(
	ctx context.Context,
	proofWithHeader *proofProducer.ProofWithHeader,
) (err error) {
	log.Info(
		"New valid block proof",
		"blockID", proofWithHeader.BlockID,
		"beneficiary", proofWithHeader.Meta.Beneficiary,
		"hash", proofWithHeader.Header.Hash(),
		"proof", common.Bytes2Hex(proofWithHeader.ZkProof),
	)
	var (
		blockID = proofWithHeader.BlockID
		header  = proofWithHeader.Header
		zkProof = proofWithHeader.ZkProof
	)

	metrics.ProverReceivedProofCounter.Inc(1)
	metrics.ProverReceivedValidProofCounter.Inc(1)

	// Get the corresponding L2 block.
	block, err := s.rpc.L2.BlockByHash(ctx, header.Hash())
	if err != nil {
		return fmt.Errorf("failed to get L2 block with given hash %s: %w", header.Hash(), err)
	}

	log.Debug(
		"Get the L2 block to prove",
		"blockID", blockID,
		"hash", block.Hash(),
		"root", header.Root.String(),
		"transactions", len(block.Transactions()),
	)

	if block.Transactions().Len() == 0 {
		return fmt.Errorf("invalid block without anchor transaction, blockID %s", blockID)
	}

	// Validate MXCL2.anchor transaction inside the L2 block.
	anchorTx := block.Transactions()[0]
	if err := s.anchorTxValidator.ValidateAnchorTx(ctx, anchorTx); err != nil {
		return fmt.Errorf("invalid anchor transaction: %w", err)
	}

	// Get and validate this anchor transaction's receipt.
	anchorTxReceipt, err := s.anchorTxValidator.GetAndValidateAnchorTxReceipt(ctx, anchorTx)
	if err != nil {
		return fmt.Errorf("failed to fetch anchor transaction receipt: %w", err)
	}

	// Generate the merkel proof (whose root is block's txRoot) of this anchor transaction.
	txRoot, anchorTxProof, err := generateTrieProof(block.Transactions(), 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor transaction proof: %w", err)
	}

	// Generate the merkel proof (whose root is block's receiptRoot) of this anchor transaction's receipt.
	receipts, err := rpc.GetReceiptsByBlock(ctx, s.rpc.L2RawRPC, block)
	if err != nil {
		return fmt.Errorf("failed to fetch block receipts: %w", err)
	}
	receiptRoot, anchorReceiptProof, err := generateTrieProof(receipts, 0)
	if err != nil {
		return fmt.Errorf("failed to generate anchor receipt proof: %w", err)
	}

	// Double check the calculated roots.
	if txRoot != block.TxHash() || receiptRoot != block.ReceiptHash() {
		return fmt.Errorf(
			"txHash or receiptHash mismatch, txRoot: %s, header.TxHash: %s, receiptRoot: %s, header.ReceiptHash: %s",
			txRoot, header.TxHash, receiptRoot, header.ReceiptHash,
		)
	}

	circuitsIdx, err := proofProducer.DegreeToCircuitsIdx(proofWithHeader.Degree)
	if err != nil {
		return err
	}

	evidence := &encoding.MXCL1Evidence{
		Meta:     *proofWithHeader.Meta,
		Header:   *encoding.FromGethHeader(header),
		Prover:   s.proverAddress,
		Proofs:   [][]byte{zkProof, anchorTxProof, anchorReceiptProof},
		Circuits: circuitsIdx,
	}

	input, err := encoding.EncodeProveBlockInput(evidence, anchorTx, anchorTxReceipt)
	if err != nil {
		return fmt.Errorf("failed to encode MXCL1.proveBlock inputs: %w", err)
	}

	// Send the MXCL1.proveBlock transaction.
	txOpts, err := getProveBlocksTxOpts(ctx, s.rpc.L1, s.rpc.L1ChainID, s.submittorPrivKey)
	if err != nil {
		return err
	}
	sendTx := func() (tx *types.Transaction, err error) {
		need, err := s.NeedNewProof(ctx, blockID)
		if err != nil {
			return nil, err
		}
		if !need {
			return nil, nil
		}
		done := make(chan bool, 1)
		go func() {
			s.mutex.Lock()
			tx, err = s.rpc.MXCL1.ProveBlock(txOpts, blockID, input)
		}()
		select {
		case <-done:
		case <-time.After(time.Second * 5):
		}
		s.mutex.Unlock()
		return tx, err
	}

	if err := sendTxWithBackoff(ctx, s.rpc, blockID, sendTx); err != nil {
		log.Error("Failed to send MXCL1.proveBlock transaction", "err", err)
		if errors.Is(err, errUnretryable) {
			return nil
		}

		return err
	}

	log.Info(
		"💰 Your block proof was accepted",
		"blockID", proofWithHeader.BlockID,
		"hash", block.Hash(), "height", block.Number(),
		"transactions", block.Transactions().Len(),
	)

	metrics.ProverSentProofCounter.Inc(1)
	metrics.ProverSentValidProofCounter.Inc(1)
	metrics.ProverLatestProvenBlockIDGauge.Update(proofWithHeader.BlockID.Int64())

	return nil
}

func (s *ValidProofSubmitter) isBlockVerified(id *big.Int) (bool, error) {
	stateVars, err := s.rpc.GetProtocolStateVariables(nil)
	if err != nil {
		return false, err
	}

	return id.Uint64() <= stateVars.LatestVerifiedId, nil
}

func (s *ValidProofSubmitter) NeedNewProof(ctx context.Context, id *big.Int) (bool, error) {
	var parentHash common.Hash
	if id.Cmp(common.Big1) == 0 {
		header, err := s.rpc.L2.HeaderByNumber(ctx, common.Big0)
		if err != nil {
			return false, err
		}

		parentHash = header.Hash()
	} else {
		parentL1Origin, err := s.rpc.WaitL1Origin(ctx, new(big.Int).Sub(id, common.Big1))
		if err != nil {
			return false, err
		}

		parentHash = parentL1Origin.L2BlockHash
	}

	fc, err := s.rpc.MXCL1.GetForkChoice(nil, id, parentHash)
	if err != nil {
		return false, err
	}

	if fc.Prover != (common.Address{}) {
		log.Info("📬 Block's proof has already been submitted", "blockID", id, "prover", fc.Prover)
		return false, nil
	}

	return true, nil
}
