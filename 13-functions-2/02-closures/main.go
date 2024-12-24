package main

import "fmt"

func WordGenerator(words []string) func() string {
	max := len(words)
	index := max // so that it is immediately reset
	return func() string {
		index++
		if index >= max {
			index = 0
		}
		return words[index]
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
