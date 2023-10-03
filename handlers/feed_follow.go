package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
	"github.com/thejasmeetsingh/rss-aggregator/models"
	"github.com/thejasmeetsingh/rss-aggregator/utils"
)

func (apiCfg *ApiConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
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
		ResponseWithError(w, 400, fmt.Sprintf("Error while try to follow a feed: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *ApiConfig) GetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollow(r.Context(), database.GetFeedFollowParams{
		UserID: user.ID,
		Limit:  10,
		Offset: utils.GetOffset(r),
	})

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Cannot fetch feed follows: %v", err))
		return
	}

	RespondWithJSON(w, 200, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Invalid feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Error caught while deleting feed follows: %v", err))
		return
	}

	RespondWithJSON(w, 200, struct{}{})
}
