package cmd

import (
	"dotty/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore [app]",
	Short: "Restore dotfiles from the Dotty folder to their original paths",
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

		sourcePath := filepath.Join(baseDir, app)

		info, err := os.Stat(sourcePath)
		if err != nil {
			fmt.Println("Error getting source path info:", err)
			return
		}

		if !info.IsDir() {
			fmt.Println("The source path is not a directory:", sourcePath)
			return
		}

		err = restoreMigration(sourcePath, baseDir, homeDir)
		if err != nil {
			fmt.Println("Error restoring migration:", err)
		} else {
			fmt.Printf("Successfully restored %s to its original location.\n", app)
		}
	},
}

func restoreMigration(src string, baseDir string, homeDir string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(baseDir, path)

		if err != nil {
			return err
		}

		destPath := filepath.Join(homeDir, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
		} else if info.Mode().IsRegular() {
			if err := utils.CopyFile(path, destPath); err != nil {
				return err
			}
		}

		return nil
	})
}
