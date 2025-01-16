package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func DownloadAndUpdate(downloadURL string) error {
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmpFile, err := os.Create("/tmp/dotty-upgrade.tar.gz")
	if err != nil {
		return err
	}

	defer tmpFile.Close()

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return err
	}

	if err := ExtractTarGz("/tmp/dotty-upgrade.tar.gz", "/tmp/dotty-upgrade/"); err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	binaryPath := filepath.Join("/tmp/dotty-upgrade", "dotty")
	if err := os.Chmod(binaryPath, 0755); err != nil {
		return err
	}

	dottyInstallPath := filepath.Join(os.Getenv("HOME"), ".local", "bin", "dotty")
	if err := exec.Command("mv", binaryPath, dottyInstallPath).Run(); err != nil {
		return err
	}

	return nil
}
