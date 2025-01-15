package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init [base-folder]",
	Short: "Initialize dotty with the base folder for your dotfiles",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		baseFolder := args[0]
		absPath, err := filepath.Abs(baseFolder)
		if err != nil {
			fmt.Println("error resolving absolute path:", err)
		}

		if err := setBaseFolder(absPath); err != nil {
			fmt.Println("error initializing dotty:", err)
		} else {
			fmt.Println("Dotty intialized with base folder:", absPath)
		}
	},
}

func setBaseFolder(baseFolder string) error {
	configPath := filepath.Join(os.Getenv("HOME"), ".dottyconfig")

	file, err := os.Create(configPath)

	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(baseFolder)
	if err != nil {
		return fmt.Errorf("failed to write config to file: %w", err)
	}

	return nil
}
