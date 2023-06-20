package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/MXCzkEVM/mxc-client/bindings"
	eventIterator "github.com/MXCzkEVM/mxc-client/pkg/chain_iterator/event_iterator"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
	"time"

	"github.com/MXCzkEVM/mxc-client/pkg/rpc"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

const (
	// RetryDelay Time to wait before the next try, when receiving subscription errors.
	RetryDelay = 10 * time.Second
)

// Relayer sync L1 proposed/proven reward to L2 and sync L2 burn to L1 token contract
type Relayer struct {
	rpc            *rpc.Client
	relayerPrivKey *ecdsa.PrivateKey

	l1CurrentBlockProposedReward uint64 // Current L1 block sync cursor
	l1CurrentBlockProvenReward   uint64

	blockProposedRewardCh chan *bindings.LPWANRewardEvent
	blockProvenRewardCh   chan *bindings.LPWANRewardEvent

	syncNotify chan struct{}

	ctx context.Context
	wg  sync.WaitGroup
}

// New initializes the given driver instance based on the command line flags.
func (r *Relayer) InitFromCli(ctx context.Context, c *cli.Context) error {
	cfg, err := NewConfigFromCliContext(c)
	if err != nil {
		return err
	}

	return InitFromConfig(ctx, r, cfg)
}

// InitFromConfig initializes the driver instance based on the given configurations.
func InitFromConfig(ctx context.Context, r *Relayer, cfg *Config) (err error) {
	r.wg = sync.WaitGroup{}
	r.syncNotify = make(chan struct{}, 1)
	r.relayerPrivKey = cfg.RelayerPrivKey
	r.blockProposedRewardCh = make(chan *bindings.LPWANRewardEvent, 10)
	r.blockProvenRewardCh = make(chan *bindings.LPWANRewardEvent, 10)
	r.ctx = ctx

	if r.rpc, err = rpc.NewClient(r.ctx, &rpc.ClientConfig{
		L1Endpoint:      cfg.L1Endpoint,
		L1HTTPEndpoint:  cfg.L1HttpEndpoint,
		L2Endpoint:      cfg.L2Endpoint,
		MxcL1Address:    cfg.MxcL1Address,
		LPWANAddress:    &cfg.LPWANAddress,
		MxcTokenAddress: &cfg.MxcTokenAddress,
	}); err != nil {
		return err
	}

	return nil
}

// Start starts the driver instance.
func (r *Relayer) Start() error {
	r.wg.Add(2)
	go r.eventLoop()
	go r.reportProtocolStatus()
	return nil
}

// Close closes the driver instance.
func (r *Relayer) Close() {
	r.wg.Wait()
}

// eventLoop starts the main loop of a L2 execution engine's driver.
func (r *Relayer) eventLoop() {
	defer r.wg.Done()
	constatnBackoff := backoff.NewConstantBackOff(12 * time.Second)

	// reqSync requests performing a synchronising operation, won't block
	// if we are already synchronising.
	reqSync := func() {
		select {
		case r.syncNotify <- struct{}{}:
		default:
		}
	}

	// doSyncWithBackoff performs a synchronising operation with a backoff strategy.
	doSyncWithBackoff := func() {
		if err := backoff.Retry(r.doSync, constatnBackoff); err != nil {
			log.Error("Sync L2 execution engine's block chain error", "error", err)
		}
	}

	// Call doSync() right away to catch up with the latest known L1 head.
	doSyncWithBackoff()

	syncL1RewardTicker := time.NewTicker(30 * time.Second)
	syncL2BurnTicker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-r.ctx.Done():
			return
		case e := <-r.blockProposedRewardCh:
			batchRewardEvent := []bindings.LPWANRewardEvent{*e}
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				for {
					select {
					case v := <-r.blockProposedRewardCh:
						batchRewardEvent = append(batchRewardEvent, *v)
					default:
						wg.Done()
						return
					}
				}
			}()
			wg.Wait()
			op := func() error {
				return r.syncProposedRewardOp(r.ctx, batchRewardEvent)
			}
			if err := backoff.Retry(op, backoff.NewConstantBackOff(time.Second)); err != nil {
				log.Info("syncProposedRewardOp error", "error", err)
			}
		case e := <-r.blockProvenRewardCh:
			batchRewardEvent := []bindings.LPWANRewardEvent{*e}
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				for {
					select {
					case v := <-r.blockProvenRewardCh:
						batchRewardEvent = append(batchRewardEvent, *v)
					default:
						wg.Done()
						return
					}
				}
			}()
			wg.Wait()
			err := r.syncProvenRewardOp(r.ctx, batchRewardEvent)
			if err != nil {
				log.Error("Sync ProposedReward error", "error", err.Error())
				continue
			}
		case <-r.syncNotify:
			doSyncWithBackoff()
		case <-syncL2BurnTicker.C:
			// get L2 zero address mxc balance and update to l1 mxc token contract
			err := r.syncL2MxcBurn(r.ctx)
			if err != nil {
				log.Error("Sync ProposedReward error", "error", err.Error())
				continue
			}
		case <-syncL1RewardTicker.C:
			reqSync()
		}
	}
}

