package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

func Init() (*Config, error) {
	mainViper := viper.New()
	mainViper.AddConfigPath("configs")
	if err := mainViper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := mainViper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
