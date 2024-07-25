package linked_list_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		name string
		head *util.ListNode
		want bool
	}{
		{
			name: "returns false for an empty list",
			head: nil,
			want: false,
		},
		{
			name: "returns false for a list with one node and no cycle",
			head: &util.ListNode{Val: 1},
			want: false,
		},
		{
			name: "returns false for a list with multiple nodes and no cycle",
			head: &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}},
			want: false,
		},
		{
			name: "returns true for a list with a cycle",
			head: func() *util.ListNode {
				head := &util.ListNode{Val: 1}
				second := &util.ListNode{Val: 2}
				third := &util.ListNode{Val: 3}
				head.Next = second
				second.Next = third
				third.Next = second // Creating a cycle
				return head
			}(),
			want: true,
		},
		{
			name: "returns true for a cycle at the head",
			head: func() *util.ListNode {
				head := &util.ListNode{Val: 1}
				second := &util.ListNode{Val: 2}
				head.Next = second
				second.Next = head // Creating a cycle at the head
				return head
			}(),
			want: true,
		},
		{
			name: "returns true for a cycle at the middle",
			head: func() *util.ListNode {
				head := &util.ListNode{Val: 1}
				second := &util.ListNode{Val: 2}
				third := &util.ListNode{Val: 3}
				fourth := &util.ListNode{Val: 4}
				head.Next = second
				second.Next = third
				third.Next = fourth
				fourth.Next = second // Creating a cycle at the second node
				return head
			}(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linked_list.LinkedListCycle(tt.head); got != tt.want {
				t.Errorf("LinkedListCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
