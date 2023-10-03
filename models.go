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

type FeedFollow struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	UserID     uuid.UUID `json:"user_id"`
	FeedID     uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:         dbFeedFollow.ID,
		CreatedAt:  dbFeedFollow.CreatedAt,
		ModifiedAt: dbFeedFollow.ModifiedAt,
		UserID:     dbFeedFollow.UserID,
		FeedID:     dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}

	return feedFollows
}

type Posts struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Posts {
	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Posts{
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

func databasePostsToPosts(dbPosts []database.Post) []Posts {
	posts := []Posts{}

	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}

	return posts
}
