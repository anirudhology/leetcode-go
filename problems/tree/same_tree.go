package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func SameTree(p *util.TreeNode, q *util.TreeNode) bool {
	// Base cases
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right)
}
