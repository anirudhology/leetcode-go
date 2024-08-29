package greedy

func MergeTripletsToFormTargetTriplet(triplets [][]int, target []int) bool {
	var validIndices []int
	for i := 0; i < len(triplets); i++ {
		if triplets[i][0] <= target[0] && triplets[i][1] <= target[1] && triplets[i][2] <= target[2] {
			validIndices = append(validIndices, i)
		}
	}
	isXPresent, isYPresent, isZPresent := false, false, false
	for _, i := range validIndices {
		if !isXPresent && triplets[i][0] == target[0] {
			isXPresent = true
		}
		if !isYPresent && triplets[i][1] == target[1] {
			isYPresent = true
		}
		if !isZPresent && triplets[i][2] == target[2] {
			isZPresent = true
		}
	}
	return isXPresent && isYPresent && isZPresent
}
