package cmd

import (
	"gochain/chain"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print [chain to print]",
	Short: "Print a specfied chain",
	Long:  `Prints a specified chain up to the config's length`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := chain.Chain{}
		c.PrintChain(args[0])
	},
}
