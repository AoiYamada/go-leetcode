package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{"Example 1", []int{1, 2, 2, 1, 1, 3}, true},
		{"Example 2", []int{1, 2}, false},
		{"Example 3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for unique-number-of-occurrences (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := uniqueOccurrences(tc.arr)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
