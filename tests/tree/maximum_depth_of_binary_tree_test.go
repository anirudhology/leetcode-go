package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	// Test case 1: Empty tree
	if tree.MaximumDepthOfBinaryTree(nil) != 0 {
		t.Error("Expected 0, got non-zero")
	}

	// Test case 2: Single node tree
	root1 := &util.TreeNode{Val: 1}
	if tree.MaximumDepthOfBinaryTree(root1) != 1 {
		t.Error("Single node tree max depth failed")
	}

	// Test case 3: Two-level tree
	root2 := &util.TreeNode{Val: 1}
	root2.Left = &util.TreeNode{Val: 2}
	if tree.MaximumDepthOfBinaryTree(root2) != 2 {
		t.Error("Two-level tree max depth failed")
	}

	// Test case 4: Full binary tree
	root3 := &util.TreeNode{Val: 1}
	root3.Left = &util.TreeNode{Val: 2}
	root3.Right = &util.TreeNode{Val: 3}
	root3.Left.Left = &util.TreeNode{Val: 4}
	root3.Left.Right = &util.TreeNode{Val: 5}
	root3.Right.Left = &util.TreeNode{Val: 6}
	root3.Right.Right = &util.TreeNode{Val: 7}
	if tree.MaximumDepthOfBinaryTree(root3) != 3 {
		t.Error("Full binary tree max depth failed")
	}
}
