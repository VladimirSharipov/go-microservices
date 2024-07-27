package main

import (
	"log"
	"my_microservice/config"
	"my_microservice/db"
	"my_microservice/models"
	"my_microservice/rabbitmq"
	"my_microservice/service"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	database, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()

	// Автомиграция модели Task
	if err := database.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Создаем сервис обработки задач
	taskService := service.NewTaskService(database)

	// Настраиваем RabbitMQ consumer
	consumer, err := rabbitmq.NewConsumer(cfg, taskService)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}

	log.Println("Listening for messages...")

	// Запускаем consumer для получения сообщений
	consumer.ListenForMessages()
}
