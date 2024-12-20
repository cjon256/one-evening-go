package main

import "fmt"

var Stats = map[string]int{}

func CreateUser(user string) {
	fmt.Println("Creating user", user)
	_, ok := Stats["create"]
	if !ok {
		Stats["create"] = 0
	}
	Stats["create"] += 1
}

func UpdateUser(user string) {
	fmt.Println("Updating user", user)
	_, ok := Stats["update"]
	if !ok {
		Stats["update"] = 0
	}
	Stats["update"] += 1
}

func PurgeStats() {
	Stats["create"] = 0
	Stats["update"] = 0
}
