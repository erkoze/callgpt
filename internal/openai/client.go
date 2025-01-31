package openai

import (
	"callgpt/configs"

	goopenai "github.com/sashabaranov/go-openai"
)

type Client struct {
	*goopenai.Client
}

func NewClient(config configs.OpenAIConfig) *Client {
	return &Client{
		goopenai.NewClient(config.APIKey),
	}
}
