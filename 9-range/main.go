package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Range")

	ids := []int{22, 33, 44, 55, 66, 11}

	// Loop through ids
	for i, id := range ids {

		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {

		fmt.Printf("Id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {

		sum += id
	}

	fmt.Println("Sum is: ", sum)

	// Range with map
	emails := map[string]string{"anil": "anilk1sagar@gmail.com", "jelly": "jelly@gmail.com"}

	for k, v := range emails {

		fmt.Printf("%s: %s\n", k, v)
	}

}
