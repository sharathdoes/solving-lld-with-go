package auth

import (
	"context"
	"errors"
	"time"
	"user-auth-service/internal/utils"

	"github.com/google/uuid"
)

type Service struct {
	repo       *Repository
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewService(r *Repository, s string, at, rt time.Duration) *Service {
	return &Service{r, s, at, rt}
}

func (s *Service) SignUp(ctx context.Context, username string, email string, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
	return err
	}
	user := &User{
		ID:       uuid.NewString(),
		Username: username, Email: email, Password: hashedPassword}
	return s.repo.CreateUser(ctx, user)
}

func (s *Service) Login(ctx context.Context, email, password string) (string, string, error) {
  user, err := s.repo.FindByEmail(ctx, email)
  if err != nil || !utils.ComparePassword(password, user.Password) {
    return "", "", err
  }

  access, _ := utils.GenerateJWT(user.ID, s.secret, s.accessTTL)

	refresh := uuid.NewString()
refreshHash := utils.HashToken(refresh)

  token := &RefreshTokenTTL{
    ID:        uuid.NewString(),
    UserId:    user.ID,
    TokenHash: refreshHash,
    ExpiresAt: time.Now().Add(s.refreshTTL),
  }

  _ = s.repo.SaveRefreshToken(ctx, token)

  return access, refresh, nil
}

func (s *Service) Refresh(
  ctx context.Context,
  refreshToken string,
) (string, string, error) {

  // hash incoming token
	hashed := utils.HashToken(refreshToken)

  // find token in DB
  stored, err := s.repo.FindRefreshToken(ctx, hashed)
  if err != nil {
    return "", "", err
  }

  if stored.ExpiresAt.Before(time.Now()) {
    return "", "", errors.New("refresh token expired")
  }

  // issue new access token
  access, _ := utils.GenerateJWT(
    stored.UserId,
    s.secret,
    s.accessTTL,
  )

  // ROTATION: revoke old token
  _ = s.repo.RevokeRefreshToken(ctx, stored.ID)

  // issue new refresh token
  newRefresh := uuid.NewString()
	newHash := utils.HashToken(newRefresh)

  _ = s.repo.SaveRefreshToken(ctx, &RefreshTokenTTL{
    ID:        uuid.NewString(),
    UserId:    stored.UserId,
    TokenHash: newHash,
    ExpiresAt: time.Now().Add(s.refreshTTL),
  })

  return access, newRefresh, nil
}
