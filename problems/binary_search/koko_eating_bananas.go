package binary_search

import "math"

func KokoEatingBananas(piles []int, h int) int {
	// Special case
	if len(piles) == 0 || h <= 0 {
		return 0
	}
	// Max element in the array
	maxNumberOfBananas := 0
	for _, pile := range piles {
		maxNumberOfBananas = max(maxNumberOfBananas, pile)
	}
	// Now, Koko can eat minimum one banana and maximum maxNumberOfBananas
	left := 1
	right := maxNumberOfBananas
	// Minimum rate
	minRate := right
	// Perform binary search in this range
	for left <= right {
		currentRate := left + (right-left)/2
		timeTaken := 0.0
		for _, pile := range piles {
			timeTaken += math.Ceil(float64(pile) / float64(currentRate))
		}
		if int(timeTaken) <= h {
			minRate = currentRate
			right = currentRate - 1
		} else {
			left = currentRate + 1
		}
	}
	return minRate
}
