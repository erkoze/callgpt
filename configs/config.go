package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Bot    BotConfig
	OpenAI OpenAIConfig
}

type BotConfig struct {
	Token string
}

type OpenAIConfig struct {
	APIKey string
}

func LoadConfig() *Config {
	godotenv.Load(".env")

	return &Config{
		BotConfig{
			Token: os.Getenv("TELEGRAM_API_TOKEN"),
		},
		OpenAIConfig{
			APIKey: os.Getenv("OPEN_AI_API_KEY"),
		},
	}
}
