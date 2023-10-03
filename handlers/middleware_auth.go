package handlers

import (
	"fmt"
	"net/http"

	"github.com/thejasmeetsingh/rss-aggregator/internal/auth"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"
)

type authHeaderHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authHeaderHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIkey(r.Header)

		if err != nil {
			ResponseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			ResponseWithError(w, 403, fmt.Sprintf("Not able to fetch the user with given: %v", err))
			return
		}

		handler(w, r, user)
	}
}
