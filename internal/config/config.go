package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	Debug         bool
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	return &Config{
		TelegramToken: mustGet("TELEGRAM_TOKEN"),
		Debug:         os.Getenv("DEBUG") == "true",
	}
}

func mustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable %s is required", key)
	}

	return val
}
