package prover

import (
	"context"
	"os"
	"time"

	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func (s *ProverTestSuite) TestNewConfigFromCliContext() {
	l1WsEndpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint := os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint := os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	mxcL1 := os.Getenv("MXC_L1_ADDRESS")
	mxcL2 := os.Getenv("MXC_L2_ADDRESS")

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L1HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.MXCL1Address.Name},
		&cli.StringFlag{Name: flags.MXCL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProverPrivKey.Name},
		&cli.BoolFlag{Name: flags.Dummy.Name},
		&cli.StringFlag{Name: flags.RandomDummyProofDelay.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1WsEndpoint, c.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, c.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, c.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, c.L2HttpEndpoint)
		s.Equal(mxcL1, c.MXCL1Address.String())
		s.Equal(mxcL2, c.MXCL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey),
		)
		s.Equal(30*time.Minute, *c.RandomDummyProofDelayLowerBound)
		s.Equal(time.Hour, *c.RandomDummyProofDelayUpperBound)
		s.True(c.Dummy)
		s.Nil(new(Prover).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"-" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"-" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"-" + flags.MXCL1Address.Name, mxcL1,
		"-" + flags.MXCL2Address.Name, mxcL2,
		"-" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"-" + flags.Dummy.Name,
		"-" + flags.RandomDummyProofDelay.Name, "30m-1h",
	}))
}
