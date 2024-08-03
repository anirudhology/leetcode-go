package trie_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/trie"
	"github.com/stretchr/testify/assert"
)

func TestWordSearchII(t *testing.T) {
	// Test Case 1: Empty board
	board := [][]byte{}
	words := []string{"word"}
	expected := []string{}
	result := trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 2: Empty words
	board = [][]byte{{'a'}}
	words = []string{}
	expected = []string{}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 3: Single word match
	board = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	words = []string{"abc"}
	expected = []string{"abc"}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 4: Multiple words match
	board = [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words = []string{"oath", "pea", "eat", "rain"}
	expected = []string{"oath", "eat"}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 5: No words match
	board = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	words = []string{"xyz", "mno"}
	expected = []string{}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 6: Word partially found
	board = [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}
	words = []string{"see", "abcced"}
	expected = []string{"see", "abcced"}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 7: Words overlapping
	board = [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}
	words = []string{"abcced", "see"}
	expected = []string{"abcced", "see"}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)

	// Test Case 8: Board with characters not in words
	board = [][]byte{
		{'x', 'y', 'z'},
		{'w', 'v', 'u'},
	}
	words = []string{"abc", "def"}
	expected = []string{}
	result = trie.WordSearchII(board, words)
	assert.ElementsMatch(t, expected, result)
}
