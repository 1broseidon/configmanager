package configmanager

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// readmegen: "Describe how to instantiate the ConfigManager struct,
// and how to use its methods. Show a code snippet."

func (cm *ConfigManager) LoadFromFile(filename string, config ConfigLoader) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	err = config.Load(file)
	if err != nil {
		return fmt.Errorf("unsupported data format or failed to parse data")
	}

	cm.data = flatten(config.GetData(), "", make(map[string]interface{}))
	return nil
}

func (cm *ConfigManager) SaveToFile(filename string, config ConfigSaver) error {
	unflattened := unflatten(cm.data)
	dynamicConfig := &DynamicConfig{Data: unflattened, Filename: filename}

	data, err := dynamicConfig.Save()
	if err != nil {
		return fmt.Errorf("failed to save configuration to file %s: %w", filename, err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("warning: error closing file %s: %v\n", filename, closeErr)
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", filename, err)
	}

	return nil
}

func (cm *ConfigManager) UpdateKey(key string, value interface{}) error {
	if _, exists := cm.data[key]; !exists {
		return fmt.Errorf("key %s does not exist", key)
	}
	cm.data[key] = value
	return nil
}

func (cm *ConfigManager) UpdateKeys(updates map[string]interface{}) error {
	for k, v := range updates {
		if _, exists := cm.data[k]; !exists {
			return fmt.Errorf("key %s does not exist", k)
		}
		cm.data[k] = v
	}
	return nil
}

// Flatten a nested map
func flatten(data interface{}, prefix string, result map[string]interface{}) map[string]interface{} {
	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)

	switch rt.Kind() {
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			strKey := fmt.Sprint(key.Interface())
			newPrefix := strKey
			if prefix != "" {
				newPrefix = prefix + "." + strKey
			}
			flatten(rv.MapIndex(key).Interface(), newPrefix, result)
		}
	case reflect.Struct:
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			newPrefix := field.Name
			if prefix != "" {
				newPrefix = prefix + "." + field.Name
			}
			flatten(rv.Field(i).Interface(), newPrefix, result)
		}
	default:
		result[prefix] = data
	}
	return result
}

// Unflatten a map
func unflatten(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		keys := strings.Split(k, ".")
		m := result
		for i, key := range keys {
			if i == len(keys)-1 {
				m[key] = v
			} else {
				if _, ok := m[key]; !ok {
					m[key] = make(map[string]interface{})
				}
				m = m[key].(map[string]interface{})
			}
		}
	}
	return result
}
