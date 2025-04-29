package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID int64  `gorm:"uniqueIndex;not null"` // ID пользователя в Telegram
	Username   string // Username (@nickname), может быть пустым
	FirstName  string // FirstName может быть пустым
	LastName   string // LastName может быть пустым
}
