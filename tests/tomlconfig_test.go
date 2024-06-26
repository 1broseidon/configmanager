package configmanager_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/1broseidon/configmanager"
	"github.com/1broseidon/configmanager/formats"
	"github.com/1broseidon/configmanager/testutils"
)

// TestLoadTOMLConfig tests loading a TOML configuration file.
func TestLoadTOMLConfig(t *testing.T) {
	// Setup
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

	// Verify
	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     int64(5432),
		"server.host":       "localhost",
		"server.port":       int64(8080),
	}
	testutils.AssertConfig(t, expected, cm.GetData())
}

// TestSaveTOMLConfig tests saving and reloading a TOML configuration file.
func TestSaveTOMLConfig(t *testing.T) {
	// Setup
	originalConfig := []byte(`
[database]
user = "dbuser"
password = "dbpass"
host = "localhost"
port = 5432

[server]
host = "localhost"
port = 8080
`)
	testutils.ResetConfigFile("../config/config.toml", originalConfig)

	cm := configmanager.New()
	config := &formats.TOMLConfig{}
	err := cm.LoadFromFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error loading TOML config: %v", err)
	}

	// Update
	err = cm.UpdateKey("database.user", "saveduser")
	if err != nil {
		t.Fatalf("Error updating key database.user: %v", err)
	}
	err = cm.UpdateKey("database.password", "savedpass")
	if err != nil {
		t.Fatalf("Error updating key database.password: %v", err)
	}

	// Verify in-memory updates before saving
	updatedExpected := map[string]interface{}{
		"database.user":     "saveduser",
		"database.password": "savedpass",
		"database.host":     "localhost",
		"database.port":     int64(5432),
		"server.host":       "localhost",
		"server.port":       int64(8080),
	}
	testutils.AssertConfig(t, updatedExpected, cm.GetData())

	// Save configuration back to the TOML file
	err = cm.SaveToFile("../config/config.toml", config)
	if err != nil {
		t.Fatalf("Error saving TOML config: %v", err)
	}

	// Read back the saved file
	savedContent, err := os.ReadFile("../config/config.toml")
	if err != nil {
		t.Fatalf("Error reading saved TOML file: %v", err)
	}
	t.Logf("Saved TOML Content: %s", string(savedContent))

	// Reload the saved configuration
	newCm := configmanager.New()
	newConfig := &formats.TOMLConfig{}
	err = newCm.LoadFromFile("../config/config.toml", newConfig)
	if err != nil {
		t.Fatalf("Error loading saved TOML config: %v", err)
	}

	// Verify the reloaded configuration
	expected := map[string]interface{}{
		"database.user":     "saveduser",
		"database.password": "savedpass",
		"database.host":     "localhost",
		"database.port":     int64(5432),
		"server.host":       "localhost",
		"server.port":       int64(8080),
	}
	testutils.AssertConfig(t, expected, newCm.GetData())
}

// TestInvalidTOMLConfig tests loading an invalid TOML configuration file.
func TestInvalidTOMLConfig(t *testing.T) {
	// Setup
	invalidConfig := []byte(`
    not valid TOML data
    `)
	testutils.ResetConfigFile("../config/invalidconfig.toml", invalidConfig)

	cm := configmanager.New()
	config := &formats.TOMLConfig{}
	err := cm.LoadFromFile("../config/invalidconfig.toml", config)
	if err == nil {
		t.Fatalf("Expected error for invalid TOML config, got nil")
	}
	expectedErr := "failed to decode TOML data"
	if !strings.Contains(err.Error(), expectedErr) {
		t.Fatalf("Unexpected error message: got %v, want %v", err, expectedErr)
	}
}

// TestNoFileTOMLConfig tests loading a non-existent TOML configuration file.
func TestNoFileTOMLConfig(t *testing.T) {
	// Setup
	cm := configmanager.New()
	config := &formats.TOMLConfig{}
	err := cm.LoadFromFile("../config/nonexistent.toml", config)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Expected error for non-existing file, got: %v", err)
	}
}
