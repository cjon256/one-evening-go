package main

import (
	"errors"
	"fmt"
)

func main() {
	err := RunSafely(func() {
		Divide(10, 0)
	})
	fmt.Println(err)
}

func Divide(x, y float64) float64 {
	if y == 0 {
		panic("division by zero")
	}

	return x / y
}

func RunSafely(f func()) (err error) {
	// execute the given function f, prevent any panics, and return an error if a panic happened.
	defer func() {
		r := recover()
		if r != nil {
			err = errors.New("Some Error")
		}
	}()
	f()
	err = nil // probably not needed
	return
}
