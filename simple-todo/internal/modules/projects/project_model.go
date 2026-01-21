package projects

import (
	"simple-todo/internal/modules/auth"

	"time"
)

type Project struct {
	ID                 string `gorm:"primaryKey"`
	Title       string
	Description string
	Status             string

	OwnerID string
	Owner   auth.User `gorm:"foreignKey:OwnerID"`

	Members []auth.User `gorm:"many2many:project_members;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}