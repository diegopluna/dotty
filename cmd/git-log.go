package cmd

import (
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var oneline bool

var GitLog = &cobra.Command{
	Use:   "git-log",
	Short: "Show the commit history of the Dotty repository",
	Run: func(cmd *cobra.Command, args []string) {
		logArgs := []string{"log"}

		if oneline {
			logArgs = append(logArgs, "--oneline")
		}

		if err := utils.RunGitCommand(logArgs...); err != nil {
			fmt.Printf("Error running git log: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	GitLog.Flags().BoolVarP(&oneline, "oneline", "o", false, "Show the commit history in one line")
}
