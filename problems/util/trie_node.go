package util

type TrieNode struct {
	Content  byte
	Children [26]*TrieNode
	IsWord   bool
}
