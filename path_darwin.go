//go:build darwin
// +build darwin

package configpath

import (
	"os"
	"path"
)

// getSystemBasePath is a function that returns the base path for system configuration files.
// This base path is typically the "/Library" directory.
func getSystemBasePath() string {
	return "/Library" // return path for system configuration files
}

// getUserBasePath is a function that gets the HOME path of the current user via the HOME environment variable.
// It constructs a string that points to the directory where user-specific configuration files are stored.
func getUserBasePath() string {
	// Get the value of the HOME environment variable
	homePath := os.Getenv("HOME")
	// Join the HOME path with ".config" to get the path for user configuration files
	// It's a common convention to store user-specific configuration files in a ".config" directory in the user's HOME directory
	return path.Join(homePath, ".config")
}
