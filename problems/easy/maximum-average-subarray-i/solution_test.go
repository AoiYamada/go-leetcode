package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumAverageSubarrayI(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{"Example 1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"Example 2", []int{5}, 1, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for maximum-average-subarray-i (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := findMaxAverage(tc.nums, tc.k)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
