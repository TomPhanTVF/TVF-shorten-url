package redis


import (
	"context"

	models "user-service/internal/models"
)

// Auth Redis repository interface
type UserRedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.User, error)
	SetUserCtx(ctx context.Context, key string, seconds int, user *models.User) error
	DeleteUserCtx(ctx context.Context, key string) error
}
