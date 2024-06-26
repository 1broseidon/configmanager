package configmanager_test

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/1broseidon/configmanager"
)

func resetConfigFile(filename string, content []byte) {
	_ = os.WriteFile(filename, content, 0644)
}

func assertConfig(t *testing.T, expected, actual map[string]interface{}) {
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			t.Errorf("Key %v missing from actual config", key)
		} else if !reflect.DeepEqual(actualValue, expectedValue) {
			t.Errorf("For key %v, expected %v (%T), got %v (%T)", key, expectedValue, expectedValue, actualValue, actualValue)
		}
	}
}

func TestLoadTOMLConfig(t *testing.T) {
	tomlConfig := []byte(`
[database]
user = "dbuser"
password = "dbpass"
host = "localhost"
port = 5432

[server]
host = "localhost"
port = 8080
`)
	resetConfigFile("../config/config.toml", tomlConfig)

	cm := configmanager.New()
	config := &configmanager.TOMLConfig{}
	err := cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     int64(5432),
		"server.host":       "localhost",
		"server.port":       int64(8080),
	}

	assertConfig(t, expected, cm.GetData())
}

func TestLoadInvalidFile(t *testing.T) {
	cm := configmanager.New()
	config := &configmanager.DynamicConfig{}
	err := cm.LoadFromFile("../config/invalidfile.txt", config)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Expected error for non-existing file, got: %v", err)
	}
}

func TestLoadInvalidFormat(t *testing.T) {
	invalidConfig := []byte(`
- invalid config
`)
	resetConfigFile("../config/invalidconfig.txt", invalidConfig)

	cm := configmanager.New()
	config := &configmanager.DynamicConfig{}
	err := cm.LoadFromFile("../config/invalidconfig.txt", config)
	if err == nil {
		t.Fatalf("Expected error for invalid config format, got nil")
	}
	expectedErr := "unsupported data format or failed to parse data"
	if !strings.Contains(err.Error(), expectedErr) {
		t.Fatalf("Unexpected error message: got %v, want %v", err, expectedErr)
	}
}

func TestSaveToFileError(t *testing.T) {
	cm := configmanager.New()

	// Trying to save with an unsupported format
	cm.UpdateKey("key", "value")
	config := &configmanager.DynamicConfig{}
	err := cm.SaveToFile("../config/config.unsupported", config)
	if err == nil {
		t.Fatalf("Expected error for unsupported file format, got nil")
	}
}

func TestUpdateKeyError(t *testing.T) {
	cm := configmanager.New()
	err := cm.UpdateKey("nonexistent.key", "newvalue")
	if err == nil {
		t.Fatalf("Expected error for updating a non-existent key, got nil")
	}
}

func TestUpdateKeysError(t *testing.T) {
	cm := configmanager.New()
	updates := map[string]interface{}{
		"nonexistent.key1": "value1",
		"nonexistent.key2": "value2",
	}
	err := cm.UpdateKeys(updates)
	if err == nil {
		t.Fatalf("Expected error for updating non-existent keys, got nil")
	}
}

