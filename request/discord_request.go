package request

type CreateDiscordChannel struct {
	Name      string `json:"name"`
	ChannelId string `json:"channel_id"`
	ServerId  string `json:"server_id"`
	DiscordId string `json:"discord_id"`
}
