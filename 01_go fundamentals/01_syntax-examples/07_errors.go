package main

import (
	"errors"
	"fmt"
)

func foo() error {
	return errors.New("An Error Occurred!  Panic time!!!!")
}
func main() {
	if err := foo(); err != nil {
		fmt.Println(err)
	}
}
