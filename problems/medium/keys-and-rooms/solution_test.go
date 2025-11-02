package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysAndRooms(t *testing.T) {
	testCases := []struct {
		name     string
		rooms    [][]int
		expected bool
	}{
		{"Example 1", [][]int{[]int{1}, []int{2}, []int{3}, []int{}}, true},
		{"Example 2", [][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for keys-and-rooms (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := canVisitAllRooms(tc.rooms)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
