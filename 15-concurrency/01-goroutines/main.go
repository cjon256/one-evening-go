package main

import (
	"fmt"
	"time"
)

/*
SignUp is a function that saves the user and sends them an email notification. However, both actions take quite long to complete, and they're executed in a sequence.

Make SaveUser and SendNotification run concurrently by using the go keyword. It will make SignUp run faster.
*/

func main() {
	SignUp("joe@example.com")

	time.Sleep(time.Second * 2)
}

func SignUp(email string) {
	go SaveUser(email)
	go SendNotification(email)
}

func SaveUser(email string) {
	fmt.Println("Saving user", email)

	// Takes a long time to process
	time.Sleep(time.Second)

	fmt.Println("User", email, "saved")
}

func SendNotification(email string) {
	fmt.Println("Sending notification to", email)

	// Takes a long time to process
	time.Sleep(time.Second)

	fmt.Println("Notification sent to", email)
}
