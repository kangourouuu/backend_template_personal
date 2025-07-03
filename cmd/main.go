package main

import (
	"backend_template_personal/cmd/generate"
	"backend_template_personal/common/config"
	"backend_template_personal/internal/sqlclient"
	"backend_template_personal/repository"
	"os"
	"path/filepath"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	err := config.InitConfig()
	if err !=nil {
		logrus.Fatal("failed to read file config")
	}

	if config.AppConfig.DB == "enabled" {
		gormsql := sqlclient.GormSqlConfig{
			Driver:   "postgresql",
			Host:     viper.GetString("db.host"),
			Port:     viper.GetInt("db.port"),
			Database: viper.GetString("db.database"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
		}
		initRepo(gormsql)
	}
	// Initialize logger
	if config.AppConfig.LogType == "FILE" {
		if err := os.MkdirAll(filepath.Dir(config.AppConfig.LogFile), 0755); err != nil {
			logrus.Fatal("Failed to create log directory: ", err)
		}
		f, err := os.OpenFile(config.AppConfig.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Fatal("Failed to open log file: ", err)
		}
		logrus.Println(f)
	}
}

func main() {
	generate.Execute()
}

func initRepo(gormSqlConfig sqlclient.GormSqlConfig) {
	repository.GormSqlClient = sqlclient.NewGormSqlClient(gormSqlConfig)
}
