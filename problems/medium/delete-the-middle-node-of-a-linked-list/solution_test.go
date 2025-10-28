package main

import (
	"leetcode-solutions/pkg/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

var deleteMiddleTestCases = []struct {
	name     string
	head     []int
	expected []int
}{
	{"Example 1", []int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
	{"Example 2", []int{1, 2, 3, 4}, []int{1, 2, 4}},
	{"Example 3", []int{2, 1}, []int{2}},
	{"Test Case 66", []int{1}, []int{}},
}

func TestDeleteMiddle(t *testing.T) {
	for _, tc := range deleteMiddleTestCases {
		t.Run(tc.name, func(t *testing.T) {
			head := common.CreateLinkedList(tc.head)
			head = deleteMiddle(head)
			result := common.LinkedListToSlice(head)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestDeleteMiddle2(t *testing.T) {
	for _, tc := range deleteMiddleTestCases {
		t.Run(tc.name, func(t *testing.T) {
			head := common.CreateLinkedList(tc.head)
			head = deleteMiddle2(head)
			result := common.LinkedListToSlice(head)
			assert.Equal(t, tc.expected, result)
		})
	}
}
