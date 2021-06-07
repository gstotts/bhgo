package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

type Result struct {
	Port           int
	ConnectionType string
}

func sendResult(ch chan Result, r Result) {
	ch <- r
}

func scanner_worker(target string, connType string, ports chan int, results chan Result) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.Dial(connType, address)
		if err != nil {
			go sendResult(results, Result{0, ""})
			continue
		}
		conn.Close()
		go sendResult(results, Result{p, connType})
	}
}

func main() {
	// Grab targets
	target := os.Args[1]
	if target == "" {
		errors.New("Error: No Target Specified")
		os.Exit(1)
	}

	ports := make(chan int, 100)
	results := make(chan Result)
	var openports []Result

	for i := 0; i < cap(ports); i++ {
		go scanner_worker(target, "tcp", ports, results)
		go scanner_worker(target, "udp", ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port.Port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	// Need to sort ports by Numbers across groups
	// Need to identify services

	for _, result := range openports {
		fmt.Printf("%d %s open\n", result.Port, result.ConnectionType)
	}
}
