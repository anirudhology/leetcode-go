package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func AddTwoNumbers(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	// Dummy head of the resultant list
	dummy := &util.ListNode{}
	// Pointer to traverse the resultant list
	temp := dummy
	// Carry
	carry := 0
	// Process both lists together
	for l1 != nil && l2 != nil {
		sum := carry + l1.Val + l2.Val
		carry = int(sum / 10)
		temp.Next = &util.ListNode{Val: sum % 10}
		temp = temp.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	// For remaining nodes
	for l1 != nil {
		sum := carry + l1.Val
		carry = int(sum / 10)
		temp.Next = &util.ListNode{Val: sum % 10}
		temp = temp.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := carry + l2.Val
		carry = int(sum / 10)
		temp.Next = &util.ListNode{Val: sum % 10}
		temp = temp.Next
		l2 = l2.Next
	}
	// Adjust the remaining carry
	if carry != 0 {
		temp.Next = &util.ListNode{Val: carry}
	}
	return dummy.Next
}
