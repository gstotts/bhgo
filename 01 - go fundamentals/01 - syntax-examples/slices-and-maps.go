package main

import (
	"fmt"
)

func main() {
	// Slice Example
	var s = make([]string, 0)
	s = append(s, "some string")
	fmt.Println(s)

	// Map Example
	var m = make(map[string]string)
	m["some key"] = "some value"
	fmt.Println(m["some key"])
}
