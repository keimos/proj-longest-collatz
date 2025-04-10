package main

import (
	"collatz/collatz"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <limit>")
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid limit: %v. Please provide an integer >= 2.", os.Args[1])

	}

	startNumber, length, err := collatz.LongestSequence(limit)
	if err != nil {
		log.Fatalf("Error finding longest Collatz sequence: %v", err)
	}
	fmt.Printf("The starting number under %d that produces the longest Collatz sequence is %d (length %d)\n", limit, startNumber, length)
}
