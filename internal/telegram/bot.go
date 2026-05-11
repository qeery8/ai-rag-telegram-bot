package telegram

import (
	"al-rag-tgbot/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api     *tgbotapi.BotAPI
	handler *Handler
}

func New(cfg *config.Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	api.Debug = cfg.Debug
	log.Printf("Authorized on account %s", api.Self.UserName)

	return &Bot{
		api:     api,
		handler: NewHandler(api),
	}, nil
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		go b.handler.Handle(update)
	}
}
