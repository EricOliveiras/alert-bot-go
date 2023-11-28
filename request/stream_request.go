package request

import "github.com/google/uuid"

type StreamRequest struct {
	StreamName string    `json:"stream_name"`
	ChannelID  uuid.UUID `json:"channel_id"`
}