func TestLoadINIConfig(t *testing.T) {
	iniConfig := []byte(`
[database]
user = dbuser
password = dbpass
host = localhost
port = 5432

[server]
host = localhost
port = 8080
`)
	resetConfigFile("../config/config.ini", iniConfig)

	cm := configmanager.New()
	config := &configmanager.INIConfig{}
	err := cm.LoadFromFile("../config/config.ini", config)
	if err != nil {
		t.Fatalf("Error loading INI config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     "5432",
		"server.host":       "localhost",
		"server.port":       "8080",
	}

	assertConfig(t, expected, cm.GetData())
}

func TestLoadJSONConfig(t *testing.T) {
	jsonConfig := []byte(`
{
    "database": {
        "user": "dbuser",
        "password": "dbpass",
        "host": "localhost",
        "port": 5432
    },
    "server": {
        "host": "localhost",
        "port": 8080
    }
}
`)
	resetConfigFile("../config/config.json", jsonConfig)

	cm := configmanager.New()
	config := &configmanager.JSONConfig{}
	err := cm.LoadFromFile("../config/config.json", config)
	if err != nil {
		t.Fatalf("Error loading JSON config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     float64(5432),
		"server.host":       "localhost",
		"server.port":       float64(8080),
	}

	assertConfig(t, expected, cm.GetData())
}

func TestLoadYAMLConfig(t *testing.T) {
	yamlConfig := []byte(`
database:
  user: dbuser
  password: dbpass
  host: localhost
  port: 5432
server:
  host: localhost
  port: 8080
`)
	resetConfigFile("../config/config.yaml", yamlConfig)

	cm := configmanager.New()
	config := &configmanager.YAMLConfig{}
	err := cm.LoadFromFile("../config/config.yaml", config)
	if err != nil {
		t.Fatalf("Error loading YAML config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     5432,
		"server.host":       "localhost",
		"server.port":       8080,
	}

	assertConfig(t, expected, cm.GetData())
}

func TestUpdateKey(t *testing.T) {
	resetConfigFile("../config/config.toml", []byte(`
[database]
user = "dbuser"
password = "dbpass"
host = "localhost"
port = 5432

[server]
host = "localhost"
port = 8080
`))
	cm := configmanager.New()
	config := &configmanager.TOMLConfig{}
	err := cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	err = cm.UpdateKey("database.user", "newuser")
	if err != nil {
		t.Fatalf("Error updating key: %v", err)
	}

	expected := map[string]interface{}{
		"database.user": "newuser",
	}
	assertConfig(t, expected, cm.GetData())
}

func TestUpdateKeys(t *testing.T) {
	resetConfigFile("../config/config.toml", []byte(`
[database]
user = "dbuser"
password = "dbpass"
host = "localhost"
port = 5432

[server]
host = "localhost"
port = 8080
`))
	cm := configmanager.New()
	config := &configmanager.TOMLConfig{}
	err := cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	updates := map[string]interface{}{
		"database.password": "newpass",
		"server.port":       9090,
	}
	err = cm.UpdateKeys(updates)
	if err != nil {
		t.Fatalf("Error updating keys: %v", err)
	}

	expected := map[string]interface{}{
		"database.password": "newpass",
		"server.port":       9090,
	}
	assertConfig(t, expected, cm.GetData())
}

func TestSaveTOMLConfig(t *testing.T) {
	resetConfigFile("../config/config.toml", []byte(`
[database]
user = "dbuser"
password = "dbpass"
host = "localhost"
port = 5432

[server]
host = "localhost"
port = 8080
`))
	cm := configmanager.New()
	config := &configmanager.TOMLConfig{}
	err := cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	err = cm.UpdateKey("database.user", "saveduser")
	if err != nil {
		t.Fatalf("Error updating key: %v", err)
	}

	err = cm.UpdateKey("database.password", "savedpass")
	if err != nil {
		t.Fatalf("Error updating key: %v", err)
	}

	err = cm.SaveToFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error saving TOML config: %v", err)
	}

	newCm := configmanager.New()
	newConfig := &configmanager.TOMLConfig{}
	err = newCm.LoadFromFile("../config/config.toml", newConfig)
	if err != nil {
		t.Fatalf("Error loading saved TOML config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "saveduser",
		"database.password": "savedpass",
	}
	assertConfig(t, expected, newCm.GetData())
}

func TestLoadEnvVariables(t *testing.T) {
	os.Setenv("DATABASE_USER", "envuser")
	defer os.Unsetenv("DATABASE_USER")

	cm := configmanager.New()
	config := &configmanager.DynamicConfig{
		Data: map[string]interface{}{
			"database.user": "dbuser",
		},
	}

	err := cm.LoadEnvVariables(config)
	if err != nil {
		t.Fatalf("Error loading environment variables: %v", err)
	}

	expected := map[string]interface{}{
		"database.user": "envuser",
	}
	assertConfig(t, expected, config.Data)
}
