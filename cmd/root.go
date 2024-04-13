package cmd

import (
	"dist/lev"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "dist [w1 w2]",
	Short: "Levenshtein Distance Calculator.\nPass two strings to calculate distance. Call with no args for interactive mode.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			RunInteractive()
		} else if len(args) == 2 {
			fmt.Println(lev.Calc(args[0], args[1]))
		} else {
			fmt.Fprintln(os.Stderr, "Error: Invalid number of arguments")
			fmt.Println()

			cmd.Help()
		}
	},
}
