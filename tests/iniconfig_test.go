package configmanager_test

import (
	"testing"

	"github.com/1broseidon/configmanager"
	"github.com/1broseidon/configmanager/formats"
	"github.com/1broseidon/configmanager/testutils" // Import the testutils package
)

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
	testutils.ResetConfigFile("../config/config.ini", iniConfig)

	cm := configmanager.New()
	config := &formats.INIConfig{}
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

	testutils.AssertConfig(t, expected, cm.GetData())
}
