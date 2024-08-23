package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestMaxProfit_SimpleCase(t *testing.T) {
	result := dynamic_programming.BestTimeToBuyAndSellStocksWithCooldown([]int{1, 2, 3, 0, 2})
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestMaxProfit_EmptyPrices(t *testing.T) {
	result := dynamic_programming.BestTimeToBuyAndSellStocksWithCooldown([]int{})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMaxProfit_OneDay(t *testing.T) {
	result := dynamic_programming.BestTimeToBuyAndSellStocksWithCooldown([]int{1})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMaxProfit_TwoDays_Profit(t *testing.T) {
	result := dynamic_programming.BestTimeToBuyAndSellStocksWithCooldown([]int{1, 2})
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestMaxProfit_TwoDays_Loss(t *testing.T) {
	result := dynamic_programming.BestTimeToBuyAndSellStocksWithCooldown([]int{2, 1})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMaxProfit_AllDecreasingPrices(t *testing.T) {
	result := dynamic_programming.BestTimeToBuyAndSellStocksWithCooldown([]int{5, 4, 3, 2, 1})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
