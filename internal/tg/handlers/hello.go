package handlers

import (
	"gopkg.in/telebot.v4"
)

type HelloHandler struct{}

func NewHelloHandler(bot *telebot.Bot) {
	handler := &HelloHandler{}

	bot.Handle("/hello", handler.handle)
}

func (h *HelloHandler) handle(c telebot.Context) error {
	return c.Send("Hello!")
}
