// Package main implements a mathematical approach to find the missing number
// in an integer sequence using O(n) time complexity and O(1) auxiliary space.
package main

import "fmt"

// findMissingOneNumber calculates the single missing integer from a sequence of 0 to n.
// It leverages the arithmetic progression sum formula: n * (n + 1) / 2.
//
// Parameters:
//   - nums: A slice of integers containing unique numbers from a range 0 to n (with one missing).
//
// Returns:
//   - int: The missing integer value.
func findMissingOneNumber(nums []int) int {
	// Determine n, which represents the length of the slice (the maximum expected value in the sequence)
	n := len(nums)

	// Calculate the mathematical sum of all numbers from 0 to n
	expectedSum := n * (n + 1) / 2

	// Initialize a variable to accumulate the sum of elements currently present in the slice
	actualSum := 0

	// Iterate through every number in the slice and add them together
	for _, num := range nums {
		actualSum += num
	}

	// The missing number is the difference between what the sum should be and what it actually is
	missingNumber := expectedSum - actualSum

	return missingNumber
}

func main() {
	// Define a test array where numbers are from 0 to 3, but '2' is missing
	array := []int{0, 1, 3}

	// Execute the function and print the result for verification
	fmt.Println("Missing Number is ", findMissingOneNumber(array)) // Expected output: 2
}
