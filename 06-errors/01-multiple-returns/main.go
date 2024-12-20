package main

import "fmt"

func Swap(a, b string) (string, string) {
	return b, a
}

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}
