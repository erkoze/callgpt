package chat

import (
	"callgpt/internal/openai"
	"context"
	"fmt"

	goopenai "github.com/sashabaranov/go-openai"
)

type ChatService interface {
	CreateChatCompletion(ctx context.Context, prompt string) (string, error)
	GenerateImage(ctx context.Context, prompt string) (string, error)
}

type OpenAIChatService struct {
	client *openai.Client
}

func NewOpenAIChatService(client *openai.Client) *OpenAIChatService {
	return &OpenAIChatService{
		client: client,
	}
}

func (s *OpenAIChatService) CreateChatCompletion(ctx context.Context, prompt string) (string, error) {
	res, err := s.client.CreateChatCompletion(ctx, goopenai.ChatCompletionRequest{
		Model: goopenai.GPT4o,
		Messages: []goopenai.ChatCompletionMessage{
			{
				Role:    goopenai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}

func (s *OpenAIChatService) GenerateImage(ctx context.Context, prompt string) (string, error) {
	req := goopenai.ImageRequest{
		ResponseFormat: goopenai.CreateImageResponseFormatURL,
		Size:           goopenai.CreateImageSize1024x1024,
		Model:          goopenai.CreateImageModelDallE3,
		Prompt:         prompt,
		N:              1,
	}

	res, err := s.client.CreateImage(ctx, req)

	if err != nil {
		fmt.Printf("Image generation error: %v \n", err)

		return "", err
	}

	return res.Data[0].URL, nil
}
