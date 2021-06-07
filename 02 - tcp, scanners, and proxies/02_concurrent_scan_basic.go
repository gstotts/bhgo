package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(target string, ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// Pull the target as an argument - Use scanme.nmap.org as test
	target := os.Args[1]

	// Make two channels - one for ports to scan, one for results returned
	ports := make(chan int, 100)
	results := make(chan int)
	// Variable to store open ports
	var openports []int

	// Start workers and loop through as you go
	for i := 0; i < cap(ports); i++ {
		go worker(target, ports, results)
	}

	// Handle ports to scan
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// Handle results
	for i := 0; i < 1024; i++ {
		port := <-results
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
