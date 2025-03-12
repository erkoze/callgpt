package handlers

import (
	"callgpt/internal/chat"
	"callgpt/pkg/ratelimiter"
	"context"
	"fmt"

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
	if !ratelimiter.GlobalLimiter.Allow(c.Sender().ID) {
		return c.Send("❌ Вы превысили лимит запросов (5 в час). Попробуйте позже.")
	}

	fmt.Printf("onText, userId: %v, content: %v \n", c.Sender().ID, c.Text())

	message, err := h.chatService.CreateChatCompletion(context.Background(), c.Text())

	if err != nil {
		return err
	}

	return c.Send(message)
}
