package util

import "github.com/spf13/viper"

// Config stores the configuration params for the application
// The values are read by viper from a config fiel or environment variables.
type Config struct {
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JwksURL       string `mapstructure:"JWKS_URL"`
	IssuerURL     string `mapstructure:"ISSUER_URL"`
}

// LoadEnv reads the configuration from a config file or environment variables.
func LoadEnv(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
