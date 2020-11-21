package app

import (
	"github.com/spf13/viper"
)

// Config contains the configurations for application.
type Config struct {
	// Application version.
	Version string

	// HTTPPrefix is http prefix for web endpoints.
	HTTPPrefix string

	// HOST CRM
	HOST_API_CRM []string

	HOST_CRM string

	// HOST CRM
	TOKEN_CRM string
}

// InitConfig load the configuration for application.
func InitConfig() (*Config, error) {
	config := &Config{
		Version:      viper.GetString("VERSION"),
		HTTPPrefix:   viper.GetString("HTTP_PREFIX"),
		HOST_CRM:     viper.GetString("HOST_CRM"),
		HOST_API_CRM: viper.GetStringSlice("HOST_API_CRM"),
		TOKEN_CRM:    viper.GetString("TOKEN_CRM"),
	}

	return config, nil
}
