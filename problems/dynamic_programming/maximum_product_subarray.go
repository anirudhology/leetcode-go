package dynamic_programming

func MaximumProductSubarray(nums []int) int {
	// Special case
	if len(nums) == 0 {
		return 0
	}
	// Both positive and negative numbers can contribute to the
	// maximum product subarray, hence we will maintain max and
	// min products calculated so far
	currentMax, currentMin := nums[0], nums[0]
	// Final maximum product
	maxProduct := nums[0]
	// Process the array
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		// Update current max and min values by considering the
		// current element
		currentMax = max(nums[i], nums[i]*currentMax)
		currentMin = min(nums[i], nums[i]*currentMin)
		// Update maxProduct until current index
		maxProduct = max(currentMax, maxProduct)
	}
	return maxProduct
}
