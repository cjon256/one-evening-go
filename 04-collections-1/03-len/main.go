package main

import "fmt"

var (
	colors  = [5]string{"red", "green", "blue"}
	systems = []string{"linux", "macos", "windows"}
)

func NumberOfColors() int {
	return len(colors)
}

func NumberOfSystems() int {
	return len(systems)
}

func main() {
	fmt.Println(NumberOfColors(), colors)
	fmt.Println(NumberOfSystems(), systems)
}
