package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	n        int
	expected []int
}{
	{"Example 1", 2, []int{0, 1, 1}},
	{"Example 1", 2, []int{0, 1, 1}},
}

func TestCountingBits(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for counting-bits (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := countBits(tc.n)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}

func TestCountingBits2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for counting-bits (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := countBits2(tc.n)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
