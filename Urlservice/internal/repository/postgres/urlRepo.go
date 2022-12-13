package repository

import (
	"context"
	"url-service/internal/models"
)


type URLRepo interface{
	Create(ctx context.Context, url *models.URL)(*models.URL, error)
	FindUrlsByEmail(ctx context.Context, email string)([]models.URL, error)
	FindUrlById(ctx context.Context, id string)(*models.URL, error)
}