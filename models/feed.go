package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

type feed struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	UserID     uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(dbFeed database.Feed) feed {
	return feed{
		ID:         dbFeed.ID,
		CreatedAt:  dbFeed.CreatedAt,
		ModifiedAt: dbFeed.ModifiedAt,
		Name:       dbFeed.Name,
		Url:        dbFeed.Url,
		UserID:     dbFeed.UserID,
	}
}

func DatabaseFeedsToFeed(dbFeeds []database.Feed) []feed {
	feeds := []feed{}

	for _, feed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(feed))
	}

	return feeds
}
