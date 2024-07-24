package linked_list_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func buildList(values []int) *util.ListNode {
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

func listToArray(head *util.ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1}, expected: []int{1}},
		{input: []int{1, 2}, expected: []int{1, 2}},
		{input: []int{1, 2, 3}, expected: []int{1, 3, 2}},
		{input: []int{1, 2, 3, 4}, expected: []int{1, 4, 2, 3}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 5, 2, 4, 3}},
		{input: []int{1, 2, 3, 4, 5, 6}, expected: []int{1, 6, 2, 5, 3, 4}},
		{input: []int{1, 6, 2, 5, 3, 4}, expected: []int{1, 4, 6, 3, 2, 5}},
	}

	for _, test := range tests {
		head := buildList(test.input)
		linked_list.ReorderList(head)
		output := listToArray(head)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, output)
		}
	}
}
