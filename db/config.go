package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Config contains the configurations for database.
type Config struct {
	// Connection string for database.
	DatabaseURI string
}

// InitConfig load the configuration for database.
func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("database_uri"),
	}
	if len(config.DatabaseURI) == 0 {
		return nil, fmt.Errorf("database_uri must be set")
	}
	return config, nil
}
