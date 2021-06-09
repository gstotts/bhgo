package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	target := os.Args[1]
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("%s:%d", target, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port closed / filtered
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
