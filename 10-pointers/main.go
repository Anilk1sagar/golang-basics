package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Pointers")

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T , %T\n", a, b)

	// Use * to read val from address
	fmt.Println(*b, *&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
