package math_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestReversePositiveNumber(t *testing.T) {
	if got := math.ReverseInteger(123); got != 321 {
		t.Errorf("reverse(123) = %d; want 321", got)
	}
}

func TestReverseNegativeNumber(t *testing.T) {
	if got := math.ReverseInteger(-123); got != -321 {
		t.Errorf("reverse(-123) = %d; want -321", got)
	}
}

func TestReverseZero(t *testing.T) {
	if got := math.ReverseInteger(0); got != 0 {
		t.Errorf("reverse(0) = %d; want 0", got)
	}
}

func TestReverseOverflow(t *testing.T) {
	if got := math.ReverseInteger(1534236469); got != 0 {
		t.Errorf("reverse(1534236469) = %d; want 0", got)
	}
}

func TestReverseSingleDigit(t *testing.T) {
	if got := math.ReverseInteger(9); got != 9 {
		t.Errorf("reverse(9) = %d; want 9", got)
	}
}
