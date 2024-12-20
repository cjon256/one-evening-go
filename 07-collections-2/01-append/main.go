package main

var users []string

func AddUser(name string) {
	users = append(users, name)
}

func main() {
	AddUser("Alice")
	AddUser("Bob")
}
