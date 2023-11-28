package repository

import (
	"context"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IDiscordStreamRepository interface {
	Create(ctx context.Context, discord_stream *models.DiscordChannelStream) error
	ChannelHasStream(ctx context.Context, channelID uuid.UUID, streamID int) (bool, error)
}

type DiscordStreamRepository struct {
	DB *sqlx.DB
}

func NewDiscordStreamRepository(db *sqlx.DB) *DiscordStreamRepository {
	return &DiscordStreamRepository{DB: db}
}

func (dsr *DiscordStreamRepository) Create(ctx context.Context, discordStream *models.DiscordChannelStream) error {
	query := `
	INSERT INTO discord_channel_streams (discord_channel_id, stream_id) 
	VALUES (:discord_channel_id, :stream_id)
`
	_, err := dsr.DB.NamedExecContext(ctx, query, discordStream)
	if err != nil {
		return err
	}

	return nil
}

func (dsr *DiscordStreamRepository) ChannelHasStream(ctx context.Context, channelID uuid.UUID, streamID int) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM discord_channel_streams
			WHERE discord_channel_id = $1
			AND stream_id = $2
		)
	`

	var exists bool
	err := dsr.DB.QueryRowContext(ctx, query, channelID, streamID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
