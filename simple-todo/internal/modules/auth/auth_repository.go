package auth

import (
	"context"
	"simple-todo/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

/*
What context actually is (plain English)

context.Context is a tiny object that carries four powers across function calls:

Cancellation

Timeout / Deadline

Request-scoped data

Signal propagation

Why Go needed context in the first place

Go loves:

Goroutines

Concurrency

Long-running operations

Without context, once you start something, you can’t tell it to stop.

That’s wasted CPU, memory, DB connections.

Context exists to say:

“Hey, this work belongs to that request. If the request dies, kill the work.”

*/

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

// we connecto db to our repo here

func (r *Repository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *Repository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *Repository) FindByIds(ctx context.Context, member_ids []string) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Where("id IN ?", member_ids).Find(&users).Error
	return users, err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) SaveRefreshToken(ctx context.Context, token *models.RefreshTokenTTL) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *Repository) FindRefreshToken(
	ctx context.Context,
	tokenHash string,
) (*models.RefreshTokenTTL, error) {
	var rt models.RefreshTokenTTL
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
		Model(&models.RefreshTokenTTL{}).
		Where("id = ?", id).
		Update("revoked", true).Error
}
