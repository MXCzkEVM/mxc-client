package rpc

import (
	"context"
	"math/big"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client contains all L1/L2 RPC clients that a driver needs.
type Client struct {
	// Geth ethclient clients
	L1           *ethclient.Client
	L2           *ethclient.Client
	L2CheckPoint *ethclient.Client
	// Geth gethclient clients
	L1GethClient *gethclient.Client
	L2GethClient *gethclient.Client
	// Geth raw RPC clients
	L1RawRPC *rpc.Client
	L2RawRPC *rpc.Client
	// Geth Engine API clients
	L2Engine *EngineClient
	// Protocol contracts clients
	MxcL1    *bindings.MxcL1Client
	MxcL2    *bindings.MxcL2Client
	LPWAN    *bindings.LPWANClient
	MxcToken *bindings.MxcTokenClient
	// Chain IDs
	L1ChainID *big.Int
	L2ChainID *big.Int
}

// ClientConfig contains all configs which will be used to initializing an
// RPC client. If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
// won't be initialized.
type ClientConfig struct {
	L1Endpoint       string
	L1HTTPEndpoint   string
	L2Endpoint       string
	L2CheckPoint     string
	MxcL1Address     common.Address
	MxcL2Address     common.Address
	LPWANAddress     *common.Address
	MxcTokenAddress  *common.Address
	L2EngineEndpoint string
	JwtSecret        string
}

// NewClient initializes all RPC clients used by Taiko client softwares.
func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	l1RPC, err := DialClientWithBackoff(ctx, cfg.L1Endpoint)
	if err != nil {
		return nil, err
	}

	mxcL1, err := bindings.NewMxcL1Client(cfg.MxcL1Address, l1RPC)
	if err != nil {
		return nil, err
	}
	var l1HTTPRPC *ethclient.Client
	if cfg.L1HTTPEndpoint != "" {
		l1HTTPRPC, err = DialClientWithBackoff(ctx, cfg.L1HTTPEndpoint)
		if err != nil {
			return nil, err
		}
		mxcL1ClientTransactor, err := bindings.NewMxcL1ClientTransactor(cfg.MxcL1Address, l1HTTPRPC)
		if err != nil {
			return nil, err
		}
		mxcL1.MxcL1ClientTransactor = *mxcL1ClientTransactor
	}

	l2RPC, err := DialClientWithBackoff(ctx, cfg.L2Endpoint)
	if err != nil {
		return nil, err
	}

	mxcL2, err := bindings.NewMxcL2Client(cfg.MxcL2Address, l2RPC)
	if err != nil {
		return nil, err
	}
	var LPWANClient *bindings.LPWANClient
	if cfg.LPWANAddress != nil {
		LPWANClient, err = bindings.NewLPWANClient(*cfg.LPWANAddress, l2RPC)
		if err != nil {
			return nil, err
		}
	}
	var MxcTokenClient *bindings.MxcTokenClient
	if cfg.MxcTokenAddress != nil {
		MxcTokenClient, err = bindings.NewMxcTokenClient(*cfg.MxcTokenAddress, l1RPC)
		if l1HTTPRPC != nil {
			MxcTokenClientTransactor, err := bindings.NewMxcTokenClientTransactor(*cfg.MxcTokenAddress, l1HTTPRPC)
			if err != nil {
				return nil, err
			}
			MxcTokenClient.MxcTokenClientTransactor = *MxcTokenClientTransactor
		}
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
		if l2AuthRPC, err = DialEngineClientWithBackoff(ctx, cfg.L2EngineEndpoint, cfg.JwtSecret); err != nil {
			return nil, err
		}
	}

	var l2CheckPoint *ethclient.Client
	if len(cfg.L2CheckPoint) != 0 {
		if l2CheckPoint, err = DialClientWithBackoff(ctx, cfg.L2CheckPoint); err != nil {
			return nil, err
		}
	}

	client := &Client{
		L1:           l1RPC,
		L2:           l2RPC,
		L2CheckPoint: l2CheckPoint,
		L1RawRPC:     l1RawRPC,
		L2RawRPC:     l2RawRPC,
		L1GethClient: gethclient.New(l1RawRPC),
		L2GethClient: gethclient.New(l2RawRPC),
		L2Engine:     l2AuthRPC,
		MxcL1:        mxcL1,
		MxcL2:        mxcL2,
		LPWAN:        LPWANClient,
		MxcToken:     MxcTokenClient,
		L1ChainID:    l1ChainID,
		L2ChainID:    l2ChainID,
	}

	if err := client.ensureGenesisMatched(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
