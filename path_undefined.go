//go:build !(windows || linux || darwin)
// +build !windows,!linux,!darwin

package configpath

// getSystemBasePath is a function that returns the base path for system configuration files.
// This base path is typically the "/etc" directory.
func getSystemBasePath() string {
	return "/etc" // return path for system configuration files
}

// getUserBasePath is a function that provides the user-specific base path.
// A relative path "./" is used to denote current execution directory.
func getUserBasePath() string {
	return "./"
}
