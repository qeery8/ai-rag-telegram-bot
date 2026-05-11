package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Handler struct {
	api *tgbotapi.BotAPI
}

func NewHandler(api *tgbotapi.BotAPI) *Handler {
	return &Handler{api: api}
}

func (h *Handler) Handle(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "start":
		h.handleStart(update.Message)
	case "help":
		h.handleHelp(update.Message)
	default:
		h.handleMessage(update.Message)
	}
}

func (h *Handler) handleStart(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "SOSI")
	h.api.Send(reply)
}

func (h *Handler) handleHelp(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "cho to cho to")
	h.api.Send(reply)
}

func (h *Handler) handleMessage(msg *tgbotapi.Message) {
	//вызов rag метода

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Processig the request...")
	h.api.Send(reply)
}
