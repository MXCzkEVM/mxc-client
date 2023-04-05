package proposer

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

// Config contains all configurations to initialize a MXC proposer.
type Config struct {
	L1Endpoint                 string
	L2Endpoint                 string
	MXCL1Address               common.Address
	MXCL2Address               common.Address
	L1ProposerPrivKey          *ecdsa.PrivateKey
	L2SuggestedFeeRecipient    common.Address
	ProposeInterval            *time.Duration
	CommitSlot                 uint64
	LocalAddresses             []common.Address
	ProposeEmptyBlocksInterval *time.Duration
	MaxProposedTxListsPerEpoch uint64
	TxMinGasPrice              *big.Int
}

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProposerPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(c.String(flags.L1ProposerPrivKey.Name)),
	)
	if err != nil {
		return nil, fmt.Errorf("invalid L1 proposer private key: %w", err)
	}

	// Proposing configuration
	var proposingInterval *time.Duration
	if c.IsSet(flags.ProposeInterval.Name) {
		interval, err := time.ParseDuration(c.String(flags.ProposeInterval.Name))
		if err != nil {
			return nil, fmt.Errorf("invalid proposing interval: %w", err)
		}
		proposingInterval = &interval
	}
	var proposeEmptyBlocksInterval *time.Duration
	if c.IsSet(flags.ProposeEmptyBlocksInterval.Name) {
		interval, err := time.ParseDuration(c.String(flags.ProposeEmptyBlocksInterval.Name))
		if err != nil {
			return nil, fmt.Errorf("invalid proposing empty blocks interval: %w", err)
		}
		proposeEmptyBlocksInterval = &interval
	}
	var txMinGasPrice *big.Int
	if c.IsSet(flags.TxMinGasPrice.Name) {
		var ok bool
		txMinGasPrice, ok = big.NewInt(0).SetString(c.String(flags.TxMinGasPrice.Name), 10)
		if !ok {
			return nil, fmt.Errorf("invalid txMinGasPrice")
		}
	}

	l2SuggestedFeeRecipient := c.String(flags.L2SuggestedFeeRecipient.Name)
	if !common.IsHexAddress(l2SuggestedFeeRecipient) {
		return nil, fmt.Errorf("invalid L2 suggested fee recipient address: %s", l2SuggestedFeeRecipient)
	}

	localAddresses := []common.Address{}
	if c.IsSet(flags.TxPoolLocals.Name) {
		for _, account := range strings.Split(c.String(flags.TxPoolLocals.Name), ",") {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				return nil, fmt.Errorf("invalid account in --txpool.locals: %s", trimmed)
			} else {
				localAddresses = append(localAddresses, common.HexToAddress(account))
			}
		}
	}

	return &Config{
		L1Endpoint:                 c.String(flags.L1WSEndpoint.Name),
		L2Endpoint:                 c.String(flags.L2HTTPEndpoint.Name),
		MXCL1Address:               common.HexToAddress(c.String(flags.MXCL1Address.Name)),
		MXCL2Address:               common.HexToAddress(c.String(flags.MXCL2Address.Name)),
		L1ProposerPrivKey:          l1ProposerPrivKey,
		L2SuggestedFeeRecipient:    common.HexToAddress(l2SuggestedFeeRecipient),
		ProposeInterval:            proposingInterval,
		CommitSlot:                 c.Uint64(flags.CommitSlot.Name),
		LocalAddresses:             localAddresses,
		ProposeEmptyBlocksInterval: proposeEmptyBlocksInterval,
		MaxProposedTxListsPerEpoch: c.Uint64(flags.MaxProposedTxListsPerEpoch.Name),
		TxMinGasPrice:              txMinGasPrice,
	}, nil
}
