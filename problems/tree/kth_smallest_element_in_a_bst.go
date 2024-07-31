package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func KthSmallestElementInABST(root *util.TreeNode, k int) int {
	count := k
	kthSmallestElement := 0

	var findKthSmallestElement func(root *util.TreeNode)

	findKthSmallestElement = func(root *util.TreeNode) {
		// Base case
		if root == nil {
			return
		}
		// Check in the left subtree
		if root.Left != nil {
			findKthSmallestElement(root.Left)
		}
		count--
		if count == 0 {
			kthSmallestElement = root.Val
		}

		// Check in the right subtree
		if root.Right != nil {
			findKthSmallestElement(root.Right)
		}
	}
	findKthSmallestElement(root)
	return kthSmallestElement
}
