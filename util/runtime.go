package util

import (
	"runtime"
)

// const ...
const (
	Linux   string = "linux"
	Windows        = "windows"
)

// IsLinux ...
func IsLinux() bool {
	return runtime.GOOS == Linux
}

// IsWindows ...
func IsWindows() bool {
	return runtime.GOOS == Windows
}
