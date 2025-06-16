package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

var RuntimeConfig AppConfig

type AppConfig struct {
	Server    Server
	Postgres  Postgres
	JWTSecret string `env:"JWT_SECRET"`
}

type Server struct {
	Addr string `env:"SERVER_ADDR"`
	Port int    `env:"SERVER_PORT"`
}

type Postgres struct {
	Host     string        `env:"DB_HOST"`
	Port     int           `env:"DB_PORT"`
	Username string        `env:"DB_USER"`
	Password string        `env:"DB_PASS"`
	Timeout  time.Duration `env:"DB_TIMEOUT"`
	DBName   string        `env:"DB_NAME"`
}

func LoadConfig(configPath string) {
	if err := godotenv.Load(configPath); err != nil {
		slog.Error("error loading .env file", "error", err)
	}

	// Parse environment variables into the Config struct

	if err := env.Parse(&RuntimeConfig); err != nil {
		slog.Error("error parsing env vars into struct", "error", err)
		os.Exit(-1)
	}

}
