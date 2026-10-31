package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if an integer reads the same forwards and backwards
// using the Two-Pointer algorithmic pattern.
func isPalindrome(x int) bool {
	// Step 1: Handle the negative number edge case.
	// Negative numbers (e.g., -121) can never be palindromes because
	// of the minus sign at the front ("121-" != "-121").
	if x < 0 {
		return false
	}

	// Step 2: Convert the integer into a string.
	// Go requires strings to index individual digits as characters.
	s := strconv.Itoa(x)

	// Step 3: Convert the string into a slice of runes.
	// Runes safely handle Unicode characters and allow index-based access.
	runes := []rune(s)

	// Step 4: Initialize the Two Pointers.
	// 'left' starts at the very beginning of the sequence (index 0).
	// 'right' starts at the very end of the sequence (last index).
	left := 0
	right := len(runes) - 1

	// Step 5: Traverse inward from both ends towards the center.
	// This runs in O(N) time complexity instead of O(N^2) nested loops.
	for left < right {
		// Compare characters at both pointers. If they don't match, it's not a palindrome.
		if runes[left] != runes[right] {
			return false
		}

		// Move the left pointer forward and the right pointer backward.
		left++
		right--
	}

	// If the loop finishes without mismatches, all outer characters matched!
	return true
}

func main() {
	// Test cases to verify our logic
	num1 := 121
	num2 := 211

	fmt.Println(num1, "is palindrome", isPalindrome(num1)) // Expected: true
	fmt.Println(num2, "is palindrome", isPalindrome(num2)) // Expected: false
}
