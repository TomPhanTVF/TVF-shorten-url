package repository

import (
	"context"
	"url-service/internal/models"
	"url-service/internal/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db *sqlx.DB
}


func NewURlRepo(db *sqlx.DB) URLRepo {
	return &UserRepository{
		db: db,
	}
}


func(u *UserRepository) Create(ctx context.Context, url *models.URL)(*models.URL, error){
	createUrl := &models.URL{}
	if err := u.db.QueryRowxContext(
		ctx,
		sql.CreateURLQuery,
		url.GenID(),
		url.Redirect,
		url.TVF,
		url.Random,
		url.UserID,
	).StructScan(createUrl); err != nil {
		return nil, errors.Wrap(err, "Create.QueryRowxContext")
	}
	return createUrl, nil
}

func(u *UserRepository) FindUrlsByEmail(ctx context.Context, email string)([]models.URL, error){
	urls := []models.URL{}
	err := u.db.SelectContext(ctx, &urls, sql.FindByOwnerIDQuery, email)
	if err != nil{
		return nil, errors.Wrap(err, "FindUrlByEmail.GetContext")
	}
	return urls, nil
}

func(u *UserRepository) FindUrlById(ctx context.Context, userId string)(*models.URL, error){
	url := &models.URL{}
	if err := u.db.GetContext(ctx, url, sql.FindByIDQuery, userId); err != nil {
		return nil, errors.Wrap(err, "FindById.GetContext")
	}
	return url, nil 
}
