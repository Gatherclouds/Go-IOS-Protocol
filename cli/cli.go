package cli

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func Run() {
	app := cli.NewApp()
	app.Name = "blockchain-tx"
	app.Usage = "Test the transaction part of blockchain."
	app.Commands = []cli.Command{
		{
			Name:    "getbalance",
			Aliases: []string{"a"},
			Usage:   "Get balance of ADDRESS",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "address",
					Value: "",
					Usage: "address of account",
				},
			},
			Action: func(c *cli.Context) error {
				address := c.String("address")
				if address == "" {
					fmt.Println("Address can't be empty!")
					return nil
				}
				getBalance(address)
				return nil
			},
		},
		{
			Name:    "createblockchain",
			Aliases: []string{"c"},
			Usage:   "Create a blockchain and send genesis block reward to ADDRESS",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "address",
					Value: "",
					Usage: "address of account",
				},
			},
			Action: func(c *cli.Context) error {
				address := c.String("address")
				if address == "" {
					fmt.Println("Address can't be empty!")
					return nil
				}
				createBlockchain(address)
				return nil
			},
		},
		{
			Name:    "printchain",
			Aliases: []string{"p"},
			Usage:   "Print all the blocks of the blockchain",
			Action: func(c *cli.Context) error {
				printChain()
				return nil
			},
		},
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "Send AMOUNT of coins from FROM address to TO",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "from",
					Value: "",
					Usage: "address of from",
				},
				cli.StringFlag{
					Name:  "to",
					Value: "",
					Usage: "address of to",
				},
				cli.IntFlag{
					Name:  "amount",
					Value: -1,
					Usage: "amount of btc to be sent",
				},
			},

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

