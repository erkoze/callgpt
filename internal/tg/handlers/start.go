package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type startHandler struct{}

func NewStartHandler(b *telebot.Bot) {
	handler := &startHandler{}

	b.Handle("/start", handler.handle)
}

func (h *startHandler) handle(c telebot.Context) error {
	user := c.Sender()

	str := fmt.Sprintf("Привет, %v! Это бот для удобного взаимодействия с нейросетью. \n Чтобы задать вопрос, просто напишите его.", user.Username)

	return c.Send(str)
}
