package handlers

import (
	"callgpt/internal/chat"
	"context"

	"gopkg.in/telebot.v4"
)

type DrawHandlerDeps struct {
	Bot         *telebot.Bot
	ChatService chat.ChatService
}

type drawHandler struct {
	chatService chat.ChatService
}

func NewDrawHandler(deps *DrawHandlerDeps) {
	handler := &drawHandler{
		chatService: deps.ChatService,
	}

	deps.Bot.Handle("/draw", handler.handle)
}

func (h *drawHandler) handle(c telebot.Context) error {
	if c.Sender().ID != 1077702537 {
		return c.Send("Вам не разрешено использовать бота")
	}

	prompt := c.Message().Payload

	if prompt == "" || len(prompt) < 3 {
		text := "Чтобы нарисовать картинку, напишите команду /draw и опишите то, что вы хотите получить.\n" +
			"Запрос должен содержать больше 3 символов.\n\n" +
			"Например: /draw кот и человек"

		return c.Send(text)
	}

	res, err := h.chatService.Draw(context.Background(), c.Message().Payload)

	if err != nil {
		return err
	}

	photo := &telebot.Photo{File: telebot.FromURL(res)}

	return c.Send(photo)
}
