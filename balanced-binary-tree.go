package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// global flag for balanced height
	balanced := true

	// deep-first-search, bottom up approach, function returns tree height
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		// base case, return zero height
		if node == nil {
			return 0
		}

		// calculate height of left and right subtrees
		lh := dfs(node.Left)
		rh := dfs(node.Right)

		// if the heights differ by more than 1
		// set global flag as false
		if abs(lh, rh) > 1 {
			balanced = false
		}

		// return the height: current node + longest subtree
		return 1 + max(lh, rh)
	}

	// start the traversal
	dfs(root)

	return balanced
}

// helper functions

// maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// absolute difference for two integers
func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// Time: O(n)
// Space: O(n)
