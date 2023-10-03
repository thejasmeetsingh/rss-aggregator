package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

type user struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	APIKey     string    `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) user {
	return user{
		ID:         dbUser.ID,
		CreatedAt:  dbUser.CreatedAt,
		ModifiedAt: dbUser.ModifiedAt,
		Name:       dbUser.Name,
		APIKey:     dbUser.ApiKey,
	}
}
