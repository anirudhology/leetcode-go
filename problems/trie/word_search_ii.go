package trie

import "github.com/anirudhology/leetcode-go/problems/util"

func WordSearchII(board [][]byte, words []string) []string {
	// List to store final output
	var result []string
	// Special case
	if len(board) == 0 || len(words) == 0 {
		return result
	}
	// Create trie with words
	root := createTrie(words)
	// Dimensions of the board
	m, n := len(board), len(board[0])
	// Process every cell on the board
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, m, n, root, &result)
		}
	}
	return result
}

func createTrie(words []string) *util.TrieNodeWithWord {
	root := &util.TrieNodeWithWord{}
	for _, word := range words {
		current := root
		for _, c := range word {
			if current.Children[int(c)-int('a')] == nil {
				current.Children[int(c)-int('a')] = &util.TrieNodeWithWord{}
			}
			current = current.Children[int(c)-int('a')]
		}
		current.Word = word
	}
	return root
}

func dfs(board [][]byte, i int, j int, m int, n int, node *util.TrieNodeWithWord, result *[]string) {
	// Handle out of bounds indices
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	c := board[i][j]
	// If we have visited this cell already
	if c == '#' {
		return
	}
	index := int(c - 'a')
	if node.Children[index] == nil {
		return
	}
	node = node.Children[index]
	if len(node.Word) > 0 {
		*result = append(*result, node.Word)
		node.Word = "" // Mark as visited
	}
	board[i][j] = '#'
	// Perform DFS
	dfs(board, i-1, j, m, n, node, result)
	dfs(board, i+1, j, m, n, node, result)
	dfs(board, i, j-1, m, n, node, result)
	dfs(board, i, j+1, m, n, node, result)
	board[i][j] = c
}
