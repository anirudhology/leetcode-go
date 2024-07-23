package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func MergeTwoSortedLists(head1 *util.ListNode, head2 *util.ListNode) *util.ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	// Head for the resultant list
	var head *util.ListNode
	if head1.Val < head2.Val {
		head = head1
		head1 = head1.Next
	} else {
		head = head2
		head2 = head2.Next
	}

	// Pointer to traverse through the resultant list
	temp := head

	// Process both lists together
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			temp.Next = head1
			head1 = head1.Next
		} else {
			temp.Next = head2
			head2 = head2.Next
		}
		temp = temp.Next
	}

	// Process remaining nodes
	if head1 != nil {
		temp.Next = head1
	} else {
		temp.Next = head2
	}

	return head
}
