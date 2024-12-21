package main

import "fmt"

func Alphabet(length int) []string {
	alpha := []string{}
	for i := 0; i < length; i++ {
		alpha = append(alpha, characterByIndex(i))
	}
	return alpha
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
