package configmanager_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/1broseidon/configmanager"
	"github.com/1broseidon/configmanager/testutils" // Import the testutils package
)

func TestLoadInvalidFile(t *testing.T) {
	cm := configmanager.New()
	err := cm.LoadFromFile("../config/invalidfile.txt")
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Expected error for non-existing file, got: %v", err)
	}
}

func TestLoadInvalidFormat(t *testing.T) {
	invalidConfig := []byte(`
- invalid config
`)
	testutils.ResetConfigFile("../config/invalidconfig.txt", invalidConfig)

	cm := configmanager.New()
	err := cm.LoadFromFile("../config/invalidconfig.txt")
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
	err := cm.SaveToFile("../config/config.unsupported")
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
