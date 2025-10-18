package main

var cache = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}

// Easy: NThTribonacciNumber
// Solution for n-th-tribonacci-number (easy)
func tribonacci(n int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)

	return cache[n]
}
