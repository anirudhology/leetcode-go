package dynamic_programming

func DistinctSubsequences(s string, t string) int {
	// Lengths of both strings
	m, n := len(s), len(t)
	// Lookup table to store distinct subsequences
	lookup := make([][]int, m+1)
	for i := range lookup {
		lookup[i] = make([]int, n+1)
		for j := range lookup[i] {
			lookup[i][j] = 0
		}
	}
	// If t is empty, it is subsequence of every string but only once
	for i := 0; i <= m; i++ {
		lookup[i][0] = 1
	}
	// If both s and t are empty
	lookup[0][0] = 1
	// Populate the remaining table
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			lookup[i][j] = lookup[i-1][j]
			if s[i-1] == t[j-1] {
				lookup[i][j] += lookup[i-1][j-1]
			}
		}
	}
	return lookup[m][n]
}
