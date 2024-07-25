package linked_list_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

// Helper function to create a linked list from a slice of integers
func createList(nums []int) *util.ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &util.ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &util.ListNode{Val: num}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice of integers
func listToSlice(head *util.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Helper function to compare two linked lists
func listsEqual(l1, l2 *util.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			l1:       []int{9, 9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 0, 1},
		},
		{
			l1:       []int{1, 8},
			l2:       []int{0},
			expected: []int{1, 8},
		},
	}

	for _, test := range tests {
		l1 := createList(test.l1)
		l2 := createList(test.l2)
		expected := createList(test.expected)

		result := linked_list.AddTwoNumbers(l1, l2)
		if !listsEqual(result, expected) {
			t.Errorf("AddTwoNumbers(%v, %v) = %v; want %v", test.l1, test.l2, listToSlice(result), test.expected)
		}
	}
}
