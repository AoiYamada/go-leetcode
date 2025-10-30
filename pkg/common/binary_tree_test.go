package common

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	tests := []struct {
		name     string
		arr      []*int
		expected *TreeNode
	}{
		{
			name:     "empty array",
			arr:      []*int{},
			expected: nil,
		},
		{
			name:     "nil root",
			arr:      []*int{nil},
			expected: nil,
		},
		{
			name: "single node",
			arr:  []*int{IntPtr(1)},
			expected: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		{
			name: "single node with zero value",
			arr:  []*int{IntPtr(0)},
			expected: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
		},
		{
			name: "tree with left and right children",
			arr:  []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "tree with nil nodes in the middle",
			arr:  []*int{IntPtr(1), nil, IntPtr(2)},
			expected: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "tree with zero values",
			arr:  []*int{IntPtr(1), IntPtr(0), IntPtr(2)},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "complete tree",
			arr:  []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "sparse tree with trailing nils",
			arr:  []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3)},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CreateTree(tt.arr)
			if !treesEqual(result, tt.expected) {
				t.Errorf("CreateTree() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to compare two trees
func treesEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treesEqual(a.Left, b.Left) && treesEqual(a.Right, b.Right)
}

func TestTreeToSlice(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []*int
	}{
		{
			name:     "nil tree",
			root:     nil,
			expected: []*int{},
		},
		{
			name:     "single node",
			root:     &TreeNode{Val: 1},
			expected: []*int{IntPtr(1)},
		},
		{
			name: "tree with left and right",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
		},
		{
			name: "tree with nil nodes",
			root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 2},
			},
			expected: []*int{IntPtr(1), nil, IntPtr(2)},
		},
		{
			name: "complete tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TreeToSlice(tt.root)
			if !arraysEqual(result, tt.expected) {
				t.Errorf("TreeToSlice() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to compare two int pointer arrays
func arraysEqual(a, b []*int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] == nil && b[i] == nil {
			continue
		}
		if a[i] == nil || b[i] == nil {
			return false
		}
		if *a[i] != *b[i] {
			return false
		}
	}
	return true
}

func TestPrintTree(t *testing.T) {
	fmt.Println("\n=== Test 1: Single node ===")
	tree1 := &TreeNode{Val: 1}
	PrintTree(tree1)

	fmt.Println("\n=== Test 2: Simple tree ===")
	tree2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	PrintTree(tree2)

	fmt.Println("\n=== Test 3: Tree with nil nodes ===")
	tree3 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: nil,
		},
	}
	PrintTree(tree3)

	fmt.Println("\n=== Test 4: Complete tree ===")
	tree4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	PrintTree(tree4)

	fmt.Println("\n=== Test 5: Empty tree ===")
	PrintTree(nil)

	fmt.Println("\n=== Test 6: Tree with zero values ===")
	tree6 := &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 5},
	}
	PrintTree(tree6)
}
