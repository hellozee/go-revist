package main

import "fmt"

func compose(a func(int) int, b func(int) int, c int) func() int {
	return func() int {
		return b(a(c))
	}
}

func main() {
	afunc := compose(func(a int) int { return a + 1 }, func(a int) int { return a + 2 }, 5)
	fmt.Println(afunc())
}
