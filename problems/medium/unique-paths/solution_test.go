package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{"Example 1", 3, 7, 28},
		{"Example 2", 3, 2, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for unique-paths (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := uniquePaths(tc.m, tc.n)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
