package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type StartHandlerDeps struct {
	Bot        *telebot.Bot
	AuthorData string
}

type startHandler struct {
	authorData string
}

func NewStartHandler(deps *StartHandlerDeps) {
	handler := &startHandler{
		authorData: deps.AuthorData,
	}

	deps.Bot.Handle("/start", handler.handle)
}

func (h *startHandler) handle(c telebot.Context) error {
	user := c.Sender()

	str := fmt.Sprintf(
		"Привет, %v! Это бот для удобного взаимодействия с нейросетью. \nБот был написан: %v\n\nЧтобы задать вопрос, просто напишите его.",
		user.Username,
		h.authorData,
	)

	return c.Send(str)
}
