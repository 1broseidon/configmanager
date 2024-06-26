package formats

import (
	"encoding/json"
	"fmt"

	"github.com/1broseidon/configmanager/internal"
)

// JSONConfig handles JSON configuration.
type JSONConfig struct {
	Data map[string]interface{}
}

// Load loads JSON configuration data.
func (jc *JSONConfig) Load(data []byte) error {
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}
	jc.Data = internal.Flatten(temp)
	return nil
}

// Save saves JSON configuration data.
func (jc *JSONConfig) Save() ([]byte, error) {
	data, err := json.MarshalIndent(internal.Unflatten(jc.Data), "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	return data, nil
}

// GetData retrieves the configuration data from JSONConfig.
func (jc *JSONConfig) GetData() map[string]interface{} {
	return jc.Data
}
