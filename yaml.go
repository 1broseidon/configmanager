package configmanager

import "gopkg.in/yaml.v2"

type YAMLConfig struct {
	Data interface{}
}

func (yc *YAMLConfig) Load(data []byte) error {
	return yaml.Unmarshal(data, &yc.Data)
}

func (yc *YAMLConfig) Save() ([]byte, error) {
	return yaml.Marshal(yc.Data)
}
