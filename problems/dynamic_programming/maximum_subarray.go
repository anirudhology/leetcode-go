package dynamic_programming

func MaximumSubArray(nums []int) int {
	localMaxima, globalMaxima := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		localMaxima = max(localMaxima+nums[i], nums[i])
		globalMaxima = max(globalMaxima, localMaxima)
	}
	return globalMaxima
}
