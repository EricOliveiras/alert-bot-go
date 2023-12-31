package service

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ericoliveiras/alert-bot-go/builder"
	"github.com/ericoliveiras/alert-bot-go/discord"
	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/request"
	"github.com/ericoliveiras/alert-bot-go/utils"
)

type StreamServiceWrapper interface {
	Create(ctx context.Context, stream *request.StreamRequest) error
}

type StreamService struct {
	StreamRepository        repository.IStreamRepository
	DiscordRepository       repository.IDiscordRepository
	DiscordStreamRepository repository.IDiscordStreamRepository
	Session                 *discordgo.Session
}

func NewStreamService(
	streamRepository repository.IStreamRepository,
	discordRepository repository.IDiscordRepository,
	discordStreamRepository repository.IDiscordStreamRepository,
	session *discordgo.Session,
) *StreamService {
	return &StreamService{
		StreamRepository:        streamRepository,
		DiscordRepository:       discordRepository,
		DiscordStreamRepository: discordStreamRepository,
		Session: session,
	}
}

func (ss *StreamService) Create(ctx context.Context, stream *request.StreamRequest) error {
	channel, err := ss.DiscordRepository.GetChannelByID(ctx, stream.ChannelID)
	if err != nil {
		return errors.New("channel not found or not exist")
	}

	if channel.StreamLimit <= 0 {
		return errors.New("stream limit exceded")
	}

	streamByName, _ := ss.StreamRepository.GetByStreamName(ctx, strings.ToLower(stream.StreamName))

	if stream.StreamName == streamByName.Name {
		exists, err := ss.DiscordStreamRepository.ChannelHasStream(ctx, channel.ID, streamByName.ID)
		if err != nil {
			return err
		}

		if exists {
			return errors.New("stream already exists in the channel")
		}

		discordStream := &models.DiscordChannelStream{
			DiscordChannelID: channel.ChannelId,
			StreamID:         streamByName.ID,
		}

		err = ss.DiscordStreamRepository.Create(ctx, discordStream)
		if err != nil {
			return err
		}

		err = ss.DiscordRepository.UpdateStreamLimit(ctx, channel.ID, channel.StreamLimit-1)
		if err != nil {
			return err
		}

		discordChannels, err := ss.DiscordStreamRepository.GetAllByStreamID(context.Background(), streamByName.ID)
		if err != nil {
			log.Println("Error when searching for Discord channels associated with the stream:", err)
			return err
		}

		if streamByName.IsLive {
			discord.SendAlert(ss.Session, *streamByName, discordChannels)
		}

		return nil
	}

	streamResponse, err := utils.GetStream(strings.ToLower(stream.StreamName))
	if err != nil {
		return err
	}

	newStream := builder.NewStreamBuilder().
		SetName(strings.ToLower(streamResponse.Data[0].BroadcasterLogin)).
		SetImageUrl(streamResponse.Data[0].ThumbnailURL).
		SetIsLive(false).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Build()

	createdStream, err := ss.StreamRepository.Create(ctx, &newStream)
	if err != nil {
		return err
	}

	err = ss.DiscordRepository.UpdateStreamLimit(ctx, channel.ID, channel.StreamLimit-1)
	if err != nil {
		return err
	}

	discordStream := &models.DiscordChannelStream{
		DiscordChannelID: channel.ChannelId,
		StreamID:         createdStream.ID,
	}

	err = ss.DiscordStreamRepository.Create(ctx, discordStream)
	if err != nil {
		return err
	}

	return nil
}
