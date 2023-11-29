//go:build linux
// +build linux

package configpath

import (
	"log"
	"os"
	"path"
)

// getSystemBasePath is a function that returns the base path for system configuration files.
// This base path is typically the "/etc" directory.
func getSystemBasePath() string {
	return "/etc" // return path for system configuration files
}

// getUserBasePath is a function that gets the HOME path of the current user via the HOME environment variable.
// It constructs a string that points to the directory where user-specific configuration files are stored.
// If the HOME environment variable is not set (empty), it falls back to the system base path, which means
// the path for system configuration files will be used for user configuration files, which could cause potential problems.
func getUserBasePath() string {
	// Get the value of the HOME environment variable
	homePath := os.Getenv("HOME")
	if homePath == "" {
		// If the HOME path is not defined (is empty), log an error message and fall back to the system base path
		log.Printf("ERRO: home path is empty, default path will be system path")
		return getSystemBasePath() // Return system base path, because we couldn't determine the user home path
	}
	// Join the HOME path with ".config" to get the path for user configuration files
	// It's a common convention to store user-specific configuration files in a ".config" directory in the user's HOME directory
	return path.Join(homePath, ".config")
}
