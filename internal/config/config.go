package config

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	RabbitMQURL      string
	RabbitMQQueue    string
}

func LoadConfig() *Config {
	return &Config{
		PostgresUser:     "postgres",
		PostgresPassword: "root",
		PostgresDB:       "go-microservice",
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		RabbitMQURL:      "amqp://guest:guest@rabbitmq:5672/",
		RabbitMQQueue:    "task_queue",
	}
}
