package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin
type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	// Instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// User reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// User writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Write %d bytes to stdout\n", s)
}
