package binary_search

func FindMinimumInRotatedSortedArray(nums []int) int {
	// Left and right pointers
	left := 0
	right := len(nums) - 1
	// Process the array
	for left < right {
		middle := left + int((right-left)/2)
		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return nums[left]
}
