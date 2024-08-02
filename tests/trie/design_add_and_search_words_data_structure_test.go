package trie

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/trie"
	"github.com/stretchr/testify/assert"
)

func TestWordDictionary(t *testing.T) {
	wordDictionary := trie.WordDictionaryConstructor()

	testCases := []struct {
		description string
		operation   string
		word        string
		expected    bool
	}{
		{"Add 'hello' and search 'hello'", "add", "hello", true},
		{"Search 'hello'", "search", "hello", true},
		{"Search 'hell'", "search", "hell", false},
		{"Search 'helloo'", "search", "helloo", false},
		{"Search 'hell.'", "search", "hell.", true},
		{"Search 'hel.o'", "search", "hel.o", true},

		{"Add 'world' and search 'world'", "add", "world", true},
		{"Search 'world'", "search", "world", true},
		{"Search 'worl'", "search", "worl", false},

		{"Search 'h.llo'", "search", "h.llo", true},
		{"Search 'he.lo'", "search", "he.lo", true},
		{"Search 'hell.'", "search", "hell.", true},
		{"Search '.....'", "search", ".....", true},
		{"Search '....'", "search", "....", false},

		{"Search empty string in an empty dictionary", "search", "", false},

		{"Add empty string and search empty string", "add", "", true},
		{"Search '.'", "search", ".", false},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			switch tc.operation {
			case "add":
				wordDictionary.AddWord(tc.word)
				assert.True(t, wordDictionary.Search(tc.word))
			case "search":
				assert.Equal(t, tc.expected, wordDictionary.Search(tc.word))
			}
		})
	}
}
