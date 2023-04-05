package proposer

import (
	"context"
	"os"
	"strconv"

	"github.com/MXCzkEVM/mxc-client/bindings"
	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	l1Endpoint := os.Getenv("L1_NODE_WS_ENDPOINT")
	l2Endpoint := os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	mxcL1 := os.Getenv("MXC_L1_ADDRESS")
	mxcL2 := os.Getenv("MXC_L2_ADDRESS")
	proposeInterval := "10s"
	commitSlot := 1024

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.MXCL1Address.Name},
		&cli.StringFlag{Name: flags.MXCL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: flags.L2SuggestedFeeRecipient.Name},
		&cli.StringFlag{Name: flags.ProposeInterval.Name},
		&cli.Uint64Flag{Name: flags.CommitSlot.Name},
		&cli.StringFlag{Name: flags.TxPoolLocals.Name},
	}
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(mxcL1, c.MXCL1Address.String())
		s.Equal(mxcL2, c.MXCL2Address.String())
		s.Equal(bindings.GoldenTouchAddress, crypto.PubkeyToAddress(c.L1ProposerPrivKey.PublicKey))
		s.Equal(bindings.GoldenTouchAddress, c.L2SuggestedFeeRecipient)
		s.Equal(float64(10), c.ProposeInterval.Seconds())
		s.Equal(uint64(commitSlot), c.CommitSlot)
		s.Equal(1, len(c.LocalAddresses))
		s.Equal(bindings.GoldenTouchAddress, c.LocalAddresses[0])
		s.Nil(new(Proposer).InitFromCli(context.Background(), ctx))

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"-" + flags.L1WSEndpoint.Name, l1Endpoint,
		"-" + flags.L2HTTPEndpoint.Name, l2Endpoint,
		"-" + flags.MXCL1Address.Name, mxcL1,
		"-" + flags.MXCL2Address.Name, mxcL2,
		"-" + flags.L1ProposerPrivKey.Name, bindings.GoldenTouchPrivKey[2:],
		"-" + flags.L2SuggestedFeeRecipient.Name, bindings.GoldenTouchAddress.Hex(),
		"-" + flags.ProposeInterval.Name, proposeInterval,
		"-" + flags.CommitSlot.Name, strconv.Itoa(commitSlot),
		"-" + flags.TxPoolLocals.Name, bindings.GoldenTouchAddress.Hex(),
	}))
}
