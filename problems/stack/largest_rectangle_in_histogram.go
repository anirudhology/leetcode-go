package stack

func LargestRectangleInHistogram(heights []int) int {
	// Special case
	if len(heights) == 0 {
		return 0
	}
	// Length of the array
	n := len(heights)
	// Monotonic stack to store indices
	stack := make([]int, 0)
	// Maximum area
	maxArea := 0
	// Process the heights
	for i := 0; i <= n; i++ {
		height := 0
		if i < n {
			height = heights[i]
		}
		for len(stack) > 0 && height < heights[stack[len(stack)-1]] {
			currentHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			currentWidth := i
			if len(stack) > 0 {
				currentWidth = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, currentHeight*currentWidth)
		}
		stack = append(stack, i)
	}
	return maxArea
}
