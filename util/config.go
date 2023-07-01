package util

import (
	"time"

	"github.com/spf13/viper"
)

// The values are read by viper from config file or environment variables

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationUrl         string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	Environment          string        `mapstructure:"ENVIRONMENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml, yml ...
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
