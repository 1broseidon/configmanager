package formats

import (
	"bytes"
	"fmt"

	"github.com/1broseidon/configmanager/internal"
	"github.com/BurntSushi/toml"
)

// TOMLConfig handles TOML configuration.
type TOMLConfig struct {
	Data map[string]interface{}
}

// Load loads TOML configuration data.
func (tc *TOMLConfig) Load(data []byte) error {
	var temp map[string]interface{}
	if _, err := toml.Decode(string(data), &temp); err != nil {
		return fmt.Errorf("failed to decode TOML data: %w", err)
	}
	tc.Data = internal.Flatten(temp)
	return nil
}

// Save saves TOML configuration data.
func (tc *TOMLConfig) Save() ([]byte, error) {
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(internal.Unflatten(tc.Data)); err != nil {
		return nil, fmt.Errorf("failed to encode TOML data: %w", err)
	}
	return buf.Bytes(), nil
}

// GetData retrieves the configuration data from TOMLConfig.
func (tc *TOMLConfig) GetData() map[string]interface{} {
	return tc.Data
}
