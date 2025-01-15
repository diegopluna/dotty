package cmd

import (
	"dotty/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var LinkCmd = &cobra.Command{
	Use:   "link [app]",
	Short: "Create symlinks for a specific app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := args[0]
		baseDir, err := utils.GetBaseFolder() // TODO get the registered baseDir for dotty later
		if err != nil {
			fmt.Println("error:", err)
		}

		homeDir, err := os.UserHomeDir()

		if err != nil {
			fmt.Println("error getting home directory:", err)
			return
		}

		if err := linkApp(app, baseDir, homeDir); err != nil {
			fmt.Println("error linking app:", err)
		} else {
			fmt.Println("Successfully linked", app)
		}

	},
}

func linkApp(appName, baseDir, homeDir string) error {
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

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		if err := os.Symlink(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to create symlink: %w", err)
		}

		return nil
	})

}
