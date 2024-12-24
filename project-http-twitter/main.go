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

type IDHolder struct {
	ID int
}

type TweetRepository interface {
	getID() int
	tweet(w http.ResponseWriter, r *http.Request)
}

type TweetMemoryRepository struct {
	ID int
}

func (repo TweetMemoryRepository) getID() int {
	repo.ID++
	return repo.ID
}

func (repo TweetMemoryRepository) tweet(w http.ResponseWriter, r *http.Request) {
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

	// each tweet has a unique ID which we implement in the simlest way possible
	id := repo.getID()
	payload, err := json.Marshal(IDHolder{ID: id})
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

func main() {
	repo := TweetMemoryRepository{}
	http.HandleFunc("/tweets", repo.tweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
