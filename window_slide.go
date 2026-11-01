// Package main implements core Data Structures and Algorithms (DSA)
// optimized for performance and minimal memory usage in Go.
package main

import "fmt"

// windowSlide calculates the maximum sum of any contiguous sub-array of a fixed size 'k'.
// Time Complexity: O(N) - Linear scan using the sliding window technique.
// Space Complexity: O(1) - Constant space overhead.
func windowSlide(arr []int, k int) int {
	// Edge case handling (if window size is larger than the array length)
	if len(arr) < k {
		return 0
	}

	currentSum := 0

	// Step 1: Calculate the sum of the very first window of size 'k'
	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}

	maxSum := currentSum

	// Step 2: Slide the window across the rest of the array.
	// Instead of recalculating from scratch, we subtract the element falling
	// off the left edge (arr[i-k]) and add the new element entering on the right edge (arr[i]).
	for i := k; i < len(arr); i++ {
		currentSum = currentSum + arr[i] - arr[i-k]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	// Sample dataset for demonstration
	number := []int{5, 1, 8, 1, 4, 4, 5, 4, 9}
	window := 3

	// Execute and print the maximum contiguous window sum
	fmt.Println("Maximum window sum set is:", windowSlide(number, window))
}
