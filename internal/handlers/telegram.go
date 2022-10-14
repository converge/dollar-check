package handlers

import (
	"dolar-check/internal/services"
)

type TelegramHandler struct {
	service *services.TelegramService
}

func NewTelegramHandler(telegramService *services.TelegramService) *TelegramHandler {
	return &TelegramHandler{service: telegramService}
}

func (handler TelegramHandler) SendMessage(text string) error {
	err := handler.service.SendMessage(text)
	if err != nil {
		return err
	}
	return nil
}
