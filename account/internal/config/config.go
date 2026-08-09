package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" required:"true" envDefault:"account-service"`
	AppEnv      string `env:"APP_ENV" required:"true" envDefault:"development"`
	Host        string `env:"HTTP_HOST" required:"true" envDefault:"localhost"`
	Port        string `env:"HTTP_PORT" required:"true" envDefault:"9000"`
	LogLevel    string `env:"LOG_LEVEL" required:"true" envDefault:"info"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warring: .env file not found, using system enviroment variables")
	}

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
