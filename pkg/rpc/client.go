package rpc

import (
	"context"
	"math/big"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client contains all L1/L2 RPC clients that a driver needs.
type Client struct {
	// Geth ethclient clients
	L1 *ethclient.Client
	L2 *ethclient.Client
	// Geth raw RPC clients
	L1RawRPC *rpc.Client
	L2RawRPC *rpc.Client
	// Geth Engine API clients
	L2Engine *EngineClient
	// Protocol contracts clients
	MXCL1            *bindings.MXCL1Client
	MXCL2            *bindings.MXCL2Client
	ArbGasInfo       *bindings.ArbGasInfo
	ArbNodeInterface *bindings.NodeInterface
	// Chain IDs
	L1ChainID *big.Int
	L2ChainID *big.Int

	cfg *ClientConfig
}

// ClientConfig contains all configs which will be used to initializing an
// RPC client. If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
// won't be initialized.
type ClientConfig struct {
	L1Endpoint       string
	L2Endpoint       string
	MXCL1Address     common.Address
	MXCL2Address     common.Address
	L2EngineEndpoint string
	JwtSecret        string
}

// NewClient initializes all RPC clients used by MXC client softwares.
func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	l1RPC, err := DialClientWithBackoff(ctx, cfg.L1Endpoint)
	if err != nil {
		return nil, err
	}

	mxcL1, err := bindings.NewMXCL1Client(cfg.MXCL1Address, l1RPC)
	if err != nil {
		return nil, err
	}

	l2RPC, err := DialClientWithBackoff(ctx, cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	mxcL2, err := bindings.NewMXCL2Client(cfg.MXCL2Address, l2RPC)
	if err != nil {
		return nil, err
	}

	arbGasInfo, err := bindings.NewArbGasInfo(common.HexToAddress("0x000000000000000000000000000000000000006c"), l1RPC)
	if err != nil {
		return nil, err
	}

	arbNodeInterface, err := bindings.NewNodeInterface(common.HexToAddress("0x00000000000000000000000000000000000000C8"), l1RPC)
	if err != nil {
		return nil, err
	}

	l1RawRPC, err := rpc.Dial(cfg.L1Endpoint)
	if err != nil {
		return nil, err
	}

	l2RawRPC, err := rpc.Dial(cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	l1ChainID, err := l1RPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	l2ChainID, err := l2RPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	// If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
	// won't be initialized.
	var l2AuthRPC *EngineClient
	if len(cfg.L2EngineEndpoint) != 0 && len(cfg.JwtSecret) != 0 {
		l2AuthRPC, err = DialEngineClientWithBackoff(ctx, cfg.L2EngineEndpoint, cfg.JwtSecret)
		if err != nil {
			return nil, err
		}
	}

	client := &Client{
		L1:               l1RPC,
		L2:               l2RPC,
		L1RawRPC:         l1RawRPC,
		L2RawRPC:         l2RawRPC,
		L2Engine:         l2AuthRPC,
		MXCL1:            mxcL1,
		MXCL2:            mxcL2,
		ArbGasInfo:       arbGasInfo,
		ArbNodeInterface: arbNodeInterface,
		L1ChainID:        l1ChainID,
		L2ChainID:        l2ChainID,
		cfg:              cfg,
	}

	if err := client.ensureGenesisMatched(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
