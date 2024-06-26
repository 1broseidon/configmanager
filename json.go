package configmanager

import "encoding/json"

type JSONConfig struct {
	Data interface{}
}

func (jc *JSONConfig) Load(data []byte) error {
	return json.Unmarshal(data, &jc.Data)
}

func (jc *JSONConfig) Save() ([]byte, error) {
	return json.Marshal(jc.Data)
}
