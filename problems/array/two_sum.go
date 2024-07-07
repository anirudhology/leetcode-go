package array

func TwoSum(nums []int, target int) []int {
	// Special case
	if len(nums) < 2 {
		return nums
	}
	// Map to store the element and its index mapping
	mappings := make(map[int]int)
	// Process every element of the array
	for i, num := range nums {
		// Check if this element already exists in the map
		if value, exists := mappings[num]; exists {
			return []int{value, i}
		}
		mappings[target-num] = i
	}
	return nil
}
