package repository

import (
	"context"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/jmoiron/sqlx"
)

type IDiscordRepository interface {
	Create(ctx context.Context, discordChannel *models.DiscordChannel) error
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
