package tasks

import (
	"simple-todo/internal/modules/projects"
	"simple-todo/internal/modules/auth"

	"time"
)


type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Status      string

	// Belongs to project
	ProjectID string
	Project   projects.Project `gorm:"foreignKey:ProjectID"`

	// Assigned users (many-to-many)
	Assignees []auth.User `gorm:"many2many:task_assignees;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}