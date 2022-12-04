package configs

import (
	"Postgres_Gin/postgres"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `envconfig:"host" required:"true"`
	Port string `envconfig:"port" required:"true"`
	postgres.Config
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
