package array

func SortColors(nums []int) []int {
	// Pointers to keep track of 0s, 1s and 2s
	left, middle, right := 0, 0, len(nums)-1
	// Process elements in the array
	for middle <= right {
		if nums[middle] == 0 {
			nums[middle], nums[left] = nums[left], nums[middle]
			left += 1
			middle += 1
		} else if nums[middle] == 1 {
			middle += 1
		} else {
			nums[middle], nums[right] = nums[right], nums[middle]
			right -= 1
		}
	}
	return nums
}
