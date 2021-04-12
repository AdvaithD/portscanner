package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	// loop over until channel is closed
	for p := range ports { // continuously receive from the ports channel (range)
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0 // send zero to results (port is closed)
			continue
		}
		conn.Close()
		results <- p // send port if (port is open)
	}
}

// entrypoint
func main() {
	// buffered channel allows workers to start immediately
	// create channel (buffered, since it has 100 items before sender blocks)
	ports := make(chan int, 100)
	results := make(chan int) // receive scan results from workers
	var openports []int

	// kickstart the workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= 1024; i++ {
		port := <-results // send an int to the ports channel
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
