package service

import (
	"context"
	model "user-service/internal/models"
)


//  User service interface
type UserService interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, email string, password string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindById(ctx context.Context, userID string) (*model.User, error)
}
