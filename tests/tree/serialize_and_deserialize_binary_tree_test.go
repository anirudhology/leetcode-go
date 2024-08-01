package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	// Create test cases
	tests := []struct {
		root *util.TreeNode
		want string
	}{
		{
			// Test case: empty tree
			root: nil,
			want: "",
		},
		{
			// Test case: single node
			root: &util.TreeNode{Val: 1},
			want: "1 n n ",
		},
		{
			// Test case: balanced tree
			root: &util.TreeNode{
				Val: 1,
				Left: &util.TreeNode{
					Val:   2,
					Left:  &util.TreeNode{Val: 4},
					Right: &util.TreeNode{Val: 5},
				},
				Right: &util.TreeNode{
					Val:   3,
					Left:  &util.TreeNode{Val: 6},
					Right: &util.TreeNode{Val: 7},
				},
			},
			want: "1 2 3 4 5 6 7 n n n n n n n n ",
		},
		{
			// Test case: tree with only right children
			root: &util.TreeNode{
				Val: 1,
				Right: &util.TreeNode{
					Val: 2,
					Right: &util.TreeNode{
						Val: 3,
					},
				},
			},
			want: "1 n 2 n 3 n n ",
		},
	}

	c := tree.Constructor()
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := c.Serialize(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeserialize(t *testing.T) {
	// Create test cases
	tests := []struct {
		data string
		want *util.TreeNode
	}{
		{
			// Test case: empty string
			data: "",
			want: nil,
		},
		{
			// Test case: single node
			data: "1 n n ",
			want: &util.TreeNode{Val: 1},
		},
		{
			// Test case: balanced tree
			data: "1 2 3 4 5 6 7 n n n n n n n n ",
			want: &util.TreeNode{
				Val: 1,
				Left: &util.TreeNode{
					Val:   2,
					Left:  &util.TreeNode{Val: 4},
					Right: &util.TreeNode{Val: 5},
				},
				Right: &util.TreeNode{
					Val:   3,
					Left:  &util.TreeNode{Val: 6},
					Right: &util.TreeNode{Val: 7},
				},
			},
		},
		{
			// Test case: tree with only right children
			data: "1 n 2 n 3 n n ",
			want: &util.TreeNode{
				Val: 1,
				Right: &util.TreeNode{
					Val: 2,
					Right: &util.TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	c := tree.Constructor()
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := c.Deserialize(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
