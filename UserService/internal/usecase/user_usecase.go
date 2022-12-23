package usecase

import (
	models "user-service/internal/models"
	"user-service/internal/service"
	"user-service/internal/repository/postgres"
	"user-service/internal/utils"
	"github.com/pkg/errors"

	"context"
)

const (
	userByIdCacheDuration = 3600
)


// User UseCase
type userUseCase struct {
	userPgRepo postgres.UserPGRepository
}

// New User UseCase
func NewUserUseCase(userRepo postgres.UserPGRepository) service.UserService {
	return &userUseCase{userPgRepo: userRepo}
}


// Register new user
func (u *userUseCase) Register(ctx context.Context, user *models.User) (*models.User, error) {
	existsUser, err := u.userPgRepo.FindByEmail(ctx, user.Email)
	if existsUser != nil || err == nil {
		return nil, utils.ErrEmailExists
	}

	return u.userPgRepo.Create(ctx, user)
}

// Find use by email address
func (u *userUseCase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	findByEmail, err := u.userPgRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "userPgRepo.FindByEmail")
	}

	findByEmail.SanitizePassword()

	return findByEmail, nil
}

// Find use by uuid
func (u *userUseCase) FindById(ctx context.Context, userID string) (*models.User, error) {
	foundUser, err := u.userPgRepo.FindById(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "userPgRepo.FindById")
	}

	return foundUser, nil
}

// Login user with email and password
func (u *userUseCase) Login(ctx context.Context, email string, password string) (*models.User, error) {
	foundUser, err := u.userPgRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "userPgRepo.FindByEmail")
	}

	if err := foundUser.ComparePasswords(password); err != nil {
		return nil, errors.Wrap(err, "user.ComparePasswords")
	}

	return foundUser, err
}
