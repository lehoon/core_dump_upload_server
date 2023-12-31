package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

var (
	config *Config
)

type Local struct {
	Address string `yaml:"address"`
}

type Alarm struct {
	Address string `yaml:"address"`
	Url     string `yaml:"url"`
	Used    bool   `yaml:"used"`
	Method  string `yaml:"method"`
}

type Logger struct {
	MaxAge      int    `yaml:"maxAge"`
	MaxSize     int    `yaml:"maxSize"`
	MaxBackup   int    `yaml:"maxBackup"`
	Compress    bool   `yaml:"compress"`
	Level       int    `yaml:"level"`
	LogPath     string `yaml:"logPath"`
	ServiceName string `yaml:"serviceName"`
}

type Config struct {
	Alarm
	Local
	Logger
}

func init() {
	newConfig()
	loadConfig()
}

func newConfig() *Config {
	config = &Config{}
	return config
}

func GetLocalAddress() string {
	return config.Local.Address
}

// alarm config
func AlarmAddress() string {
	return config.Alarm.Address
}

// alarm url
func AlarmUrl() string {
	return config.Alarm.Url
}

func UsedAlarm() bool {
	return config.Alarm.Used
}

func AlarmMethod() string {
	return strings.ToLower(config.Alarm.Method)
}

func GetLoggerLevel() int {
	return config.Logger.Level
}
func GetLoggerPath() string {
	return config.Logger.LogPath
}

func GetLoggerServiceName() string {
	return config.Logger.ServiceName
}

func GetLoggerMaxAge() int {
	return config.Logger.MaxAge
}

func GetLoggerMaxSize() int {
	return config.Logger.MaxSize
}

func GetLoggerMaxBackup() int {
	return config.Logger.MaxBackup
}

func GetLoggerCompress() bool {
	return config.Logger.Compress
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(errors.New("反序列化配置文件出错"))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(errors.New("反序列化配置文件出错"))
	}
}
