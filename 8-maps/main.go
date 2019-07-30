package main

import (
	"fmt"
)

func main() {

	// fmt.Println("Welcome to Maps")

	// // Define map
	// emails := make(map[string]string)

	// // Assign Key Value to map
	// emails["anil"] = "anilk1sagar@gmail.com"
	// emails["jelly"] = "jelly@gmail.com"
	// emails["parshu"] = "parshu@gmail.com"

	// Declare map and Assign Key Value
	emails := map[string]string{"anil": "anilk1sagar@gmail.com", "jelly": "jelly@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["anil"])

	// Delete from map
	delete(emails, "jelly")
	fmt.Println(emails)
	fmt.Println(len(emails))
}
