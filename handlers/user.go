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

func (apiCfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name:       params.Name,
	})

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Cannot create a user: %v", err))
		return
	}

	RespondWithJSON(w, 201, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) GetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
		Offset: utils.GetOffset(r),
	})

	if err != nil {
		ResponseWithError(w, 400, fmt.Sprintf("Cannot fetch posts: %v", err))
		return
	}

	RespondWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
