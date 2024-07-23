package linked_list_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 5}}}
	list2 := &util.ListNode{Val: 2, Next: &util.ListNode{Val: 4, Next: &util.ListNode{Val: 6}}}
	expected := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3, Next: &util.ListNode{Val: 4, Next: &util.ListNode{Val: 5, Next: &util.ListNode{Val: 6}}}}}}

	result := linked_list.MergeTwoSortedLists(list1, list2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMergeWithNullList(t *testing.T) {
	list1 := (*util.ListNode)(nil)
	list2 := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}}
	expected := &util.ListNode{Val: 1, Next: &util.ListNode{Val: 2, Next: &util.ListNode{Val: 3}}}

	result := linked_list.MergeTwoSortedLists(list1, list2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMergeBothNull(t *testing.T) {
	list1 := (*util.ListNode)(nil)
	list2 := (*util.ListNode)(nil)
	var expected *util.ListNode

	result := linked_list.MergeTwoSortedLists(list1, list2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
