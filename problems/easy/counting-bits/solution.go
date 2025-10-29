package main

import "math/bits"

// Easy: CountingBits
// Solution for counting-bits (easy)
func countBits(n int) []int {
	ans := make([]int, n+1)

	for ; n >= 0; n-- {
		ans[n] = bitCount(n)
	}

	return ans
}

func bitCount(n int) int {
	count := 0

	for n > 0 {
		count++
		// for a number n in the form of ...10...(repeating 0s at the end)
		// n - 1 is ...01...(repeating 1s at the end with the same length)
		// n & (n - 1) is ...00...(repeating 0s)
		// so it removes the least significant 1 in the binary number
		// so this process can count all 1s in the number n
		n = n & (n - 1)
	}

	return count
}

// Copy from sol
func countBits2(n int) []int {
	result := make([]int, n+1)
	for i := range result {
		// LOL, use lib
		result[i] = bits.OnesCount(uint(i))
	}
	return result
}
