package testutils

import (
	"os"
	"reflect"
	"testing"
)

// ResetConfigFile resets the content of a file for testing.
func ResetConfigFile(filename string, content []byte) {
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		panic(err)
	}
}

// AssertConfig checks if the expected configuration matches the actual configuration.
func AssertConfig(t *testing.T, expected, actual map[string]interface{}) {
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			t.Errorf("Key %v missing from actual config", key)
		} else if !reflect.DeepEqual(actualValue, expectedValue) {
			t.Errorf("For key %v, expected %v (%T), got %v (%T)", key, expectedValue, expectedValue, actualValue, actualValue)
		}
	}
}
