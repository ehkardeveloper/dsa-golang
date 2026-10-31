package main

import "fmt"

func main() {

	// Declaring a array of 5 integers (all initialized to 0 by default)
	var numbers [5]int

	// output: [0, 0, 0, 0, 0]
	fmt.Println(numbers)

	// Declaring a array of 3 strigs and initializing with values using array literals
	fruits := [3]string{"apple", "banana", "guava"}

	// output: [apple banana guava]
	fmt.Println(fruits)

	// accessing and modifying elements
	fruits[1] = "Pinaple"

	// output: [apple Pinaple guava]
	fmt.Println(fruits)

	// getting the length of an array
	fmt.Println(len(fruits))

}
