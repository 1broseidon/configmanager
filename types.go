package configmanager

type ConfigLoader interface {
	Load([]byte) error
	GetData() map[string]interface{}
}

type ConfigSaver interface {
	Save() ([]byte, error)
}

type ConfigManager struct {
	data map[string]interface{}
}

func New() *ConfigManager {
	return &ConfigManager{
		data: make(map[string]interface{}),
	}
}

type DynamicConfig struct {
	Data     map[string]interface{}
	Filename string
}

func (dc *DynamicConfig) GetData() map[string]interface{} {
	return dc.Data
}

func (cm *ConfigManager) GetData() map[string]interface{} {
	return cm.data
}
