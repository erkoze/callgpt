package tg

import (
	"callgpt/configs"
	"callgpt/internal/chat"

	"gopkg.in/telebot.v4"
)

type Bot struct {
	tele        *telebot.Bot
	chatService *chat.ChatService
}

type BotDeps struct {
	Config      *configs.BotConfig
	ChatService *chat.ChatService
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
		tele:        tele,
		chatService: deps.ChatService,
	}

	b.initHandlers()

	return b, nil
}

func (b *Bot) initHandlers() {
	b.tele.Handle("/hello", b.helloCommand)
}

func (b *Bot) helloCommand(c telebot.Context) error {
	return c.Send("Hello!")
}

func (b *Bot) Start() {
	b.tele.Start()
}
