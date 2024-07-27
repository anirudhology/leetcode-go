package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func MergeKSortedLists(lists []*util.ListNode) *util.ListNode {
	// Special case
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*util.ListNode, left int, right int) *util.ListNode {
	// Boundary condition
	if left >= right {
		return lists[left]
	}
	// Divide the lists in two halves
	middle := left + int((right-left)/2)
	leftHalf := mergeLists(lists, left, middle)
	rightHalf := mergeLists(lists, middle+1, right)
	// Merge lists
	return merge(leftHalf, rightHalf)
}

func merge(left *util.ListNode, right *util.ListNode) *util.ListNode {
	// Dummy head
	dummy := &util.ListNode{}
	// Pointer to traverse the lists
	temp := dummy
	// Process both lists
	for left != nil && right != nil {
		if left.Val < right.Val {
			temp.Next = &util.ListNode{Val: left.Val}
			left = left.Next
		} else {
			temp.Next = &util.ListNode{Val: right.Val}
			right = right.Next
		}
		temp = temp.Next
	}
	for left != nil {
		temp.Next = &util.ListNode{Val: left.Val}
		left = left.Next
		temp = temp.Next
	}
	for right != nil {
		temp.Next = &util.ListNode{Val: right.Val}
		right = right.Next
		temp = temp.Next
	}
	return dummy.Next
}
