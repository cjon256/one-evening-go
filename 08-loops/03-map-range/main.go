package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func Keys(p map[int]string) []int {
	keys := []int{}
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

func Values(p map[int]string) []string {
	vals := []string{}
	for k := range p {
		vals = append(vals, p[k])
	}
	return vals
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}
