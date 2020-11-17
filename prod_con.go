package main

import (
	"fmt"
	"time"
)

func consumer(c chan int) {
	for {
		fmt.Println("Consumed: ", <-c)
		time.Sleep(1 * time.Second)
	}
}

func producer(c chan int) {
	i := 0
	for {
		fmt.Println("Produced: ", i)
		c <- i
		i++
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	c := make(chan int, 10)
	go producer(c)
	consumer(c)
}
