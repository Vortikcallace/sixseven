package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sixseven",
	Short: "just autism",
	Long:  `just autism`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the last autism's castle")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
