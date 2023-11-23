package repository

import (
	"context"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByDiscordID(ctx context.Context, discordId string) (*models.User, error)
	UpdateChannelLimit(ctx context.Context, userId uuid.UUID, limit int) error
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
		(id, discord_id, username, email, avatar, channel_limit, created_at, updated_at) 
	VALUES 
		(:id, :discord_id, :username, :email, :avatar, :channel_limit, :created_at, :updated_at)
	`

	_, err := ur.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email = $1"
	err := ur.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (ur *UserRepository) GetByDiscordID(ctx context.Context, discordId string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE discord_id = $1"
	err := ur.DB.GetContext(ctx, &user, query, discordId)
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateChannelLimit(ctx context.Context, userId uuid.UUID, limit int) error {
	query := "UPDATE users SET channel_limit = $1 WHERE id = $2"

	_, err := ur.DB.ExecContext(ctx, query, limit, userId)
	if err != nil {
		return err
	}

	return nil
}
