package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{2, 2, 1}, 1},
		{"Example 2", []int{4, 1, 2, 1, 2}, 4},
		{"Example 3", []int{1}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for single-number (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := singleNumber(tc.nums)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
