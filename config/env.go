package config

import (
	"github.com/spf13/viper"
)

// EnvSetup sets up the environment variables.
func EnvSetup() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
