package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// AppConfig holds the application configuration
type AppConfig struct {
	Port     string `envconfig:"PORT" default:"8123"`
	MongoURI string `envconfig:"MONGO_URI" default:"mongodb://user:password@localhost:27017"`
	DbName   string `envconfig:"DB_NAME" default:"test"`
}

// LoadConfig reads the configuration from environment variables
func LoadConfig() *AppConfig {
	var config AppConfig
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Error loading the configuration: %s", err)
	}
	return &config
}

// create global variable CFG in config.go
var CFG *AppConfig = LoadConfig()
