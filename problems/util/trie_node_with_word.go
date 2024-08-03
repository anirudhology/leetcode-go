package util

type TrieNodeWithWord struct {
	Word     string
	Children [26]*TrieNodeWithWord
}
