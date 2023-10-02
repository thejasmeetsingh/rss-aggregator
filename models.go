package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:         dbUser.ID,
		CreatedAt:  dbUser.CreatedAt,
		ModifiedAt: dbUser.ModifiedAt,
		Name:       dbUser.Name,
	}
}
