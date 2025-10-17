package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheHighestAltitude(t *testing.T) {
	testCases := []struct {
		name     string
		gain     []int
		expected int
	}{
		{"Example 1", []int{-5, 1, 5, 0, -7}, 1},
		{"Example 2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for find-the-highest-altitude (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := largestAltitude(tc.gain)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
