package main

import "leetcode-solutions/pkg/common"

type TreeNode = common.TreeNode

type Info struct {
	depth int
	node  *TreeNode
}

// Medium: BinaryTreeRightSideView
// Solution for binary-tree-right-side-view (medium)
func rightSideView(root *TreeNode) []int {
	result := []int{}
	queue := []Info{
		{
			depth: 1,
			node:  root,
		},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node == nil {
			continue
		}

		if len(result) < current.depth {
			result = append(result, current.node.Val)
		} else {
			result[current.depth-1] = current.node.Val
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

	return result
}
