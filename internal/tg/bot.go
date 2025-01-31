package tg

import (
	"callgpt/configs"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	tele *telebot.Bot
}

func NewBot(config configs.BotConfig) (*Bot, error) {
	settings := telebot.Settings{
		Token: config.Token,
	}

	tele, err := telebot.NewBot(settings)

	if err != nil {
		return nil, err
	}

	bot := &Bot{
		tele: tele,
	}

	bot.initHandlers()

	return bot, nil
}

func (bot *Bot) initHandlers() {
	bot.tele.Handle("/hello", bot.HelloCommand)
}

func (bot *Bot) HelloCommand(c telebot.Context) error {
	return c.Send("Hello!")
}

func (bot *Bot) Start() {
	bot.tele.Start()
}
