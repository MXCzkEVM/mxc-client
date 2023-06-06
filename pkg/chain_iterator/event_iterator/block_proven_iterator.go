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

// EndBlockProvenEventIterFunc ends the current iteration.
type EndBlockProvenEventIterFunc func()

// OnBlockProvenEvent represents the callback function which will be called when a MxcL1.BlockProven event is
// iterated.
type OnBlockProvenEvent func(context.Context, *bindings.MxcL1ClientBlockProven, EndBlockProvenEventIterFunc) error

// BlockProvenIterator iterates the emitted MxcL1.BlockProven events in the chain,
// with the awareness of reorganization.
type BlockProvenIterator struct {
	ctx                context.Context
	mxcL1              *bindings.MxcL1Client
	blockBatchIterator *chainIterator.BlockBatchIterator
	filterQuery        []*big.Int
	isEnd              bool
}

// BlockProvenIteratorConfig represents the configs of a BlockProven event iterator.
type BlockProvenIteratorConfig struct {
	Client                *ethclient.Client
	MxcL1                 *bindings.MxcL1Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	FilterQuery           []*big.Int
	Reverse               bool
	OnBlockProvenEvent    OnBlockProvenEvent
}

// NewBlockProvenIterator creates a new instance of BlockProven event iterator.
func NewBlockProvenIterator(ctx context.Context, cfg *BlockProvenIteratorConfig) (*BlockProvenIterator, error) {
	if cfg.OnBlockProvenEvent == nil {
		return nil, errors.New("invalid callback")
	}

	iterator := &BlockProvenIterator{
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
		OnBlocks: assembleBlockProvenIteratorCallback(
			cfg.Client,
			cfg.MxcL1,
			cfg.FilterQuery,
			cfg.OnBlockProvenEvent,
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
// will call the callback when a BlockProven event is iterated.
func (i *BlockProvenIterator) Iter() error {
	return i.blockBatchIterator.Iter()
}

// end ends the current iteration.
func (i *BlockProvenIterator) end() {
	i.isEnd = true
}

// assembleBlockProvenIteratorCallback assembles the callback which will be used
// by a event iterator's inner block iterator.
func assembleBlockProvenIteratorCallback(
	client *ethclient.Client,
	mxcL1Client *bindings.MxcL1Client,
	filterQuery []*big.Int,
	callback OnBlockProvenEvent,
	eventIter *BlockProvenIterator,
) chainIterator.OnBlocksFunc {
	return func(
		ctx context.Context,
		start, end *types.Header,
		updateCurrentFunc chainIterator.UpdateCurrentFunc,
		endFunc chainIterator.EndIterFunc,
	) error {
		endHeight := end.Number.Uint64()
		iter, err := mxcL1Client.FilterBlockProven(
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
