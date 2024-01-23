package driver

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/MXCzkEVM/mxc-client/pkg/jwt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Mxc driver.
type Config struct {
	L1Endpoint            string
	L2Endpoint            string
	L2EngineEndpoint      string
	L2CheckPoint          string
	MxcL1Address          common.Address
	MxcL2Address          common.Address
	JwtSecret             string
	P2PSyncVerifiedBlocks bool
	P2PSyncTimeout        time.Duration
	IPFSGateways          []string
}

// NewConfigFromCliContext creates a new config instance from
// the command line inputs.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	jwtSecret, err := jwt.ParseSecretFromFile(c.String(flags.JWTSecret.Name))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT secret file: %w", err)
	}

	var (
		p2pSyncVerifiedBlocks = c.Bool(flags.P2PSyncVerifiedBlocks.Name)
		l2CheckPoint          = c.String(flags.CheckPointSyncUrl.Name)
	)

	if p2pSyncVerifiedBlocks && len(l2CheckPoint) == 0 {
		return nil, errors.New("empty L2 check point URL")
	}

	return &Config{
		L1Endpoint:            c.String(flags.L1WSEndpoint.Name),
		L2Endpoint:            c.String(flags.L2WSEndpoint.Name),
		L2EngineEndpoint:      c.String(flags.L2AuthEndpoint.Name),
		L2CheckPoint:          l2CheckPoint,
		MxcL1Address:          common.HexToAddress(c.String(flags.MxcL1Address.Name)),
		MxcL2Address:          common.HexToAddress(c.String(flags.MxcL2Address.Name)),
		JwtSecret:             string(jwtSecret),
		P2PSyncVerifiedBlocks: p2pSyncVerifiedBlocks,
		P2PSyncTimeout:        time.Duration(int64(time.Second) * int64(c.Uint(flags.P2PSyncTimeout.Name))),
		IPFSGateways:          strings.Split(c.String(flags.IPFSGateways.Name), ","),
	}, nil
}
