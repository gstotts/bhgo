package main

import (
	"fmt"
)

// Person Struct
type Person struct {
	// Struct fields
	Name string
	Age  string
}

// A Function of the Type Person
func (p *Person) SayHello() {
	fmt.Println("Hello, ", p.Name)
}

func main() {
	var bob = new(Person)
	bob.Name = "Bob"
	bob.SayHello()
}
