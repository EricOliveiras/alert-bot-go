package response

import "github.com/ericoliveiras/alert-bot-go/request"

type UserGuildsResponse struct {
	User   request.CreateUser `json:"user"`
	Guilds []GuildResponse    `json:"guilds"`
}
