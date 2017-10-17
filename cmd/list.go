package cmd

import (
	"fmt"

	"GoChain/chain"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all Chains",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		chains, err := chain.GetAllChains()
		if err != nil {
			fmt.Println("Error getting chains")
			return
		}
		for _, f := range chains {
			fmt.Println(f)
		}
	},
}
