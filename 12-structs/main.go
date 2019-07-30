package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	name   string
	city   string
	gender string
	age    int
}

// Greeting method (value reciever)
func (p Person) greet() string {

	return "Hello, my name is " + p.name + " and i am  " + strconv.Itoa(p.age)
}

// hasBithday method (pointer reciever)
func (p *Person) hasBithday() {

	p.age++
}

func main() {

	fmt.Println("Welcome to struct")

	// Init person using struct
	person1 := Person{name: "anil", city: "bangalore", gender: "m", age: 22}
	// Alternative
	// person1 := Person{"anil", "bangalore", "m", 22}

	// fmt.Println(person1)
	// fmt.Println(person1.name)
	// person1.city = "jaipur"
	// fmt.Println(person1)

	person1.hasBithday()
	fmt.Println(person1.greet())
}
