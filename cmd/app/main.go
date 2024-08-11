package main

import (
	"github.com/VladimirSharipov/go-microservices/internal/config"
	"github.com/VladimirSharipov/go-microservices/internal/db"
	"github.com/VladimirSharipov/go-microservices/internal/models"
	"github.com/VladimirSharipov/go-microservices/rabbitmq"
	"github.com/VladimirSharipov/go-microservices/service"
	"log"
)

func main() {

	cfg := config.LoadConfig()

	database, err := db.NewDatabase(cfg)
	if err != nil {

		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()

	if err := database.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	taskService := service.NewTaskService(database)

	consumer, err := rabbitmq.NewConsumer(cfg, taskService)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}

	log.Println("Listening for messages...")

	consumer.ListenForMessages()
}
