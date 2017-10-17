package cmd

import (
	"fmt"

	"GoChain/chain"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [name of chain to create]",
	Short: "Add a new chain",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newChain := chain.CreateChain(args[0])
		if newChain != nil {
			fmt.Println("Creation Failed!")
			fmt.Println(newChain)
			return
		}
		fmt.Println(args[0] + " Chain Created.  Get To Work!")
	},
}
