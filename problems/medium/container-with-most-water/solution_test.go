package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	testCases := []struct {
		name     string
		heights  []int
		expected int
	}{
		{"Example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"Example 2", []int{1, 1}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for container-with-most-water (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := maxArea(tc.heights)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
