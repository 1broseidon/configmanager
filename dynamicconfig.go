package configmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/1broseidon/configmanager/internal"
	"github.com/BurntSushi/toml"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

// DynamicConfig dynamically loads and saves configuration based on file extension.
type DynamicConfig struct {
	Data     map[string]interface{}
	Filename string
}

// Load dynamically loads configuration based on file extension.
func (dc *DynamicConfig) Load(data []byte) error {
	var temp map[string]interface{}
	var err error

	switch ext := filepath.Ext(dc.Filename); ext {
	case ".json":
		err = json.Unmarshal(data, &temp)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &temp)
	case ".toml":
		_, err = toml.Decode(string(data), &temp)
	case ".ini":
		cfg, iniErr := ini.Load(data)
		if iniErr == nil {
			temp = map[string]interface{}{}
			for _, section := range cfg.Sections() {
				for _, key := range section.Keys() {
					if section.Name() == ini.DefaultSection {
						temp[key.Name()] = key.Value()
					} else {
						temp[fmt.Sprintf("%s.%s", section.Name(), key.Name())] = key.Value()
					}
				}
			}
			err = nil
		} else {
			err = fmt.Errorf("unsupported data format or failed to parse data: %w", iniErr)
		}
	default:
		err = fmt.Errorf("unsupported file format")
	}

	if err != nil {
		return fmt.Errorf("failed to load data: %w", err)
	}

	// Flatten the loaded configuration data
	dc.Data = internal.Flatten(temp)

	return nil
}

// Save dynamically saves configuration based on file extension.
func (dc *DynamicConfig) Save() ([]byte, error) {
	// Unflatten the data
	data := internal.Unflatten(dc.Data)
	fmt.Printf("Saving Data: %+v\n", data)

	switch ext := filepath.Ext(dc.Filename); ext {
	case ".json":
		return json.MarshalIndent(data, "", "  ")
	case ".yaml", ".yml":
		return yaml.Marshal(data)
	case ".toml":
		var buf bytes.Buffer
		if err := toml.NewEncoder(&buf).Encode(data); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	case ".ini":
		cfg := ini.Empty()
		for k, v := range data {
			cfg.Section("").Key(k).SetValue(fmt.Sprint(v))
		}
		var buf bytes.Buffer
		if _, err := cfg.WriteTo(&buf); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	default:
		return nil, fmt.Errorf("unsupported file format")
	}
}

// GetData retrieves the configuration data from DynamicConfig.
func (dc *DynamicConfig) GetData() map[string]interface{} {
	return dc.Data
}
