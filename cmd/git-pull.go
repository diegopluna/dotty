package cmd

import (
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pullRemote, pullBranch string
var rebase bool

var GitPull = &cobra.Command{
	Use:   "git-pull",
	Short: "Pull changes from a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		pullArgs := []string{"pull"}

		if rebase {
			pullArgs = append(pullArgs, "--rebase")
		}
		if pullRemote != "" {
			pullArgs = append(pullArgs, pullRemote)
		}

		if pullBranch != "" {
			pullArgs = append(pullArgs, pullBranch)
		}

		if err := utils.RunGitCommand(pullArgs...); err != nil {
			fmt.Printf("Error pulling changes: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Changes pulles successfully.")
	},
}

func init() {
	GitPull.Flags().BoolVarP(&rebase, "rebase", "", false, "Rebase local commits on top of the remote branch")
	GitPull.Flags().StringVarP(&pullRemote, "remote", "r", "", "Specify the remote repository (default: origin)")
	GitPull.Flags().StringVarP(&pullBranch, "branch", "b", "", "Specify the branch to pull from (default: current branch)")
}
