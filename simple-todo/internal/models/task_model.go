package models

import "time"

type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string

	ProjectID string 

	Assignees []User `gorm:"many2many:task_assignees;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}