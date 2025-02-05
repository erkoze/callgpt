package tg

import (
	"callgpt/configs"
	"callgpt/internal/chat"
	"callgpt/internal/tg/handlers"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	Tele        *telebot.Bot
	ChatService chat.ChatService
}

type BotDeps struct {
	Config      *configs.BotConfig
	ChatService chat.ChatService
}

func NewBot(deps *BotDeps) (*Bot, error) {
	settings := telebot.Settings{
		Token: deps.Config.Token,
	}

	tele, err := telebot.NewBot(settings)

	if err != nil {
		return nil, err
	}

	b := &Bot{
		Tele:        tele,
		ChatService: deps.ChatService,
	}

	b.initHandlers()

	return b, nil
}

func (b *Bot) initHandlers() {
	handlers.NewHelloHandler(b.Tele)
	handlers.NewStartHandler(b.Tele)

	handlers.NewTextHandler(&handlers.TextHandlerDeps{
		Bot:         b.Tele,
		ChatService: b.ChatService,
	})
}

func (b *Bot) Start() {
	b.Tele.Start()
}
