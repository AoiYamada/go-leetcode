package main

import "leetcode-solutions/pkg/common"

type TreeNode = common.TreeNode

type Info struct {
	depth int
	node  *TreeNode
}

// Easy: MaximumDepthOfBinaryTree
// Solution for maximum-depth-of-binary-tree (easy)
func maxDepth(root *TreeNode) int {
	queue := []Info{
		{
			depth: 1,
			node:  root,
		},
	}
	maxDepth := 0

	// BFS
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node == nil {
			continue
		}

		if current.depth > maxDepth {
			maxDepth = current.depth
		}

		if current.node.Left != nil {
			queue = append(queue, Info{
				current.depth + 1,
				current.node.Left,
			})
		}

		if current.node.Right != nil {
			queue = append(queue, Info{
				current.depth + 1,
				current.node.Right,
			})
		}
	}

	return maxDepth
}

// recursive, copy from most memory efficient sol
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth2(root.Left), maxDepth2(root.Right))
}
