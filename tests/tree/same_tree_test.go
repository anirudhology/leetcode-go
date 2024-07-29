package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestSameTree(t *testing.T) {
	// Test case 1: Both trees are null
	if !tree.SameTree(nil, nil) {
		t.Error("Expected true, got false")
	}

	// Test case 2: One tree is null
	p1 := &util.TreeNode{Val: 1}
	if tree.SameTree(p1, nil) {
		t.Error("Expected false, got true")
	}
	if tree.SameTree(nil, p1) {
		t.Error("Expected false, got true")
	}

	// Test case 3: Both trees have one node with the same value
	p2 := &util.TreeNode{Val: 1}
	q2 := &util.TreeNode{Val: 1}
	if !tree.SameTree(p2, q2) {
		t.Error("Expected true, got false")
	}

	// Test case 4: Both trees have one node with different values
	p3 := &util.TreeNode{Val: 1}
	q3 := &util.TreeNode{Val: 2}
	if tree.SameTree(p3, q3) {
		t.Error("Expected false, got true")
	}

	// Test case 5: Both trees have multiple nodes with the same structure and values
	p4 := &util.TreeNode{Val: 1}
	p4.Left = &util.TreeNode{Val: 2}
	p4.Right = &util.TreeNode{Val: 3}
	q4 := &util.TreeNode{Val: 1}
	q4.Left = &util.TreeNode{Val: 2}
	q4.Right = &util.TreeNode{Val: 3}
	if !tree.SameTree(p4, q4) {
		t.Error("Expected true, got false")
	}

	// Test case 6: Both trees have multiple nodes with different structures
	p5 := &util.TreeNode{Val: 1}
	p5.Left = &util.TreeNode{Val: 2}
	q5 := &util.TreeNode{Val: 1}
	q5.Right = &util.TreeNode{Val: 2}
	if tree.SameTree(p5, q5) {
		t.Error("Expected false, got true")
	}
}
