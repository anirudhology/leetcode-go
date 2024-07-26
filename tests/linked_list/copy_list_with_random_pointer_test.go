package linked_list_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

// Helper function to create a linked list with given values and random pointers
func createListWithRandom(vals []int, randomIndexes []int) *util.ListNodeWithRandom {
	if len(vals) == 0 {
		return nil
	}

	// Create nodes
	nodes := make([]*util.ListNodeWithRandom, len(vals))
	for i, val := range vals {
		nodes[i] = &util.ListNodeWithRandom{Val: val}
	}

	// Set next pointers
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Set random pointers
	for i, index := range randomIndexes {
		if index != -1 {
			nodes[i].Random = nodes[index]
		} else {
			nodes[i].Random = nil
		}
	}

	return nodes[0]
}

// Helper function to compare two lists
func compareLists(l1, l2 *util.ListNodeWithRandom) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		if l1.Random != nil && l2.Random != nil {
			if l1.Random.Val != l2.Random.Val {
				return false
			}
		} else if l1.Random != l2.Random {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func TestCopyListWithRandomPointer(t *testing.T) {
	// Test case 1: Empty list
	var head *util.ListNodeWithRandom
	copied := linked_list.CopyListWithRandomPointer(head)
	if copied != nil {
		t.Errorf("TestCopyRandomList failed: expected nil, got %v", copied)
	}

	// Test case 2: Single node with no random pointer
	head = createListWithRandom([]int{1}, []int{-1})
	copied = linked_list.CopyListWithRandomPointer(head)
	if !compareLists(head, copied) {
		t.Errorf("TestCopyRandomList failed: expected %v, got %v", head, copied)
	}

	// Test case 3: Single node with self-referencing random pointer
	head = createListWithRandom([]int{1}, []int{0})
	copied = linked_list.CopyListWithRandomPointer(head)
	if !compareLists(head, copied) {
		t.Errorf("TestCopyRandomList failed: expected %v, got %v", head, copied)
	}

	// Test case 4: Multi-node list with no random pointers
	head = createListWithRandom([]int{1, 2, 3}, []int{-1, -1, -1})
	copied = linked_list.CopyListWithRandomPointer(head)
	if !compareLists(head, copied) {
		t.Errorf("TestCopyRandomList failed: expected %v, got %v", head, copied)
	}

	// Test case 5: Multi-node list with random pointers
	head = createListWithRandom([]int{1, 2, 3}, []int{2, 0, -1})
	copied = linked_list.CopyListWithRandomPointer(head)
	if !compareLists(head, copied) {
		t.Errorf("TestCopyRandomList failed: expected %v, got %v", head, copied)
	}

	// Test case 6: Cyclic list
	head = createListWithRandom([]int{1, 2, 3}, []int{1, 2, 0})
	copied = linked_list.CopyListWithRandomPointer(head)
	if !compareLists(head, copied) {
		t.Errorf("TestCopyRandomList failed: expected %v, got %v", head, copied)
	}
}
