// Package configpath is designed to provide an OS-independent method
// for determining the appropriate paths for configuration folders and files.

package configpath

import (
	"path/filepath"
)

// ConfigPath acts as a struct for user-specific and system-specific configuration path details.
type ConfigPath struct {
	Vendor           string
	Application      string
	CustomFolderPath string
}

// UserFolder generates a user-specific configuration folder path
func (c ConfigPath) UserFolder() string {
	return c.folderPath(getUserBasePath)
}

// SystemFolder generates a system-wide configuration folder path.
// It works similar to UserFolder, but for system-wide configuration.
func (c ConfigPath) SystemFolder() string {
	return c.folderPath(getSystemBasePath)
}

// CustomFolder generates a custom folder path.
func (c ConfigPath) CustomFolder() string {
	return c.folderPath(func() string {
		return c.CustomFolderPath
	})
}

func (c ConfigPath) folderPath(getBasePath func() string) string {
	basePath := getBasePath()
	return filepath.Join(basePath, c.Vendor, c.Application)
}

// UserFile returns a file-specific user path based on a relative path.
// It is useful for locating a specific file in user-specific folder.
func (c ConfigPath) UserFile(relativePath string) string {
	// Join the user configuration path with the relative path.
	return filepath.Join(c.UserFolder(), relativePath)
}

// SystemFile returns a file-specific system path based on a relative path.
// This is a system-wide equivalent of UserFile function.
func (c ConfigPath) SystemFile(relativePath string) string {
	// Join the system configuration path with the relative path.
	return filepath.Join(c.SystemFolder(), relativePath)
}

// CustomFolderFile returns a file-specific system path based on a custom path.
func (c ConfigPath) CustomFolderFile(relativePath string) string {
	// Join the system configuration path with the relative path.
	return filepath.Join(c.CustomFolder(), relativePath)
}
