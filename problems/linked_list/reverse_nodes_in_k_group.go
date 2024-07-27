package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func ReverseNodesInKGroup(head *util.ListNode, k int) *util.ListNode {
	// Special case
	if head == nil || k <= 0 {
		return head
	}
	// Dummy node
	dummy := &util.ListNode{}
	dummy.Next = head
	// Pointer the traverse the list
	temp := dummy
	// Traverse the list
	for {
		// Current temp
		currentTemp := temp
		// Check if there are k nodes to reverse
		for i := 0; i < k && currentTemp != nil; i++ {
			currentTemp = currentTemp.Next
		}
		if currentTemp == nil {
			break
		}
		// Reverse the k nodes in the list
		var previous *util.ListNode
		current := temp.Next
		var next *util.ListNode
		for i := 0; i < k; i++ {
			next = current.Next
			current.Next = previous
			previous = current
			current = next
		}
		// Tail of the reversed list
		tail := temp.Next
		tail.Next = current
		temp.Next = previous
		temp = tail
	}
	return dummy.Next
}
