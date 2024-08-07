package quick_select

func FindKthLargestElementInAnArray(nums []int, k int) int {
	k = len(nums) - k
	left, right := 0, len(nums)-1
	for left < right {
		pivotIndex := partition(nums, left, right)
		if pivotIndex < k {
			left = pivotIndex + 1
		} else if pivotIndex > k {
			right = pivotIndex - 1
		} else {
			break
		}
	}
	return nums[k]
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	pivotIndex := left
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[pivotIndex], nums[i] = nums[i], nums[pivotIndex]
			pivotIndex++
		}
	}
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
	return pivotIndex
}
