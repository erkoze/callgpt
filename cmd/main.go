package main

import (
	"callgpt/configs"
	"callgpt/internal/chat"
	"callgpt/internal/openai"
	"callgpt/internal/tg"
	"fmt"
)

func main() {
	conf := configs.LoadConfig()

	openaiClient := openai.NewClient(conf.OpenAI)
	openaiService := chat.NewOpenAIChatService(openaiClient)

	b, err := tg.NewBot(&tg.BotDeps{
		BotConfig:   &conf.Bot,
		ChatService: openaiService,
		AuthorData:  conf.AuthorData,
	})

	if err != nil {
		panic(err.Error())
	}

	b.Start()

	fmt.Println("Bot is online!")
}
