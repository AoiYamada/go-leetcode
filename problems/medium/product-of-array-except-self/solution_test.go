package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		// {"Example 2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for product-of-array-except-self (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := productExceptSelf(tc.nums)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
