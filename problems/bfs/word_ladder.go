package bfs

import "container/list"

func WordLadder(beginWord string, endWord string, wordList []string) int {
	// Convert the list of words to set for O(1) lookup
	words := make(map[string]bool)
	for _, word := range wordList {
		words[word] = true
	}
	// Check if the end word is in the set or not
	if _, ok := words[endWord]; !ok {
		return 0
	}
	// Queue to perform BFS
	nodes := list.New()
	nodes.PushBack(beginWord)
	// Since we are going to start from beginword, we will be
	// at the first rung of the ladder
	level := 1
	// Process the queue until it is empty or we find the endWord
	for nodes.Len() > 0 {
		size := nodes.Len()
		level++
		for i := 0; i < size; i++ {
			// Get the current word
			currentWord := nodes.Remove(nodes.Front()).(string)
			for j := 0; j < len(currentWord); j++ {
				// Try for every combination of lowercase characters
				for c := 'a'; c <= 'z'; c++ {
					newWord := currentWord[:j] + string(c) + currentWord[j+1:]
					// Check if this word is present in the set
					if _, ok := words[newWord]; ok {
						if newWord == endWord {
							return level
						}
						delete(words, newWord)
						nodes.PushBack(newWord)
					}
				}
			}
		}
	}
	return 0
}
