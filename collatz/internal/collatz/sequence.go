package collatz

import (
	"errors"
)

func sequenceLength(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	if val, exists := memo[n]; exists {
		return val
	}

	var next int
	if n%2 == 0 {
		next = n / 2
	} else {
		next = 3*n + 1
	}

	length := 1 + sequenceLength(next, memo)
	memo[n] = length
	return length
}

// LongestSequence returns the starting number under `limit` that produces the longest Collatz sequence
func LongestSequence(limit int) (int, int, error) {
	if limit < 2 {
		return 0, 0, errors.New("limit must be at least 2")
	}

	maxLen := 0
	number := 0
	memo := make(map[int]int)

	for i := 1; i < limit; i++ {
		length := sequenceLength(i, memo)
		if length > maxLen {
			maxLen = length
			number = i
		}
	}

	return number, maxLen, nil
}
