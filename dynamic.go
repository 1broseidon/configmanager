package configmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

func (dc *DynamicConfig) Load(data []byte) error {
	var temp map[string]interface{}
	var err error

	// Try JSON
	if err = json.Unmarshal(data, &temp); err == nil {
		dc.Data = temp
		return nil
	}
	fmt.Printf("failed to unmarshal JSON data: %v\n", err)

	// Reset temp for next attempt
	temp = nil

	// Try YAML
	if err = yaml.Unmarshal(data, &temp); err == nil {
		dc.Data = temp
		return nil
	}
	fmt.Printf("failed to unmarshal YAML data: %v\n", err)

	// Reset temp for next attempt
	temp = nil

	// Try TOML
	if _, err = toml.Decode(string(data), &temp); err == nil {
		dc.Data = temp
		return nil
	}
	fmt.Printf("failed to decode TOML data: %v\n", err)

	// Reset temp for next attempt
	temp = nil

	// Try INI
	var cfg *ini.File
	if cfg, err = ini.Load(data); err == nil {
		temp = make(map[string]interface{})
		for _, section := range cfg.Sections() {
			for _, key := range section.Keys() {
				sectionName := section.Name()
				if sectionName == ini.DefaultSection {
					temp[key.Name()] = key.Value()
				} else {
					temp[fmt.Sprintf("%s.%s", sectionName, key.Name())] = key.Value()
				}
			}
		}
		dc.Data = temp
		return nil
	}
	fmt.Printf("failed to load INI data: %v\n", err)

	return fmt.Errorf("unsupported data format or failed to parse data")
}

func (dc *DynamicConfig) Save() ([]byte, error) {
	ext := filepath.Ext(dc.Filename)

	switch ext {
	case ".json":
		data, err := json.MarshalIndent(dc.Data, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON data: %w", err)
		}
		return data, nil
	case ".yaml", ".yml":
		data, err := yaml.Marshal(dc.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal YAML data: %w", err)
		}
		return data, nil
	case ".toml":
		var buf bytes.Buffer
		if err := toml.NewEncoder(&buf).Encode(dc.Data); err != nil {
			return nil, fmt.Errorf("failed to encode TOML data: %w", err)
		}
		return buf.Bytes(), nil
	case ".ini":
		cfg := ini.Empty()
		for k, v := range dc.Data {
			cfg.Section("").Key(k).SetValue(fmt.Sprint(v))
		}
		var buf bytes.Buffer
		if _, err := cfg.WriteTo(&buf); err != nil {
			return nil, fmt.Errorf("failed to write INI data: %w", err)
		}
		return buf.Bytes(), nil
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
}
