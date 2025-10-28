package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var removingStarsFromAStringTestCases = []struct {
	name     string
	s        string
	expected string
}{
	{"Example 1", "leet**cod*e", "lecoe"},
	{"Example 2", "erase*****", ""},
}

func TestRemovingStarsFromAString(t *testing.T) {
	for _, tc := range removingStarsFromAStringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeStars(tc.s)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestRemovingStarsFromAString2(t *testing.T) {
	for _, tc := range removingStarsFromAStringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeStars2(tc.s)
			assert.Equal(t, tc.expected, result)
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
