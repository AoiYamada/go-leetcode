package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type constructor func() StockSpanner

func produceResult(fn constructor, prices []int) []int {
	obj := fn()
	result := make([]int, len(prices))

	for pos, price := range prices {
		result[pos] = obj.Next(price)
	}

	return result
}

func TestOnlineStockSpan(t *testing.T) {
	testCases := []struct {
		name     string
		prices   []int
		expected []int
	}{
		{"Example 1", []int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for online-stock-span (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := produceResult(Constructor, tc.prices)
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
