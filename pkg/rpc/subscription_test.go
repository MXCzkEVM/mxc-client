package rpc

import (
	"context"
	"testing"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/require"
)

func TestSubscribeEvent(t *testing.T) {
	require.NotNil(t, SubscribeEvent("test", func(ctx context.Context) (event.Subscription, error) {
		return event.NewSubscription(func(c <-chan struct{}) error { return nil }), nil
	}))
}

func TestSubscribeBlockVerified(t *testing.T) {
	require.NotNil(t, SubscribeBlockVerified(
		newTestClient(t).MXCL1,
		make(chan *bindings.MXCL1ClientBlockVerified, 1024)),
	)
}

func TestSubscribeBlockProposed(t *testing.T) {
	require.NotNil(t, SubscribeBlockProposed(
		newTestClient(t).MXCL1,
		make(chan *bindings.MXCL1ClientBlockProposed, 1024)),
	)
}

func TestSubscribeSubscribeHeaderSynced(t *testing.T) {
	require.NotNil(t, SubscribeHeaderSynced(
		newTestClient(t).MXCL1,
		make(chan *bindings.MXCL1ClientHeaderSynced, 1024)),
	)
}

func TestSubscribeBlockProven(t *testing.T) {
	require.NotNil(t, SubscribeBlockProven(
		newTestClient(t).MXCL1,
		make(chan *bindings.MXCL1ClientBlockProven, 1024)),
	)
}

func TestSubscribeChainHead(t *testing.T) {
	require.NotNil(t, SubscribeChainHead(
		newTestClient(t).L1,
		make(chan *types.Header, 1024)),
	)
}
