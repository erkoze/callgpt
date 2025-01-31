package main

import (
	"callgpt/configs"
	"callgpt/internal/tg"
)

func main() {
	config := configs.LoadConfig()

	b, err := tg.NewBot(config.Bot)

	if err != nil {
		panic(err.Error())
	}

	b.Start()
}
