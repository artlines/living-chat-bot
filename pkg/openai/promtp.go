package openai

import (
	"living-chat-bot/pkg/database"
	"living-chat-bot/pkg/database/models"
	"log"
)

func GetPrompt() string {
	var prompt models.SystemMessage
	err := database.DB.First(&prompt, "code = ?", "openai_prompt").Error
	if err != nil {
		log.Println("❌ Не удалось получить openai_prompt из базы:", err)
	}

	return prompt.Text
}
