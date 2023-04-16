package config

import "github.com/caarlos0/env/v6"

type Config struct {
	BotToken string `env:"BOT_TOKEN"`
	DBLToken string `env:"DBL_TOKEN"`
}

func Load() (Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
