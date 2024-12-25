package server

import (
	"encoding/json"
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

func (lis *TweetsList) Append(t Tweet) {
	lis.Tweets = append(lis.Tweets, t)
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

func (repo *TweetMemoryRepository) addTweet(t Tweet) int {
	repo.id++
	repo.tweets.Append(t)
	return repo.id
}

func (repo *TweetMemoryRepository) GetTweets() TweetsList {
	return repo.tweets
}

type TweetServer struct {
	Repo TweetMemoryRepository
}

func (srv *TweetServer) ListTweets(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payload, err := json.Marshal(srv.Repo.GetTweets())
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Printf("%+v", string(payload))
	w.Write(payload)
}

func (srv *TweetServer) AddTweet(w http.ResponseWriter, r *http.Request) {
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

	// each tweet has a unique ID which we implement in the simplest way possible
	id := srv.Repo.addTweet(u)

	payload, err := json.Marshal(IDHolder{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}
