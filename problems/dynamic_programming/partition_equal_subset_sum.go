package dynamic_programming

func PartitionEqualSubsetSum(nums []int) bool {
	// Special case
	if len(nums) == 0 {
		return false
	}
	// Sum of all the elements in the array
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// Check if partition is possible
	if sum%2 != 0 {
		return false
	}
	// Target to reach
	target := sum / 2
	// Lookup table to check if partition is possible
	lookup := make([]bool, target+1)
	lookup[0] = true
	// Populate the table
	for _, num := range nums {
		for j := target; j >= num; j-- {
			lookup[j] = lookup[j] || lookup[j-num]
		}
	}
	return lookup[target]
}
