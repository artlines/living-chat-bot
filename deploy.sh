#!/bin/bash

set -e
echo "🚀 Запуск деплоя Living Chat Bot..."

PROJECT_DIR=./
COMPOSE_FILE=.infrastructure/docker-compose.prod.yml

cd $PROJECT_DIR || {
  echo "❌ Не удалось найти папку проекта $PROJECT_DIR"
  exit 1
}

echo "🛑 Останавливаем старые контейнеры..."
docker compose -f $COMPOSE_FILE down

echo "🔨 Собираем образы..."
docker compose -f $COMPOSE_FILE build

echo "🚀 Запускаем проект..."
docker compose -f $COMPOSE_FILE up -d

echo "✅ Деплой завершён успешно!"
