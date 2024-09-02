package dynamic_programming

func BurstBalloons(nums []int) int {
	// Array to store balloons
	balloons := make([]int, len(nums)+2)
	n := len(balloons)
	balloons[0] = 1
	balloons[n-1] = 1
	for i := 1; i < n-1; i++ {
		balloons[i] = nums[i-1]
	}
	// Lookup table
	lookup := make([][]int, n-1)
	for i := range lookup {
		lookup[i] = make([]int, n-1)
		for j := range lookup[i] {
			lookup[i][j] = -1
		}
	}
	// Burst all balloons between 1 and n - 2
	return burst(balloons, 1, n-2, lookup)
}

func burst(balloons []int, i int, j int, lookup [][]int) int {
	// Base case
	if i > j {
		return 0
	}
	// If there is only one balloon
	if i == j {
		temp := balloons[i]
		if i-1 >= 0 {
			temp = temp * balloons[i-1]
		}
		if j+1 < len(balloons) {
			temp = temp * balloons[j+1]
		}
		return temp
	}
	// If we have already calculated this result
	if lookup[i][j] != -1 {
		return lookup[i][j]
	}
	// Score
	score := 0
	// Burst all balloons between range i...j
	for k := i; k <= j; k++ {
		temp := balloons[k]
		if i-1 >= 0 {
			temp = temp * balloons[i-1]
		}
		if j+1 < len(balloons) {
			temp = temp * balloons[j+1]
		}
		temp += (burst(balloons, i, k-1, lookup)) + (burst(balloons, k+1, j, lookup))
		score = max(score, temp)
	}
	lookup[i][j] = score
	return score
}
