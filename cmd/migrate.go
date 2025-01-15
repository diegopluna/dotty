package cmd

import (
	"dotty/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate [source-folder]",
	Short: "Migrate a folder to the Dotty base folder and replace it with a symlink",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		srcPath, err := filepath.Abs(args[0])
		if err != nil {
			fmt.Println("Error resolving source path:", err)
			return
		}

		baseDir, err := utils.GetBaseFolder()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		appName := filepath.Base(srcPath)
		destPath := filepath.Join(baseDir, appName)

		if err := migrateApp(srcPath, destPath); err != nil {
			fmt.Println("Error migrating folder:", err)
		} else {
			fmt.Printf("Sucessfully migrate %s to Dotty.\n", appName)
		}
	},
}

func migrateApp(src, destBase string) error {
	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Calculate the relative path from the home directory
	relPath, err := filepath.Rel(homeDir, src)
	if err != nil {
		return fmt.Errorf("failed to calculate relative path: %w", err)
	}

	// Append the relative path to the destination path
	dest := filepath.Join(destBase, relPath)

	// Ensure the destination parent directory exist
	if err := os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create destination directories: %w", err)
	}

	// Move the source directory to the destination
	if err := os.Rename(src, dest); err != nil {
		return fmt.Errorf("failed to move %s to %s: %w", src, dest, err)
	}

	// Create a symlink from the source to the destination
	if err := os.Symlink(dest, src); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	return nil
}
