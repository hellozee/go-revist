package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 9
	fmt.Println(<-c)
	close(c)
	// this should panic
	c <- 3
}
