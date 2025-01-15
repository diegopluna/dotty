package cmd

import (
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var GitStatus = &cobra.Command{
	Use:   "git-status",
	Short: "Show the status of the git repository in the Dotty base folder.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.RunGitCommand("status"); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}
