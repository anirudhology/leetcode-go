package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func InvertBinaryTree(root *util.TreeNode) *util.TreeNode {
	// Base case
	if root == nil {
		return root
	}
	// Left and right subtrees
	left := root.Left
	right := root.Right
	// Invert
	root.Right = InvertBinaryTree(left)
	root.Left = InvertBinaryTree(right)
	return root
}
