package tree

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/anirudhology/leetcode-go/problems/util"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *util.TreeNode) string {
	// Special case
	if root == nil {
		return ""
	}
	// String to store the serialized tree
	serialized := ""
	// Queue to perform BFS
	nodes := list.New()
	nodes.PushBack(root)
	// Process the tree
	for nodes.Len() > 0 {
		// Get node from the front
		node := nodes.Remove(nodes.Front()).(*util.TreeNode)
		if node == nil {
			serialized += "n" + " "
		} else {
			serialized += strconv.Itoa(node.Val) + " "
			nodes.PushBack(node.Left)
			nodes.PushBack(node.Right)
		}
	}
	return serialized
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *util.TreeNode {
	// Special case
	if data == "" {
		return nil
	}
	// Get values for all the nodes to be constructed
	values := strings.Split(data, " ")
	val, _ := strconv.Atoi(values[0])
	root := &util.TreeNode{Val: val}
	// Queue to perform BFS
	nodes := list.New()
	nodes.PushBack(root)
	// Process the remaining values
	for i := 1; i < len(values); i += 2 {
		if nodes.Len() == 0 {
			break
		}
		// Get node from the front of the queue
		node := nodes.Remove(nodes.Front()).(*util.TreeNode)
		if values[i] != "n" {
			val, _ := strconv.Atoi(values[i])
			node.Left = &util.TreeNode{Val: val}
			nodes.PushBack(node.Left)
		}
		if i+1 < len(values) && values[i+1] != "n" {
			val, _ := strconv.Atoi(values[i+1])
			node.Right = &util.TreeNode{Val: val}
			nodes.PushBack(node.Right)
		}
	}
	return root
}
