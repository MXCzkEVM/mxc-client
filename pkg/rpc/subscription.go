package rpc

import (
	"context"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

// SubscribeEvent creates a event subscription, will retry if the established subscription failed.
func SubscribeEvent(
	eventName string,
	handler func(ctx context.Context) (event.Subscription, error),
) event.Subscription {
	return event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe protocol event, try resubscribing", "event", eventName, "error", err)
			}

			return handler(ctx)
		},
	)
}

// SubscribeBlockVerified subscribes the protocol's BlockVerified events.
func SubscribeBlockVerified(
	mxcL1 *bindings.MXCL1Client,
	ch chan *bindings.MXCL1ClientBlockVerified,
) event.Subscription {
	return SubscribeEvent("BlockVerified", func(ctx context.Context) (event.Subscription, error) {
		sub, err := mxcL1.WatchBlockVerified(nil, ch, nil)
		if err != nil {
			log.Error("Create MXCL1.BlockVerified subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeBlockProposed subscribes the protocol's BlockProposed events.
func SubscribeBlockProposed(
	mxcL1 *bindings.MXCL1Client,
	ch chan *bindings.MXCL1ClientBlockProposed,
) event.Subscription {
	return SubscribeEvent("BlockProposed", func(ctx context.Context) (event.Subscription, error) {
		sub, err := mxcL1.WatchBlockProposed(nil, ch, nil)
		if err != nil {
			log.Error("Create MXCL1.BlockProposed subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeHeaderSynced subscribes the protocol's HeaderSynced events.
func SubscribeHeaderSynced(
	mxcL1 *bindings.MXCL1Client,
	ch chan *bindings.MXCL1ClientHeaderSynced,
) event.Subscription {
	return SubscribeEvent("HeaderSynced", func(ctx context.Context) (event.Subscription, error) {
		sub, err := mxcL1.WatchHeaderSynced(nil, ch, nil)
		if err != nil {
			log.Error("Create MXCL1.HeaderSynced subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeBlockProven subscribes the protocol's BlockProven events.
func SubscribeBlockProven(
	mxcL1 *bindings.MXCL1Client,
	ch chan *bindings.MXCL1ClientBlockProven,
) event.Subscription {
	return SubscribeEvent("BlockProven", func(ctx context.Context) (event.Subscription, error) {
		sub, err := mxcL1.WatchBlockProven(nil, ch, nil)
		if err != nil {
			log.Error("Create MXCL1.BlockProven subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeChainHead subscribes the new chain heads.
func SubscribeChainHead(
	client *ethclient.Client,
	ch chan *types.Header,
) event.Subscription {
	return SubscribeEvent("ChainHead", func(ctx context.Context) (event.Subscription, error) {
		sub, err := client.SubscribeNewHead(ctx, ch)
		if err != nil {
			log.Error("Create chain head subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// waitSubErr keeps waiting until the given subscription failed.
func waitSubErr(ctx context.Context, sub event.Subscription) (event.Subscription, error) {
	for {
		select {
		case err := <-sub.Err():
			return sub, err
		case <-ctx.Done():
			return sub, nil
		}
	}
}
