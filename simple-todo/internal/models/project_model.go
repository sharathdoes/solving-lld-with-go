package models

import "time"

type Project struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string

	OwnerID string
	Owner   User `gorm:"foreignKey:OwnerID"`

	Members []User `gorm:"many2many:project_members;"`
	Tasks   []Task `gorm:"foreignKey:ProjectID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}