package response

import (
	"time"

	"github.com/google/uuid"
)

type GuildResponse struct {
	Icon     string `json:"icon"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Owner    bool   `json:"owner"`
	Channels []GuildChannelsResponse
}

type GuildChannelsResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

type ChannelResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ChannelId   string    `json:"channel_id"`
	ServerId    string    `json:"server_id"`
	StreamLimit int       `json:"stream_limit"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
