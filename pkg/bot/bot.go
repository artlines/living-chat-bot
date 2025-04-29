package bot

import (
	"log"

	"github.com/artlines/living-chat-bot/pkg/bot/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Updater struct {
	Bot *tgbotapi.BotAPI
}

func New(apiToken string) (*Updater, error) {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Updater{Bot: bot}, nil
}

func (u *Updater) Run() error {
	u.Bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := u.Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.ChannelPost != nil {
			handler.HandleChannelPost(u.Bot, update)
		} else if update.Message != nil {
			if update.Message.IsCommand() {
				handler.HandleStart(u.Bot, update)
			} else {
				handler.HandleMessage(u.Bot, update)
			}
		}
	}

	return nil
}
