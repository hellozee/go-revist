package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numWorkers = 2

func someWork() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("Work Done")
}

func doWork(c chan func(), wg *sync.WaitGroup, worker int) {
	fmt.Println("Worker ", worker, " Launched")
	// tries to read from channel if not closed, raising a deadlock
	for f := range c {
		f()
	}
	wg.Done()
}

func main() {
	jobs := make(chan func(), 5)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		jobs <- someWork
	}
	close(jobs)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go doWork(jobs, &wg, i)
	}

	wg.Wait()
}
