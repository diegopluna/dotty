package cmd

import (
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var remote, branch string

var GitPush = &cobra.Command{
	Use:   "git-push",
	Short: "Push local commits to the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		pushArgs := []string{"push"}

		if remote != "" {
			pushArgs = append(pushArgs, remote)
		}

		if branch != "" {
			pushArgs = append(pushArgs, branch)
		}

		if err := utils.RunGitCommand(pushArgs...); err != nil {
			fmt.Printf("Error pushing changes: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Changes pushed successfully.")
	},
}

func init() {
	GitPush.Flags().StringVarP(&remote, "remote", "r", "", "Remote repository to push to")
	GitPush.Flags().StringVarP(&branch, "branch", "b", "", "Branch to push to")
}
