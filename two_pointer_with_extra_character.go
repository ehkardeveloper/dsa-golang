// Package main implements a demonstration of the Two Pointers algorithm
// in Go to check for valid palindromes.
package main

import (
	"fmt"
	"unicode"
)

// isPalindrome checks whether a given string 's' is a valid palindrome,
// considering only alphanumeric characters and ignoring case differences.
// It uses the Two Pointers pattern moving from both ends inward.
func isPalindrome(s string) bool {
	// Initialize pointers at the start and end of the string
	left := 0
	right := len(s) - 1

	for left < right {
		// Skip non-alphanumeric characters from the left side
		if !isArithMetic(rune(s[left])) {
			left++
			continue
		}

		// Skip non-alphanumeric characters from the right side
		if !isArithMetic(rune(s[right])) {
			right--
			continue
		}

		// Compare characters case-insensitively; return false if a mismatch is found
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		// Move both pointers inward
		left++
		right--
	}

	return true
}

// isArithMetic helper function verifies if a given rune is an alphanumeric character
// (a-z, A-Z, or 0-9).
func isArithMetic(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

// main executes test cases to validate the palindrome function output.
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Expected: true
	fmt.Println(isPalindrome("race a car"))                     // Expected: false
}
