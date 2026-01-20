package auth

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
  db *gorm.DB
}


func NewRepository(db *gorm.DB) *Repository {
  return &Repository{db}
}

func (r *Repository) CreateUser(ctx context.Context, user *User) error {
 	 return r.db.WithContext(ctx).Create(user).Error
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err:=r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) SaveRefreshToken(ctx context.Context, token *RefreshTokenTTL) error {
	return r.db.WithContext(ctx).Create(token).Error
}


func (r *Repository) FindRefreshToken(
  ctx context.Context,
  tokenHash string,
) (*RefreshTokenTTL, error) {
  var rt RefreshTokenTTL
  err := r.db.WithContext(ctx).
    Where("token_hash = ? AND revoked = false", tokenHash).
    First(&rt).Error
  return &rt, err
}

func (r *Repository) RevokeRefreshToken(
  ctx context.Context,
  id string,
) error {
  return r.db.WithContext(ctx).
    Model(&RefreshTokenTTL{}).
    Where("id = ?", id).
    Update("revoked", true).Error
}
