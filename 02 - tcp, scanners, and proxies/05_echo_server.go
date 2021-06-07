package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// Buffer to store received data
	b := make([]byte, 512)
	for {
		// Receive data
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// Bind to Port 20080 tcp
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Wait for connection
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}
}
