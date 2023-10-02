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
	APIKey     string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:         dbUser.ID,
		CreatedAt:  dbUser.CreatedAt,
		ModifiedAt: dbUser.ModifiedAt,
		Name:       dbUser.Name,
		APIKey:     dbUser.ApiKey,
	}
}

type Feed struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	UserID     uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:         dbFeed.ID,
		CreatedAt:  dbFeed.CreatedAt,
		ModifiedAt: dbFeed.ModifiedAt,
		Name:       dbFeed.Name,
		Url:        dbFeed.Url,
		UserID:     dbFeed.UserID,
	}
}

func databaseFeedsToFeed(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}

	return feeds
}
