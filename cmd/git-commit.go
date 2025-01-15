package cmd

import (
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var commitMessage string

var GitCommit = &cobra.Command{
	Use:   "git-commit",
	Short: "Stage and commit all changes in the Dotty base folder.",
	Run: func(cmd *cobra.Command, args []string) {
		if commitMessage == "" {
			fmt.Println("Error: commit message is required.")
			os.Exit(1)
		}

		// Stage all changes
		if err := utils.RunGitCommand("add", "."); err != nil {
			fmt.Printf("Error staging files: %v\n", err)
			os.Exit(1)
		}

		// Commit with the provided message
		if err := utils.RunGitCommand("commit", "-m", commitMessage); err != nil {
			fmt.Printf("Error committing changes: %v\n", err)
			os.Exit(1)
		}
		commitMessage = ""
		fmt.Println("Changes committed successfully.")
	},
}

func init() {
	GitCommit.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message")
}
