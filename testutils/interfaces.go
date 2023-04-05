package testutils

import (
	"context"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/MXCzkEVM/mxc-client/cmd/utils"
	"github.com/ethereum/go-ethereum/core/types"
)

type CalldataSyncer interface {
	ProcessL1Blocks(ctx context.Context, l1End *types.Header) error
}

type Proposer interface {
	utils.SubcommandApplication
	ProposeOp(ctx context.Context) error
	ProposeEmptyBlockOp(ctx context.Context) error
	CommitTxList(ctx context.Context, txListBytes []byte, gasLimit uint64, splittedIdx int) (
		*bindings.MXCDataBlockMetadata,
		*types.Transaction,
		error,
	)
	ProposeTxList(
		ctx context.Context,
		meta *bindings.MXCDataBlockMetadata,
		commitTx *types.Transaction,
		txListBytes []byte,
		txNum uint,
	) error
}
