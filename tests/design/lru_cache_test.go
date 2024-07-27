package design_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/design"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	cache := design.Constructor(2)

	// Test initial state
	assert.Equal(t, -1, cache.Get(1))

	// Test put operation
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))

	// Test eviction of least recently used item
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))

	// Test further put operations and evictions
	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))

	// Test updating existing key
	cache.Put(4, 40)
	assert.Equal(t, 40, cache.Get(4))

	// Test capacity constraints
	cache.Put(5, 5)
	assert.Equal(t, -1, cache.Get(3))
	assert.Equal(t, 40, cache.Get(4))
	assert.Equal(t, 5, cache.Get(5))
}

func TestLRUCacheEdgeCases(t *testing.T) {
	cache := design.Constructor(1)

	// Test single item capacity
	cache.Put(1, 1)
	assert.Equal(t, 1, cache.Get(1))
	cache.Put(2, 2)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 2, cache.Get(2))
}

func TestLRUCacheUpdateHead(t *testing.T) {
	cache := design.Constructor(2)

	// Test internal workings of updateHead
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(1)) // 1 should be evicted
	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(2)) // 2 should be evicted
	assert.Equal(t, 3, cache.Get(3))  // 3 should be present
	assert.Equal(t, 4, cache.Get(4))  // 4 should be present
}

func TestLRUCacheDeleteEntry(t *testing.T) {
	cache := design.Constructor(2)

	// Test internal workings of deleteEntry
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(1)) // 1 should be evicted
	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(2)) // 2 should be evicted
	assert.Equal(t, 3, cache.Get(3))  // 3 should be present
	assert.Equal(t, 4, cache.Get(4))  // 4 should be present
}

func TestLRUCacheFullCoverage(t *testing.T) {
	cache := design.Constructor(3)

	// Testing full coverage
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	assert.Equal(t, 1, cache.Get(1))  // Access 1
	cache.Put(4, 4)                   // This should evict key 2
	assert.Equal(t, -1, cache.Get(2)) // 2 should be evicted
	cache.Put(5, 5)                   // This should evict key 3
	assert.Equal(t, -1, cache.Get(3)) // 3 should be evicted
	assert.Equal(t, 4, cache.Get(4))  // 4 should be present
	assert.Equal(t, 5, cache.Get(5))  // 5 should be present
}
