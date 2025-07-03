package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Dir     string `env:"CONFIG_DIR" envDefault:"config/config.json"`
	Port    string
	LogType string
	LogFile string
	DB      string
	Redis   string
}

var AppConfig Config

func InitConfig() error {
	// Parse environment variables
	if err := env.Parse(&AppConfig); err != nil {
		logrus.Fatal("Failed to parse environment variables: ", err)
	}

	// Initialize viper for configuration
	viper.SetConfigFile(AppConfig.Dir)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Failed to read config file: ", err)
	}

	AppConfig.Port = viper.GetString("main.port")
	AppConfig.LogFile = viper.GetString("main.log_file")
	AppConfig.LogType = viper.GetString("main.log_type")
	AppConfig.DB = viper.GetString("main.db")

	return nil
}
