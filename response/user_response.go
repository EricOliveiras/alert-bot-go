package response

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	DiscordID    string    `json:"discord_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Avatar       string    `json:"avatar"`
	ChannelLimit int       `json:"channel_limit"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserGuildsResponse struct {
	User   UserResponse    `json:"user"`
	Guilds []GuildResponse `json:"guilds"`
}

type UserChannelResponse struct {
	User    UserResponse    `json:"user"`
	Channel ChannelResponse `json:"channel"`
}
