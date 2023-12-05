package discord

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/ericoliveiras/alert-bot-go/repository"
	"github.com/ericoliveiras/alert-bot-go/utils"
	"github.com/jmoiron/sqlx"
)

func Bot(ctx context.Context, dg *discordgo.Session, db *sqlx.DB) {
	var streamRepository = repository.NewStreamRepository(db)
	
	streams, err := streamRepository.GetAllStreams(ctx)
	if err != nil {
		log.Println("Error getting streams:", err)
		return
	}

	for _, stream := range streams {
		fetchStream, err := utils.GetStream(stream.Name)
		if err != nil {
			log.Println("Error when fetching stream data:", err)
			continue
		}

		CheckAndNotify(dg, stream, fetchStream.Data[0].IsLive, db)
	}
}

func CheckAndNotify(dg *discordgo.Session, stream models.Stream, streamerIsLive bool, db *sqlx.DB) {
	var discordStreamRepository = repository.NewDiscordStreamRepository(db)
	var streamRepository = repository.NewStreamRepository(db)

	discordChannels, err := discordStreamRepository.GetAllByStreamID(context.Background(), stream.ID)
	if err != nil {
		log.Println("Error when searching for Discord channels associated with the stream:", err)
		return
	}

	if streamerIsLive && !stream.IsLive {
		sendAlert(dg, stream, discordChannels)
		err := streamRepository.UpdateStreamIsLive(context.Background(), stream.ID, true)
		if err != nil {
			log.Println("Error updating stream status in database:", err)
		}
	} else if !streamerIsLive && stream.IsLive {
		err := streamRepository.UpdateStreamIsLive(context.Background(), stream.ID, false)
		if err != nil {
			log.Println("Error updating stream status in database:", err)
		}
	}
}

func sendAlert(dg *discordgo.Session, stream models.Stream, discordChannels []models.DiscordChannelStream) {
	for _, discordChannel := range discordChannels {
		channel, err := dg.State.Channel(discordChannel.DiscordChannelID)
		if err != nil {
			log.Println("Error retrieving Discord channel:", err)
			continue
		}

		embed := NewEmbedMessage(stream.Name, stream.ImageUrl)

		_, err = dg.ChannelMessageSendEmbed(channel.ID, embed)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}
