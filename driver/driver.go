package driver

import (
	"context"
	"sync"
	"time"

	chainSyncer "github.com/MXCzkEVM/mxc-client/driver/chain_syncer"
	"github.com/MXCzkEVM/mxc-client/driver/state"
	"github.com/MXCzkEVM/mxc-client/pkg/rpc"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

const (
	// Time to wait before the next try, when receiving subscription errors.
	RetryDelay = 10 * time.Second
)

// Driver keeps the L2 execution engine's local block chain in sync with the MXCL1
// contract.
type Driver struct {
	rpc           *rpc.Client
	l2ChainSyncer *chainSyncer.L2ChainSyncer
	state         *state.State

	l1HeadCh   chan *types.Header
	l1HeadSub  event.Subscription
	syncNotify chan struct{}

	ctx context.Context
	wg  sync.WaitGroup
}

// New initializes the given driver instance based on the command line flags.
func (d *Driver) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, d, cfg)
}

// InitFromConfig initializes the driver instance based on the given configurations.
func InitFromConfig(ctx context.Context, d *Driver, cfg *Config) (err error) {
	d.l1HeadCh = make(chan *types.Header, 1024)
	d.wg = sync.WaitGroup{}
	d.syncNotify = make(chan struct{}, 1)
	d.ctx = ctx

	if d.rpc, err = rpc.NewClient(d.ctx, &rpc.ClientConfig{
		L1Endpoint:       cfg.L1Endpoint,
		L2Endpoint:       cfg.L2Endpoint,
		MXCL1Address:     cfg.MXCL1Address,
		MXCL2Address:     cfg.MXCL2Address,
		L2EngineEndpoint: cfg.L2EngineEndpoint,
		JwtSecret:        cfg.JwtSecret,
	}); err != nil {
		return err
	}

	if d.state, err = state.New(d.ctx, d.rpc); err != nil {
		return err
	}

	peers, err := d.rpc.L2.PeerCount(d.ctx)
	if err != nil {
		return err
	}

	if cfg.P2PSyncVerifiedBlocks && peers == 0 {
		log.Warn("P2P syncing verified blocks enabled, but no connected peer found in L2 execution engine")
	}

	if d.l2ChainSyncer, err = chainSyncer.New(
		d.ctx,
		d.rpc,
		d.state,
		cfg.ThrowawayBlocksBuilderPrivKey,
		cfg.P2PSyncVerifiedBlocks,
		cfg.P2PSyncTimeout,
	); err != nil {
		return err
	}

	d.l1HeadSub = d.state.SubL1HeadsFeed(d.l1HeadCh)

	return nil
}

// Start starts the driver instance.
func (d *Driver) Start() error {
	d.wg.Add(2)
	go d.eventLoop()
	go d.reportProtocolStatus()

	return nil
}

// Close closes the driver instance.
func (d *Driver) Close() {
	d.state.Close()
	d.wg.Wait()
}

// eventLoop starts the main loop of a L2 execution engine's driver.
func (d *Driver) eventLoop() {
	defer d.wg.Done()
	// reqSync requests performing a synchronising operation, won't block
	// if we are already synchronising.
	reqSync := func() {
		d.syncNotify <- struct{}{}
		log.Info("d.syncNotify <- struct{}{}")
	}

	doSyncWithBackoff := func() {
		log.Info("doSyncWithBackoff")
		if err := backoff.Retry(d.doSync, backoff.NewExponentialBackOff()); err != nil {
			log.Error("Sync L2 execution engine's block chain error", "error", err)
		}
	}

	// Call doSync() right away to catch up with the latest known L1 head.
	reqSync()
	for {
		forceSyncTicker := time.NewTicker(time.Second * 3)
		select {
		case <-d.ctx.Done():
			log.Info("Driver context error", d.ctx.Err())
			return
		case <-d.syncNotify:
			log.Info("<-d.syncNotify")
			done := make(chan bool, 1)
			go func() {
				defer func() { done <- true }()
				doSyncWithBackoff()
			}()
			select {
			case <-done:
				continue
			case <-time.After(time.Second * 30):
				log.Error("doSyncWithBackoff timeout")
				panic("doSyncWithBackoff timeout")
				continue
			}
		case <-d.l1HeadCh:
			log.Info("<-d.l1HeadCh")
			reqSync()
			forceSyncTicker.Reset(time.Second * 3)
		case <-forceSyncTicker.C:
			log.Info("<-forceSyncTicker.C")
			reqSync()
		}
	}
}

// doSync fetches all `BlockProposed` events emitted from local
// L1 sync cursor to the L1 head, and then applies all corresponding
// L2 blocks into node's local block chain.
func (d *Driver) doSync() error {
	// Check whether the application is closing.
	if d.ctx.Err() != nil {
		log.Warn("Driver context error", "error", d.ctx.Err())
		return nil
	}

	l1Head := d.state.GetL1Head()

	if err := d.l2ChainSyncer.Sync(l1Head); err != nil {
		log.Error("Process new L1 blocks error", "error", err)
		return err
	}

	return nil
}

// ChainSyncer returns the driver's chain syncer.
func (d *Driver) ChainSyncer() *chainSyncer.L2ChainSyncer {
	return d.l2ChainSyncer
}

// reportProtocolStatus reports some protocol status intervally.
func (d *Driver) reportProtocolStatus() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		d.wg.Done()
	}()

	var maxNumBlocks uint64
	if err := backoff.Retry(
		func() error {
			configs, err := d.rpc.MXCL1.GetConfig(nil)
			if err != nil {
				return err
			}

			maxNumBlocks = configs.MaxNumBlocks.Uint64()
			return nil
		},
		backoff.NewConstantBackOff(RetryDelay),
	); err != nil {
		log.Error("Failed to get protocol state variables", "error", err)
		return
	}

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			vars, err := d.rpc.GetProtocolStateVariables(nil)
			if err != nil {
				log.Error("Failed to get protocol state variables", "error", err)
				continue
			}

			log.Info(
				"📖 Protocol status",
				"latestVerifiedId", vars.LatestVerifiedId,
				"latestVerifiedHeight", vars.LatestVerifiedHeight,
				"pendingBlocks", vars.NextBlockId-vars.LatestVerifiedId-1,
				"availableSlots", vars.LatestVerifiedId+maxNumBlocks-vars.NextBlockId,
			)
		}
	}
}

// Name returns the application name.
func (d *Driver) Name() string {
	return "driver"
}
