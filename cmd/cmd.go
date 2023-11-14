package cmd

import (
	"log"

	"debitllama.com/relayer/workers"
	cli "github.com/urfave/cli"
)

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "debitrelayer"
	app.Usage = "Start the debitllama relayer and process subscription payments"

	passwdFlag := &cli.StringFlag{
		Name:  "password",
		Value: "",
		Usage: "Password to access the keystore",
	}

	xrelayerFlag := &cli.StringFlag{
		Name:  "xrelayer",
		Value: "",
		Usage: "The X-Relayer authentication header for the relayer",
	}

	authkeyFlag := &cli.StringFlag{
		Name:  "authorization",
		Value: "",
		Usage: "The bearer token to authenticate the debitllama API",
	}

	devenvFlag := &cli.BoolFlag{
		Name:  "dev",
		Usage: "Add this flag when developing using localhost",
	}

	networkFlag := &cli.StringSliceFlag{
		Name:  "chain",
		Usage: "List the networks you want to process transactions on. Example: --chain btt --chain btt_testnet etc..",
	}

	app.Commands = []cli.Command{
		{
			Name:    "balances",
			Aliases: []string{"b"},
			Usage:   "Get the relayer balances",
			Action: func(c *cli.Context) error {
				workers.GetRelayerBalances()
				return nil
			},
		},
		{
			Name:    "relay",
			Aliases: []string{"r"},
			Usage:   "Start relaying transactions",
			Action: func(c *cli.Context) error {
				xrelayer := c.String("xrelayer")
				authkey := c.String("authorization")
				devenv := c.Bool("dev")
				password := c.String("password")

				chains := c.StringSlice("chain")

				if xrelayer == "" {
					log.Fatalln("Missing XRelayer token")
				}
				if authkey == "" {
					log.Fatalln("Missing authkey")
				}
				if password == "" {
					log.Fatalln("Missing password")
				}

				if len(chains) == 0 {
					log.Fatalln("You need to use at least 1 --chain")
				}

				workers.StartWorkers(workers.CliArgs{
					Devenv:   devenv,
					Xrelayer: xrelayer,
					Authkey:  authkey,
					Password: password,
					Chains:   chains,
				})
				return nil
			},
			Flags: []cli.Flag{
				passwdFlag,
				xrelayerFlag,
				authkeyFlag,
				devenvFlag,
				networkFlag,
			},
		},
		{
			Name:    "new_keystore",
			Aliases: []string{"ks"},
			Usage:   "Create a new keystore",
			Action: func(c *cli.Context) error {
				if c.NumFlags() == 0 {
					log.Fatal("Need to enter a password")
				} else {
					password := c.String("password")
					workers.CreateKs(password)
				}

				return nil
			},
			Flags: []cli.Flag{
				passwdFlag,
			},
		},
		{
			Name:    "list_chains",
			Aliases: []string{"lsch"},
			Usage:   "List all available chains",
			Action: func(c *cli.Context) error {
				workers.PrintChainIdsHelp()
				return nil
			},
		},
	}

	return app
}
