package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type TweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

// just used for response... would inline work?
type IDHolder struct {
	ID int `json:"ID"`
}

type TweetRepository interface {
	addTweet() int
	getTweets() (TweetsList, error)
}

type TweetMemoryRepository struct {
	id     int
	tweets TweetsList
}

type TweetServer struct {
	repo TweetMemoryRepository
}

func (repo *TweetMemoryRepository) addTweet(tweet Tweet) int {
	repo.id++
	repo.tweets.Tweets = append(repo.tweets.Tweets, tweet)
	return repo.id
}

func (repo TweetMemoryRepository) getTweets() (TweetsList, error) {
	return repo.tweets, nil
}

func (srv *TweetServer) tweets(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		srv.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		srv.listTweets(w, r)
	}
}

func (srv TweetServer) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, _ := srv.repo.getTweets()
	payload, err := json.Marshal(tweets)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

func (srv *TweetServer) addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	u := Tweet{}

	if err := json.Unmarshal(body, &u); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("Tweet: `%s` from %s\n", u.Message, u.Location)

	// each tweet has a unique ID which we implement in the simplest way possible
	id := srv.repo.addTweet(u)
	payload, err := json.Marshal(IDHolder{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

func main() {
	server := TweetServer{}
	http.HandleFunc("/tweets", server.tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
