package config

import (
	"github.com/go-yaml/yaml"
)

type Configuration struct {
	Database string `yaml:"database"` // 数据库类型，决定Schema和驱动的位置，支持sqlite3和mysql
	DSN      string `yaml:"dsn"`      // 数据库DSN
	Port     string `yaml:"port"`     // HTTP服务端口号
}

var configuration *Configuration

func LoadConfiguration(data []byte) error {
	var config Configuration
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
