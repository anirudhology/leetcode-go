package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func ReverseLinkedList(head *util.ListNode) *util.ListNode {
	// Special case
	if head == nil || head.Next == nil {
		return head
	}
	// Three pointers to traverse the list
	var previous *util.ListNode
	current := head
	var next *util.ListNode
	// Traverse the list
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
