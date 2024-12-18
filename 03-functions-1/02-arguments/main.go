package main

import "fmt"

func Greet(s string) {
	fmt.Println("Hello, " + s)
}

func main() {
	Greet("Alice")
	Greet("Bob")
}
