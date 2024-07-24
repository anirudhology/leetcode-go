package linked_list_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func convertListToArray(head *util.ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func arrayToList(array []int) *util.ListNode {
	if len(array) == 0 {
		return nil
	}
	head := &util.ListNode{Val: array[0]}
	current := head
	for _, val := range array[1:] {
		current.Next = &util.ListNode{Val: val}
		current = current.Next
	}
	return head
}

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{"remove first node", []int{1, 2, 3}, 3, []int{2, 3}},
		{"remove last node", []int{1, 2, 3}, 1, []int{1, 2}},
		{"remove middle node", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := arrayToList(tt.input)
			result := linked_list.RemoveNthFromEnd(head, tt.n)
			resultArray := convertListToArray(result)
			if !reflect.DeepEqual(resultArray, tt.expected) {
				t.Errorf("got %v, expected %v", resultArray, tt.expected)
			}
		})
	}
}
