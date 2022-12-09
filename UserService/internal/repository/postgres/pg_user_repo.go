package postgres

import (
	"context"
	models "user-service/internal/models"
)

// User pg repository interfaces
type UserPGRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindById(ctx context.Context, userID string) (*models.User, error)
}
