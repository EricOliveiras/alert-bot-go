package models

type DiscordChannelStream struct {
	DiscordChannelID string `db:"discord_channel_id,fk"`
	StreamID         int    `db:"stream_id,fk"`
}
