package rabbitmq

import (
	"encoding/json"
	"github.com/VladimirSharipov/go-microservices/internal/config"
	"github.com/VladimirSharipov/go-microservices/internal/models"
	"github.com/VladimirSharipov/go-microservices/service"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	config      *config.Config
	taskService *service.TaskService
}

func NewConsumer(cfg *config.Config, taskService *service.TaskService) (*Consumer, error) {
	return &Consumer{config: cfg, taskService: taskService}, nil
}

func (c *Consumer) ListenForMessages() {
	conn, err := amqp.Dial(c.config.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		c.config.RabbitMQQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var task models.Task
			if err := json.Unmarshal(d.Body, &task); err != nil {
				log.Printf("Error decoding JSON: %v", err)
				continue
			}
			if err := c.taskService.CreateTask(&task); err != nil {
				log.Printf("Error saving task to database: %v", err)
				continue
			}
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
