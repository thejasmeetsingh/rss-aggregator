package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

type posts struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DatabasePostToPost(dbPost database.Post) posts {
	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return posts{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		ModifiedAt:  dbPost.ModifiedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func DatabasePostsToPosts(dbPosts []database.Post) []posts {
	posts := []posts{}

	for _, post := range dbPosts {
		posts = append(posts, DatabasePostToPost(post))
	}

	return posts
}
