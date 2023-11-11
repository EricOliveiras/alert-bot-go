package models

import (
	"github.com/google/uuid"
)

type DiscordChannelStream struct {
	ID               int64     `db:"id,pk,autoincr"`
	DiscordChannelID uuid.UUID `db:"discord_channel_id,fk"`
	StreamID         int64     `db:"stream_id,fk"`
}
