package main

import (
	"errors"
	"fmt"
)

func Divide(q, d float64) (float64, error) {
	if d == 0 {
		return 0, errors.New("Illegal divide by zero error")
	}
	return q / d, nil
}

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}
