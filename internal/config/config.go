package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	SubscriptionID string `mapstructure:"subscription"`
	OutputFormat   string `mapstructure:"output"`
}

func Load() (*Config, error) {
	viper.SetDefault("output", "table")
	if sub := os.Getenv("AZURE_SUBSCRIPTION_ID"); sub != "" {
		viper.SetDefault("subscription", sub)
	}
	var cfg Config
	return &cfg, viper.Unmarshal(&cfg)
}
