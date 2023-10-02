package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thejasmeetsingh/rss-aggregator/internal/auth"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Name:       params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Cannot create a user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIkey(r.Header)

	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Not able to fetch the user with given: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
