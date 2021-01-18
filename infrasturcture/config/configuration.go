package config

import (
	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"
	"strings"
)

type Configuration struct {
	Database string `yaml:"database"` // 数据库类型，决定Schema和驱动的位置，支持sqlite3和mysql
	DSN      string `yaml:"dsn"`      // 数据库DSN
	Port     string `yaml:"port"`     // HTTP服务端口号
	Logger   struct {
		Level           string `yaml:"level"`  // 日志级别
		TimestampFormat string `yaml:"format"` // 日志时间格式
	}
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

func (c *Configuration) LoggerLevel() logrus.Level {
	switch strings.ToLower(c.Logger.Level) {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.DebugLevel
	}
}
