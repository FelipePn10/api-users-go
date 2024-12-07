package env

import (
	"log/slog"

	"github.com/spf13/viper"
)

var Env *config

type config struct {
	GoEnv       string `mapstructure:"GO_ENV"`
	GoPort      string `mapstructure:"GO_PORT=8080"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadingConfig(path string) (*config, error) {
	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.SetDefault("GO_PORT", "8080")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		slog.Warn("No .env file found, using default values")
		return nil, err
	}
	err = viper.Unmarshal(&Env)
	if err != nil {
		return nil, err
	}
	return Env, nil
}
