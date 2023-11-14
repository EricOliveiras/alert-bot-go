package repository

import (
	"context"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users 
		(id, username, email, channel_limit, created_at, updated_at) 
	VALUES 
		(:id, :username, :email, :channel_limit, :created_at, :updated_at)
	`

	_, err := ur.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}
