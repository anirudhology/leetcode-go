package sliding_window

func MaxSlidingWindow(nums []int, k int) []int {
	// Special case
	if len(nums) == 0 || k < 0 {
		return nums
	}
	// Length of the array
	n := len(nums)
	// Slice to store the sliding window elements
	slidingWindowMax := make([]int, 0, n-k+1)
	// Deque to store the elements of the sliding window
	deque := make([]int, 0)
	// Process the array
	for i := 0; i < n; i++ {
		// Remove elements from the left
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		// Remove elements from the deque that are less than nums[i] as
		// they can never become max for this window
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		// Add current index to the deque
		deque = append(deque, i)
		if i >= k-1 {
			slidingWindowMax = append(slidingWindowMax, nums[deque[0]])
		}
	}
	return slidingWindowMax
}
