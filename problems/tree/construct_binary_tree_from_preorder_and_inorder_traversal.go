package tree

import "github.com/anirudhology/leetcode-go/problems/util"

func ConstructBinaryTreeFromPreorderAndInorderTraversal(preorder []int, inorder []int) *util.TreeNode {
	// Map to store indices of elements in the inorder array
	indexMappings := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		indexMappings[inorder[i]] = i
	}
	// Perform DFS on the binary tree
	return helper(0, len(preorder)-1, 0, len(inorder)-1, preorder, inorder, indexMappings)
}

func helper(preStart int, preEnd int, inStart int, inEnd int, preorder []int, inorder []int, indexMappings map[int]int) *util.TreeNode {
	// Base case
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	// Create root of the tree
	root := &util.TreeNode{Val: preorder[preStart]}
	// Get index of this value in the inorder
	index := indexMappings[root.Val]
	// All the elements to the left of the index in the inorder array
	// will make the left subtree and all the elements to the right of
	// the index in the inorder array will make the right subtree
	root.Left = helper(preStart+1, preEnd, inStart, index-1, preorder, inorder, indexMappings)
	root.Right = helper(preStart+index-inStart+1, preEnd, index+1, inEnd, preorder, inorder, indexMappings)
	return root
}
