# Установим базовый образ
FROM golang:1.22-alpine

# Установим рабочую директорию внутри контейнера
WORKDIR /app

# Скопируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Установим зависимости
RUN go mod download

# Скопируем все остальные файлы в контейнер
COPY . .
    
RUN go build -o main ./cmd/app

# Запуск контейнера
CMD ["./main"]

