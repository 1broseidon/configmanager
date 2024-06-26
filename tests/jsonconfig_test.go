package configmanager_test

import (
	"testing"

	"github.com/1broseidon/configmanager"
	"github.com/1broseidon/configmanager/formats"
	"github.com/1broseidon/configmanager/testutils" // Import the testutils package
)

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
	testutils.ResetConfigFile("../config/config.json", jsonConfig)

	cm := configmanager.New()
	config := &formats.JSONConfig{}
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

	testutils.AssertConfig(t, expected, cm.GetData())
}
