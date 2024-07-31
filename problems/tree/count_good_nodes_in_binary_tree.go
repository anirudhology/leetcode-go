package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func CountGoodNodesInBinaryTree(root *util.TreeNode) int {
	var count int
	var dfs func(node *util.TreeNode, max int)

	dfs = func(node *util.TreeNode, max int) {
		if node == nil {
			return
		}
		if node.Val >= max {
			max = node.Val
			count++
		}
		dfs(node.Left, max)
		dfs(node.Right, max)
	}

	dfs(root, root.Val)
	return count
}
