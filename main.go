package main

import (
	"collatz/collatz"
	"fmt"
	"log"
)

func main() {
	limit := 1000000

	startNumber, length, err := collatz.LongestSequence(limit)
	if err != nil {
		log.Fatalf("Error finding longest Collatz sequence: %v", err)
	}
	fmt.Printf("The starting number under %d that produces the longest Collatz sequence is %d (length %d)\n", limit, startNumber, length)
}
