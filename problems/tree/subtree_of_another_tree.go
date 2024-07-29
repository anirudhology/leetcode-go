package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func SubtreeOfAnotherTree(root *util.TreeNode, subRoot *util.TreeNode) bool {
	if root == nil {
		return false
	}
	if isSame(root, subRoot) {
		return true
	}
	return SubtreeOfAnotherTree(root.Left, subRoot) || SubtreeOfAnotherTree(root.Right, subRoot)
}

func isSame(s *util.TreeNode, t *util.TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil || s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
