package configmanager_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/1broseidon/configmanager"
	"github.com/1broseidon/configmanager/formats"
	"github.com/1broseidon/configmanager/testutils"
)

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

	expectedErr := "key nonexistent.key1 does not exist"
	if err.Error() != expectedErr {
		t.Fatalf("Expected error `%s`, got `%s`", expectedErr, err.Error())
	}
}

func TestVariadicLoadFromFile(t *testing.T) {
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
	testutils.ResetConfigFile("../config/config.toml", tomlConfig)

	cm := configmanager.New()

	// Test default (DynamicConfig)
	err := cm.LoadFromFile("../config/config.toml")
	if err != nil {
		t.Fatalf("Error loading config with default DynamicConfig: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     int64(5432),
		"server.host":       "localhost",
		"server.port":       int64(8080),
	}
	testutils.AssertConfig(t, expected, cm.GetData())

	// Test with specific format (TOMLConfig)
	config := &formats.TOMLConfig{}
	err = cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	testutils.AssertConfig(t, expected, cm.GetData())
}

func TestVariadicSaveToFile(t *testing.T) {
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
	testutils.ResetConfigFile("../config/config.toml", tomlConfig)

	cm := configmanager.New()
	config := &formats.TOMLConfig{}
	err := cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	// Update a configuration key
	err = cm.UpdateKey("database.user", "newuser")
	if err != nil {
		t.Fatalf("Error updating key: %v", err)
	}

	// Test default (DynamicConfig)
	err = cm.SaveToFile("../config/config.toml")
	if err != nil {
		t.Fatalf("Error saving config with default DynamicConfig: %v", err)
	}

	// Read the saved content for verification
	savedContent, err := os.ReadFile("../config/config.toml")
	if err != nil {
		t.Fatalf("Failed to read saved configuration: %v", err)
	}
	fmt.Printf("Saved TOML Content: %s\n", string(savedContent))

	// Create a new ConfigManager to load the saved configuration
	newCm := configmanager.New()
	err = newCm.LoadFromFile("../config/config.toml", &formats.TOMLConfig{})
	if err != nil {
		t.Fatalf("Error loading saved config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "newuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     int64(5432),
		"server.host":       "localhost",
		"server.port":       int64(8080),
	}
	testutils.AssertConfig(t, expected, newCm.GetData())
}
