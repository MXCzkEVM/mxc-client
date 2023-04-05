package submitter

import (
	"context"

	"github.com/MXCzkEVM/mxc-client/bindings"
	proofProducer "github.com/MXCzkEVM/mxc-client/prover/proof_producer"
)

type ProofSubmitter interface {
	RequestProof(ctx context.Context, event *bindings.MXCL1ClientBlockProposed) error
	SubmitProof(ctx context.Context, proofWithHeader *proofProducer.ProofWithHeader) error
}
