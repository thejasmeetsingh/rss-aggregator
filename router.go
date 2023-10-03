package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/thejasmeetsingh/rss-aggregator/handlers"
)

func GetRouter(apiCfg handlers.ApiConfig) chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready/", handlers.Readiness)
	v1Router.Get("/err/", handlers.Error)
	v1Router.Post("/user/", apiCfg.CreateUser)
	v1Router.Get("/user/", apiCfg.MiddlewareAuth(apiCfg.GetUser))
	v1Router.Post("/feed/", apiCfg.MiddlewareAuth(apiCfg.CreateFeed))
	v1Router.Get("/feed/", apiCfg.GetFeeds)
	v1Router.Post("/feed-follow/", apiCfg.MiddlewareAuth(apiCfg.CreateFeedFollow))
	v1Router.Get("/feed-follow/", apiCfg.MiddlewareAuth(apiCfg.GetFeedFollow))
	v1Router.Delete("/feed-follow/{feedFollowID}/", apiCfg.MiddlewareAuth(apiCfg.DeleteFeedFollow))
	v1Router.Get("/posts/", apiCfg.MiddlewareAuth(apiCfg.GetPostsForUser))

	router.Mount("/v1", v1Router)

	return router
}
