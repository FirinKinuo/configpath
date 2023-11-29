//go:build windows
// +build windows

package configpath

import (
	"os"
)

// getSystemBasePath is a function that returns the base path for system-wide configuration files.
// On a typical Windows system, this would be the "ProgramData" directory.
func getSystemBasePath() string {
	// Return the value of the "PROGRAMDATA" environment variable, which points to the directory for system-wide configuration files
	return os.Getenv("PROGRAMDATA")
}

// getUserBasePath is a function that returns the base path for the current user's configuration files.
// On a typical Windows system, this would be the "AppData" directory within the user's profile.
func getUserBasePath() string {
	// Return the value of the "APPDATA" environment variable, which points to the directory for the current user's configuration files
	return os.Getenv("APPDATA")
}
