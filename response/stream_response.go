package response

type StreamResponse struct {
	Data []struct {
		BroadcasterLanguage string   `json:"broadcaster_language"`
		BroadcasterLogin    string   `json:"broadcaster_login"`
		DisplayName         string   `json:"display_name"`
		GameID              string   `json:"game_id"`
		GameName            string   `json:"game_name"`
		ID                  string   `json:"id"`
		IsLive              bool     `json:"is_live"`
		TagIDs              []string `json:"tag_ids"`
		Tags                []string `json:"tags"`
		ThumbnailURL        string   `json:"thumbnail_url"`
		Title               string   `json:"title"`
		StartedAt           string   `json:"started_at"`
	} `json:"data"`
}