// doSync fetches all `BlockProposed` events emitted from local
// L1 sync cursor to the L1 head, and then applies all corresponding
// L2 blocks into node's local block chain.
func (r *Relayer) doSync() error {
	// Check whether the application is closing.
	if r.l1CurrentBlockProposedReward == 0 || r.l1CurrentBlockProvenReward == 0 {
		err := r.resetL1Current()
		if err != nil {
			return err
		}
	}
	latestBlockNumber, err := r.rpc.L1.BlockNumber(r.ctx)
	if err != nil {
		return err
	}
	err = r.ProcessL1BlocksProposedReward(r.ctx, latestBlockNumber)
	if err != nil {
		return err
	}
	err = r.ProcessL1BlocksProvenReward(r.ctx, latestBlockNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *Relayer) resetL1Current() error {
	stateVars, err := r.rpc.GetProtocolStateVariables(nil)
	if err != nil {
		return err
	}
	syncStatus, err := r.rpc.LPWAN.GetRelaySyncStatus(nil)
	if err != nil {
		return err
	}
	if syncStatus.ProposedRewardEventHeight.Uint64() == 0 {
		r.l1CurrentBlockProposedReward = stateVars.GenesisHeight
	} else {
		l1Height, err := r.rpc.L2.L1OriginByID(r.ctx, syncStatus.ProposedRewardEventHeight)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				log.Warn("Failed to find proposedReward L1Origin for blockID", "blockID", syncStatus.ProposedRewardEventHeight.Uint64())
				return nil
			}
			return err
		}
		r.l1CurrentBlockProposedReward = l1Height.L1BlockHeight.Uint64() + 1
	}
	if syncStatus.ProvedRewardEventHeight.Uint64() == 0 {
		r.l1CurrentBlockProvenReward = stateVars.GenesisHeight
	} else {
		l1Height, err := r.rpc.L2.L1OriginByID(r.ctx, syncStatus.ProvedRewardEventHeight)
		if err != nil {
			if err.Error() == ethereum.NotFound.Error() {
				log.Warn("Failed to find provenReward L1Origin for blockID", "blockID", syncStatus.ProvedRewardEventHeight.Uint64())
				return nil
			}
			return err
		}
		r.l1CurrentBlockProvenReward = l1Height.L1BlockHeight.Uint64() + 1
	}
	return nil
}

func (r *Relayer) syncProposedRewardOp(ctx context.Context, input []bindings.LPWANRewardEvent) error {
	log.Info("Start execute LPWAN syncProposedRewardOp")
	opts, err := getTxOpts(ctx, r.rpc.L2, r.rpc.L2ChainID, r.relayerPrivKey)
	tx, err := r.rpc.LPWAN.SyncProposedRewardEvent(opts, input, false)
	if err != nil {
		log.Error("failed to sync proposed", "err", err)
		return fmt.Errorf("failed to sync proposed %v", err)
	}
	if _, err := rpc.WaitReceipt(ctx, r.rpc.L2, tx); err != nil {
		return err
	}

	log.Info("sync reward transactions succeeded", "hash", tx.Hash().String())
	return nil
}

