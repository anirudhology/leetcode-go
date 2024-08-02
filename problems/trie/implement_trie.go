package trie

import "github.com/anirudhology/leetcode-go/problems/util"

type Trie struct {
	Root *util.TrieNode
}

func Constructor() Trie {
	return Trie{Root: &util.TrieNode{}}
}

func (this *Trie) Insert(word string) {
	current := this.Root
	for i := 0; i < len(word); i++ {
		charIndex := int(word[i]) - int('a')
		if current.Children[charIndex] == nil {
			current.Children[charIndex] = &util.TrieNode{Content: word[i]}
		}
		current = current.Children[charIndex]
	}
	current.IsWord = true
}

func (this *Trie) Search(word string) bool {
	current := this.Root
	for i := 0; i < len(word); i++ {
		charIndex := int(word[i]) - int('a')
		if current.Children[charIndex] == nil {
			return false
		}
		current = current.Children[charIndex]
	}
	return current.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.Root
	for i := 0; i < len(prefix); i++ {
		charIndex := int(prefix[i]) - int('a')
		if current.Children[charIndex] == nil {
			return false
		}
		current = current.Children[charIndex]
	}
	return true
}
