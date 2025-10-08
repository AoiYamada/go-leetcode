package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWordsInAString(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{"Example 1", "the sky is blue", "blue is sky the"},
		{"Example 2", "  hello world  ", "world hello"},
		{"Example 3", "a good   example", "example good a"},
		{"Test Case 1", "EPY2giL", "EPY2giL"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseWords(tc.s)
			assert.Equal(t, tc.expected, result)
		})
	}
}
