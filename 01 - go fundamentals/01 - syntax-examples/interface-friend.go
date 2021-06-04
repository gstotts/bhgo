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

// Dog Struct
type Dog struct{}

// A Function of the Type Person
func (p *Person) SayHello() {
	fmt.Println("Hello, ", p.Name)
}

// A Function of the Type Dog
func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

// Interface Friend
type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	var bob = new(Person)
	bob.Name = "Bob"
	bob.SayHello()

	// Since Dave can SayHello, Dave is a Friend so we can use the Greet Function
	var dave = new(Person)
	dave.Name = "Dave"
	Greet(dave)

	// Since dog also uses SayHello, it is a Friend and can also use Greet
	var dog = new(Dog)
	Greet(dog)
}
