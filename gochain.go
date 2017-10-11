package main

import (
	"fmt"
	"os"

	"gochain/chain"
	"gochain/cmd"
)

func main() {
	chain.CreateChain("cchhaaiinn")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
