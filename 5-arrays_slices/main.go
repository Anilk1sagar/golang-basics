package main

import (
	"fmt"
)

func main() {

	// fmt.Println("Helo")

	// // Arrays
	// var fruitArr [2]string

	// // Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "orange"

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println("fruitArr: ", fruitArr)
	fmt.Println(fruitArr[1])

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Water", "grapes"}
	fmt.Println("fruitSlice: ", fruitSlice)

	//Get Length
	fmt.Println("Length of fruitslice Arr: ", len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
