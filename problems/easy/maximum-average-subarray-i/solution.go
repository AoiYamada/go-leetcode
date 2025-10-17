package main

import "runtime/debug"

// Easy: MaximumAverageSubarrayI
// Solution for maximum-average-subarray-i (easy)
func findMaxAverage(nums []int, k int) float64 {
	var totalInWindow int
	for i := 0; i < k; i++ {
		totalInWindow += nums[i]
	}

	maxTotal := totalInWindow

	for i := 0; i < len(nums)-k; i++ {
		totalInWindow = totalInWindow + nums[k+i] - nums[i]
		if totalInWindow > maxTotal {
			maxTotal = totalInWindow
		}
	}

	return float64(maxTotal) / float64(k)
}

func init() {
	debug.SetMemoryLimit(0)
}
