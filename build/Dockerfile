FROM golang:latest

WORKDIR /app

# Устанавливаем CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Добавляем /go/bin в PATH (где CompileDaemon)
ENV PATH="/go/bin:$PATH"

# Копируем код
COPY . .

# Устанавливаем зависимости
RUN go mod download

EXPOSE 8888

# Запускаем CompileDaemon
CMD ["CompileDaemon", "--build=go build -o main .", "--command=./main"]
