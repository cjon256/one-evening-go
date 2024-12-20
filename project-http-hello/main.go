package main

import (
	"fmt"
	"log"
	"net/http"
)

func yourFunction(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Hello, ", name)
}

func main() {
	http.HandleFunc("/hello", yourFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
