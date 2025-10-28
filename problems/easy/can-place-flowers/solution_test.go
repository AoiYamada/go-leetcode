package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var canPlaceFlowersTestCases = []struct {
	name      string
	flowerbed []int
	n         int
	expected  bool
}{
	{"Example 1", []int{1, 0, 0, 0, 1}, 1, true},
	{"Example 2", []int{1, 0, 0, 0, 1}, 2, false},
	{"Test Case 13", []int{0, 0, 1, 0, 1}, 1, true},
	{"Test Case 16", []int{1, 0, 0, 0, 1, 0, 0}, 2, true},
	{"Test Case 123", []int{1, 0}, 1, false},
}

func TestCanPlaceFlowers(t *testing.T) {
	for _, tc := range canPlaceFlowersTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canPlaceFlowers(tc.flowerbed, tc.n)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestCanPlaceFlowers2(t *testing.T) {
	for _, tc := range canPlaceFlowersTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canPlaceFlowers2(tc.flowerbed, tc.n)
			assert.Equal(t, tc.expected, result)
		})
	}
}
