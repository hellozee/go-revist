package main

func product(nums ...int) int {
	result := 1

	for num := range nums {
		result *= num
	}

	return result
}

func main() {
	product(0, 1)
	product(0, 1, 2)
}
