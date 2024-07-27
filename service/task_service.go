package service

import (
	"gorm.io/gorm"
	"my_microservice/models"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.db.Create(task).Error
}
