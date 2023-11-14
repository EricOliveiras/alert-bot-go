package builder

import (
	"time"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/google/uuid"
)

type UserBuilder struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Password     string
	ChannelLimit int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (userBuilder *UserBuilder) SetID(id uuid.UUID) *UserBuilder {
	userBuilder.ID = id
	return userBuilder
}

func (userBuilder *UserBuilder) SetUsername(username string) *UserBuilder {
	userBuilder.Username = username
	return userBuilder
}

func (userBuilder *UserBuilder) SetEmail(email string) *UserBuilder {
	userBuilder.Email = email
	return userBuilder
}

func (userBuilder *UserBuilder) SetChannelLimit(channelLimit int) *UserBuilder {
	userBuilder.ChannelLimit = channelLimit
	return userBuilder
}

func (userBuilder *UserBuilder) SetCreatedAt(createdAt time.Time) *UserBuilder {
	userBuilder.CreatedAt = createdAt
	return userBuilder
}

func (userBuilder *UserBuilder) SetUpdatedAt(updatedAt time.Time) *UserBuilder {
	userBuilder.UpdatedAt = updatedAt
	return userBuilder
}

func (userBuilder *UserBuilder) Build() models.User {
	user := models.User{
		ID:           userBuilder.ID,
		Username:     userBuilder.Username,
		Email:        userBuilder.Email,
		ChannelLimit: userBuilder.ChannelLimit,
		CreatedAt:    userBuilder.CreatedAt,
		UpdatedAt:    userBuilder.UpdatedAt,
	}

	return user
}
