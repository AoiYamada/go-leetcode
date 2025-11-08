package main

// Medium: UniquePaths
// Solution for unique-paths (medium)
var cache = make(map[int]map[int]int)

func uniquePaths(m int, n int) int {
	if cache[m] != nil && cache[m][n] > 0 {
		return cache[m][n]
	}

	if m <= 1 {
		return 1
	}

	if n <= 1 {
		return 1
	}

	if cache[m] == nil {
		cache[m] = make(map[int]int)
	}

	cache[m][n] = uniquePaths(m-1, n) + uniquePaths(m, n-1)

	return cache[m][n]
}
