package main

import (
	"fmt"
	"gochain/chain"
	"os"

	"github.com/urfave/cli"
)

func main() {
	/*
		args := os.Args[1]
		if len(args) != 0 {
			fmt.Println(args)
		}
	*/

	chain.CreateChain("cchhaaiinn")
	app := cli.NewApp()
	app.Name = "gochain"
	app.Usage = "Don't break the chain"
	app.Action = func(c *cli.Context) error {
		fmt.Println("chains")
		return nil
	}

	app.Run(os.Args)
}
