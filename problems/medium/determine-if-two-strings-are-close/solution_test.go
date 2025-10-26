package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineIfTwoStringsAreClose(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected bool
	}{
		{"Example 1", "abc", "bca", true},
		{"Example 2", "a", "aa", false},
		{"Example 3", "cabbba", "abbccc", true},
		{"My test 1", "aabbcccc", "aabbccdd", false}, // letters count doesn't match
		{"My test 2", "abbbb", "aabbb", false},       // counts aren't 1 to 1
		{"My test 3", "abbccc", "aabbcc", false},     // length of counts doesn't match
		{"Test case 147", "uau", "ssx", false},       // not all letters in word1 occurs in word2
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for determine-if-two-strings-are-close (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := closeStrings(tc.word1, tc.word2)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
