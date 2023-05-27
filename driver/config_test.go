package driver

import (
	"context"
	"os"
	"time"

	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/urfave/cli/v2"
)

func (s *DriverTestSuite) TestNewConfigFromCliContext() {
	l1Endpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l2Endpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2EngineEndpoint := os.Getenv("L2_EXECUTION_ENGINE_AUTH_ENDPOINT")
	mxcL1 := os.Getenv("MXC_L1_ADDRESS")
	mxcL2 := os.Getenv("MXC_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2AuthEndpoint.Name},
		&cli.StringFlag{Name: flags.MxcL1Address.Name},
		&cli.StringFlag{Name: flags.MxcL2Address.Name},
		&cli.StringFlag{Name: flags.JWTSecret.Name},
		&cli.UintFlag{Name: flags.P2PSyncTimeout.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(l2EngineEndpoint, c.L2EngineEndpoint)
		s.Equal(mxcL1, c.MxcL1Address.String())
		s.Equal(mxcL2, c.MxcL2Address.String())
		s.Equal(120*time.Second, c.P2PSyncTimeout)
		s.NotEmpty(c.JwtSecret)
		s.Nil(new(Driver).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1Endpoint,
		"-" + flags.L2WSEndpoint.Name, l2Endpoint,
		"-" + flags.L2AuthEndpoint.Name, l2EngineEndpoint,
		"-" + flags.MxcL1Address.Name, mxcL1,
		"-" + flags.MxcL2Address.Name, mxcL2,
		"-" + flags.JWTSecret.Name, os.Getenv("JWT_SECRET"),
		"-" + flags.P2PSyncTimeout.Name, "120",
	}))
}
