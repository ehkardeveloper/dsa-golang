// Package main implements a demonstration of duplicate detection
// using Go hash maps for optimal O(n) time complexity.
package main

import (
	"fmt"
)

// checkDuplicate takes a slice of integers and checks if any value appears
// more than once. It utilizes a map[int]bool to track seen elements in O(1) lookup time.
//
// Parameters:
//   - nums: A slice of integers to evaluate.
//
// Returns:
//   - bool: true if a duplicate is found, false if all elements are unique.
func checkDuplicate(nums []int) bool {
	// Initialize a hash map to act as a memory checklist (a Go set equivalent)
	seen := make(map[int]bool)

	// Iterate through each number in the slice sequentially
	for _, num := range nums {
		// Check if the current number has already been recorded in our map
		if seen[num] {
			// If yes, we found a duplicate; terminate early and return true
			return true
		}
		// If no, record this number in our map as seen
		seen[num] = true
	}

	// If the loop completes without early exit, all elements are unique
	return false
}

func main() {
	// Define test datasets: list1 contains all unique items, list2 contains a duplicate (7)
	list1 := []int{5, 2, 7, 8, 3}
	list2 := []int{5, 2, 7, 8, 3, 7}

	// Execute the function and print the results for verification
	fmt.Println("is list1 duplicate?", checkDuplicate(list1)) // Expected: false
	fmt.Println("is list2 duplicate?", checkDuplicate(list2)) // Expected: true
}
