package binary_search

func SearchInRotatedSortedArray(nums []int, target int) int {
	// Special case
	if len(nums) == 0 {
		return -1
	}
	// Left and right pointers
	left := 0
	right := len(nums) - 1
	// Process the array from both sides
	for left <= right {
		middle := left + int((right-left)/2)
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] {
			if target >= nums[left] && target <= nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target >= nums[middle] && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}
