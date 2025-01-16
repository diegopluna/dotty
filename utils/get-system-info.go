package utils

import (
	"fmt"
	"runtime"
)

func GetSystemInfo() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	return fmt.Sprintf("%s_%s", os, arch)
}
