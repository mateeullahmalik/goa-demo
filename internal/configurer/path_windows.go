// +build windows

package configurer

import (
	"os"
	"path"
	"path/filepath"
	"syscall"
)

const (
	beforeVistaAppDir = "Application Data"
	sinceVistaAppDir  = "AppData/Roaming"
)

var defaultConfigPaths = []string{
	path.Join("$HOME", beforeVistaAppDir, "goa-demo"),
	path.Join("$HOME", sinceVistaAppDir, "goa-demo"),
	".",
}

// DefaultPath returns the default config path for Windows OS.
func DefaultPath() string {
	homeDir, _ := os.UserHomeDir()
	appDir := beforeVistaAppDir

	v, _ := syscall.GetVersion()
	if v&0xff > 5 {
		appDir = sinceVistaAppDir
	}
	return filepath.Join(homeDir, filepath.FromSlash(appDir), "goa-demo")
}
