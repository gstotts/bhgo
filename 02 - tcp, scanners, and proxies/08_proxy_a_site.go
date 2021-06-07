package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func handle(src net.Conn, address string) {
	fmt.Printf("Received request, sending to %s\n", address)
	dst, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	fmt.Printf("Returning Request from %s\n", address)
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	address := fmt.Sprintf("%s:443", os.Args[1])
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn, address)
	}
}
