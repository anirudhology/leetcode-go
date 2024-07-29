package tree

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {
	// Test case 1: Both nodes are null
	if tree.LowestCommonAncestorOfABinarySearchTree(nil, nil, nil) != nil {
		t.Error("Expected nil, got non-nil")
	}

	// Test case 2: One node is null
	root := &util.TreeNode{Val: 3}
	p := &util.TreeNode{Val: 1}
	if tree.LowestCommonAncestorOfABinarySearchTree(root, nil, p) != nil {
		t.Error("Expected nil, got non-nil")
	}
	if tree.LowestCommonAncestorOfABinarySearchTree(root, p, nil) != nil {
		t.Error("Expected nil, got non-nil")
	}

	// Test case 3: Both nodes are the same
	if tree.LowestCommonAncestorOfABinarySearchTree(p, p, p) != p {
		t.Error("Expected p, got different node")
	}

	// Test case 4: LCA is root
	q := &util.TreeNode{Val: 5}
	root.Left = p
	root.Right = q
	if tree.LowestCommonAncestorOfABinarySearchTree(root, p, q) != root {
		t.Error("Expected root, got different node")
	}

	// Test case 5: LCA is in the left subtree
	p2 := &util.TreeNode{Val: 0}
	p.Left = p2
	if tree.LowestCommonAncestorOfABinarySearchTree(root, p2, p) != p {
		t.Error("Expected p, got different node")
	}

	// Test case 6: LCA is in the right subtree
	q2 := &util.TreeNode{Val: 6}
	q.Right = q2
	if tree.LowestCommonAncestorOfABinarySearchTree(root, q, q2) != q {
		t.Error("Expected q, got different node")
	}

	// Test case 7: Complex tree
	root2 := &util.TreeNode{Val: 6}
	p3 := &util.TreeNode{Val: 2}
	q3 := &util.TreeNode{Val: 8}
	root2.Left = &util.TreeNode{Val: 0}
	root2.Right = &util.TreeNode{Val: 4}
	root2.Left.Left = &util.TreeNode{Val: 3}
	root2.Left.Right = &util.TreeNode{Val: 5}
	if tree.LowestCommonAncestorOfABinarySearchTree(root2, p3, q3) != root2 {
		t.Error("Expected root2, got different node")
	}
}
