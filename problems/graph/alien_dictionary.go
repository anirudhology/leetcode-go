package graph

import "container/list"

func AlienDictionary(words []string) string {
	// Total number of words
	n := len(words)
	// Graph to represent the order of characters
	graph := make([][]bool, 26)
	for i := range graph {
		graph[i] = make([]bool, 26)
		for j := range graph[i] {
			graph[i][j] = false
		}
	}
	// Array to store the characters present in the dictionary
	characters := make([]bool, 26)
	// Unique character count
	uniqueCharacterCount := 0

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if uniqueCharacterCount == 26 {
				break
			}
			index := int(word[i]) - int('a')
			if !characters[index] {
				characters[index] = true
				uniqueCharacterCount++
			}
		}
	}
	// Populate the graph now
	for i := 0; i < n-1; i++ {
		for j := 0; j < 26; j++ {
			// If next word is a prefix of current word, we have
			// invalid order
			if j >= len(words[i+1]) {
				return ""
			}
			// Compare characters of two words
			current, next := words[i][j], words[i+1][j]
			if current == next {
				continue
			}
			x, y := int(current)-int('a'), int(next)-int('a')
			if graph[y][x] {
				return ""
			}
			graph[x][y] = true
			break
		}
	}
	// Now, the graph is created, we will now start the process of
	// topological sorting

	// 1. Find indegrees of every node
	indegrees := make([]int, 26)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i != j && characters[i] && characters[j] && graph[i][j] {
				indegrees[j]++
			}
		}
	}
	// 2. Queue to store nodes with zero indegrees
	zeroIndegreeNodes := list.New()
	for i := 0; i < 26; i++ {
		if characters[i] && indegrees[i] == 0 {
			zeroIndegreeNodes.PushBack(i)
		}
	}
	// 3. Perform topological sort
	order := ""
	for zeroIndegreeNodes.Len() > 0 {
		current := zeroIndegreeNodes.Remove(zeroIndegreeNodes.Front()).(int)
		order += string(rune(current + 'a'))
		for i := 0; i < 26; i++ {
			if i != current && characters[i] && graph[current][i] {
				indegrees[i]--
				if indegrees[i] == 0 {
					zeroIndegreeNodes.PushBack(i)
				}
			}
		}
	}
	if len(order) < uniqueCharacterCount {
		return ""
	}
	return order
}
