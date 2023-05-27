package submitter

import (
	"context"
	"math/big"

	"github.com/MXCzkEVM/mxc-client/bindings"
	proofProducer "github.com/MXCzkEVM/mxc-client/prover/proof_producer"
)

type ProofSubmitter interface {
	RequestProof(ctx context.Context, event *bindings.MxcL1ClientBlockProposed) error
	SubmitProof(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) error
	CancelProof(ctx context.Context, blockID *big.Int) error
}
