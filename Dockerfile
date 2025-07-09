# ---------- Этап 1: Сборка ----------
FROM golang:1.23.1-alpine AS builder

RUN apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o marketflow ./cmd/main.go

# ---------- Этап 2: Запуск ----------
FROM alpine:latest

WORKDIR /app

# Копируем только бинарник
COPY --from=builder /build/marketflow .

EXPOSE 8080

CMD ["./marketflow"]
