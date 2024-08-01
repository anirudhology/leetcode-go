package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func BinaryTreeMaximumPathSum(root *util.TreeNode) int {
	maxSum := ^(1 << 31)

	var helper func(node *util.TreeNode) int
	helper = func(node *util.TreeNode) int {
		// Base case
		if node == nil {
			return 0
		}
		// Calculate non-negative sum for left and right subtrees
		leftSum := max(helper(node.Left), 0)
		rightSum := max(helper(node.Right), 0)
		// Sum for the current node
		currentSum := node.Val + leftSum + rightSum
		// Update maxSum, if required
		maxSum = max(currentSum, maxSum)
		// Pick the subtree with greater value
		return node.Val + max(leftSum, rightSum)
	}
	helper(root)
	return maxSum
}
