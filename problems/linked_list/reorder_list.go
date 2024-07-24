package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func ReorderList(head *util.ListNode) {
	// Special case
	if head == nil || head.Next == nil {
		return
	}
	// Find the middle of the list
	middle := getMiddle(head)
	head1, head2 := head, middle.Next
	// Detach both halves
	middle.Next = nil
	// Reverse the second half
	reverseHead2 := reverse(head2)
	// Merge the two halves
	for head1 != nil && reverseHead2 != nil {
		next1, next2 := head1.Next, reverseHead2.Next
		head1.Next = reverseHead2
		reverseHead2.Next = next1
		head1 = next1
		reverseHead2 = next2
	}
}

func getMiddle(head *util.ListNode) *util.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *util.ListNode) *util.ListNode {
	var previous *util.ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
