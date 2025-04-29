package database

import (
	"fmt"
	"log"
	"os"

	"github.com/artlines/living-chat-bot/pkg/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect устанавливает подключение к базе данных
func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST"),
		getEnv("DB_USER"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"),
		getEnv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к базе данных: %v", err)
	}

	fmt.Println("✅ Успешное подключение к базе данных")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Message{},
		&models.SystemMessage{},
	)
	if err != nil {
		log.Fatalf("❌ Ошибка миграции базы данных: %v", err)
	}

	fmt.Println("✅ Миграция базы данных завершена")
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("❌ Переменная окружения %s не установлена", key)
	}
	return value
}
