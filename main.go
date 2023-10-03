package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/thejasmeetsingh/rss-aggregator/handlers"
	"github.com/thejasmeetsingh/rss-aggregator/utils"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not configured in the enviorment")
	}

	dbConn := GetDBConn()
	apiCfg := handlers.ApiConfig{
		DB: dbConn,
	}

	go utils.StartScraping(dbConn, 10, time.Minute)

	router := GetRouter(apiCfg)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting at port: %s", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
