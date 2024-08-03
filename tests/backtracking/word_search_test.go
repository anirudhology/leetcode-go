package backtracking_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestExist(t *testing.T) {
	testCases := []struct {
		board  [][]byte
		word   string
		result bool
	}{
		// Test cases
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
		{[][]byte{{'A'}}, "A", true},
		{[][]byte{{'A'}}, "B", false},
		{[][]byte{}, "A", false},
		{[][]byte{{'A', 'B'}, {'C', 'D'}}, "ABCD", false},
	}

	for _, tc := range testCases {
		result := backtracking.WordSearch(tc.board, tc.word)
		if result != tc.result {
			t.Errorf("expected %v, got %v for board: %v and word: %s", tc.result, result, tc.board, tc.word)
		}
	}
}
