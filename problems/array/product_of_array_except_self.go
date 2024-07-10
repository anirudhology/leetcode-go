package array

func ProductOfArrayExceptSelf(nums []int) []int {
	// Special case
	if len(nums) < 2 {
		return nums
	}
	// Length of the array
	n := len(nums)
	// Slice to store the final output
	product := make([]int, n)
	// Cumulative product
	cumulativeProduct := 1
	// Process the array from left to right
	for i := 0; i < n; i++ {
		product[i] = cumulativeProduct
		cumulativeProduct *= nums[i]
	}
	// Reset the cumulativeProduct
	cumulativeProduct = 1
	// Process the array from right to left
	for i := n - 1; i >= 0; i-- {
		product[i] *= cumulativeProduct
		cumulativeProduct *= nums[i]
	}
	return product
}
