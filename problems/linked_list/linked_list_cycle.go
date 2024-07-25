package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func LinkedListCycle(head *util.ListNode) bool {
	// Special case
	if head == nil {
		return false
	}
	// Slow and fast pointers
	slow, fast := head, head
	// Traverse the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
