# configmanager-go

## Project Description

configmanager-go is an open-source project that provides a flexible and powerful configuration management system for Go applications. It offers support for various configuration file formats such as JSON, TOML, YAML, and INI, as well as environment variables. With configmanager-go, developers can easily load, save, and update configuration data, making their applications more dynamic and adaptable.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [Project Structure](#project-structure)
- [Acknowledgements](#acknowledgements)

## Features

- Support for multiple configuration file formats: JSON, TOML, YAML, and INI.
- Dynamic loading and saving of configuration data.
- Environment variable support for runtime configuration updates.
- Flattening and unflattening of nested configuration data.
- Easy-to-use API for managing configuration options.

## Installation

To install configmanager-go, you need to have Go installed on your system. You can then use `go get` to install the package:

```bash
go get github.com/1broseidon/configmanager-go
```

Alternatively, you can add it to your Go module's `go.mod` file:

```go
require (
    "github.com/1broseidon/configmanager-go" v1.0.0
)
```

## Usage

Here's an example of how to use configmanager-go to load and update configuration data:

```go
package main

import (
    "fmt"

    "github.com/1broseidon/configmanager-go/pkg/configmanager"
)

func main() {
    // Create a new ConfigManager instance
    cm, err := configmanager.New("config.toml")
    if err != nil {
        panic(err)
    }

    // Load configuration data from file
    err = cm.LoadFromFile("config.toml")
    if err != nil {
        panic(err)
    }

    // Update a configuration value
    cm.UpdateKey("server.host", "localhost")

    // Save updated configuration to file
    err = cm.SaveToFile("updated_config.toml")
    if err != nil {
        panic(err)
    }

    // Access configuration data
    serverHost, ok := cm.GetData()["server.host"]
    if !ok {
        panic("server host not found in configuration")
    }
    fmt.Println("Server host:", serverHost)
}
```

## Configuration

configmanager-go provides several configuration options that can be customized:

- `ConfigFile`: The path to the configuration file to be loaded.
- `EnvironmentPrefix`: A prefix to use for environment variables. This can be useful to avoid conflicts with other environment variables.
- `Flatten`: A boolean flag indicating whether to flatten the configuration data or keep it nested.

## Contributing

Contributions are welcome! Please refer to the [Contributing Guidelines](CONTRIBUTING.md) for information on how to propose changes and submit patches. We value and appreciate all contributions, and we're happy to help with any issues or questions you may have.

## Project Structure

The project is organized into several directories:

- `cmd`: Contains the main package and the entry point for the application.
- `config`: Stores the configuration files in different formats.
- `pkg`: Houses the main logic of the project, including the `configmanager` package.
- `tests`: Includes test suites to ensure the proper functioning of the package.

## Acknowledgements

We would like to thank the open-source community for their contributions and support. Special thanks to the developers of the external libraries used in this project:

- BurntSushi for the TOML library: [https://github.com/BurntSushi/toml](https://github.com/BurntSushi/toml)
- gopkg.in for the INI and YAML libraries: [https://gopkg.in/ini.v1](https://gopkg.in/ini.v1), [https://gopkg.in/yaml.v2](https://gopkg.in/yaml.v2)

We also appreciate the feedback and suggestions from our users, which help us improve configmanager-go and make it more useful for the Go community.

## License

This project is licensed under the MIT License License.
