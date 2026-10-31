package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	// converting to lowercase to make it uniform
	s = strings.ToLower(s)

	// converting stings to slice of runes/bytes to use pointers easily
	// (In go, strings are read-only byte slices)
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	// runnig for loop util left variable less than right variable
	for left < right {
		if runes[left] != runes[right] {
			return false
		}

		// move the pointers closer to the center
		left++
		right--
	}

	return true

}

func main() {
	word1 := "racecar"
	word2 := "hello"

	// output true
	fmt.Println(word1, "is palindrome?", isPalindrome(word1))
	// outpur false
	fmt.Println(word2, "is palindrome?", isPalindrome(word2))
}
