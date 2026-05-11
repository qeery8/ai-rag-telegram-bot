package main

import (
	"al-rag-tgbot/internal/config"
	"al-rag-tgbot/internal/telegram"
	"log"
)

func main() {
	cfg := config.Load()

	b, err := telegram.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	b.Run()

	log.Println("Bot started...")
}
