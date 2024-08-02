package trie

import "github.com/anirudhology/leetcode-go/problems/util"

type WordDictionary struct {
	Root *util.TrieNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{Root: &util.TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	if this.Search(word) {
		return
	}
	current := this.Root
	for i := 0; i < len(word); i++ {
		if current.Children[int(word[i])-int('a')] == nil {
			current.Children[int(word[i])-int('a')] = &util.TrieNode{Content: word[i]}
		}
		current = current.Children[int(word[i])-int('a')]
	}
	current.IsWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchHelper(word, 0, this.Root)
}

func (this *WordDictionary) SearchHelper(word string, index int, node *util.TrieNode) bool {
	if index == len(word) {
		return node.IsWord
	}
	if word[index] == '.' {
		for _, child := range node.Children {
			if child != nil && this.SearchHelper(word, index+1, child) {
				return true
			}
		}
	} else {
		return node.Children[int(word[index])-int('a')] != nil && this.SearchHelper(word, index+1, node.Children[int(word[index])-int('a')])
	}
	return false
}
