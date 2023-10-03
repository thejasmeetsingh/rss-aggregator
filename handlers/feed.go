package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
	"github.com/thejasmeetsingh/rss-aggregator/models"
	"github.com/thejasmeetsingh/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
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
		ResponseWithError(w, 400, fmt.Sprintf("Cannot create a feed: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context(), database.GetFeedsParams{
		Limit:  10,
		Offset: utils.GetOffset(r),
	})

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Cannot fetch the feeds: %v", err))
		return
	}

	RespondWithJSON(w, 200, models.DatabaseFeedsToFeed(feeds))
}
