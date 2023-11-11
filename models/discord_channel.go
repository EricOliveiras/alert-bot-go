package models

import (
	"time"

	"github.com/google/uuid"
)

type DiscordChannel struct {
	ID          uuid.UUID `db:"id,uuid,pk"`
	Name        string    `db:"name"`
	ChannelId   string    `db:"channel_id,unique"`
	ServerId    string    `db:"server_id,unique"`
	StreamLimit int64     `db:"stream_limit"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
