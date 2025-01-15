package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetBaseFolder() (string, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".dottyconfig")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	baseFolder := strings.TrimSpace(string(data))
	if baseFolder == "" {
		return "", fmt.Errorf("base folder is not set in the config file")
	}

	return baseFolder, nil
}
