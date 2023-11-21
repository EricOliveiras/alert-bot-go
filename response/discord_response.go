package response

type GuildResponse struct {
	Icon     string `json:"icon"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Owner    bool   `json:"owner"`
	Channels []GuildChannelsResponse
}

type GuildChannelsResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}
