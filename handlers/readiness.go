package handlers

import "net/http"

func Readiness(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, struct{}{})
}
