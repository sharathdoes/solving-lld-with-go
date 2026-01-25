package database

import (
	"simple-todo/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) (*gorm.DB, error) {
  db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
  
  if err != nil {
    return nil, err
  }
   db.AutoMigrate(
    &models.User{},
    &models.RefreshTokenTTL{},
    &models.Project{},
    &models.Task{},
    )

   return db,nil
}