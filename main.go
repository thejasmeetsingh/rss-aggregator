package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/thejasmeetsingh/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not configured in the enviorment")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB credentials is not configured in the enviorment")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

	dbConn := database.New(conn)
	apiCfg := apiConfig{
		DB: dbConn,
	}

	go startScraping(dbConn, 10, time.Minute)

	router := chi.NewRouter()
	address := ":" + port

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready/", handlerReadiness)
	v1Router.Get("/err/", handlerError)
	v1Router.Post("/user/", apiCfg.handlerCreateUser)
	v1Router.Get("/user/", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feed/", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed/", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed-follow/", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed-follow/", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Router.Delete("/feed-follow/{feedFollowID}/", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    address,
	}

	log.Printf("Server starting at port: %s", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
