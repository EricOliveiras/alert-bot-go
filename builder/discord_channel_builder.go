package builder

import (
	"time"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/google/uuid"
)

type DiscordChannelBuilder struct {
	ID          uuid.UUID
	Name        string
	ChannelId   string
	ServerId    string
	StreamLimit int
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewDiscordChannelBuilder() *DiscordChannelBuilder {
	return &DiscordChannelBuilder{}
}

func (discordChannelBuilder *DiscordChannelBuilder) SetID(id uuid.UUID) *DiscordChannelBuilder {
	discordChannelBuilder.ID = id
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetName(name string) *DiscordChannelBuilder {
	discordChannelBuilder.Name = name
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetChannelId(channelId string) *DiscordChannelBuilder {
	discordChannelBuilder.ChannelId = channelId
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetServerId(serverId string) *DiscordChannelBuilder {
	discordChannelBuilder.ServerId = serverId
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetStreamLimit(streamLimit int) *DiscordChannelBuilder {
	discordChannelBuilder.StreamLimit = streamLimit
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetUserId(userId uuid.UUID) *DiscordChannelBuilder {
	discordChannelBuilder.UserID = userId
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetCreatedAt(createdAt time.Time) *DiscordChannelBuilder {
	discordChannelBuilder.CreatedAt = createdAt
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) SetUpdatedAt(updatedAt time.Time) *DiscordChannelBuilder {
	discordChannelBuilder.UpdatedAt = updatedAt
	return discordChannelBuilder
}

func (discordChannelBuilder *DiscordChannelBuilder) Build() models.DiscordChannel {
	discord := models.DiscordChannel{
		ID:          discordChannelBuilder.ID,
		Name:        discordChannelBuilder.Name,
		ChannelId:   discordChannelBuilder.ChannelId,
		ServerId:    discordChannelBuilder.ServerId,
		StreamLimit: discordChannelBuilder.StreamLimit,
		UserID:      discordChannelBuilder.UserID,
		CreatedAt:   discordChannelBuilder.CreatedAt,
		UpdatedAt:   discordChannelBuilder.UpdatedAt,
	}

	return discord
}
