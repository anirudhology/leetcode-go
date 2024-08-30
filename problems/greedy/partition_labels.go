package greedy

func PartitionLabels(s string) []int {
	// Slice to store the sizes of the partitions
	var sizes []int
	// Special case
	if len(s) == 0 {
		return sizes
	}
	// Array to store last positions of the characters
	lastPositions := make([]int, 26)
	for i := 0; i < len(s); i++ {
		lastPositions[int(s[i])-int('a')] = i
	}
	// Pointers for sliding window
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, lastPositions[int(s[i])-int('a')])
		if right == i {
			sizes = append(sizes, right-left+1)
			left = i + 1
		}
	}
	return sizes
}
