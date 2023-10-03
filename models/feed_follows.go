package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

type feedFollow struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	UserID     uuid.UUID `json:"user_id"`
	FeedID     uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) feedFollow {
	return feedFollow{
		ID:         dbFeedFollow.ID,
		CreatedAt:  dbFeedFollow.CreatedAt,
		ModifiedAt: dbFeedFollow.ModifiedAt,
		UserID:     dbFeedFollow.UserID,
		FeedID:     dbFeedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []feedFollow {
	feedFollows := []feedFollow{}

	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(feedFollow))
	}

	return feedFollows
}
