package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func MaximumDepthOfBinaryTree(root *util.TreeNode) int {
	// Special case
	if root == nil {
		return 0
	}
	return 1 + max(MaximumDepthOfBinaryTree(root.Left), MaximumDepthOfBinaryTree(root.Right))
}
