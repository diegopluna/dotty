package main

import (
	"dotty/cmd"
	"dotty/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	commit      string
	date        string
	version     string = "dev"
	buildSource        = "unkown"
)

var Upgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade dotty to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for updates...")

		latestVersion, downloadURL, err := utils.CheckForUpdates()
		if err != nil {
			fmt.Printf("Error checking for updates: %v\n", err)
			os.Exit(1)
		}

		if latestVersion > version {
			fmt.Printf("New version available: %s\n", latestVersion)
			fmt.Println("Downloading the latest version")

			if err := utils.DownloadAndUpdate(downloadURL); err != nil {
				fmt.Printf("Error downloading the update: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Update successful.")
		} else {
			fmt.Println("You're already using the latest version.")
		}
	},
}

func main() {
	var versionFlag bool
	rootCmd := &cobra.Command{
		Use:   "dotty",
		Short: "Manage your dotfiles like a pro.",
		Run: func(cmd *cobra.Command, args []string) {
			if versionFlag {
				fmt.Printf("dotty version %s, built on commit %s at %s with build source %s\n", version, commit, date, buildSource)
			} else {
				fmt.Println("Please provide a subcommand. Run 'dotty --help' for more information.")
			}
		},
	}

	rootCmd.AddCommand(cmd.LinkCmd, cmd.InitCmd, cmd.UnlinkCmd, cmd.MigrateCmd, cmd.RestoreCmd, cmd.GitInit, cmd.GitStatus, cmd.GitCommit, cmd.GitLog, cmd.GitPush, cmd.GitPull)
	rootCmd.AddCommand(Upgrade)

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Show the current version of dotty")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
