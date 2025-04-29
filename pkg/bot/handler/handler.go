package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"living-chat-bot/pkg/database"
	"living-chat-bot/pkg/database/models"
	"living-chat-bot/pkg/openai"
	"log"
)

func HandleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	user := update.Message.From

	var existing models.User
	result := database.DB.First(&existing, "telegram_id = ?", user.ID)
	if result.Error != nil {
		createUser(user)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🌿 Привет. Я рядом")
	bot.Send(msg)
}

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Received message: %s", update.Message.Text)
	user := update.Message.From

	var existing models.User
	result := database.DB.First(&existing, "telegram_id = ?", user.ID)
	if result.Error != nil {
		createUser(user)
	}

	response, err := openai.SendMessage(update.Message.Text)
	if err != nil {
		log.Printf("Error sending message to OpenAI: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func HandleChannelPost(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	post := update.ChannelPost
	channelName := func() string {
		if update.Message.Chat.Title != "" {
			return update.Message.Chat.Title
		}
		return fmt.Sprintf("%d", update.Message.Chat.ID)
	}()

	log.Printf("Новый пост в канале: %s", channelName)

	reply, err := openai.SendMessage(post.Text)
	if err != nil {
		log.Printf("Ошибка OpenAI: %v", err)
		return
	}

	if post.AuthorSignature != "" {
		log.Printf("Автор поста: %s", post.AuthorSignature)

		var dbUser models.User
		result := database.DB.First(&dbUser, "username = ?", post.AuthorSignature)
		if result.Error == nil {
			msg := tgbotapi.NewMessage(dbUser.TelegramID, reply)
			bot.Send(msg)

			message := models.Message{
				UserID:   dbUser.ID,
				Text:     post.Text,
				BotReply: reply,
				Source:   "channel",
			}
			if err := database.DB.Create(&message).Error; err != nil {
				log.Printf("Ошибка сохранения сообщения: %v", err)
			}

			return
		} else {
			log.Printf("Автор поста не найден в базе: %v", result.Error)
		}
	}

	log.Println("Автор не определён. Ответ не отправлен.")
}

func createUser(user *tgbotapi.User) models.User {
	newUser := models.User{
		TelegramID: user.ID,
		Username:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		log.Printf("Error saving user: %v", err)
	}

	return newUser
}
