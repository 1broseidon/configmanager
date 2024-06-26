# ConfigManager-Go

[![Go Report Card](https://goreportcard.com/badge/github.com/1broseidon/configmanager)](https://goreportcard.com/report/github.com/1broseidon/configmanager)
[![GoDoc](https://godoc.org/github.com/1broseidon/configmanager?status.svg)](https://godoc.org/github.com/1broseidon/configmanager)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents

- [Project Description](#project-description)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [Project Structure](#project-structure)
- [Acknowledgements](#acknowledgements)

## Project Description

`configmanager-go` is a robust and flexible configuration management library for Go applications. It provides a unified interface for loading configuration data from various sources, including:

- JSON, YAML, TOML, and INI files
- Environment variables

The library prioritizes ease of use, flexibility, and robust error handling. It is designed to simplify the process of managing application settings, allowing developers to focus on core application logic.

## Features

- **Support for multiple configuration formats:** Seamlessly load configuration from JSON, YAML, TOML, and INI files.
- **Environment variable overrides:** Override file-based configurations with environment variables for dynamic adjustments.
- **Dynamic format detection:** Automatically determine the configuration format based on file extensions.
- **Flattened data representation:** Access nested configuration values easily using dot-separated keys.
- **Robust error handling:** Provides clear error messages for common configuration issues.
- **Extensible design:** Easily add support for new configuration formats.
- **Thoroughly tested:** Includes a comprehensive suite of unit tests to ensure reliability.

## Installation

```bash
go get github.com/1broseidon/configmanager
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/1broseidon/configmanager"
)

func main() {
	// Create a new ConfigManager instance
	cm := configmanager.New()

	// Load configuration from a file (e.g., config.toml)
	err := cm.LoadFromFile("config.toml")
	if err != nil {
		panic(err)
	}

	// Access configuration values
	databaseHost := cm.GetData()["database.host"].(string)
	serverPort := cm.GetData()["server.port"].(int)

	fmt.Println("Database Host:", databaseHost)
	fmt.Println("Server Port:", serverPort)
}
```

## Configuration

### Configuration File Example (`config.toml`):

```toml
[database]
host = "localhost"
port = 5432
user = "myuser"
password = "mypassword"

[server]
host = "0.0.0.0"
port = 8080
```

### Environment Variable Overrides:

Environment variables can be used to override configuration values loaded from files. The environment variable names should follow a specific pattern:

```
CONFIG_SECTION_KEY=value
```

For example, to override the `host` value in the `database` section, you would use the following environment variable:

```
CONFIG_DATABASE_HOST=my-database-host
```

## Contributing

We welcome contributions from the community! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute code, report issues, and suggest enhancements.

### Code of Conduct

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Project Structure

```
configmanager-go/
├── config/                 # Sample configuration files
│   ├── config.json
│   ├── config.toml
│   ├── config.yaml
│   ├── invalidconfig.toml
│   └── invalidconfig.txt
├── formats/                # Format-specific configuration loaders and savers
│   ├── jsonconfig.go
│   ├── tomlconfig.go
│   └── yamlconfig.go
├── internal/               # Internal utility functions
│   └── flatten.go
├── configmanager.go         # Core configuration manager implementation
├── dynamicconfig.go        # Dynamic configuration loading logic
├── iniconfig.go            # INI configuration handler
└── tests/                   # Unit tests
    ├── configmanager_test.go
    ├── iniconfig_test.go
    ├── jsonconfig_test.go
    ├── testutils/
    │   └── utils.go
    ├── tomlconfig_test.go
    └── yamlconfig_test.go
```

## Acknowledgements

- This project utilizes the following excellent third-party libraries:
  - `github.com/BurntSushi/toml`: For TOML parsing and encoding.
  - `gopkg.in/ini.v1`: For INI file handling.
  - `gopkg.in/yaml.v2`: For YAML parsing and encoding.
- Thanks to all contributors who have helped make this project possible!

## License

This project is licensed under the MIT License License.
