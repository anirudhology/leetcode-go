package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

// Helper function to compare two trees
func equal(t1, t2 *util.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return equal(t1.Left, t2.Left) && equal(t1.Right, t2.Right)
}

func TestInvertBinaryTree(t *testing.T) {
	// Test case 1: Empty tree
	if tree.InvertBinaryTree(nil) != nil {
		t.Error("Expected nil, got non-nil")
	}

	// Test case 2: Full binary tree
	root2 := &util.TreeNode{Val: 1}
	root2.Left = &util.TreeNode{Val: 2}
	root2.Right = &util.TreeNode{Val: 3}
	root2.Left.Left = &util.TreeNode{Val: 4}
	root2.Left.Right = &util.TreeNode{Val: 5}
	root2.Right.Left = &util.TreeNode{Val: 6}
	root2.Right.Right = &util.TreeNode{Val: 7}

	invertedRoot2 := tree.InvertBinaryTree(root2)
	expectedRoot2 := &util.TreeNode{Val: 1}
	expectedRoot2.Left = &util.TreeNode{Val: 3}
	expectedRoot2.Right = &util.TreeNode{Val: 2}
	expectedRoot2.Left.Left = &util.TreeNode{Val: 7}
	expectedRoot2.Left.Right = &util.TreeNode{Val: 6}
	expectedRoot2.Right.Left = &util.TreeNode{Val: 5}
	expectedRoot2.Right.Right = &util.TreeNode{Val: 4}

	if !equal(invertedRoot2, expectedRoot2) {
		t.Error("Full binary tree inversion failed")
	}
}
