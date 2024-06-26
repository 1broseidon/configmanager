# configmanager

## Project Description

`configmanager` is a versatile and easy-to-use configuration management library for Go applications. It provides a unified interface for loading, accessing, modifying, and saving configuration data from various sources, including:

- **File-based Configurations:** JSON, YAML, TOML, INI
- **Environment Variables:** Overrides configuration with values from the environment.
- **Dynamic Loading:** Supports reloading configurations from files without application restarts.

This library simplifies the process of managing configuration settings in your Go projects, allowing you to focus on building core application logic.

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

- **Support for Multiple Formats:** Seamlessly load and save configurations in JSON, YAML, TOML, and INI formats.
- **Environment Variable Integration:** Override configuration values using environment variables for flexible deployments.
- **Dynamic Configuration Reloading:** Update configuration settings on-the-fly without restarting your application.
- **Simple and Intuitive API:** Easily integrate configuration management into your Go projects.
- **Extensible Design:** Add support for new configuration formats or sources through a well-defined interface.

## Installation

To use `configmanager` in your project, follow these simple steps:

1. **Install Go:** Ensure that you have Go installed on your system. If not, download and install it from the official [Go website](https://golang.org/).

2. **Get the Package:** Use the `go get` command to fetch the `configmanager` package:

   ```bash
   go get github.com/1broseidon/configmanager
   ```

3. **Import:** Import the package into your Go code:

   ```go
   import "github.com/1broseidon/configmanager"
   ```

## Usage

Here's a basic example demonstrating how to use `configmanager`:

```go
package main

import (
	"fmt"
	"github.com/1broseidon/configmanager"
)

func main() {
	// Create a new ConfigManager
	cm := configmanager.New()

	// Load configuration from a TOML file
	err := cm.LoadFromFile("config.toml", &configmanager.DynamicConfig{})
	if err != nil {
		panic(err)
	}

	// Access configuration values
	databaseHost := cm.GetString("database.host")
	serverPort := cm.GetInt("server.port")

	fmt.Println("Database Host:", databaseHost)
	fmt.Println("Server Port:", serverPort)

	// Update a configuration value
	cm.UpdateKey("server.port", 8081)

	// Save the updated configuration back to the file
	err = cm.SaveToFile("config.toml", &configmanager.DynamicConfig{})
	if err != nil {
		panic(err)
	}
}
```

## Configuration

`configmanager` supports loading configuration from various sources:

- **Configuration Files:** You can use JSON, YAML, TOML, or INI files to store your application's settings. See the `config` directory for sample configuration files in each format.
- **Environment Variables:** The library allows you to override configuration values using environment variables. Environment variable names are derived from configuration keys by converting them to uppercase and replacing dots (`.`) with underscores (`_`). For example, the configuration key `database.host` would be overridden by the environment variable `DATABASE_HOST`.

## Contributing

We welcome contributions from the community! If you'd like to contribute to `configmanager`, please follow these guidelines:

1. **Fork the Repository:** Fork the project repository to your own GitHub account.

2. **Create a Branch:** Create a new branch for your feature or bug fix. Use a descriptive branch name that reflects the changes you're making.

3. **Make Your Changes:** Implement your changes, ensuring to follow Go coding conventions and write clear, concise code.

4. **Write Tests:** Add appropriate unit tests to cover your changes and ensure the library's functionality remains intact.

5. **Submit a Pull Request:** Once your changes are ready, submit a pull request to the main repository. Clearly describe your changes and the problem they solve.

6. **Code of Conduct:** Please adhere to our [Code of Conduct](CODE_OF_CONDUCT.md) when contributing to the project.

## Project Structure

The project has the following directory structure:

```
configmanager/
├── config/             # Sample configuration files
│   ├── config.json
│   ├── config.toml
│   ├── config.yaml
│   └── config.ini
├── configmanager.go    # Core ConfigManager implementation
├── dynamic.go          # Dynamic configuration loading
├── env.go              # Environment variable loading
├── ini.go              # INI format support
├── json.go             # JSON format support
├── configmanager_test.go # Unit tests
├── toml.go             # TOML format support
├── types.go            # Core types and interfaces
└── yaml.go             # YAML format support
```

## Acknowledgements

We'd like to express our gratitude to the developers of the following libraries, which `configmanager` relies on:

- `github.com/BurntSushi/toml`: For TOML parsing and encoding.
- `gopkg.in/ini.v1`: For INI file parsing and writing.
- `gopkg.in/yaml.v2`: For YAML encoding and decoding.
- The Go standard library: For JSON handling, file I/O, and other core functionalities.

## License

This project is licensed under the MIT License License.
