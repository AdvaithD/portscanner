package main

import (
	"fmt"
	"sync"
)

// use a pool of goroutines to manage the concurrent work being perfomed
// take two args: chan int and a WaitGroup pointer, channel used to receive
// work and WaitGroup used to track work completion state
func worker(ports chan int, wg *sync.WaitGroup) {
	// loop over until channel is closed
	for p := range ports { // continuously receive from the ports channel (range)
		fmt.Println(p) // interesting note: the numbers are not printed in order
		wg.Done()
	}
}

// entrypoint
func main() {
	// buffered channel allows workers to start immediately
	// create channel (buffered, since it has 100 items before sender blocks)
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	// kickstart the workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i // send an int to the ports channel
	}
	wg.Wait()
	close(ports)
}
