package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create your own project",
	Long:  `You can create your project. Enter`,
	Run: func(cmd *cobra.Command, args []string) {
		sixseven()
	},
}

func sixseven() {
	fmt.Println("sixseven")

}
