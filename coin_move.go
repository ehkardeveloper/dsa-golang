package main

import "fmt"

func winningPlayer(x int, y int) string {
	// Find how many total rounds can be played based on whichever coin runs out first
	maxRounds := x
	if y/4 < maxRounds {
		maxRounds = y / 4
	}

	// If the total rounds played is odd, Alice wins. If even, Bob wins.
	if maxRounds%2 != 0 {
		return "Alice"
	} else {
		return "Bob"
	}
}

func main() {
	fmt.Println(winningPlayer(2, 7))  // Output: Alice
	fmt.Println(winningPlayer(4, 11)) // Output: Bob
}
