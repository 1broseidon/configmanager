package configmanager

import (
	"fmt"
	"os"
	"strings"
)

func (cm *ConfigManager) LoadEnvVariables(config *DynamicConfig) error {
	for key := range config.Data {
		envKey := strings.ToUpper(strings.Replace(key, ".", "_", -1))
		if value, exists := os.LookupEnv(envKey); exists {
			// Update both the ConfigManager's data and the DynamicConfig's Data
			cm.data[key] = value
			config.Data[key] = value
		} else {
			return fmt.Errorf("environment variable %s does not exist", envKey)
		}
	}
	return nil
}
