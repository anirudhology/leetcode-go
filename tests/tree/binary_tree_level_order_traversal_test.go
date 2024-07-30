package tree_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestLevelOrder(t *testing.T) {
	// Test case 1: Empty tree
	var root1 *util.TreeNode = nil
	result1 := tree.BinaryTreeLevelOrderTraversal(root1)
	if len(result1) != 0 {
		t.Errorf("Expected empty list, got %v", result1)
	}

	// Test case 2: Single node tree
	root2 := &util.TreeNode{Val: 1}
	expected2 := [][]int{{1}}
	result2 := tree.BinaryTreeLevelOrderTraversal(root2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Tree with two levels
	root3 := &util.TreeNode{Val: 1}
	root3.Left = &util.TreeNode{Val: 2}
	root3.Right = &util.TreeNode{Val: 3}
	expected3 := [][]int{{1}, {2, 3}}
	result3 := tree.BinaryTreeLevelOrderTraversal(root3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, got %v", expected3, result3)
	}

	// Test case 4: Tree with three levels
	root4 := &util.TreeNode{Val: 1}
	root4.Left = &util.TreeNode{Val: 2}
	root4.Right = &util.TreeNode{Val: 3}
	root4.Left.Left = &util.TreeNode{Val: 4}
	root4.Left.Right = &util.TreeNode{Val: 5}
	root4.Right.Left = &util.TreeNode{Val: 6}
	root4.Right.Right = &util.TreeNode{Val: 7}
	expected4 := [][]int{{1}, {2, 3}, {4, 5, 6, 7}}
	result4 := tree.BinaryTreeLevelOrderTraversal(root4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Expected %v, got %v", expected4, result4)
	}
}
