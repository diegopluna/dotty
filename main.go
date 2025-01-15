package main

import (
	"dotty/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dotty",
		Short: "Manage your dotfiles like a pro.",
	}

	rootCmd.AddCommand(cmd.LinkCmd, cmd.InitCmd, cmd.UnlinkCmd, cmd.MigrateCmd, cmd.RestoreCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
