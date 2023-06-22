package flags

import "github.com/urfave/cli/v2"

var (
	MxcTokenAddress = &cli.StringFlag{
		Name:     "mxctoken",
		Usage:    "Mxc Token Contract Address on L1",
		Required: true,
		Category: relayerCategory,
	}
	LPWANAddress = &cli.StringFlag{
		Name:     "lpwan",
		Usage:    "LPWAN Contract Address on L2",
		Required: true,
		Category: relayerCategory,
	}
	RelayerPrivKey = &cli.StringFlag{
		Name:     "relayerPrivKey",
		Usage:    "Private key of the sync mxc token supply relayer",
		Required: true,
		Category: relayerCategory,
	}
)

var RelayerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1HTTPEndpoint,
	L2WSEndpoint,
	L2HTTPEndpoint,
	MxcTokenAddress,
	LPWANAddress,
	RelayerPrivKey,
})
