package testutils

import (
	"context"

	"github.com/MXCzkEVM/mxc-client/bindings/encoding"
	"github.com/MXCzkEVM/mxc-client/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CalldataSyncer interface {
	ProcessL1Blocks(ctx context.Context, l1End *types.Header) error
}

type Proposer interface {
	utils.SubcommandApplication
	ProposeOp(ctx context.Context) error
	ProposeEmptyBlockOp(ctx context.Context) error
	L2SuggestedFeeRecipient() common.Address
	ProposeTxList(
		ctx context.Context,
		meta *encoding.MxcL1BlockMetadataInput,
		txListBytes []byte,
		txNum uint,
		nonce *uint64,
	) error
}
