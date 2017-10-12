package cmd

import (
	"gochain/chain"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(linkCmd)
}

var linkCmd = &cobra.Command{
	Use:   "link [chain to add link to]",
	Short: "Add a link to a specified chain",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := chain.Chain{}
		c.CreateLink(args[0])
	},
}
