package linked_list_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

// Helper function to create a linked list from an array
func createLinkedList(arr []int) *util.ListNode {
	dummy := &util.ListNode{}
	current := dummy
	for _, num := range arr {
		current.Next = &util.ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to convert a linked list to an array
func linkedListToArray(list *util.ListNode) []int {
	var arr []int
	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}
	return arr
}

func TestMergeKSortedLists(t *testing.T) {
	tests := []struct {
		lists    [][]int
		expected []int
	}{
		{lists: [][]int{}, expected: []int{}},
		{lists: [][]int{nil}, expected: []int{}},
		{lists: [][]int{{1, 2, 3}}, expected: []int{1, 2, 3}},
		{lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, expected: []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{lists: [][]int{{1, 3}, {2}, {4, 5, 6}}, expected: []int{1, 2, 3, 4, 5, 6}},
		{lists: [][]int{{1, 3, 5}, {}, {2, 4, 6}}, expected: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range tests {
		var listNodes []*util.ListNode
		for _, list := range test.lists {
			listNodes = append(listNodes, createLinkedList(list))
		}
		result := linked_list.MergeKSortedLists(listNodes)
		resultArr := linkedListToArray(result)
		if len(resultArr) != len(test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, resultArr)
			continue
		}
		for i, v := range resultArr {
			if v != test.expected[i] {
				t.Errorf("Expected %v, but got %v", test.expected, resultArr)
				break
			}
		}
	}
}
