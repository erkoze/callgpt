package handlers

import (
	"callgpt/internal/chat"
	"callgpt/pkg/ratelimiter"
	"context"
	"fmt"

	"gopkg.in/telebot.v4"
)

type ImgHandlerDeps struct {
	Bot         *telebot.Bot
	ChatService chat.ChatService
}

type imgHandler struct {
	chatService chat.ChatService
	Bot         *telebot.Bot
}

func NewImgHandler(deps *ImgHandlerDeps) {
	handler := &imgHandler{
		chatService: deps.ChatService,
		Bot:         deps.Bot,
	}

	deps.Bot.Handle("/img", handler.handle)
}

func (h *imgHandler) handle(c telebot.Context) error {
	prompt := c.Message().Payload

	if prompt == "" || len(prompt) < 3 {
		return c.Send("Чтобы нарисовать картинку, напишите команду /img и опишите то, что вы хотите получить.\n" +
			"Запрос должен содержать больше 3 символов.\n\n" +
			"Например: /img кот и человек")
	}

	if !ratelimiter.GlobalLimiter.Allow(c.Sender().ID) {
		return c.Send("❌ Вы превысили лимит запросов (5 в час). Попробуйте позже.")
	}

	fmt.Printf("New /img command, userId: %v, prompt: %v \n", c.Sender().ID, prompt)

	msg, err := h.Bot.Send(c.Chat(), "Генерация изображения...")

	if err != nil {
		return err
	}

	res, err := h.chatService.GenerateImage(context.Background(), prompt)

	if err != nil {
		h.Bot.Edit(msg, "Произошла ошибка с генерацией изображения")

		return err
	}

	h.Bot.Delete(msg)

	photo := &telebot.Photo{
		File:    telebot.FromURL(res),
		Caption: fmt.Sprintf("Запрос: %v", prompt),
	}

	return c.Send(photo)
}
