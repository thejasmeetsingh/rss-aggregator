package handlers

import "net/http"

func Error(w http.ResponseWriter, r *http.Request) {
	ResponseWithError(w, 400, "Something went wrong")
}
