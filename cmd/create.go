package cmd

import (
	"fmt"

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
		fmt.Println("add new chain")
	},
}
