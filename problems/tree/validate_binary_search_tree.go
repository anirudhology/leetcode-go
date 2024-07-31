package tree

import (
	"math"

	"github.com/anirudhology/leetcode-go/problems/util"
)

func ValidateBinarySearchTree(root *util.TreeNode) bool {
	return dfsTraversal(root, math.MinInt, math.MaxInt)
}

func dfsTraversal(root *util.TreeNode, minValue int, maxValue int) bool {
	// Base case
	if root == nil {
		return true
	}
	if root.Val >= maxValue || root.Val <= minValue {
		return false
	}
	return dfsTraversal(root.Left, minValue, root.Val) && dfsTraversal(root.Right, root.Val, maxValue)
}
