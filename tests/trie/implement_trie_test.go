package trie_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/trie"
)

func TestTrie_InsertAndSearch(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Search for 'apple' should return true")
	}
	if trie.Search("app") {
		t.Errorf("Search for 'app' should return false")
	}
	if !trie.StartsWith("app") {
		t.Errorf("StartsWith for 'app' should return true")
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Search for 'app' should return true after insertion")
	}
}

func TestTrie_InsertPrefixOfExistingWord(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("apple")
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Search for 'app' should return true after inserting prefix of an existing word")
	}
	if !trie.Search("apple") {
		t.Errorf("Search for 'apple' should return true")
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("banana")
	if !trie.StartsWith("ban") {
		t.Errorf("StartsWith for 'ban' should return true")
	}
	if !trie.StartsWith("bana") {
		t.Errorf("StartsWith for 'bana' should return true")
	}
	if trie.StartsWith("bananaa") {
		t.Errorf("StartsWith for 'bananaa' should return false")
	}
	if trie.StartsWith("xyz") {
		t.Errorf("StartsWith for 'xyz' should return false")
	}
}

func TestTrie_InsertDuplicateWords(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("car")
	trie.Insert("car")
	trie.Insert("car")
	if !trie.Search("car") {
		t.Errorf("Search for 'car' should return true after inserting duplicate words")
	}
	if !trie.StartsWith("ca") {
		t.Errorf("StartsWith for 'ca' should return true")
	}
}

func TestTrie_HandleEmptyString(t *testing.T) {
	trie := trie.Constructor()

	trie.Insert("")
	if !trie.Search("") {
		t.Errorf("Search for empty string should return true after insertion")
	}
	if !trie.StartsWith("") {
		t.Errorf("StartsWith for empty string should return true")
	}
}
