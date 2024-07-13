package array

import "sort"

func ThreeSum(nums []int) [][]int {
	// List to store the final triplets
	var triplets [][]int
	// Special case
	if len(nums) < 3 {
		return triplets
	}
	// Length of the array
	n := len(nums)
	// Sort the array
	sort.Ints(nums)
	// Process the array
	for i := 0; i < n-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Left and right pointers of the remaining array
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// If the sum is zero
			if sum == 0 {
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// Skip duplicates for `j`
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// Skip duplicates for `k`
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return triplets
}
