package dynamic_programming

func LongestCommonSubsequence(text1 string, text2 string) int {
	// Lengths of both strings
	m, n := len(text1), len(text2)
	// Lookup table to store the longest length until the index i
	// of string text and index j of string text 2
	lookup := make([][]int, m+1)
	for i := range lookup {
		lookup[i] = make([]int, n+1)
		for j := range lookup[i] {
			lookup[i][j] = 0
		}
	}
	// Populate the table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				lookup[i][j] = 1 + lookup[i-1][j-1]
			} else {
				lookup[i][j] = max(lookup[i-1][j], lookup[i][j-1])
			}
		}
	}
	return lookup[m][n]
}
