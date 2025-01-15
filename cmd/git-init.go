package cmd

import (
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var GitInit = &cobra.Command{
	Use:   "git-init",
	Short: "Initialize a git repository in the Dotty base folder.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.RunGitCommand("init"); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Initialized a new Git repository in the Dotty base folder.")
	},
}