func (r *Relayer) syncProvenRewardOp(ctx context.Context, input []bindings.LPWANRewardEvent) error {
	log.Info("Start execute LPWAN syncProvenRewardOp")
	opts, err := getTxOpts(ctx, r.rpc.L2, r.rpc.L2ChainID, r.relayerPrivKey)
	tx, err := r.rpc.LPWAN.SyncProvenRewardEvent(opts, input, false)
	if err != nil {
		return fmt.Errorf("failed to sync proposed %v", err)
	}
	if _, err := rpc.WaitReceipt(ctx, r.rpc.L2, tx); err != nil {
		return err
	}

	log.Info("sync reward transactions succeeded", "hash", tx.Hash().String())
	return nil
}

func (r *Relayer) syncL2MxcBurn(ctx context.Context) error {
	if err := r.rpc.WaitTillL2ExecutionEngineSynced(ctx); err != nil {
		return fmt.Errorf("failed to wait until L2 execution engine synced: %w", err)
	}
	blockNumber, err := r.rpc.L2.BlockNumber(r.ctx)
	if err != nil {
		return err
	}
	balance, err := r.rpc.L2.BalanceAt(ctx, common.BigToAddress(common.Big0), big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}
	log.Info("l2 zero address mxc balance", "balance", balance.Uint64())
	opts, err := getTxOpts(ctx, r.rpc.L1, r.rpc.L1ChainID, r.relayerPrivKey)
	burnTx, err := r.rpc.MxcToken.SyncL2Burn(opts, balance)
	if err != nil {
		return fmt.Errorf("sync L2 mxc token burn error %w", err)
	}
	if _, err := rpc.WaitReceipt(ctx, r.rpc.L1, burnTx); err != nil {
		return err
	}
	log.Info("sync l2 mxc zero address burn success", "hash", burnTx.Hash().String())
	return nil
}

func (r *Relayer) ProcessL1BlocksProposedReward(ctx context.Context, latestBlockNumber uint64) error {
	iter, err := eventIterator.NewBlockProposeRewardIterator(ctx, &eventIterator.BlockProposeRewardIteratorConfig{
		Client:                    r.rpc.L1,
		MxcL1:                     r.rpc.MxcL1,
		StartHeight:               new(big.Int).SetUint64(r.l1CurrentBlockProposedReward),
		EndHeight:                 new(big.Int).SetUint64(latestBlockNumber),
		FilterQuery:               nil,
		OnBlockProposeRewardEvent: r.onBlockProposedReward,
	})
	if err != nil {
		return err
	}
	if err := iter.Iter(); err != nil {
		return err
	}
	return nil
}

func (r *Relayer) onBlockProposedReward(
	ctx context.Context,
	event *bindings.MxcL1ClientBlockProposeReward,
	endIter eventIterator.EndBlockProposeRewardEventIterFunc,
) error {
	if event.Id.Cmp(common.Big0) == 0 {
		return nil
	}
	// sort and send to channel
	tx, pending, err := r.rpc.L1.TransactionByHash(ctx, event.Raw.TxHash)
	if err != nil {
		return err
	}
	if pending {
		return fmt.Errorf("transaction pending")
	}
	rewardEvent := &bindings.LPWANRewardEvent{
		RewardHeight: event.Id,
		Account:      event.Proposer,
		Amount:       event.Reward,
		Cost:         tx.Cost(),
	}
	r.blockProposedRewardCh <- rewardEvent
	log.Info("New BlockProposedReward Event", "RewardEvent", rewardEvent)
	r.l1CurrentBlockProposedReward = event.Raw.BlockNumber
	if len(r.blockProposedRewardCh) == cap(r.blockProposedRewardCh) {
		endIter()
	}
	return nil
}

