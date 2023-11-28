package builder

import (
	"time"

	"github.com/ericoliveiras/alert-bot-go/models"
)

type StreamBuilder struct {
	Name      string
	ImageUrl  string
	IsLive    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStreamBuilder() *StreamBuilder {
	return &StreamBuilder{}
}

func (streamBuilder *StreamBuilder) SetName(name string) *StreamBuilder {
	streamBuilder.Name = name
	return streamBuilder
}

func (streamBuilder *StreamBuilder) SetImageUrl(imageUrl string) *StreamBuilder {
	streamBuilder.ImageUrl = imageUrl
	return streamBuilder
}

func (streamBuilder *StreamBuilder) SetIsLive(isLive bool) *StreamBuilder {
	streamBuilder.IsLive = isLive
	return streamBuilder
}

func (streamBuilder *StreamBuilder) SetCreatedAt(createdAt time.Time) *StreamBuilder {
	streamBuilder.CreatedAt = createdAt
	return streamBuilder
}

func (streamBuilder *StreamBuilder) SetUpdatedAt(updatedAt time.Time) *StreamBuilder {
	streamBuilder.UpdatedAt = updatedAt
	return streamBuilder
}

func (streamBuilder *StreamBuilder) Build() models.Stream {
	stream := models.Stream{
		Name:      streamBuilder.Name,
		ImageUrl:  streamBuilder.ImageUrl,
		IsLive:    streamBuilder.IsLive,
		CreatedAt: streamBuilder.CreatedAt,
		UpdatedAt: streamBuilder.UpdatedAt,
	}

	return stream
}
