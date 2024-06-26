package formats

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/1broseidon/configmanager/internal"
)

// YAMLConfig handles YAML configuration.
type YAMLConfig struct {
	Data map[string]interface{}
}

// Load loads YAML configuration data.
func (yc *YAMLConfig) Load(data []byte) error {
	var temp map[string]interface{}
	if err := yaml.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal YAML data: %w", err)
	}

	yc.Data = internal.Flatten(temp)

	return nil
}

// Save saves YAML configuration data.
func (yc *YAMLConfig) Save() ([]byte, error) {
	unflattenedData := internal.Unflatten(yc.Data)
	data, err := yaml.Marshal(unflattenedData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal YAML data: %w", err)
	}
	return data, nil
}

// GetData retrieves the configuration data from YAMLConfig.
func (yc *YAMLConfig) GetData() map[string]interface{} {
	return yc.Data
}
