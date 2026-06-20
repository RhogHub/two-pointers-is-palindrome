package main

import (
	"fmt"
	"time"
)

const Iterations = 100_000_000

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	//word := "arara"
	word := "rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr"

	start := time.Now()

	var result bool

	for i := 0; i < Iterations; i++ {
		result = isPalindrome(word)
	}

	elapsed := time.Since(start)

	fmt.Printf("Palavra: %s\n", word)
	fmt.Printf("É palíndromo? %v\n", result)
	fmt.Printf("Iterações: %d\n", Iterations)
	fmt.Printf("Tempo total: %v\n", elapsed)
	fmt.Printf("Tempo médio: %.2f ns/op\n",
		float64(elapsed.Nanoseconds())/float64(Iterations))
}
