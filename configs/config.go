package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Bot BotConfig
}

type BotConfig struct {
	Token string
}

func LoadConfig() *Config {
	godotenv.Load(".env")

	return &Config{
		BotConfig{
			Token: os.Getenv("TELEGRAM_API_TOKEN"),
		},
	}
}
