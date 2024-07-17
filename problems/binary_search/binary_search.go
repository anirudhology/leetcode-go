package binary_search

func BinarySearch(nums []int, target int) int {
	// Special case
	if len(nums) == 0 {
		return -1
	}
	// Left and right pointers
	left := 0
	right := len(nums) - 1
	// Process the array from both ends
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
