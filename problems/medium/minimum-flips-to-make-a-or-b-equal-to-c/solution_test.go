package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumFlipsToMakeAOrBEqualToC(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		c        int
		expected int
	}{
		{"Example 1", 2, 6, 5, 3},
		{"Example 2", 4, 2, 7, 1},
		{"Example 3", 1, 2, 3, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for minimum-flips-to-make-a-or-b-equal-to-c (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := minFlips(tc.a, tc.b, tc.c)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
