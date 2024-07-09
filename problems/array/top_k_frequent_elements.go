package array

func TopKFrequent(nums []int, k int) []int {
	// Special case
	if len(nums) == 0 || k < 0 {
		return make([]int, 0)
	}
	// Map to store the frequencies of elements in the array
	frequencies := make(map[int]int)
	for _, num := range nums {
		frequencies[num]++
	}
	// List to store the buckets
	buckets := make([][]int, len(nums)+1)
	// Populate the buckets
	for key, value := range frequencies {
		buckets[value] = append(buckets[value], key)
	}
	// Array to store the final output
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}
	return result
}
