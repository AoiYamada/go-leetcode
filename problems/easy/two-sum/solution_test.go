package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumBruteForce(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSumBruteForce(tc.nums, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestTwoSumHashMap(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSumHashMap(tc.nums, tc.target)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// Benchmark tests
func BenchmarkTwoSumBruteForce(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	for i := 0; i < b.N; i++ {
		twoSumBruteForce(nums, target)
	}
}

func BenchmarkTwoSumHashMap(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	for i := 0; i < b.N; i++ {
		twoSumHashMap(nums, target)
	}
}
