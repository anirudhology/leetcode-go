package binary_search

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Make sure that nums1 is the bigger array
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	// Length of the arrays
	m, n := len(nums1), len(nums2)
	// Left and right pointers for the binary search
	left, right := 0, m
	// Process the array from both ends
	for left <= right {
		partitionM := (left + right) / 2
		partitionN := (m+n+1)/2 - partitionM
		// Find boundaries for both arrays
		maxLeftM := math.MinInt32
		if partitionM != 0 {
			maxLeftM = nums1[partitionM-1]
		}
		minRightM := math.MaxInt32
		if partitionM != m {
			minRightM = nums1[partitionM]
		}
		maxLeftN := math.MinInt32
		if partitionN != 0 {
			maxLeftN = nums2[partitionN-1]
		}
		minRightN := math.MaxInt32
		if partitionN != n {
			minRightN = nums2[partitionN]
		}
		// Check if we have found the correct boundaries
		if maxLeftM <= minRightN && maxLeftN <= minRightM {
			if (m+n)%2 == 0 {
				return float64(math.Max(float64(maxLeftM), float64(maxLeftN))+math.Min(float64(minRightM), float64(minRightN))) / 2.0
			} else {
				return math.Max(float64(maxLeftM), float64(maxLeftN))
			}
		} else if maxLeftM > minRightN {
			right = partitionM - 1
		} else {
			left = partitionM + 1
		}
	}
	panic("Invalid input")
}
