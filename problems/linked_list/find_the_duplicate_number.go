package linked_list

func FindTheDuplicateNumber(nums []int) int {
	// Slow and fast pointers
	slow, fast := nums[0], nums[nums[0]]
	// Traverse until the two pointers meet
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// Reset fast pointer
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
