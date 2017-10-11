package cmd

import (
	"fmt"

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
		fmt.Println("this is where i would print the chain homie")
		printChain(args[0])
	},
}

func printChain(name string) error {
	// TODO: print stuff
	fmt.Println(name)
	return nil
}
