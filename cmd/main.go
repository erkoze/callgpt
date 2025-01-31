package main

import (
	"callgpt/configs"
	"callgpt/internal/tg"
)

func main() {
	config := configs.LoadConfig()

	b := tg.NewBot(config.Bot)

	b.Start()
}
