package main

import (
	"fmt"
	"log"
)

func main() {

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("\n%T %T %T %T\n", i, f, b, s) //Prints types of the variable

	fmt.Printf("%v   %v       %v %q   \n", i, f, b, s) //prints initial value of the variable

	/* Static and Dynamic Variables */
	var x float64 = 20.0

	log.Println("\nD is: ", x)

	var age int32 = 22
	const isCool = true

	//Shorthand
	// name := "anil"
	// email := "anilk1sagar@gmail.com"

	name, email := "anil", "anilk1sagar@gmail.com"

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T", isCool)

}
