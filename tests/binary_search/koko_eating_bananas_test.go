package binary_search_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestKokoEatingBananas(t *testing.T) {
	tests := []struct {
		piles []int
		h     int
		want  int
	}{
		// Edge cases
		{piles: []int{}, h: 5, want: 0},
		{piles: []int{3, 6, 7, 11}, h: 0, want: 0},
		// Normal cases
		{piles: []int{3, 6, 7, 11}, h: 8, want: 4},
		{piles: []int{30, 11, 23, 4, 20}, h: 5, want: 30},
		{piles: []int{30, 11, 23, 4, 20}, h: 6, want: 23},
		// Special cases
		{piles: []int{1000000000}, h: 2, want: 500000000},
		{piles: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, h: 10, want: 1},
		{piles: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, h: 5, want: 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := binary_search.KokoEatingBananas(tt.piles, tt.h)
			if got != tt.want {
				t.Errorf("KokoEatingBananas() = %v, want %v", got, tt.want)
			}
		})
	}
}
