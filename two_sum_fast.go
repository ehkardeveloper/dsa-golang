// Package main implements the Two Sum algorithm using a hash map
// to achieve optimal O(n) time complexity.
package main

import (
	"fmt"
)

// twoSum finds the indices of two numbers in a slice that add up to a specific target.
// It uses a hash map to store numbers and their indices as it iterates through the slice.
//
// Parameters:
//   - nums: A slice of integers to search through.
//   - target: The target sum we are trying to reach.
//
// Returns:
//   - []int: A slice containing the two indices whose values add up to the target.
func twoSum(nums []int, target int) []int {
	// Initialize a hash map where the key is the number and the value is its index
	numMap := make(map[int]int)

	// Iterate through the slice, capturing both the current index and the number itself
	for currentIndex, currentNum := range nums {
		// Calculate the complement needed to reach the target sum
		complement := target - currentNum

		// Check if the complement already exists in our map
		if missingIndex, found := numMap[complement]; found {
			// If found, return the current index and the index of the complement
			return []int{currentIndex, missingIndex}
		}

		// Otherwise, record the current number and its index into the map
		numMap[currentNum] = currentIndex
	}

	// Return an empty slice if no such pair exists
	return []int{}
}

func main() {
	// Define a test dataset and a target sum (e.g., 7 + 2 = 9)
	list1 := []int{11, 36, 5, 7, 2}
	target := 9

	// Execute the function and print the resulting indices
	fmt.Println(twoSum(list1, target)) // Expected output: [3, 4] (indices of 7 and 2)
}
