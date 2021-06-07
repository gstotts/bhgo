package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	target := fmt.Sprintf("%s:%d", os.Args[1], os.Args[2])
	_, err := net.Dial("tcp", target)
	if err == nil {
		fmt.Println("Connection successful")
	}
}
