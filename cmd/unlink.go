package cmd

import (
	"dotty/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var UnlinkCmd = &cobra.Command{
	Use:   "unlink [app]",
	Short: "Remove symlinks for a specific app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := args[0]

		baseDir, err := utils.GetBaseFolder()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}

		if err := unlinkApp(app, baseDir, homeDir); err != nil {
			fmt.Println("Error unlink app:", err)
		} else {
			fmt.Println("Successfully unlinked", app)
		}
	},
}

func unlinkApp(appName, baseDir, homeDir string) error {
	appPath := filepath.Join(baseDir, appName)

	return filepath.Walk(appPath, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(appPath, srcPath)
		dstPath := filepath.Join(homeDir, relPath)

		// Check if it's a symlink before removing
		fileInfo, err := os.Lstat(dstPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("skipping non-existent file:", dstPath)
				return nil
			}
			return fmt.Errorf("error checking file: %w", err)
		}

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if err := os.Remove(dstPath); err != nil {
				return fmt.Errorf("failed to remove symlink: %w", err)
			}
			fmt.Println("unlinked:", dstPath)
		} else {
			fmt.Println("skipping non-symlink file:", dstPath)
		}

		return nil
	})

}
