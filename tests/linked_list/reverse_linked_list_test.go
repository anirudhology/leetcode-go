package linked_list

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func createLinkedList(values []int) *util.ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &util.ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &util.ListNode{Val: val}
		current = current.Next
	}
	return head
}

func linkedListToSlice(head *util.ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{input: nil, output: nil},
		{input: []int{1}, output: []int{1}},
		{input: []int{1, 2, 3}, output: []int{3, 2, 1}},
		{input: []int{1, 2}, output: []int{2, 1}},
	}

	for _, test := range tests {
		head := createLinkedList(test.input)
		reversedHead := linked_list.ReverseLinkedList(head)
		result := linkedListToSlice(reversedHead)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.output, result)
		}
	}
}
