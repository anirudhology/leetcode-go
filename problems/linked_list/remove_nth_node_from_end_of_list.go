package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func RemoveNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	// Slow and fast pointers
	slow, fast := head, head
	// Move fast pointer n steps ahead
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}
	// If fast is null, it means n >= number of nodes in the list
	if fast == nil && head != nil {
		return head.Next
	}
	// Move both pointers with the same pace
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow != nil && slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return head
}
