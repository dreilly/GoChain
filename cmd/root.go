package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gochain",
	Short: "Go Chain is a toy app to motivate the user to continue working daily",
	Long: `Go Chain will add a link to the chain everyday some work is done.
				  The goal is to not break the chain.
				  `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please enter a valid command.  For a listing use -h or --help")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

}
