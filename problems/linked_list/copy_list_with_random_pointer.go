package linked_list

import "github.com/anirudhology/leetcode-go/problems/util"

func CopyListWithRandomPointer(head *util.ListNodeWithRandom) *util.ListNodeWithRandom {
	// Special case
	if head == nil {
		return nil
	}
	// Map to store original and cloned nodes' mappings
	mappings := map[*util.ListNodeWithRandom]*util.ListNodeWithRandom{}
	// Head of the cloned list
	var clone *util.ListNodeWithRandom
	// Pointers for both lists
	temp := head
	var cloneTemp *util.ListNodeWithRandom
	// Traverse through the list
	for temp != nil {
		// Create copy of the current node
		copy := &util.ListNodeWithRandom{Val: temp.Val}
		if cloneTemp == nil {
			clone = copy
			cloneTemp = clone
		} else {
			cloneTemp.Next = copy
			cloneTemp = cloneTemp.Next
		}
		// Add mappings
		mappings[temp] = cloneTemp
		temp = temp.Next
	}
	// Reset both pointers
	temp, cloneTemp = head, clone
	// Again traverse the list
	for temp != nil {
		if temp.Random != nil {
			cloneTemp.Random = mappings[temp.Random]
		}
		temp = temp.Next
		cloneTemp = cloneTemp.Next
	}
	return clone
}
