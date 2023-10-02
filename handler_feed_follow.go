package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		UserID:     user.ID,
		FeedID:     params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error while try to follow a feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

// func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request) {
// 	feeds, err := apiCfg.DB.GetFeeds(r.Context())

// 	if err != nil {
// 		responseWithError(w, 400, fmt.Sprintf("Cannot fetch the feeds: %v", err))
// 		return
// 	}

// 	respondWithJSON(w, 200, databaseFeedsToFeed(feeds))
// }
