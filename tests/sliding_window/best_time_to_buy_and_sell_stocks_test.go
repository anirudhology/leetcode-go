package sliding_window_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestBestTimeToBuyAndSellStocks(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"Empty prices", []int{}, 0},
		{"Single price", []int{10}, 0},
		{"All increasing", []int{1, 2, 3, 4, 5}, 4},
		{"All decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"Mixed prices", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Prices with no profit", []int{7, 6, 4, 3, 1}, 0},
		{"Prices with multiple peaks", []int{3, 3, 5, 0, 0, 3, 1, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliding_window.BestTimeToBuyAndSellStocks(tt.prices)
			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
