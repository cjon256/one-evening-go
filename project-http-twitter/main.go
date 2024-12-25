package main

import (
	"log"
	"net/http"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	server := server.TweetServer{}

	r.Use(middleware.Logger)
	r.Get("/tweets", server.ListTweets)
	r.Post("/tweets", server.AddTweet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
