package configmanager

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type TOMLConfig struct {
	Data interface{}
}

func (tc *TOMLConfig) Load(data []byte) error {
	_, err := toml.Decode(string(data), tc.Data)
	return err
}

func (tc *TOMLConfig) Save() ([]byte, error) {
	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(tc.Data)
	return buf.Bytes(), err
}
