package heaps_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/heaps"
)

func TestAddNumAndFindMedian(t *testing.T) {
	finder := heaps.MedianFinderConstructor()

	finder.AddNum(1)
	if median := finder.FindMedian(); median != 1.0 {
		t.Errorf("Expected median 1.0, but got %f", median)
	}

	finder.AddNum(2)
	if median := finder.FindMedian(); median != 1.5 {
		t.Errorf("Expected median 1.5, but got %f", median)
	}

	finder.AddNum(3)
	if median := finder.FindMedian(); median != 2.0 {
		t.Errorf("Expected median 2.0, but got %f", median)
	}

	finder.AddNum(4)
	if median := finder.FindMedian(); median != 2.5 {
		t.Errorf("Expected median 2.5, but got %f", median)
	}

	finder.AddNum(5)
	if median := finder.FindMedian(); median != 3.0 {
		t.Errorf("Expected median 3.0, but got %f", median)
	}
}

func TestSingleElement(t *testing.T) {
	finder := heaps.MedianFinderConstructor()

	finder.AddNum(10)
	if median := finder.FindMedian(); median != 10.0 {
		t.Errorf("Expected median 10.0, but got %f", median)
	}
}
