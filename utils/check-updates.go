package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const GithubAPIURL = "https://api.github.com/repos/diegopluna/dotty/releases/latest"

func CheckForUpdates() (string, string, error) {
	resp, err := http.Get(GithubAPIURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", "", fmt.Errorf("failed to fetch release info: %s", resp.Status)
	}

	var release struct {
		TagName string `json:"tag_name"`
		Assets  []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", "", err
	}

	system := GetSystemInfo()

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, system) && strings.HasSuffix(asset.Name, ".tar.gz") {
			return release.TagName, asset.BrowserDownloadURL, nil
		}
	}

	return "", "", fmt.Errorf("no release found for %s", system)
}
