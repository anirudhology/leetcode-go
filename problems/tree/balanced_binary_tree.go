package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func BalancedBinaryTree(root *util.TreeNode) bool {
	// Find heights of subtrees in bottom up manner
	return dfs(root) != -1
}

func dfs(root *util.TreeNode) int {
	// Base case
	if root == nil {
		return 0
	}
	// Height of the right subtree
	left := dfs(root.Left)
	if left == -1 {
		return -1
	}
	// Height of the right subtree
	right := dfs(root.Right)
	if right == -1 {
		return -1
	}
	// Check for balancing property
	if abs(left-right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
