package bfs_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bfs"
)

func TestLadderLengthWordExists(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result := bfs.WordLadder("hit", "cog", wordList)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestLadderLengthWordDoesNotExist(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log"}
	result := bfs.WordLadder("hit", "cog", wordList)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestLadderLengthBeginWordEqualsEndWord(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result := bfs.WordLadder("hit", "hit", wordList)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestLadderLengthEmptyWordList(t *testing.T) {
	wordList := []string{}
	result := bfs.WordLadder("hit", "cog", wordList)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
