package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func InitDiscord(token string) (*discordgo.Session, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	err = discord.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Discord connection established")
	return discord, nil
}
