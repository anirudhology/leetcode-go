package dynamic_programming

func RegularExpressionMatching(s string, p string) bool {
	// Length of strings
	m, n := len(s), len(p)
	// Lookup table to store if s[0...i] matches p[0...j]
	lookup := make([][]bool, m+1)
	for i := range lookup {
		lookup[i] = make([]bool, n+1)
		for j := range lookup[i] {
			lookup[i][j] = false
		}
	}
	// Empty string matches with empty patten
	lookup[0][0] = true
	// For empty string and star pattern
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			lookup[0][j] = lookup[0][j-2]
		}
	}
	// Populate remaining table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If characters are same or pattern is period
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				lookup[i][j] = lookup[i-1][j-1]
			} else if j > 1 && p[j-1] == '*' {
				lookup[i][j] = lookup[i][j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					lookup[i][j] = lookup[i][j] || lookup[i-1][j]
				}
			}
		}
	}
	return lookup[m][n]
}
