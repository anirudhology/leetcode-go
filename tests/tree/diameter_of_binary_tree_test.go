package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	// Helper function to create a new TreeNode
	newNode := func(val int) *util.TreeNode {
		return &util.TreeNode{Val: val}
	}

	// Test case 1: Empty tree
	if tree.DiameterOfBinaryTree(nil) != 0 {
		t.Errorf("Expected 0, got %d", tree.DiameterOfBinaryTree(nil))
	}

	// Test case 2: Single node tree
	root1 := newNode(1)
	if tree.DiameterOfBinaryTree(root1) != 0 {
		t.Errorf("Expected 0, got %d", tree.DiameterOfBinaryTree(root1))
	}

	// Test case 3: Two-level tree
	root2 := newNode(1)
	root2.Left = newNode(2)
	if tree.DiameterOfBinaryTree(root2) != 1 {
		t.Errorf("Expected 1, got %d", tree.DiameterOfBinaryTree(root2))
	}

	// Test case 4: Full binary tree
	root3 := newNode(1)
	root3.Left = newNode(2)
	root3.Right = newNode(3)
	root3.Left.Left = newNode(4)
	root3.Left.Right = newNode(5)
	if tree.DiameterOfBinaryTree(root3) != 3 {
		t.Errorf("Expected 3, got %d", tree.DiameterOfBinaryTree(root3))
	}

	// Test case 5: Larger tree
	root4 := newNode(1)
	root4.Left = newNode(2)
	root4.Right = newNode(3)
	root4.Left.Left = newNode(4)
	root4.Left.Right = newNode(5)
	root4.Right.Right = newNode(6)
	root4.Left.Left.Left = newNode(7)
	if tree.DiameterOfBinaryTree(root4) != 5 {
		t.Errorf("Expected 4, got %d", tree.DiameterOfBinaryTree(root4))
	}
}
