package models

import (
	"github.com/google/uuid"
)

type DiscordChannelStream struct {
	DiscordChannelID uuid.UUID `db:"discord_channel_id,fk"`
	StreamID         int       `db:"stream_id,fk"`
}
