package dynamic_programming

func InterleavingString(s1 string, s2 string, s3 string) bool {
	// Lengths of strings
	m, n := len(s1), len(s2)
	// Special case
	if m+n != len(s3) {
		return false
	}
	// Lookup table to check if s3[0...k] can be formed by s1[0...i] and s2[0...j]
	lookup := make([][]bool, m+1)
	for i := range lookup {
		lookup[i] = make([]bool, n+1)
		for j := range lookup[i] {
			lookup[i][j] = false
		}
	}
	// For empty strings
	lookup[0][0] = true
	// If s2 is empty
	for i := 1; i <= m; i++ {
		lookup[i][0] = lookup[i-1][0] && s1[i-1] == s3[i-1]
	}
	// If s1 is empty
	for j := 1; j <= n; j++ {
		lookup[0][j] = lookup[0][j-1] && s2[j-1] == s3[j-1]
	}
	// Populate the remaining table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			lookup[i][j] = (lookup[i-1][j] && s1[i-1] == s3[i+j-1]) || (lookup[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return lookup[m][n]
}
