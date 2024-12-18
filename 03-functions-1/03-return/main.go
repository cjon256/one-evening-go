package main

import "fmt"

func Sum(i1, i2, i3, i4, i5 int) int {
	return i1 + i2 + i3 + i4 + i5
}

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}
