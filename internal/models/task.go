package models

import "time"

type Task struct {
	ID          string `gorm:"primaryKey"`
	CreatedAt   time.Time
	Name        string
	Description string
}
