package configmanager

import (
	"bytes"
	"fmt"

	"gopkg.in/ini.v1"
)

type INIConfig struct {
	Data map[string]interface{}
}

func (ic *INIConfig) Load(data []byte) error {
	cfg, err := ini.Load(data)
	if err != nil {
		return err
	}

	ic.Data = make(map[string]interface{})
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			if section.Name() == ini.DefaultSection {
				ic.Data[key.Name()] = key.Value()
			} else {
				ic.Data[fmt.Sprintf("%s.%s", section.Name(), key.Name())] = key.Value()
			}
		}
	}
	return nil
}

func (ic *INIConfig) Save() ([]byte, error) {
	cfg := ini.Empty()
	for k, v := range ic.Data {
		cfg.Section("").Key(k).SetValue(fmt.Sprint(v))
	}

	var buf bytes.Buffer
	_, err := cfg.WriteTo(&buf)
	return buf.Bytes(), err
}

func (ic *INIConfig) GetData() map[string]interface{} {
	return ic.Data
}
