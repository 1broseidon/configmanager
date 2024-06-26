package formats

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/1broseidon/configmanager/internal"
	"gopkg.in/ini.v1"
)

// INIConfig handles INI configuration.
type INIConfig struct {
	Data map[string]interface{}
}

// Load loads INI configuration data.
func (ic *INIConfig) Load(data []byte) error {
	cfg, err := ini.Load(data)
	if err != nil {
		return fmt.Errorf("failed to load INI data: %w", err)
	}

	temp := make(map[string]interface{})
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			if section.Name() == ini.DefaultSection {
				temp[key.Name()] = key.Value()
			} else {
				temp[fmt.Sprintf("%s.%s", section.Name(), key.Name())] = key.Value()
			}
		}
	}
	ic.Data = internal.Flatten(temp)
	return nil
}

// Save saves INI configuration data.
func (ic *INIConfig) Save() ([]byte, error) {
	cfg := ini.Empty()
	unflattenedData := internal.Unflatten(ic.Data)
	for k, v := range unflattenedData {
		sectionKey := strings.Split(k, ".")
		section, key := ini.DefaultSection, sectionKey[0]
		if len(sectionKey) > 1 {
			section, key = sectionKey[0], sectionKey[1]
		}
		cfg.Section(section).Key(key).SetValue(fmt.Sprint(v))
	}
	var buf bytes.Buffer
	if _, err := cfg.WriteTo(&buf); err != nil {
		return nil, fmt.Errorf("failed to write INI data: %w", err)
	}
	return buf.Bytes(), nil
}

// GetData retrieves the configuration data from INIConfig.
func (ic *INIConfig) GetData() map[string]interface{} {
	return ic.Data
}
