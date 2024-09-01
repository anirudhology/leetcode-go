package dynamic_programming

func EditDistance(word1 string, word2 string) int {
	// Lengths of strings
	m, n := len(word1), len(word2)
	// Lookup table to store the minimum operations count
	lookup := make([][]int, m+1)
	for i := range lookup {
		lookup[i] = make([]int, n+1)
		for j := range lookup[i] {
			lookup[i][j] = 0
		}
	}
	// Base initialization
	for i := 0; i <= m; i++ {
		lookup[i][0] = i
	}
	for j := 0; j <= n; j++ {
		lookup[0][j] = j
	}
	// Populate the remaining table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If both characters are same, we don't have to do anything
			cost := 0
			if word1[i-1] != word2[j-1] {
				cost = 1
			}
			lookup[i][j] = min(cost+lookup[i-1][j-1], 1+min(lookup[i-1][j], lookup[i][j-1]))
		}
	}
	return lookup[m][n]
}
