BINARY_NAME=living-chat-bot

# Работа с DEV окружением
dev-up:
	docker compose -f docker-compose.dev.yml up --build -d && APP_ENV=dev go run ./cmd/main.go

dev-down:
	docker compose -f docker-compose.dev.yml down

dev-build:
	docker compose -f docker-compose.dev.yml build

dev-logs:
	docker compose -f docker-compose.dev.yml logs -f bot

# Работа с PROD окружением
prod-up:
	docker compose -f docker-compose.production.yml up --build -d && APP_ENV=prod go run ./cmd/main.go

prod-down:
	docker compose -f docker-compose.production.yml down

prod-build:
	docker compose -f docker-compose.production.yml build

prod-logs:
	docker compose -f docker-compose.production.yml logs -f bot