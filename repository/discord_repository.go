package repository

import (
	"context"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IDiscordRepository interface {
	Create(ctx context.Context, discordChannel *models.DiscordChannel) error
	GetChannelByUserID(ctx context.Context, userID uuid.UUID) (*models.DiscordChannel, error)
}

type DiscordRepository struct {
	DB *sqlx.DB
}

func NewDiscordRepository(db *sqlx.DB) *DiscordRepository {
	return &DiscordRepository{DB: db}
}

func (dr *DiscordRepository) Create(ctx context.Context, discordChannel *models.DiscordChannel) error {
	query := `INSERT INTO discord_channels 
		(id, name, channel_id, server_id, stream_limit, user_id, created_at, updated_at) 
	VALUES 
		(:id, :name, :channel_id, :server_id, :stream_limit, :user_id, :created_at, :updated_at)`

	_, err := dr.DB.NamedExecContext(ctx, query, discordChannel)
	if err != nil {
		return err
	}

	return nil
}

func (dr *DiscordRepository) GetChannelByUserID(ctx context.Context, userID uuid.UUID) (*models.DiscordChannel, error) {
	var channel models.DiscordChannel

	query := "SELECT * FROM discord_channels WHERE user_id = $1"
	if err := dr.DB.GetContext(ctx, &channel, query, userID); err != nil {
		return &models.DiscordChannel{}, err
	}

	return &channel, nil
}
