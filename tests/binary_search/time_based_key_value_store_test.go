package binary_search_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestTimeMap(t *testing.T) {
	timeMap := binary_search.Constructor()

	// Test setting and getting a single value
	timeMap.Set("foo", "bar", 1)
	result := timeMap.Get("foo", 1)
	if result != "bar" {
		t.Errorf("Expected 'bar', but got '%s'", result)
	}

	// Test getting a value with a timestamp greater than the set timestamp
	result = timeMap.Get("foo", 3)
	if result != "bar" {
		t.Errorf("Expected 'bar', but got '%s'", result)
	}

	// Test setting a new value with a higher timestamp
	timeMap.Set("foo", "bar2", 4)
	result = timeMap.Get("foo", 4)
	if result != "bar2" {
		t.Errorf("Expected 'bar2', but got '%s'", result)
	}

	// Test getting the latest value with a timestamp greater than the latest set timestamp
	result = timeMap.Get("foo", 5)
	if result != "bar2" {
		t.Errorf("Expected 'bar2', but got '%s'", result)
	}

	// Test getting a value with a timestamp less than the set timestamp
	result = timeMap.Get("foo", 2)
	if result != "bar" {
		t.Errorf("Expected 'bar', but got '%s'", result)
	}

	// Test getting a value for a key that does not exist
	result = timeMap.Get("bar", 1)
	if result != "" {
		t.Errorf("Expected '', but got '%s'", result)
	}

	// Test setting multiple values with different keys
	timeMap.Set("foo1", "bar1", 1)
	timeMap.Set("foo2", "bar2", 2)

	// Test getting values for different keys
	result = timeMap.Get("foo1", 1)
	if result != "bar1" {
		t.Errorf("Expected 'bar1', but got '%s'", result)
	}

	result = timeMap.Get("foo2", 2)
	if result != "bar2" {
		t.Errorf("Expected 'bar2', but got '%s'", result)
	}
}
