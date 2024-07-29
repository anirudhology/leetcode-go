package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestIsSubtree(t *testing.T) {
	// Test case 1: Both trees are null
	if tree.SubtreeOfAnotherTree(nil, nil) {
		t.Error("Expected false, got true")
	}

	// Test case 2: One tree is null
	root1 := &util.TreeNode{Val: 1}
	if tree.SubtreeOfAnotherTree(root1, nil) {
		t.Error("Expected false, got true")
	}
	if tree.SubtreeOfAnotherTree(nil, root1) {
		t.Error("Expected false, got true")
	}

	// Test case 3: Both trees have one node with the same value
	root2 := &util.TreeNode{Val: 1}
	subRoot2 := &util.TreeNode{Val: 1}
	if !tree.SubtreeOfAnotherTree(root2, subRoot2) {
		t.Error("Expected true, got false")
	}

	// Test case 4: Subtree with different values
	root3 := &util.TreeNode{Val: 1}
	subRoot3 := &util.TreeNode{Val: 2}
	if tree.SubtreeOfAnotherTree(root3, subRoot3) {
		t.Error("Expected false, got true")
	}

	// Test case 5: Subtree is an actual subtree
	root4 := &util.TreeNode{Val: 3}
	root4.Left = &util.TreeNode{Val: 4}
	root4.Right = &util.TreeNode{Val: 5}
	root4.Left.Left = &util.TreeNode{Val: 1}
	root4.Left.Right = &util.TreeNode{Val: 2}

	subRoot4 := &util.TreeNode{Val: 4}
	subRoot4.Left = &util.TreeNode{Val: 1}
	subRoot4.Right = &util.TreeNode{Val: 2}
	if !tree.SubtreeOfAnotherTree(root4, subRoot4) {
		t.Error("Expected true, got false")
	}

	// Test case 6: Subtree is not an actual subtree
	subRoot5 := &util.TreeNode{Val: 4}
	subRoot5.Left = &util.TreeNode{Val: 1}
	subRoot5.Right = &util.TreeNode{Val: 3}
	if tree.SubtreeOfAnotherTree(root4, subRoot5) {
		t.Error("Expected false, got true")
	}
}
