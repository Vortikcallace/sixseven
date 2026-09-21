package cli

import (
	"bufio"
	"fmt"
	"os"
	"sixseven/wiki"
	"strings"

	"github.com/spf13/cobra"
)

var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "learn sixseven",
	Long:  `It's command for upgrade yout knowledge about sixseven`,
	Run: func(cmd *cobra.Command, args []string) {
		sixseven()
	},
}

func sixseven() {
	fmt.Print("Do you want to know more about sixseven? [Y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ans := strings.ToLower(strings.TrimSpace(scanner.Text()))
		switch ans {
		case "", "y", "yes":
			wiki.Main()
		case "n", "no":
			fmt.Println("You're fucking bezdar")
		default:
			fmt.Println("Invalid action.")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(learnCmd)
}
