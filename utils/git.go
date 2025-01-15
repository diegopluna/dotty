package utils

import (
	"fmt"
	"os/exec"
)

func RunGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	basePath, err := GetBaseFolder()
	cmd.Dir = basePath
	if err != nil {
		return fmt.Errorf("error getting base folder: %s", err)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git command failed: %s\n%s", err, output)
	}
	fmt.Printf("git command output: %s\n", output)
	return nil
}
