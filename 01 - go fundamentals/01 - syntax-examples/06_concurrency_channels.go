package main

import (
	"fmt"
)

func strlen(s string, c chan int) {
	// Pass the lengths to channel c
	c <- len(s)
}

func main() {
	// Define c as a channel of type int
	c := make(chan int)

	// Run concurrently using go
	go strlen("Salutations", c)
	go strlen("World", c)

	// Pass the output of c to x and y
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
