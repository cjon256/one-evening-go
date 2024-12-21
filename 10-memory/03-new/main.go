package main

import "fmt"

var BuffersAllocated int

func AllocateBuffer() *string {
	if BuffersAllocated > 2 {
		return nil
	}
	BuffersAllocated++
	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
