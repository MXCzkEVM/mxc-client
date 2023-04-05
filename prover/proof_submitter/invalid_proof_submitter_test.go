package submitter

import (
	"context"
	"time"

	"github.com/MXCzkEVM/mxc-client/bindings"
	proofProducer "github.com/MXCzkEVM/mxc-client/prover/proof_producer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *ProofSubmitterTestSuite) TestProveBlockInvalidL1OriginTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s.ErrorContains(
		s.invalidProofSubmitter.RequestProof(ctx, &bindings.MXCL1ClientBlockProposed{Id: common.Big256}),
		"context deadline exceeded",
	)
}

func (s *ProofSubmitterTestSuite) TestSubmitInvalidBlockProofThrowawayBlockNotFound() {
	s.Error(
		s.invalidProofSubmitter.SubmitProof(
			context.Background(), &proofProducer.ProofWithHeader{
				BlockID: common.Big256,
				Meta:    &bindings.MXCDataBlockMetadata{},
				Header:  &types.Header{},
				ZkProof: []byte{0xff},
			},
		),
	)
}
