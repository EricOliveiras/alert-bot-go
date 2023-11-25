package service

import (
	"context"
	"errors"
	"time"

	"github.com/ericoliveiras/alert-bot-go/builder"
	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/google/uuid"
)

type DiscordChannelServiceWrapper interface {
	Create(ctx context.Context, discordChannel *request.CreateDiscordChannel) error
}

type DiscordChannelService struct {
	DiscordRepository repository.IDiscordRepository
	UserRepository    repository.IUserRepository
}

func NewDiscordChannelService(
	discordRepository repository.IDiscordRepository,
	userRepository repository.IUserRepository,
) *DiscordChannelService {
	return &DiscordChannelService{
		DiscordRepository: discordRepository,
		UserRepository:    userRepository,
	}
}

func (ds *DiscordChannelService) Create(ctx context.Context, discordChannel *request.CreateDiscordChannel) error {
	user, err := ds.UserRepository.GetByDiscordID(ctx, discordChannel.DiscordId)
	if err != nil {
		return err
	}

	if user.ChannelLimit <= 0 {
		return errors.New("channel limit exceeded")
	}

	discordChannelBuilder := builder.NewDiscordChannelBuilder().
		SetID(uuid.New()).
		SetName(discordChannel.Name).
		SetChannelId(discordChannel.ChannelId).
		SetServerId(discordChannel.ServerId).
		SetStreamLimit(3).
		SetUserId(user.ID).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	err = ds.UserRepository.UpdateChannelLimit(ctx, user.ID, user.ChannelLimit-1)
	if err != nil {
		return err
	}

	err = ds.DiscordRepository.Create(ctx, &discordChannelBuilder)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DiscordChannelService) GetChannelByUserID(ctx context.Context, userID uuid.UUID) (*models.DiscordChannel, error) {
	channel, err := ds.DiscordRepository.GetChannelByUserID(ctx, userID)
	if err != nil {
		return &models.DiscordChannel{}, err
	}

	return channel, nil
}
