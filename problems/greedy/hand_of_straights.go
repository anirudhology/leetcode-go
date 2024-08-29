package greedy

import (
	"sort"
)

func HandOfStraights(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false // Cannot divide into groups evenly
	}

	// Count the occurrences of each card
	countMap := make(map[int]int)
	for _, card := range hand {
		countMap[card]++
	}

	// Sort the hand slice
	sort.Ints(hand)

	// Try to form groups
	for _, card := range hand {
		if countMap[card] > 0 { // If the current card is still available
			// Form a group starting with 'card'
			for i := 0; i < groupSize; i++ {
				currentCard := card + i
				if countMap[currentCard] > 0 {
					// Decrement the count of the current card
					countMap[currentCard]--
				} else {
					// If a card needed to complete the group is missing, return false
					return false
				}
			}
		}
	}
	return true
}
