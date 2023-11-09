package cmd

import (
	"fmt"
	"log"

	eth "debitllama.com/relayer/eth"
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

	keyStorePathFlag := &cli.StringFlag{
		Name:  "path",
		Value: "",
		Usage: "The path to the keystore",
	}

	app.Commands = []cli.Command{
		{
			Name:    "balances",
			Aliases: []string{"b"},
			Usage:   "Get the relayer balances",
			Action: func(c *cli.Context) error {
				eth.GetRelayerBalances()
				return nil
			},
		},
		{
			Name:    "relay",
			Aliases: []string{"r"},
			Usage:   "Start relaying transactions",
			Action: func(c *cli.Context) error {
				fmt.Println("wanna start relaying txs")
				return nil
			},
			Flags: []cli.Flag{
				passwdFlag,
				keyStorePathFlag,
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
					eth.CreateKs(password)
				}

				return nil
			},
			Flags: []cli.Flag{
				passwdFlag,
			},
		},
	}

	return app
}

//Hello returns a greeting for something as a test
func Hello(name string) {
	fmt.Printf("hi, %v. Welcome!", name)
}
