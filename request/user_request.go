package request

type CreateUser struct {
	DiscordID string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
