# Configuration Path Package (configpath)

[![GoDoc Widget]][GoDoc]

`configpath` package is designed to serve as an Operating System (OS) - independent way to determine appropriate paths
for configuration folders and files. It supports different paths depending on the OS (Unix, Darwin, Windows) and follows
the conventions specific to each system.

The package provides `ConfigPath` type which is a structure containing all the necessary path details both on a user and
system level. This type comprises two strings - `Vendor` and `Application`, these are used in constructing the
generated paths.

## Example

```go
package main

import "github.com/FirinKinuo/configpath"

func main() {
	configPath := configpath.ConfigPath{
		Vendor:      "TestVendor",
		Application: "TestApplication",
	}

	// /home/user/TestVendor/TestApplication/test.conf
	userConfigPath := configPath.UserFile("test.conf")
	// /etc/TestVendor/TestApplication/test.conf
	systemConfigPath := configPath.SystemFile("test.conf")
}
```

## Typical paths for OS

- [Linux](#specific-paths-for-linux-systems)
- [Windows](#specific-paths-for-windows-systems)
- [macOS](#specific-paths-for-macos-darwin-systems)
- [Other](#specific-paths-for-undefined-operating-systems)
- [Custom](#custom-defined-paths)

### Specific Paths for Linux Systems

On Linux systems, the configpath package uses the conventions generally expected.

#### System Base Path

System level configuration files are typically stored in the `/etc` directory.

Example: `/etc/vendor/application`

#### User Base Path

For user level configurations, typically a .config directory under the user's home directory is used.

Example: `/home/user/.config/vendor/application`

> Please note that if the HOME environment variable isn't set,
> the package will default to using system level configuration paths,
> which may lead to potential issues.
> The package does not create directories or files,
> but only helps in constructing the appropriate paths.
> Always ensure to handle necessary rights and checks when your program interacts with these paths.

### Specific Paths for Windows Systems

On Windows, the configpath package respects the Windows specific conventions for storing user specific and system level
configurations

#### System Base Path

The system level configuration files are typically stored in the ProgramData directory.

Example: `C:\ProgramData\Vendor\Application`

#### User Base Path

For user specific configurations, the AppData directory within the user's profile directory is typically used.

Example: `C:\Users\User\AppData\Roaming\Vendor\Application`

### Specific Paths for macOS (Darwin) Systems

On macOS (Darwin), the configpath package follows the macOS specific conventions.

#### System Base Path

System level configuration files are typically stored in the /Library directory.

Example: `/Library/Vendor/Application`

#### User Base Path

For user level configurations, typically a .config directory under the user's home directory is used.

Example: `/Users/user/Vendor/Application`

### Specific Paths for Undefined Operating Systems

In those rare instances where the operating system is not detected as one of the three major ones (Windows, Unix/Linux,
or macOS),
the configpath package defaults to specific values.

#### System Base Path

The system configuration files are stored in the /etc directory, similar to Unix conventions.

Example: `/etc/vendor/application`

#### User Base Path

User specific configuration files are expected in the current execution directory of the concerning application.
This is achieved by using a relative path "./".

Example: `./vendor/application`

This makes configpath flexible and applicable even on unconventional or lesser-known OS setups.

### Custom Defined Paths

While configpath is designed to provide an OS-agnostic way to handle paths,
there may be times when you want to use a custom path as the base for your application's configuration files.

```go
package main

import "github.com/FirinKinuo/configpath"

func main() {
	cp := configpath.ConfigPath{
		Application:      "TestApplication",
		CustomFolderPath: "/custom/path",
	}

	cp.CustomFolderFile("test.conf") // /custom/path/TestApplication/test.conf
}
```

## License

Copyright (c) [Firin Kinuo](https://fkinuo.ru)

Licensed under [MIT License](./LICENSE)

[GoDoc]: https://pkg.go.dev/github.com/FirinKinuo/configpath

[GoDoc Widget]: https://godoc.org/github.com/FirinKinuo/configpath?status.svg