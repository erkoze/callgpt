package handlers

import (
	"callgpt/internal/chat"
	"context"

	"gopkg.in/telebot.v4"
)

type TextHandlerDeps struct {
	Bot         *telebot.Bot
	ChatService chat.ChatService
}

type textHandler struct {
	chatService chat.ChatService
}

func NewTextHandler(deps *TextHandlerDeps) {
	handler := &textHandler{
		chatService: deps.ChatService,
	}

	deps.Bot.Handle(telebot.OnText, handler.handle)
}

func (h *textHandler) handle(c telebot.Context) error {
	if c.Sender().ID != 1077702537 {
		return c.Send("Вам не разрешено использовать бота")
	}

	res, err := h.chatService.GetAnswer(context.Background(), c.Text())

	if err != nil {
		return err
	}

	return c.Send(res)
}
