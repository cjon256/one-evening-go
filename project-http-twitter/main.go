package main

import (
	"log"
	"net/http"
	"twitter/server"
)

func main() {
	server := server.TweetServer{}
	http.HandleFunc("/tweets", server.ServeTweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
