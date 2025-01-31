package chat

import "callgpt/internal/openai"

type ChatService struct {
	openaiClient *openai.Client
}

func NewChatService(openaiClient *openai.Client) *ChatService {
	return &ChatService{
		openaiClient: openaiClient,
	}
}

//TODO: Make chatting methods
