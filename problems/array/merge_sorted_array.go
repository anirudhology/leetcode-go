package array

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	// Indices to keep track of nums1 and nums2
	i, j := m-1, n-1
	// Current index
	index := m + n - 1
	// Process both arrays
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i -= 1
		} else {
			nums1[index] = nums2[j]
			j -= 1
		}
		index -= 1
	}
	for i >= 0 {
		nums1[index] = nums1[i]
		i -= 1
		index -= 1
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		j -= 1
		index -= 1
	}
	return nums1
}
