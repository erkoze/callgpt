package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type IdHandlerDeps struct {
	Bot *telebot.Bot
}

type idHandler struct {
}

func NewIdHandler(b *telebot.Bot) {
	handler := &idHandler{}

	b.Handle("/id", handler.handle)
}

func (h *idHandler) handle(c telebot.Context) error {
	return c.Send(fmt.Sprintf("Ваш айди: %v", c.Sender().ID))
}
