package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not configured in the enviorment")
	}

	router := chi.NewRouter()
	address := ":" + port

	srv := &http.Server{
		Handler: router,
		Addr:    address,
	}

	log.Printf("Server starting at port: %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
