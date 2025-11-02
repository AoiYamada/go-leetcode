package main

import (
	"leetcode-solutions/pkg/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

var intPtr = common.IntPtr

func TestRightSideView(t *testing.T) {
	testCases := []struct {
		name       string
		nodeValues []*int
		expected   []int
	}{
		{"Example 1", []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(5), nil, intPtr(4)}, []int{1, 3, 4}},
		{"Example 2", []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), nil, nil, nil, intPtr(5)}, []int{1, 3, 4, 5}},
		{"Example 3", []*int{intPtr(1), nil, intPtr(3)}, []int{1, 3}},
		{"Example 4", []*int{}, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for binary-tree-right-side-view (medium)
			t.Run(tc.name, func(t *testing.T) {
				result := rightSideView(common.CreateTree(tc.nodeValues))
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
