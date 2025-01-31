package tg

import (
	"callgpt/configs"

	"gopkg.in/telebot.v4"
)

func NewBot(config configs.BotConfig) *telebot.Bot {
	settings := telebot.Settings{
		Token: config.Token,
	}

	b, err := telebot.NewBot(settings)

	if err != nil {
		panic(err.Error())
	}

	b.Handle("/hello", func(c telebot.Context) error {
		return c.Send("Hello!")
	})

	return b
}
