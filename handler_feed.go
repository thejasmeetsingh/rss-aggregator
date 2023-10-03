package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name:       params.Name,
		Url:        params.URL,
		UserID:     user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot create a feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context(), database.GetFeedsParams{
		Limit:  10,
		Offset: GetOffset(r),
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot fetch the feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeed(feeds))
}
