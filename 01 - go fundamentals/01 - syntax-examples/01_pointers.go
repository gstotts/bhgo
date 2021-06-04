package main

import (
	"fmt"
)

func main() {
	var count = int(42)
	fmt.Println("Count Variable:		", count)

	// Pointer
	ptr := &count
	fmt.Println("Pointer Location:	", ptr)
	fmt.Println("Value at Location:	", *ptr)

	// Set New Value at Location
	*ptr = 100
	fmt.Println("New Value of Count:	", count)
}
