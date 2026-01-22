package database

import (
	"simple-todo/internal/modules/projects"
    "simple-todo/internal/modules/auth"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) (*gorm.DB, error) {
  db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
  
  if err != nil {
    return nil, err
  }
   db.AutoMigrate(
    &auth.User{},
    &auth.RefreshTokenTTL{},
    &projects.Project{},
    )

   return db,nil
}