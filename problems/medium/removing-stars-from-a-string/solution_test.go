package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovingStarsFromAString(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{"Example 1", "leet**cod*e", "lecoe"},
		{"Example 2", "erase*****", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for removing-stars-from-a-string (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := removeStars(tc.s)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}

func TestRemovingStarsFromAString2(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{"Example 1", "leet**cod*e", "lecoe"},
		{"Example 2", "erase*****", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for removing-stars-from-a-string (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := removeStars2(tc.s)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}

func BenchmarkRemoveStarsBySliceAppendAndPop(b *testing.B) {
	s := "leet**cod*e_leet**cod*e_leet**cod*e_leet**cod*e_leet**cod*e"

	for i := 0; i < b.N; i++ {
		removeStars(s)
	}
}

func BenchmarkRemoveStarsByInPlaceReplacement(b *testing.B) {
	s := "leet**cod*e_leet**cod*e_leet**cod*e_leet**cod*e_leet**cod*e"

	for i := 0; i < b.N; i++ {
		removeStars2(s)
	}
}
