package env

import (
	"log/slog"

	"github.com/spf13/viper"
)

var Env *config

type config struct {
	GoEnv       string `mapstructure:"GO_ENV"`
	GoPort      string `mapstructure:"GO_PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadingConfig(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	viper.SetDefault("GO_PORT", "8080")

	err := viper.ReadInConfig()
	if err != nil {
		slog.Warn("No app_config.env file found, using default values")
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		return nil, err
	}
	return Env, nil
}
