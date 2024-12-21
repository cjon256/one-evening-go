package main

func Sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}
