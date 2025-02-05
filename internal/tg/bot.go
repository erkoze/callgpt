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

	b.Tele.SetCommands(b.getCommands())

	b.initHandlers()

	return b, nil
}

func (b *Bot) initHandlers() {
	handlers.NewStartHandler(b.Tele)

	handlers.NewTextHandler(&handlers.TextHandlerDeps{
		Bot:         b.Tele,
		ChatService: b.ChatService,
	})

	handlers.NewImgHandler(&handlers.ImgHandlerDeps{
		Bot:         b.Tele,
		ChatService: b.ChatService,
	})
}

func (b *Bot) getCommands() []telebot.Command {
	return []telebot.Command{
		{Text: "start", Description: "Перезапуск"},
		{Text: "img", Description: "Сгенерировать изображение"},
	}
}

func (b *Bot) Start() {
	b.Tele.Start()
}
