package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID   uint   // #Связь с таблицей пользователей
	Text     string // Текст сообщения пользователя
	BotReply string // Ответ бота
	Source   string // Источник: "channel" или "private"
}
