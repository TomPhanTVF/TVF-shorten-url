package service

import (
	"context"
	models "user-service/internal/models"
)


//  User service interface
type UserService interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindById(ctx context.Context, userID string) (*models.User, error)
}
