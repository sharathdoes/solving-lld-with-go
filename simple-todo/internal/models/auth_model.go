package models

import "time"

type User struct {
	ID        string   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}

type RefreshTokenTTL struct {
  ID     string   `gorm:"primaryKey"`
  UserId string `gorm:"index"`
  TokenHash string `gorm:"not null"`
  ExpiresAt time.Time
  Revoked   bool
  CreatedAt time.Time
}
