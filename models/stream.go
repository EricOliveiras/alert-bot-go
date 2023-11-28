package models

import "time"

type Stream struct {
	ID              int              `db:"id,autoincr"`
	Name            string           `db:"name,unique"`
	ImageUrl        string           `db:"image_url"`
	IsLive          bool             `db:"is_live"`
	DiscordChannels []DiscordChannel `db:"-"`
	CreatedAt       time.Time        `db:"created_at"`
	UpdatedAt       time.Time        `db:"updated_at"`
}
