package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type constructor func() RecentCounter

var numberOfRecentCallsTestCases = []struct {
	name     string
	times    []int
	expected []int
}{
	{"Example 1", []int{1, 100, 3001, 3002}, []int{1, 2, 3, 3}},
}

func produceResultWithPing(fn constructor, times []int) []int {
	obj := fn()
	result := make([]int, len(times))

	for pos, t := range times {
		result[pos] = obj.Ping(t)
	}

	return result
}

func produceResultWithPing2(fn constructor, times []int) []int {
	obj := fn()
	result := make([]int, len(times))

	for pos, t := range times {
		result[pos] = obj.Ping2(t)
	}

	return result
}

func TestNumberOfRecentCalls(t *testing.T) {
	for _, tc := range numberOfRecentCallsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := produceResultWithPing(Constructor, tc.times)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestNumberOfRecentCalls2(t *testing.T) {
	for _, tc := range numberOfRecentCallsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := produceResultWithPing2(Constructor, tc.times)
			assert.Equal(t, tc.expected, result)
		})
	}
}
