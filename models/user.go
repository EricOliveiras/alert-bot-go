package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id,uuid,pk"`
	UserName     string    `db:"username"`
	Email        string    `db:"email,unique"`
	Password     string    `db:"password"`
	ChannelLimit int       `db:"channel_limit"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
