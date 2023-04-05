package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/MXCzkEVM/mxc-client/cmd/flags"
	"github.com/MXCzkEVM/mxc-client/cmd/utils"
	"github.com/MXCzkEVM/mxc-client/driver"
	"github.com/MXCzkEVM/mxc-client/proposer"
	"github.com/MXCzkEVM/mxc-client/prover"
	"github.com/MXCzkEVM/mxc-client/version"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := cli.NewApp()

	app.Name = "MXC Clients"
	app.Usage = "The mxc client softwares command line interface"
	app.Copyright = "Copyright 2021-2022 MXC Foundation"
	app.Version = version.VersionWithCommit()
	app.Description = "Client softwares implementation in Golang for MXC protocol"
	app.Authors = []*cli.Author{{Name: "MXC Foundation", Email: "luanxu@mxc.org"}}
	app.EnableBashCompletion = true

	// All supported sub commands.
	app.Commands = []*cli.Command{
		{
			Name:        "driver",
			Flags:       flags.DriverFlags,
			Usage:       "Starts the driver software",
			Description: "MXC driver software",
			Action:      utils.SubcommandAction(new(driver.Driver)),
		},
		{
			Name:        "proposer",
			Flags:       flags.ProposerFlags,
			Usage:       "Starts the proposer software",
			Description: "MXC proposer software",
			Action:      utils.SubcommandAction(new(proposer.Proposer)),
		},
		{
			Name:        "prover",
			Flags:       flags.ProverFlags,
			Usage:       "Starts the prover software",
			Description: "MXC prover software",
			Action:      utils.SubcommandAction(new(prover.Prover)),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Crit("Failed to start MXC client", "error", err)
	}
}
