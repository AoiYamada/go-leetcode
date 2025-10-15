package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{"Example 1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"Example 2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"Example 3", []int{30, 60, 90}, []int{1, 1, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for daily-temperatures (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := dailyTemperatures(tc.temperatures)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}

func TestDailyTemperatures2(t *testing.T) {
	testCases := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{"Example 1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"Example 2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"Example 3", []int{30, 60, 90}, []int{1, 1, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for daily-temperatures (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := dailyTemperatures2(tc.temperatures)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
