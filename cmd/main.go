package main

import (
	"living-chat-bot/pkg/config"
	"living-chat-bot/pkg/database"
	"log"
	"os"

	"github.com/artlines/living-chat-bot/pkg/bot"
)

func main() {
	config.LoadEnv()

	database.Connect()
	database.Migrate()
	database.Seed()

	apiToken := os.Getenv("TELEGRAM_TOKEN")
	if apiToken == "" {
		log.Fatal("TELEGRAM_TOKEN not set in environment")
	}

	b, err := bot.New(apiToken)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	if err := b.Run(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
