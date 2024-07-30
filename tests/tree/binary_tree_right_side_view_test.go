package tree_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestBinaryTreeRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *util.TreeNode
		want []int
	}{
		{"single node tree", &util.TreeNode{Val: 1}, []int{1}},
		{"right leaning tree", &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 2, Right: &util.TreeNode{Val: 3}}}, []int{1, 2, 3}},
		{"left leaning tree", &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 3}}}, []int{1, 2, 3}},
		{"balanced tree", &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 4}, Right: &util.TreeNode{Val: 5}}, Right: &util.TreeNode{Val: 3, Right: &util.TreeNode{Val: 6}}}, []int{1, 3, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.BinaryTreeRightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
