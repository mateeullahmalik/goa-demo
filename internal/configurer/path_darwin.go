// +build darwin

package configurer

import (
	"os"
	"path/filepath"
)

var defaultConfigPaths = []string{
	"$HOME/Library/Application Support/GoaDemo",
	".",
}

// DefaultPath returns the default config path for darwin OS.
func DefaultPath() string {
	homeDir, _ := os.UserConfigDir()
	return filepath.Join(homeDir, "App")
}
