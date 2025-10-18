package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNThTribonacciNumber(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"Example 1", 4, 4},
		{"Example 2", 25, 1389537},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for n-th-tribonacci-number (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := tribonacci(tc.n)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
