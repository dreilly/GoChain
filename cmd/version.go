package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Go Chain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Chain v0.1")
	},
}
