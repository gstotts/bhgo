package main

import (
	"fmt"
)

// Type Switch
func foo(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Println("Uknown type!")
	}
}

func main() {

	// Basic If/Else Structure
	x := int(1)

	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	// Switch
	y := "flubber"

	switch y {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	// Type Switch Use-Case
	foo(x)
	foo(y)

	// Basic For Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Loop over a Collection
	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}
