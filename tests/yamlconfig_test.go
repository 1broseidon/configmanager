package configmanager_test

import (
	"testing"

	"github.com/1broseidon/configmanager"
	"github.com/1broseidon/configmanager/formats"
	"github.com/1broseidon/configmanager/testutils"
)

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

	testutils.ResetConfigFile("../config/config.yaml", yamlConfig)

	cm := configmanager.New()
	config := &formats.YAMLConfig{}
	err := cm.LoadFromFile("../config/config.yaml", config)
	if err != nil {
		t.Fatalf("Error loading YAML config: %v", err)
	}

	expected := map[string]interface{}{
		"database.user":     "dbuser",
		"database.password": "dbpass",
		"database.host":     "localhost",
		"database.port":     5432, // YAML unmarshaled numeric values should be of type int
		"server.host":       "localhost",
		"server.port":       8080,
	}

	testutils.AssertConfig(t, expected, cm.GetData())
}
