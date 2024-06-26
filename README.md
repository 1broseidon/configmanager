# configmanager-go

## Project Description

`configmanager-go` is a robust and extensible configuration management library for Go applications. It provides a unified interface for loading, accessing, modifying, and saving configuration data from various sources, including:

- **Configuration files:** Supports JSON, YAML, TOML, and INI formats.
- **Environment variables:** Allows overriding configuration values.

The library is designed with flexibility and ease of use in mind, making it simple to integrate into new or existing Go projects.

## Table of Contents

- [Project Description](#project-description)
- [Table of Contents](#table-of-contents)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [Project Structure](#project-structure)
- [Acknowledgements](#acknowledgements)

## Features

- **Support for multiple configuration formats:** JSON, YAML, TOML, INI
- **Dynamic format detection:** Automatically determine the format based on file extension.
- **Environment variable overrides:** Easily override configuration values using environment variables.
- **Data flattening and unflattening:** Simplifies access to nested configuration values.
- **Concurrency-safe access:** Uses mutexes to ensure thread safety when reading and writing configuration data.
- **Extensible design:** Supports custom configuration loaders and savers through interfaces.
- **Thoroughly tested:** Includes a comprehensive suite of unit tests for reliability.

## Installation

1. Make sure you have Go installed on your system.

2. Use `go get` to install the `configmanager-go` package:

   ```bash
   go get github.com/1broseidon/configmanager
   ```

3. Import the package into your Go project:

   ```go
   import "github.com/1broseidon/configmanager"
   ```

## Usage

Here's a basic example demonstrating how to load configuration from a TOML file:

```go
package main

import (
	"fmt"
	"log"

	"github.com/1broseidon/configmanager"
)

func main() {
	// Create a new ConfigManager
	cm := configmanager.New()

	// Load configuration from a TOML file
	if err := cm.LoadFromFile("config.toml"); err != nil {
		log.Fatal(err)
	}

	// Access configuration values
	databaseHost := cm.GetData()["database.host"].(string)
	serverPort := int(cm.GetData()["server.port"].(float64))

	fmt.Println("Database Host:", databaseHost)
	fmt.Println("Server Port:", serverPort)
}
```

**config.toml:**

```toml
[database]
host = "localhost"
port = 5432

[server]
host = "localhost"
port = 8080
```

For more detailed examples and advanced usage, please refer to the [documentation](link-to-documentation).

## Configuration

`configmanager-go` primarily relies on configuration files for loading settings. The library supports the following formats:

- JSON (.json)
- YAML (.yaml, .yml)
- TOML (.toml)
- INI (.ini)

You can specify the configuration file to load using the `LoadFromFile` method. The library will automatically determine the format based on the file extension.

**Environment Variable Overrides:**

You can override configuration values using environment variables. The library will look for environment variables in uppercase with underscores separating nested keys. For example, to override the `database.host` value, you would set the environment variable `DATABASE_HOST`.

## Contributing

We welcome contributions from the community! If you'd like to contribute to `configmanager-go`, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure the tests pass.
4. Submit a pull request.

**Code of Conduct:**

Please be respectful and considerate of others when contributing to the project. We follow the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct/).

## Project Structure

The project is structured as follows:

- **`configmanager.go`:** Contains the core `ConfigManager` struct and related functions.
- **`dynamicconfig.go`:** Implements dynamic configuration loading based on file extensions.
- **`formats/`:** Contains format-specific configuration loaders and savers (JSON, YAML, TOML, INI).
- **`internal/`:** Holds internal utility functions used by the library.
- **`tests/`:** Contains unit tests for the library.

## Acknowledgements

We would like to thank the developers of the following libraries for their contributions:

- `github.com/BurntSushi/toml`: TOML parsing and encoding.
- `gopkg.in/ini.v1`: INI file parsing.
- `gopkg.in/yaml.v2`: YAML encoding and decoding.

## License

This project is licensed under the MIT License License.
