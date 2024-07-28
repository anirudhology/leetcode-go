package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func DiameterOfBinaryTree(root *util.TreeNode) int {
	// Diameter
	diameter := 0
	// Calculate maximum height and update diameter
	maxDepth(root, &diameter)
	return diameter
}

func maxDepth(root *util.TreeNode, diameter *int) int {
	// Base case
	if root == nil {
		return 0
	}
	// Calculate left and right height
	leftHeight, rightHeight := maxDepth(root.Left, diameter), maxDepth(root.Right, diameter)
	// Update diameter, if required
	*diameter = max(*diameter, leftHeight+rightHeight)
	return 1 + max(leftHeight, rightHeight)
}
