package main

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"backend_template_personal/api"
	slog "backend_template_personal/common/log"
	"backend_template_personal/internal/sqlclient"
	"backend_template_personal/middleware/auth"
	"backend_template_personal/repository"
	server "backend_template_personal/server/http"
)

type Config struct {
	Dir     string `env:"CONFIG_DIR" envDefault:"config/config.json"`
	Port    string
	LogType string
	LogFile string
	DB      string
	Redis   string
}

var config Config

func init() {
	// Parse environment variables
	if err := env.Parse(&config); err != nil {
		slog.Fatal("Failed to parse environment variables: ", err)
	}

	// Initialize viper for configuration
	viper.SetConfigFile(config.Dir)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		slog.Fatal("Failed to read config file: ", err)
	}

	// Set config values from viper
	cfg := Config{
		Dir:     config.Dir,
		Port:    viper.GetString("main.port"),
		LogType: viper.GetString("main.log_type"),
		LogFile: viper.GetString("main.log_file"),
		DB:      viper.GetString("main.db"),
		Redis:   viper.GetString("main.redis"),
	}
	if cfg.DB == "enabled" {
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
	if config.LogType == "FILE" {
		if err := os.MkdirAll(filepath.Dir(config.LogFile), 0755); err != nil {
			slog.Fatal("Failed to create log directory: ", err)
		}
		f, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			slog.Fatal("Failed to open log file: ", err)
		}
		slog.Println(f)
	}
	config = cfg
}

func main() {
	// Initialize Gin router
	engine := server.NewEngine()

	// Apply middleware
	engine.Use(auth.AuthMiddleWare())

	apiV1 := engine.Group("/api/v1")
	api.SetUpRoutes(apiV1)

	// Start the server
	appServer := server.New(config.Port, engine)
	if err := appServer.Run(); err != nil {
		logrus.Fatalf("Server could not be started: %v", err)
	}
}

func initRepo(gormSqlConfig sqlclient.GormSqlConfig) {
	repository.GormSqlClient = sqlclient.NewGormSqlClient(gormSqlConfig)
}
