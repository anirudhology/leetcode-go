package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func LowestCommonAncestorOfABinarySearchTree(root, p, q *util.TreeNode) *util.TreeNode {
	// Special case
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		// If value of root is greater than values at both p and q,
		// it means LCA can only be in the left subtree
		return LowestCommonAncestorOfABinarySearchTree(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		// If value of root is smaller than values at both p and q,
		// it means LCA can only be in the right subtree
		return LowestCommonAncestorOfABinarySearchTree(root.Right, p, q)
	} else {
		//We will reach here if value is root is in between values at
		// p and q. If this happens, then only root can be the LCA
		return root
	}
}
