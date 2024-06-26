package configmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/1broseidon/configmanager/internal"
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
	mu   sync.RWMutex
}

// New creates a new instance of ConfigManager.
func New() *ConfigManager {
	return &ConfigManager{
		data: make(map[string]interface{}),
	}
}

// GetData retrieves the configuration data from ConfigManager.
func (cm *ConfigManager) GetData() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.data
}

// LoadFromFile loads configuration data from a file, using DynamicConfig by default if no config loader is provided.
func (cm *ConfigManager) LoadFromFile(filename string, config ...ConfigLoader) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	var loader ConfigLoader
	if len(config) > 0 {
		loader = config[0]
	} else {
		loader = &DynamicConfig{Filename: filename}
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	err = loader.Load(file)
	if err != nil {
		return fmt.Errorf("unsupported data format or failed to parse data: %w", err)
	}

	cm.data = internal.Flatten(loader.GetData())
	return nil
}

// SaveToFile saves configuration data to a file, using DynamicConfig by default if no config saver is provided.
func (cm *ConfigManager) SaveToFile(filename string, config ...ConfigSaver) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	var saver ConfigSaver
	if len(config) > 0 {
		saver = config[0]
		// Update the config's data with the current ConfigManager data
		if loader, ok := saver.(ConfigLoader); ok {
			unflattenedData := internal.Unflatten(cm.data)
			unflattenedBytes, err := serializeData(filename, unflattenedData)
			if err != nil {
				return err
			}
			loader.Load(unflattenedBytes)
		}
	} else {
		saver = &DynamicConfig{Data: internal.Unflatten(cm.data), Filename: filename}
	}

	data, err := saver.Save()
	if err != nil {
		return fmt.Errorf("failed to save configuration to file %s: %w", filename, err)
	}

	// Debug: Print the data to be saved
	fmt.Printf("Saving Data: %s\n", string(data))

	// Write the data to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", filename, err)
	}

	return nil
}

// UpdateKey updates a specific key in the configuration.
func (cm *ConfigManager) UpdateKey(key string, value interface{}) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, exists := cm.data[key]; !exists {
		return fmt.Errorf("key %s does not exist", key)
	}
	cm.data[key] = value
	return nil
}

// UpdateKeys updates multiple keys in the configuration.
func (cm *ConfigManager) UpdateKeys(updates map[string]interface{}) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	for k, v := range updates {
		if _, exists := cm.data[k]; !exists {
			return fmt.Errorf("key %s does not exist", k)
		}
		cm.data[k] = v
	}
	return nil
}

// LoadEnvVariables loads configuration data from environment variables.
func (cm *ConfigManager) LoadEnvVariables(config *DynamicConfig) error {
	for key := range config.Data {
		envKey := strings.ToUpper(strings.Replace(key, ".", "_", -1))
		if value, exists := os.LookupEnv(envKey); exists {
			// Update both the ConfigManager's data and the DynamicConfig's Data
			cm.data[key] = value
			config.Data[key] = value
		}
	}
	return nil
}

// serializeData serializes unflattened data based on filename extension
func serializeData(filename string, data map[string]interface{}) ([]byte, error) {
	switch ext := filepath.Ext(filename); ext {
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
			sectionKey := strings.Split(k, ".")
			section, key := ini.DefaultSection, sectionKey[0]
			if len(sectionKey) > 1 {
				section, key = sectionKey[0], sectionKey[1]
			}
			cfg.Section(section).Key(key).SetValue(fmt.Sprint(v))
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
