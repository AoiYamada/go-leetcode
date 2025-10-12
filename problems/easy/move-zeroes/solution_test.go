package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"Example 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Example 2", []int{0}, []int{0}},
		{"Test case 27", []int{1, 0, 1}, []int{1, 1, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run(tc.name, func(t *testing.T) {
				moveZeroes(tc.nums)
				assert.Equal(t, tc.expected, tc.nums)
			})
		})
	}
}
