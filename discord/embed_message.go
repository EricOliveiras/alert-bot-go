package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func NewEmbedMessage(name, image string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:     0x0099FF,
		Title:     fmt.Sprintf("%s está online", name),
		URL:       fmt.Sprintf("https://twitch.tv/%s", name),
		Image:     &discordgo.MessageEmbedImage{URL: image},
		Timestamp: time.Now().Format(time.RFC3339),
		Footer:    &discordgo.MessageEmbedFooter{Text: "Alert Bot Go"},
	}
}