func (r *Relayer) ProcessL1BlocksProvenReward(ctx context.Context, latestBlockNumber uint64) error {
	iter, err := eventIterator.NewBlockProvenRewardIterator(ctx, &eventIterator.BlockProvenRewardIteratorConfig{
		Client:                   r.rpc.L1,
		MxcL1:                    r.rpc.MxcL1,
		StartHeight:              new(big.Int).SetUint64(r.l1CurrentBlockProvenReward),
		EndHeight:                new(big.Int).SetUint64(latestBlockNumber),
		FilterQuery:              nil,
		OnBlockProvenRewardEvent: r.onBlockProvenReward,
	})
	if err != nil {
		return err
	}
	if err := iter.Iter(); err != nil {
		return err
	}
	return nil
}

func (r *Relayer) onBlockProvenReward(
	ctx context.Context,
	event *bindings.MxcL1ClientBlockProvenReward,
	endIter eventIterator.EndBlockProvenRewardEventIterFunc,
) error {
	if event.Id.Cmp(common.Big0) == 0 {
		return nil
	}
	// sort and send to channel
	tx, pending, err := r.rpc.L1.TransactionByHash(ctx, event.Raw.TxHash)
	if err != nil {
		return err
	}
	if pending {
		return fmt.Errorf("transaction pending")
	}
	rewardEvent := &bindings.LPWANRewardEvent{
		RewardHeight: event.Id,
		Account:      event.Prover,
		Amount:       event.Reward,
		Cost:         tx.Cost(),
	}
	r.blockProvenRewardCh <- rewardEvent
	r.l1CurrentBlockProvenReward = event.Raw.BlockNumber
	if len(r.blockProvenRewardCh) == cap(r.blockProvenRewardCh) {
		endIter()
	}
	return nil
}

// reportProtocolStatus reports some protocol status intervally.
func (r *Relayer) reportProtocolStatus() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		r.wg.Done()
	}()
	var relaySyncStatus bindings.LPWANRelaySyncStatus
	var protocolVars *bindings.MxcDataStateVariables
	if err := backoff.Retry(
		func() error {
			var err error
			protocolVars, err = r.rpc.GetProtocolStateVariables(nil)
			if err != nil {
				log.Error("Failed to get protocol state variables", "error", err)
			}
			relaySyncStatus, err = r.rpc.LPWAN.GetRelaySyncStatus(nil)
			if err != nil {
				return err
			}
			return nil
		},
		backoff.NewConstantBackOff(RetryDelay),
	); err != nil {
		log.Error("Failed to get relay sync status", "error", err)
		return
	}
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-ticker.C:
			log.Info(
				"Sync Relayer status",
				"L2 Height", protocolVars.NumBlocks,
				"l1CurrentBlockProposedReward", r.l1CurrentBlockProposedReward,
				"l1CurrentBlockProvenReward", r.l1CurrentBlockProvenReward,
				"ProposedRewardEventHeight", relaySyncStatus.ProposedRewardEventHeight.Uint64(),
				"ProvedRewardEventHeight", relaySyncStatus.ProvedRewardEventHeight.Uint64(),
			)
		}
	}
}

func getTxOpts(
	ctx context.Context,
	cli *ethclient.Client,
	chainID *big.Int,
	proverPrivKey *ecdsa.PrivateKey,
) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(proverPrivKey, chainID)
	if err != nil {
		return nil, err
	}
	gasTipCap, err := cli.SuggestGasTipCap(ctx)
	if err != nil {
		if rpc.IsMaxPriorityFeePerGasNotFoundError(err) {
			gasTipCap = rpc.FallbackGasTipCap
		} else {
			return nil, err
		}
	}

	opts.GasTipCap = gasTipCap

	return opts, nil
}

// Name returns the application name.
func (r *Relayer) Name() string {
	return "relayer"
}
