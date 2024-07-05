package eventiterator

import (
	"context"
	"errors"
	"github.com/MXCzkEVM/mxc-client/bindings"
	chainIterator "github.com/MXCzkEVM/mxc-client/pkg/chain_iterator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

// EndBlockProposedEventIterFunc ends the current iteration.
type EndBlockProposedEventIterFunc func()

// OnBlockProposedEvent represents the callback function which will be called when a MxcL1.BlockProposed event is
// iterated.
type OnBlockProposedEvent func(
	context.Context,
	*bindings.MxcL1ClientBlockProposed,
	EndBlockProposedEventIterFunc,
) error

// BlockProposedIterator iterates the emitted MxcL1.BlockProposed events in the chain,
// with the awareness of reorganization.
type BlockProposedIterator struct {
	ctx                context.Context
	mxcL1              *bindings.MxcL1Client
	blockBatchIterator *chainIterator.BlockBatchIterator
	filterQuery        []*big.Int
	isEnd              bool
}

// BlockProposedIteratorConfig represents the configs of a BlockProposed event iterator.
type BlockProposedIteratorConfig struct {
	Client                *ethclient.Client
	MxcL1                 *bindings.MxcL1Client
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	FilterQuery           []*big.Int
	Reverse               bool
	OnBlockProposedEvent  OnBlockProposedEvent
}

// NewBlockProposedIterator creates a new instance of BlockProposed event iterator.
func NewBlockProposedIterator(ctx context.Context, cfg *BlockProposedIteratorConfig) (*BlockProposedIterator, error) {
	if cfg.OnBlockProposedEvent == nil {
		return nil, errors.New("invalid callback")
	}

	iterator := &BlockProposedIterator{
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
		OnBlocks: assembleBlockProposedIteratorCallback(
			cfg.Client,
			cfg.MxcL1,
			cfg.FilterQuery,
			cfg.OnBlockProposedEvent,
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
// will call the callback when a BlockProposed event is iterated.
func (i *BlockProposedIterator) Iter() error {
	return i.blockBatchIterator.Iter()
}

// end ends the current iteration.
func (i *BlockProposedIterator) end() {
	i.isEnd = true
}

// assembleBlockProposedIteratorCallback assembles the callback which will be used
// by a event iterator's inner block iterator.
func assembleBlockProposedIteratorCallback(
	client *ethclient.Client,
	mxcL1Client *bindings.MxcL1Client,
	filterQuery []*big.Int,
	callback OnBlockProposedEvent,
	eventIter *BlockProposedIterator,
) chainIterator.OnBlocksFunc {
	return func(
		ctx context.Context,
		start, end *types.Header,
		updateCurrentFunc chainIterator.UpdateCurrentFunc,
		endFunc chainIterator.EndIterFunc,
	) error {
		var iter *bindings.MxcL1ClientBlockProposedIterator
		var err error
		startHeight := start.Number.Uint64()
		endHeight := end.Number.Uint64()
		log.Info("create iteratorCallback", "startHeight", startHeight, "end", endHeight)
		for {
			ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
			go func() {
				log.Info("FilterBlockProposed", "startHeight", startHeight, "end", endHeight)
				defer cancel()
				iter, err = mxcL1Client.FilterBlockProposed(
					&bind.FilterOpts{Start: startHeight, End: &endHeight, Context: ctxWithTimeout},
					filterQuery,
				)
				if err != nil {
					return
				}
				defer iter.Close()
			}()
			select {
			case <-ctxWithTimeout.Done():
				if err != nil {
					log.Error("FilterBlockProposed filterLog", "error", err)
					return err
				}
				if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
					log.Warn("FilterBlockProposed filterLog deadline exceeded")
					continue
				}
				break
			}
			break
		}

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
