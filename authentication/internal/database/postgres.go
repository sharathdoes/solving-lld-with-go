package database

import (
  "gorm.io/driver/postgres"	
  "gorm.io/gorm"
    "user-auth-service/internal/modules/auth"

)

func Connect(url string) (*gorm.DB, error) {
  db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
  if err != nil {
    return nil, err
  }

  db.AutoMigrate(
    &auth.User{},
    &auth.RefreshTokenTTL{},
  )

   return db,nil
}