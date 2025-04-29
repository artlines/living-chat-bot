# ========================
# Stage 1 — Builder (общий)
# ========================
FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY ../go.mod go.sum ../
RUN go mod download

COPY .. ..

# По умолчанию сборка PROD бинарника с оптимизацией
RUN go build -ldflags="-s -w" -o app ../cmd/

# ========================
# Stage 2 — DEV образ
# ========================
FROM golang:1.24.2-alpine AS dev

RUN apk add --no-cache git

WORKDIR /app

COPY --from=builder /app /app

RUN go install github.com/air-verse/air@latest
ENV PATH=$PATH:/go/bin

CMD ["air"]

# ========================
# Stage 3 — PROD образ
# ========================
FROM alpine:latest AS prod

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/app ..

CMD ["../app"]
