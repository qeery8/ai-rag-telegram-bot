package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Telegram TelegramConfig
	Postgres PostgresConfig
}

type TelegramConfig struct {
	Token string
	Debug bool
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	return &Config{
		Env: getEnv("ENVIRONMENT", "development"),
		Telegram: TelegramConfig{
			Token: mustGet("TELEGRAM_TOKEN"),
			Debug: getEnv("TELEGRAM_DEBUG", "false") == "true",
		},
		Postgres: PostgresConfig{
			Username: getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "password"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			DbName:   getEnv("POSTGRES_DB", "dbname"),
			SslMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
		},
	}
}

func mustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable %s is required", key)
	}

	return val
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
