package main

import (
	"callgpt/configs"
	"callgpt/internal/chat"
	"callgpt/internal/openai"
	"callgpt/internal/tg"
)

func main() {
	conf := configs.LoadConfig()

	openaiClient := openai.NewClient(conf.OpenAI)
	chatService := chat.NewChatService(openaiClient)

	b, err := tg.NewBot(&tg.BotDeps{
		Config:      &conf.Bot,
		ChatService: chatService,
	})

	if err != nil {
		panic(err.Error())
	}

	b.Start()
}
