package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	calls = []string{}
	stats = map[string]int{}
)

func updateServerStats(name string) {
	calls = append(calls, name)
	_, ok := stats[name]
	if !ok {
		stats[name] = 0
	}
	stats[name] += 1
}

func yourFunction(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	updateServerStats(name)
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello, ", name)
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
}

func main() {
	http.HandleFunc("/hello", yourFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
