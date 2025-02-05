package chat

import (
	"callgpt/internal/openai"
	"context"
	"fmt"

	goopenai "github.com/sashabaranov/go-openai"
)

type ChatService interface {
	GetAnswer(ctx context.Context, prompt string) (string, error)
}

type OpenAIChatService struct {
	client *openai.Client
}

func NewOpenAIChatService(client *openai.Client) *OpenAIChatService {
	return &OpenAIChatService{
		client: client,
	}
}

func (s *OpenAIChatService) GetAnswer(ctx context.Context, prompt string) (string, error) {
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

	fmt.Println(res.Choices[0].Message.Content)

	return res.Choices[0].Message.Content, nil
}
