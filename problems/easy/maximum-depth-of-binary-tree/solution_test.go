package main

import (
	"leetcode-solutions/pkg/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

var intPtr = common.IntPtr

var testCases = []struct {
	name      string
	treeNodes []*int
	expected  int
}{
	{"Example 1", []*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)}, 3},
	{"Example 2", []*int{intPtr(1), nil, intPtr(2)}, 2},
	{"Test Case 38", []*int{}, 0},
}

func TestMaxDepth(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for maximum-depth-of-binary-tree (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := maxDepth(common.CreateTree(tc.treeNodes))
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}

func TestMaxDepth2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test implementation for maximum-depth-of-binary-tree (easy)
			t.Run(tc.name, func(t *testing.T) {
				result := maxDepth2(common.CreateTree(tc.treeNodes))
				assert.Equal(t, tc.expected, result)
			})
		})
	}
}
