package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Bot    BotConfig
	OpenAI OpenAIConfig
	Db     DbConfig
}

type BotConfig struct {
	Token string
}

type OpenAIConfig struct {
	APIKey string
}

type DbConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SslMode  string
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
		DbConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			Port:     os.Getenv("DATABASE_PORT"),
			SslMode:  os.Getenv("DATABASE_SSL_MODE"),
		},
	}
}
