# 1. Билдим приложение в Go
FROM golang:1.24 AS builder

WORKDIR /app

# Копируем все файлы проекта
COPY . .

# Устанавливаем зависимости
RUN go mod download

# Компилируем Go-приложение
RUN go build -o main .

# 2. Финальный контейнер
FROM debian:bookworm-slim
WORKDIR /app

# Копируем все файлы из билдера (включая бинарник)
COPY --from=builder /app .

# Открываем порт
EXPOSE 8888

# Запускаем приложение
CMD ["./main"]
