package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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

func (srv *TweetServer) ServeTweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodPost {
		srv.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		srv.listTweets(w)
	}
}

func (srv TweetServer) listTweets(w http.ResponseWriter) {
	tweets, _ := srv.repo.getTweets()
	payload, err := json.Marshal(tweets)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

/*
* Log requests using the following format:

Call Printf inside of a anonymous function passed to defer.

To measure duration (how long a request took), call start := time.Now() at the beginning of the handler and duration := time.Since(start) in defer.

*/

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
