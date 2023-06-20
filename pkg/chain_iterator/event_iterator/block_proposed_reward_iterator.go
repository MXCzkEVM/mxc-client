package eventiterator

import (
	"context"
	"errors"
	"math/big"

	"github.com/MXCzkEVM/mxc-client/bindings"
	chainIterator "github.com/MXCzkEVM/mxc-client/pkg/chain_iterator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EndBlockProposeRewardEventIterFunc ends the current iteration.
type EndBlockProposeRewardEventIterFunc func()

// OnBlockProposeRewardEvent represents the callback function which will be called when a MxcL1.BlockProposeReward event is
// iterated.
type OnBlockProposeRewardEvent func(
	context.Context,
	*bindings.MxcL1ClientBlockProposeReward,
	EndBlockProposeRewardEventIterFunc,
) error

// BlockProposeRewardIterator iterates the emitted MxcL1.BlockProposeReward events in the chain,
// with the awareness of reorganization.
type BlockProposeRewardIterator struct {
	ctx                context.Context
	mxcL1              *bindings.MxcL1Client
	blockBatchIterator *chainIterator.BlockBatchIterator
	filterQuery        []*big.Int
	isEnd              bool
}

// BlockProposeRewardIteratorConfig represents the configs of a BlockProposeReward event iterator.
type BlockProposeRewardIteratorConfig struct {
	Client                    *ethclient.Client
	MxcL1                     *bindings.MxcL1Client
	MaxBlocksReadPerEpoch     *uint64
	StartHeight               *big.Int
	EndHeight                 *big.Int
	FilterQuery               []*big.Int
	Reverse                   bool
	OnBlockProposeRewardEvent OnBlockProposeRewardEvent
}

// NewBlockProposeRewardIterator creates a new instance of BlockProposeReward event iterator.
func NewBlockProposeRewardIterator(ctx context.Context, cfg *BlockProposeRewardIteratorConfig) (*BlockProposeRewardIterator, error) {
	if cfg.OnBlockProposeRewardEvent == nil {
		return nil, errors.New("invalid callback")
	}

	iterator := &BlockProposeRewardIterator{
		ctx:         ctx,
		mxcL1:       cfg.MxcL1,
		filterQuery: cfg.FilterQuery,
	}

	// Initialize the inner block iterator.
	blockIterator, err := chainIterator.NewBlockBatchIterator(ctx, &chainIterator.BlockBatchIteratorConfig{
		Client:                cfg.Client,
		MaxBlocksReadPerEpoch: cfg.MaxBlocksReadPerEpoch,
		StartHeight:           cfg.StartHeight,
		EndHeight:             cfg.EndHeight,
		Reverse:               cfg.Reverse,
		OnBlocks: assembleBlockProposeRewardIteratorCallback(
			cfg.Client,
			cfg.MxcL1,
			cfg.FilterQuery,
			cfg.OnBlockProposeRewardEvent,
			iterator,
		),
	})
	if err != nil {
		return nil, err
	}

	iterator.blockBatchIterator = blockIterator

	return iterator, nil
}

// Iter iterates the given chain between the given start and end heights,
// will call the callback when a BlockProposeReward event is iterated.
func (i *BlockProposeRewardIterator) Iter() error {
	return i.blockBatchIterator.Iter()
}

// end ends the current iteration.
func (i *BlockProposeRewardIterator) end() {
	i.isEnd = true
}

// assembleBlockProposeRewardIteratorCallback assembles the callback which will be used
// by a event iterator's inner block iterator.
func assembleBlockProposeRewardIteratorCallback(
	client *ethclient.Client,
	mxcL1Client *bindings.MxcL1Client,
	filterQuery []*big.Int,
	callback OnBlockProposeRewardEvent,
	eventIter *BlockProposeRewardIterator,
) chainIterator.OnBlocksFunc {
	return func(
		ctx context.Context,
		start, end *types.Header,
		updateCurrentFunc chainIterator.UpdateCurrentFunc,
		endFunc chainIterator.EndIterFunc,
	) error {
		endHeight := end.Number.Uint64()
		iter, err := mxcL1Client.FilterBlockProposeReward(
			&bind.FilterOpts{Start: start.Number.Uint64(), End: &endHeight, Context: ctx},
			filterQuery,
		)
		if err != nil {
			return err
		}
		defer iter.Close()

		for iter.Next() {
			event := iter.Event

			if err := callback(ctx, event, eventIter.end); err != nil {
				return err
			}

			if eventIter.isEnd {
				endFunc()
				return nil
			}

			current, err := client.HeaderByHash(ctx, event.Raw.BlockHash)
			if err != nil {
				return err
			}

			updateCurrentFunc(current)
		}

		return nil
	}
}
