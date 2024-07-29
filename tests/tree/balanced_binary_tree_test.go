package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestBalancedBinaryTree(t *testing.T) {

	// Test case 1: Empty tree
	if !tree.BalancedBinaryTree(nil) {
		t.Error("Expected true, got false")
	}

	// Test case 2: Single node tree
	root1 := &util.TreeNode{Val: 1}
	if !tree.BalancedBinaryTree(root1) {
		t.Error("Expected true, got false")
	}

	// Test case 3: Balanced tree
	root2 := &util.TreeNode{Val: 1}
	root2.Left = &util.TreeNode{Val: 2}
	root2.Right = &util.TreeNode{Val: 3}
	if !tree.BalancedBinaryTree(root2) {
		t.Error("Expected true, got false")
	}

	// Test case 4: Unbalanced tree
	root3 := &util.TreeNode{Val: 1}
	root3.Left = &util.TreeNode{Val: 2}
	root3.Left.Left = &util.TreeNode{Val: 3}
	if tree.BalancedBinaryTree(root3) {
		t.Error("Expected false, got true")
	}

	// Test case 5: Larger balanced tree
	root4 := &util.TreeNode{Val: 1}
	root4.Left = &util.TreeNode{Val: 2}
	root4.Right = &util.TreeNode{Val: 3}
	root4.Left.Left = &util.TreeNode{Val: 4}
	root4.Left.Right = &util.TreeNode{Val: 5}
	root4.Right.Left = &util.TreeNode{Val: 6}
	root4.Right.Right = &util.TreeNode{Val: 7}
	if !tree.BalancedBinaryTree(root4) {
		t.Error("Expected true, got false")
	}
}
