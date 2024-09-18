package util

import "github.com/spf13/viper"

// Config stores the configuration params for the application
// The values are read by viper from a config fiel or environment variables.
type Config struct {
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadEnv reads the configuration from a config file or environment variables.
func LoadEnv(path string, file string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
