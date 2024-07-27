package linked_list_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

// Helper function to create a list from a slice
func create(values []int) *util.ListNode {
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

// Helper function to convert a list to a slice
func listSlice(head *util.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "Empty list",
			input:    []int{},
			k:        2,
			expected: []int{},
		},
		{
			name:     "k <= 0",
			input:    []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "k == 1",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "List length < k",
			input:    []int{1, 2},
			k:        3,
			expected: []int{1, 2},
		},
		{
			name:     "List length == k",
			input:    []int{1, 2, 3},
			k:        3,
			expected: []int{3, 2, 1},
		},
		{
			name:     "List length > k",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "List length > k with remainder",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{3, 2, 1, 6, 5, 4, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := create(tt.input)
			resultList := linked_list.ReverseNodesInKGroup(inputList, tt.k)
			result := listSlice(resultList)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Helper function to compare two slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
