package models

import (
	"gorm.io/gorm"
)

type SystemMessage struct {
	gorm.Model
	Code string `gorm:"uniqueIndex;not null"` // Код системного сообщения
	Text string `gorm:"type:text;not null"`   // Содержимое сообщения
}
