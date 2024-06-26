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

// ConfigLoader interface represents the ability to load configuration data.
type ConfigLoader interface {
	Load([]byte) error
	GetData() map[string]interface{}
}

// ConfigSaver interface represents the ability to save configuration data.
type ConfigSaver interface {
	Save() ([]byte, error)
}

// ConfigManager is the primary struct for managing configuration data.
type ConfigManager struct {
	data map[string]interface{}
}

// New creates a new instance of ConfigManager.
func New() *ConfigManager {
	return &ConfigManager{
		data: make(map[string]interface{}),
	}
}

// GetData retrieves the configuration data from ConfigManager.
func (cm *ConfigManager) GetData() map[string]interface{} {
	return cm.data
}

// DynamicConfig dynamically loads and saves configuration based on file extension.
type DynamicConfig struct {
	Data     map[string]interface{}
	Filename string
}

// Load dynamically loads configuration based on file extension.
func (dc *DynamicConfig) Load(data []byte) error {
	var temp map[string]interface{}
	var err error

	if err = json.Unmarshal(data, &temp); err == nil {
		dc.Data = temp
		return nil
	}
	if err = yaml.Unmarshal(data, &temp); err == nil {
		dc.Data = temp
		return nil
	}
	if _, err = toml.Decode(string(data), &temp); err == nil {
		dc.Data = temp
		return nil
	}
	cfg, iniErr := ini.Load(data)
	if iniErr == nil {
		temp = make(map[string]interface{})
		for _, section := range cfg.Sections() {
			for _, key := range section.Keys() {
				temp[key.Name()] = key.Value()
			}
		}
		dc.Data = temp
		return nil
	}
	return fmt.Errorf("unsupported data format or failed to parse data: %w", err)
}

// Save dynamically saves configuration based on file extension.
func (dc *DynamicConfig) Save() ([]byte, error) {
	ext := filepath.Ext(dc.Filename)

	switch ext {
	case ".json":
		return json.MarshalIndent(dc.Data, "", "  ")
	case ".yaml", ".yml":
		return yaml.Marshal(dc.Data)
	case ".toml":
		var buf bytes.Buffer
		if err := toml.NewEncoder(&buf).Encode(dc.Data); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	case ".ini":
		cfg := ini.Empty()
		for k, v := range dc.Data {
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

// TOMLConfig handles TOML configuration.
type TOMLConfig struct {
	Data map[string]interface{}
}

// Load loads TOML configuration data.
func (tc *TOMLConfig) Load(data []byte) error {
	if _, err := toml.Decode(string(data), &tc.Data); err != nil {
		return err
	}
	return nil
}

// Save saves TOML configuration data.
func (tc *TOMLConfig) Save() ([]byte, error) {
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(tc.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetData retrieves the configuration data from TOMLConfig.
func (tc *TOMLConfig) GetData() map[string]interface{} {
	return tc.Data
}

// JSONConfig handles JSON configuration.
type JSONConfig struct {
	Data map[string]interface{}
}

// Load loads JSON configuration data.
func (jc *JSONConfig) Load(data []byte) error {
	return json.Unmarshal(data, &jc.Data)
}

// Save saves JSON configuration data.
func (jc *JSONConfig) Save() ([]byte, error) {
	return json.MarshalIndent(jc.Data, "", "  ")
}

// GetData retrieves the configuration data from JSONConfig.
func (jc *JSONConfig) GetData() map[string]interface{} {
	return jc.Data
}

// YAMLConfig handles YAML configuration.
type YAMLConfig struct {
	Data map[string]interface{}
}

// Load loads YAML configuration data.
func (yc *YAMLConfig) Load(data []byte) error {
	return yaml.Unmarshal(data, &yc.Data)
}

// Save saves YAML configuration data.
func (yc *YAMLConfig) Save() ([]byte, error) {
	return yaml.Marshal(yc.Data)
}

// GetData retrieves the configuration data from YAMLConfig.
func (yc *YAMLConfig) GetData() map[string]interface{} {
	return yc.Data
}

// INIConfig handles INI configuration.
type INIConfig struct {
	Data map[string]interface{}
}

// Load loads INI configuration data.
func (ic *INIConfig) Load(data []byte) error {
	cfg, err := ini.Load(data)
	if err != nil {
		return err
	}

	ic.Data = make(map[string]interface{})
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			if section.Name() == ini.DefaultSection {
				ic.Data[key.Name()] = key.Value()
			} else {
				ic.Data[fmt.Sprintf("%s.%s", section.Name(), key.Name())] = key.Value()
			}
		}
	}
	return nil
}

// Save saves INI configuration data.
func (ic *INIConfig) Save() ([]byte, error) {
	cfg := ini.Empty()
	for k, v := range ic.Data {
		cfg.Section("").Key(k).SetValue(fmt.Sprint(v))
	}
	var buf bytes.Buffer
	if _, err := cfg.WriteTo(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetData retrieves the configuration data from INIConfig.
func (ic *INIConfig) GetData() map[string]interface{} {
	return ic.Data
}
