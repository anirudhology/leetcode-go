package array

func containsDuplicate(nums []int) bool {
	// Special case
	if nums == nil || len(nums) == 0 {
		return false
	}
	// Hash set to store elements
	elements := make(map[int]struct{})
	// Process each element in the array one by one
	for _, num := range nums {
		if _, exists := elements[num]; exists {
			return true
		}
		elements[num] = struct{}{}
	}
	return false
}
