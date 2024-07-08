package strings

import (
	"fmt"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	// List to store the final output
	groups := make([][]string, 0)
	// Special case
	if len(strs) == 0 {
		return groups
	}
	// Map to store the frequency representation of string and the
	// related strings
	mappings := make(map[string][]string)
	// Process each string one by one
	for _, s := range strs {
		// Calculate frequencies for the current string
		var frequencies [26]int
		for _, c := range s {
			frequencies[int(c)-int('a')]++
		}
		var key strings.Builder
		for _, f := range frequencies {
			key.WriteString(fmt.Sprintf("%d#", f))
		}
		mappings[key.String()] = append(mappings[key.String()], s)
	}
	// Add related groups together
	for _, value := range mappings {
		groups = append(groups, value)
	}
	return groups
}
