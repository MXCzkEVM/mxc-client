package relayer

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/ethereum/go-ethereum/common"
	crypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

// Config contains the configurations to initialize a Mxc driver.
type Config struct {
	L1Endpoint      string
	L1HttpEndpoint  string
	L2Endpoint      string
	MxcL1Address    common.Address
	LPWANAddress    common.Address
	MxcTokenAddress common.Address
	RelayerPrivKey  *ecdsa.PrivateKey
}

// NewConfigFromCliContext creates a new config instance from
// the command line inputs.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	relayerPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(c.String(flags.RelayerPrivKey.Name)))
	if err != nil {
		return nil, fmt.Errorf("invalid relayer private key: %w", err)
	}
	return &Config{
		L1Endpoint:      c.String(flags.L1WSEndpoint.Name),
		L1HttpEndpoint:  c.String(flags.L1HTTPEndpoint.Name),
		L2Endpoint:      c.String(flags.L2WSEndpoint.Name),
		MxcL1Address:    common.HexToAddress(c.String(flags.MxcL1Address.Name)),
		LPWANAddress:    common.HexToAddress(c.String(flags.LPWANAddress.Name)),
		MxcTokenAddress: common.HexToAddress(c.String(flags.MxcTokenAddress.Name)),
		RelayerPrivKey:  relayerPrivKey,
	}, nil
}
